package structs

import (
				"github.com/google/uuid"
				)
type Comment struct {
Annotation *string
Completed bool
CompletedAt *string
CompleterId uuid.UUID
HasReplies bool
Id uuid.UUID
LikeCount int
Owner *User
OwnerId uuid.UUID
Text *string
Timestamp int
}
