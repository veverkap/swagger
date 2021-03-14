package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Webhook struct {
AccountId uuid.UUID
Active bool
AppId uuid.UUID
DeletedAt time.Time
Events []interface{}
Id uuid.UUID
InsertedAt time.Time
Name *string
ProjectId uuid.UUID
Secret *string
Team *Team
TeamId uuid.UUID
UpdatedAt time.Time
Url *string
}
