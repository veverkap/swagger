package structs

import (
				"github.com/google/uuid"
				)
type LibrarySearch struct {
AccountId uuid.UUID
CustomFields interface{}
Filter interface{}
IncludeDeleted bool
Opts interface{}
Page int
PageSize int
Q *string
Query *string
Sort *string
}
