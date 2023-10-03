package modules

import (
	"github.com/DisgoOrg/disgohook/api"
	"github.com/tnyim/jungletv/proto"
	"github.com/tnyim/jungletv/server/components/chatmanager"
	"github.com/tnyim/jungletv/server/components/mediaqueue"
	"github.com/tnyim/jungletv/server/components/pointsmanager"
)

// Dependencies is a "everything and the kitchen sink" struct used for injection of singleton dependencies in modules
type Dependencies struct {
	ModLogWebhook          api.WebhookClient
	ChatManager            *chatmanager.Manager
	PointsManager          *pointsmanager.Manager
	MediaQueue             *mediaqueue.MediaQueue
	OtherMediaQueueMethods OtherMediaQueueMethods
}

type OtherMediaQueueMethods interface {
	MediaEnqueuingRestriction() proto.AllowedMediaEnqueuingType
	SetMediaEnqueuingRestriction(restriction proto.AllowedMediaEnqueuingType, password string)
	NewQueueEntriesAllUnskippable() bool
	SetNewQueueEntriesAllUnskippable(bool)
}
