package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type AnonymousUser struct {
DeletedAt time.Time
Email *string
Id uuid.UUID
InsertedAt time.Time
Name *string
NotificationPreferences *ProjectPreferences
UpdatedAt time.Time
UserDefaultColor *string
}
