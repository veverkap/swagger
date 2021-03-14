package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Presentation struct {
AccessControl *ShareableEntityAccessControl
Asset *Asset
AssetId uuid.UUID
Assets []interface{}
Autoplay bool
AvailableFeatures interface{}
BackgroundColor *string
CanDownload bool
Color *string
DeletedAt time.Time
Description *string
Enabled bool
ExpiresAt time.Time
Format *string
Id uuid.UUID
IncludeExt bool
IncludeUploadDate bool
InsertedAt time.Time
Layout *string //blog,reel
Name *string
OwnerId uuid.UUID
Password *string
PresentationItems []interface{}
Project *Project
ProjectId uuid.UUID
ReviewerCount int
Secure bool
ShortUrl *string
Style *string
Team *Team
TextColor *string
Title *string
UpdatedAt time.Time
Vanity *string
ViewCount int
}
