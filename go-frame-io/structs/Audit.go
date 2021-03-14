package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Audit struct {
AccountId uuid.UUID
Action *string
Actor *interface{}
ActorId uuid.UUID
ApplicationId uuid.UUID
InsertedAt time.Time
IpAddress *string
ItemId uuid.UUID
ItemType *string
RequestId uuid.UUID
Resource interface{}
TeamId uuid.UUID
UpdatedAt time.Time
}
