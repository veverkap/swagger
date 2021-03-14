package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type AccountMember struct {
AcceptedAt time.Time
AccountId uuid.UUID
DeclinedAt time.Time
DeletedAt time.Time
InsertedAt time.Time
Role *string //admin,billing_manager,account_manager,member
UpdatedAt time.Time
UserId uuid.UUID
}
