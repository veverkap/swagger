package structs

import (
				"github.com/google/uuid"
				)
type AssetSubtitle struct {
Type *string //asset_subtitle
Id uuid.UUID
SubtitleTracks []interface{}
}
