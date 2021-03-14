package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Installation struct {
AccountId uuid.UUID
AppId uuid.UUID
AppVersion int
BotUserId uuid.UUID
DeletedAt time.Time
Id uuid.UUID
InstalledAt time.Time
}
