package soundcloud

import (
	"context"
	"math/big"
	"time"

	"github.com/bytedance/sonic"
	"github.com/palantir/stacktrace"
	"github.com/tnyim/jungletv/proto"
	"github.com/tnyim/jungletv/server/auth"
	"github.com/tnyim/jungletv/server/components/payment"
	"github.com/tnyim/jungletv/server/media"
	"github.com/tnyim/jungletv/types"
	"google.golang.org/protobuf/types/known/durationpb"
)

type queueEntrySoundCloudTrack struct {
	media.CommonQueueEntry
	media.CommonInfo
	id           string
	uploader     string
	artist       string
	permalink    string
	thumbnailURL string
}

func (e *queueEntrySoundCloudTrack) ProduceMediaQueueEntry(requestedBy auth.User, requestCost payment.Amount, unskippable, concealed bool, queueID string) media.QueueEntry {
	e.FillMediaQueueEntryFields(requestedBy, requestCost, unskippable, concealed, queueID)
	return e
}

func (e *queueEntrySoundCloudTrack) MediaID() (types.MediaType, string) {
	return types.MediaTypeSoundCloudTrack, e.id
}

func (e *queueEntrySoundCloudTrack) SerializeForAPIQueue(ctx context.Context) proto.IsQueueEntry_MediaInfo {
	info := &proto.QueueEntry_SoundcloudTrackData{
		SoundcloudTrackData: &proto.QueueSoundCloudTrackData{
			Id:           e.id,
			Title:        e.Title(),
			ThumbnailUrl: e.thumbnailURL,
			Uploader:     e.uploader,
			Artist:       e.artist,
			Permalink:    e.permalink,
		},
	}
	return info
}

type queueEntrySoundCloudTrackJsonRepresentation struct {
	QueueID      string
	Type         string
	ID           string
	Title        string
	Uploader     string
	Artist       string
	Permalink    string
	ThumbnailURL string
	Duration     time.Duration
	Offset       time.Duration
	RequestedBy  string
	RequestCost  *big.Int
	RequestedAt  time.Time
	Unskippable  bool
	Concealed    bool
	MovedBy      []string
}

func (e *queueEntrySoundCloudTrack) MarshalJSON() ([]byte, error) {
	j, err := sonic.Marshal(queueEntrySoundCloudTrackJsonRepresentation{
		QueueID:      e.PerformanceID(),
		Type:         string(types.MediaTypeSoundCloudTrack),
		ID:           e.id,
		Title:        e.Title(),
		Uploader:     e.uploader,
		Artist:       e.artist,
		Permalink:    e.permalink,
		ThumbnailURL: e.thumbnailURL,
		Duration:     e.Length(),
		Offset:       e.Offset(),
		RequestedBy:  e.RequestedBy().Address(),
		RequestCost:  e.RequestCost().Int,
		RequestedAt:  e.RequestedAt(),
		Unskippable:  e.Unskippable(),
		Concealed:    e.Concealed(),
		MovedBy:      e.MovedBy(),
	})
	if err != nil {
		return nil, stacktrace.Propagate(err, "error serializing queue entry %s", e.PerformanceID())
	}
	return j, nil
}

func (e *queueEntrySoundCloudTrack) UnmarshalJSON(b []byte) error {
	var t queueEntrySoundCloudTrackJsonRepresentation
	if err := sonic.Unmarshal(b, &t); err != nil {
		return stacktrace.Propagate(err, "error deserializing queue entry")
	}

	e.InitializeBase(e, e)
	e.SetQueueID(t.QueueID)
	e.id = t.ID
	e.SetTitle(t.Title)
	e.uploader = t.Uploader
	e.artist = t.Artist
	e.permalink = t.Permalink
	e.thumbnailURL = t.ThumbnailURL
	e.SetLength(t.Duration)
	e.SetOffset(t.Offset)
	e.SetRequestedBy(auth.NewAddressOnlyUser(t.RequestedBy))
	e.SetRequestCost(payment.NewAmount(t.RequestCost))
	e.SetRequestedAt(t.RequestedAt)
	e.SetUnskippable(t.Unskippable)
	e.SetConcealed(t.Concealed)
	for _, m := range t.MovedBy {
		e.SetAsMovedBy(auth.NewAddressOnlyUser(m))
	}
	return nil
}

func (e *queueEntrySoundCloudTrack) FillAPITicketMediaInfo(ticket *proto.EnqueueMediaTicket) {
	ticket.Length = durationpb.New(e.Length())
	ticket.Offset = durationpb.New(e.Offset())
	ticket.MediaInfo = &proto.EnqueueMediaTicket_SoundcloudTrackData{
		SoundcloudTrackData: &proto.QueueSoundCloudTrackData{
			Id:           e.id,
			Title:        e.Title(),
			Uploader:     e.uploader,
			Artist:       e.artist,
			Permalink:    e.permalink,
			ThumbnailUrl: e.thumbnailURL,
		},
	}
}

func (e *queueEntrySoundCloudTrack) ProduceCheckpointForAPI(ctx context.Context) *proto.MediaConsumptionCheckpoint {
	return &proto.MediaConsumptionCheckpoint{
		MediaInfo: &proto.MediaConsumptionCheckpoint_SoundcloudTrackData{
			SoundcloudTrackData: &proto.NowPlayingSoundCloudTrackData{
				Id: e.id,
			},
		},
	}
}
