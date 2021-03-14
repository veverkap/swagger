package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Action struct {
Active bool
AllowCollaborators bool
CreatorId uuid.UUID
DeletedAt time.Time
Description *string
Event *string
Id uuid.UUID
Image *string
InsertedAt time.Time
Name *string
Team *Team
TeamId uuid.UUID
UpdatedAt time.Time
UploadUrl *string
Url *string
Webhook *Webhook
}
