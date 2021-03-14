package structs

import (
				"github.com/google/uuid"
				)
type User struct {
AccountId uuid.UUID
Email *string
Id uuid.UUID
Name *string
}
