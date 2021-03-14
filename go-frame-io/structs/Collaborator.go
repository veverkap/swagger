package structs

import (
				"github.com/google/uuid"
				)
type Collaborator struct {
Type *string //collaborator,pending_collaborator
Email *string
ProjectId uuid.UUID
User *User
UserId uuid.UUID
}
