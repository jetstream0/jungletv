package server

import (
	"context"
	"strings"

	"github.com/tnyim/jungletv/proto"
	"github.com/tnyim/jungletv/server/auth"
	authinterceptor "github.com/tnyim/jungletv/server/interceptors/auth"
)

func (s *grpcServer) serializeUserForAPI(ctx context.Context, user auth.User) *proto.User {
	userAddress := user.Address()
	fetchedUser, err := s.nicknameCache.GetOrFetchUser(ctx, userAddress)
	if err == nil && fetchedUser != nil && !fetchedUser.IsUnknown() {
		user = fetchedUser
	}

	s.vipUsersMutex.RLock()
	vipUserAppearance, isVip := s.vipUsers[userAddress]
	s.vipUsersMutex.RUnlock()

	roles := []proto.UserRole{}
	if isVip {
		switch vipUserAppearance {
		case vipUserAppearanceModerator:
			roles = append(roles, proto.UserRole_MODERATOR)
		case vipUserAppearanceVIP:
			roles = append(roles, proto.UserRole_VIP)
		case vipUserAppearanceVIPModerator:
			roles = append(roles, proto.UserRole_VIP, proto.UserRole_MODERATOR)
		}
	} else if user.ApplicationID() != "" {
		roles = append(roles, proto.UserRole_APPLICATION)
	} else if auth.UserPermissionLevelIsAtLeast(user, auth.AdminPermissionLevel) {
		roles = append(roles, proto.UserRole_MODERATOR)
	}

	mediaCount, requestedCurrent, err := s.mediaQueue.CountEnqueuedOrRecentlyPlayedMediaRequestedBy(ctx, user)
	if err == nil {
		switch {
		case mediaCount >= 10:
			roles = append(roles, proto.UserRole_TIER_3_REQUESTER)
		case mediaCount >= 5:
			roles = append(roles, proto.UserRole_TIER_2_REQUESTER)
		case mediaCount > 0:
			roles = append(roles, proto.UserRole_TIER_1_REQUESTER)
		}
		if requestedCurrent {
			roles = append(roles, proto.UserRole_CURRENT_ENTRY_REQUESTER)
		}
	}

	var nickname *string
	bannedFromChat, err := s.moderationStore.LoadUserBannedFromChat(ctx, userAddress, "")
	serializingForUser := authinterceptor.UserClaimsFromContext(ctx)
	if err == nil && (!bannedFromChat || (serializingForUser != nil && serializingForUser.Address() == userAddress)) {
		nickname = user.Nickname()
		if nickname != nil && strings.TrimSpace(*nickname) == "" {
			nickname = nil
		}
	}
	if id := user.ApplicationID(); nickname == nil && id != "" {
		nickname = &id
	}

	var status proto.UserStatus
	if appID := user.ApplicationID(); appID != "" {
		isRunning, _, _ := s.appRunner.IsRunning(appID)
		status = proto.UserStatus_USER_STATUS_OFFLINE
		if isRunning {
			status = proto.UserStatus_USER_STATUS_WATCHING
		}
	} else {
		status = s.rewardsHandler.GetSpectatorActivityStatus(userAddress)
	}

	return &proto.User{
		Address:  userAddress,
		Roles:    roles,
		Nickname: nickname,
		Status:   status,
	}
}

type vipUserAppearance int

const (
	vipUserAppearanceNormal vipUserAppearance = iota
	vipUserAppearanceModerator
	vipUserAppearanceVIP
	vipUserAppearanceVIPModerator
)

func (s *grpcServer) isVIPUser(user auth.User) bool {
	s.vipUsersMutex.RLock()
	defer s.vipUsersMutex.RUnlock()
	if user != nil && !user.IsUnknown() {
		_, present := s.vipUsers[user.Address()]
		return present
	}
	return false
}
