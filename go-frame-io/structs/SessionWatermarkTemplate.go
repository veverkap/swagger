package structs

import (
				"github.com/google/uuid"
				)
type SessionWatermarkTemplate struct {
AccountId uuid.UUID
AppDefault bool
CreatorId uuid.UUID
Id uuid.UUID
Name *string
WatermarkBlocks []interface{}
}
