package structs

import (
				"github.com/google/uuid"
				)
type AssetSearch struct {
AccountId uuid.UUID
CustomFields interface{}
Filter interface{}
Include *string
IncludeDeleted bool
Opts interface{}
Page int
PageSize int
ProjectId uuid.UUID
Properties interface{}
Q *string
Query *string
SharedProjects bool
Sort *string
TeamId *interface{}
Type *string
}
