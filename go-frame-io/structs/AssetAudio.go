package structs

import (
				"github.com/google/uuid"
				)
type AssetAudio struct {
Type *string //asset_audio
AudioTracks []interface{}
Id uuid.UUID
}
