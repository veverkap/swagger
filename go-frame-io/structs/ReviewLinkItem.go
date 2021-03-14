package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type ReviewLinkItem struct {
Asset *Asset
AssetId uuid.UUID
DeletedAt time.Time
Index int
InsertedAt time.Time
ReviewLink *ReviewLink
ReviewLinkId uuid.UUID
UpdatedAt time.Time
}
