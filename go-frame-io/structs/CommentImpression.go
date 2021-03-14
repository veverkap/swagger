package structs

import (
				"github.com/google/uuid"
				)
type CommentImpression struct {
AssetId uuid.UUID
Count int
DeletedAt *string
Id uuid.UUID
InsertedAt *string
ReviewLinkId uuid.UUID
Type *string
UpdatedAt *string
User *User
UserId uuid.UUID
}
