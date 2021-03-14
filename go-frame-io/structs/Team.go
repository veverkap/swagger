package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Team struct {
Access *string //private,public,restricted
Account *Account
AccountId uuid.UUID
AdminOnlyActions *AdminOnlyAction
ArchivedStorage int
AssetLifecyclePolicy int
BackgroundColor *string
Bio *string
CollaboratorCount int
Color *string
CreatorId uuid.UUID
DarkTheme bool
DefaultBackgroundColor *string
DefaultColor *string
DefaultFontColor *string
DefaultSessionWatermarkTemplate *SessionWatermarkTemplate
DefaultSessionWatermarkTemplateId uuid.UUID
DeletedAt time.Time
DisableSbwmInternally bool
Duration int
EmailBranding *EmailBranding
FileCount int
FolderCount int
FontColor *string
Frames int
Id uuid.UUID
Image128 *string
Image256 *string
Image32 *string
Image64 *string
InsertedAt time.Time
Link *string
Location *string
MemberCount int
MemberLimit int
Name *string
ProjectCount int
SessionWatermarkTemplates []interface{}
SlackWebhook interface{}
Solo bool
Storage int
StorageLimit int
TeamImage *string
UpdatedAt time.Time
UploadUrl *string
UserRole interface{}
Watermark *Watermark
}
