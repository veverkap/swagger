package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Subscription struct {
AccountId uuid.UUID
ArchivedStorageLimit int
Balance int
CancellationOption *string
CancellationReason *string
CancelledAt time.Time
DeletedAt time.Time
Id uuid.UUID
InsertedAt time.Time
LastPaymentAt time.Time
MemberLimit int
NextBillAt time.Time
OnTrial bool
Plan *Plan
PlanId uuid.UUID
PromotionExpiresAt time.Time
PromotionId uuid.UUID
RequiresAutoscaling bool
StorageLimit int
SubscriptionEndAt time.Time
TotalArchivedStorageLimit int
TotalLifetimeFileLimit int
TotalMemberLimit int
TotalProjectLimit int
TotalStorageLimit int
TotalUserLimit int
UpdatedAt time.Time
UserLimit int
}
