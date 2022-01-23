package server

import (
	"context"
	"errors"
	"time"

	"github.com/palantir/stacktrace"
	"github.com/tnyim/jungletv/proto"
	"github.com/tnyim/jungletv/server/auth"
	"github.com/tnyim/jungletv/utils/event"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *grpcServer) MonitorQueue(r *proto.MonitorQueueRequest, stream proto.JungleTV_MonitorQueueServer) error {
	ctx := stream.Context()
	user := auth.UserClaimsFromContext(ctx)

	unregister := s.statsHandler.RegisterQueueSubscriber(user != nil && !user.IsUnknown())
	defer unregister()

	send := func() error {
		tokensExhausted := false
		if user != nil && !user.IsUnknown() {
			used, remaining, err := s.ownEntryRemovalRateLimiter.Get(ctx, user.Address())
			if err != nil {
				return stacktrace.Propagate(err, "")
			}
			// rate limiter memory store returns 0, 0 when it doesn't find a key, instead of returning the maximum for remaining...
			tokensExhausted = remaining == 0 && used != 0
		}

		queue := &proto.Queue{
			IsHeartbeat:            false,
			OwnEntryRemovalEnabled: !tokensExhausted && s.mediaQueue.RemovalOfOwnEntriesAllowed(),
		}
		entries := s.mediaQueue.Entries()
		queue.Entries = make([]*proto.QueueEntry, len(entries))
		for i, entry := range entries {
			queue.Entries[i] = entry.SerializeForAPI(ctx, s.userSerializer)
		}

		insertCursor, hasInsertCursor := s.mediaQueue.InsertCursor()
		if hasInsertCursor {
			queue.InsertCursor = &insertCursor
		}

		playingSince := s.mediaQueue.PlayingSince()
		if !playingSince.IsZero() {
			queue.PlayingSince = timestamppb.New(playingSince)
		}

		return stacktrace.Propagate(stream.Send(queue), "")
	}

	onQueueChanged, queueUpdatedU := s.mediaQueue.queueUpdated.Subscribe(event.AtLeastOnceGuarantee)
	defer queueUpdatedU()

	err := send()
	if err != nil {
		return stacktrace.Propagate(err, "")
	}

	heartbeat := time.NewTicker(5 * time.Second)
	defer heartbeat.Stop()

	for {
		var err error
		select {
		case <-onQueueChanged:
			err = send()
			if err != nil {
				return stacktrace.Propagate(err, "")
			}
		case <-heartbeat.C:
			err = stream.Send(&proto.Queue{
				IsHeartbeat: true,
			})
		case <-ctx.Done():
			return nil
		}
		if err != nil {
			return stacktrace.Propagate(err, "")
		}
	}
}

func (s *grpcServer) RemoveOwnQueueEntry(ctx context.Context, r *proto.RemoveOwnQueueEntryRequest) (*proto.RemoveOwnQueueEntryResponse, error) {
	user := auth.UserClaimsFromContext(ctx)
	if user == nil {
		// this should never happen, as the auth interceptors should have taken care of this for us
		return nil, status.Error(codes.Unauthenticated, "missing user claims")
	}

	err := s.mediaQueue.RemoveOwnEntry(r.Id, user)
	if err != nil {
		if errors.Is(err, ErrInsufficientPermissionsToRemoveEntry) {
			return nil, status.Error(codes.PermissionDenied, "insufficient permissions")
		}
		return nil, stacktrace.Propagate(err, "failed to remove queue entry")
	}

	_, _, _, ok, err := s.ownEntryRemovalRateLimiter.Take(ctx, user.Address())
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	if !ok {
		return nil, status.Errorf(codes.ResourceExhausted, "rate limit reached")
	}

	s.log.Printf("Queue entry with ID %s removed by its requester", r.Id)

	return &proto.RemoveOwnQueueEntryResponse{}, nil
}
