package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Project struct {
ArchiveStatus *string //standard,archiving,unarchiving,archived
ArchivedAt time.Time
ArchivedFileCount int
ArchivedStorage int
CollaboratorCount int
Collaborators []interface{}
DeletedAt time.Time
Description *string
FileCount int
FolderCount int
Id uuid.UUID
IgnoreArchive bool
InsertedAt time.Time
InviteUrl *string
Name *string
OwnerId uuid.UUID
Private bool
ProjectPreferences *ProjectPreferences
ReadOnly bool
RootAsset *Asset
RootAssetId uuid.UUID
Shared bool
Storage int
Team *Team
TeamId uuid.UUID
UpdatedAt time.Time
UserPermissions *UserPermissions
UserPreferences *ProjectPreferences
Watermark *Watermark
}
