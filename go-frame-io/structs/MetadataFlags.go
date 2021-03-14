package structs

import (
				"github.com/google/uuid"
				)
type MetadataFlags struct {
Id uuid.UUID
Is360 bool
IsHdr bool
IsOriginalPlayerCompatible bool
}
