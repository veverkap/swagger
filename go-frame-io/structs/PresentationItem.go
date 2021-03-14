package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type PresentationItem struct {
Asset *Asset
Id uuid.UUID
Index int
InsertedAt time.Time
Presentation *Presentation
UpdatedAt time.Time
}
