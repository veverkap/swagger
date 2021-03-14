package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Account struct {
AccountDefaultColor *string
AccountMembers []interface{}
ArchivedStorage int
BillingEmails *string
City *string
CollaboratorCount int
CollaboratorRoleCount int
CompanyAddress *string
CompanyName *string
Country *string
DeletedAt time.Time
DelinquentAt time.Time
DisplayName *string
Duration int
FileCount int
FolderCount int
Frames int
Id uuid.UUID
Image *string
Image128 *string
Image256 *string
Image32 *string
Image64 *string
InsertedAt time.Time
InvoiceEmails []interface{}
LifetimeFileCount int
Line1 *string
LockedAt time.Time
MemberCount int
Owner *User
OwnerId uuid.UUID
Plan *Plan
PostalCode *string
ProjectCount int
State *string
Storage int
Subscription *Subscription
TeamCount int
Teams []interface{}
UnpaidAt time.Time
UpdatedAt time.Time
UploadUrl *string
UserCount int
Vat *string
Watermark *Watermark
}
