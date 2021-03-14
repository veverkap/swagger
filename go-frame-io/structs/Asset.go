package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Asset struct {
AccountId uuid.UUID
ArchiveFrom uuid.UUID
AssetType *string //stream,image,document
Bundle bool
BundleView *string
CoverAssetId uuid.UUID
Creator *User
Frames int
HardDeletedAt time.Time
Id uuid.UUID
Index float64
IsBundleChild bool
IsHlsRequired bool
IsSessionWatermarked bool
ItemCount int
Label *string //approved,needs_review,in_progress
Metadata interface{}
MetadataFlags *MetadataFlags
Name *string
Original *string
ProjectId uuid.UUID
Properties interface{}
RequiredTranscodes *RequiredTranscodes
TeamId uuid.UUID
Type *string //file,folder,version_stack,bundle
UserPermissions *UserPermissions
ViewCount int
}
