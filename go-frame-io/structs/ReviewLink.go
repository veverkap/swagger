package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type ReviewLink struct {
AccessControl *ShareableEntityAccessControl
AllowApprovals bool
AvailableFeatures interface{}
BundleCount int
CurrentVersionOnly bool
DeletedAt time.Time
EnableComments bool
EnableDownloading bool
ExpiresAt time.Time
FolderItemCount int
HasPassword bool
Id uuid.UUID
InsertedAt time.Time
IsActive bool
IsAutoplay bool
ItemCount int
Items []interface{}
Name *string
NotifyOnView bool
Owner *User
OwnerId uuid.UUID
Password *string
Project *Project
ProjectId uuid.UUID
ReviewerCount int
ShortUrl *string
Team *Team
UpdatedAt time.Time
ViewCount int
}
