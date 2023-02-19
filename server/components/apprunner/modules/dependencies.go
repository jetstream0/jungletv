package modules

import "github.com/tnyim/jungletv/server/components/chatmanager"

// Dependencies is a "everything and the kitchen sink" struct used for injection of singleton dependencies in modules
type Dependencies struct {
	ChatManager *chatmanager.Manager
}
