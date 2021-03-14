package structs

import (
				"time"
				
				"github.com/google/uuid"
				)
type Plan struct {
ArchivedStorageLimit int
Autoscaling bool
AvailableFeatures *AvailableFeatures
CollaboratorLimit int
Cost int
DefaultPlan bool
DeletedAt time.Time
Enterprise bool
FileLimit int
Id uuid.UUID
InsertedAt time.Time
LifetimeFileLimit int
MemberLimit int
Name *string
PaymentMethod *string //stripe,check
Period *string //monthly,yearly,semiannually
ProjectLimit int
StorageLimit int
TeamLimit int
Tier *string //free,starter,pro,team,business,enterprise
Title *string
UpdatedAt time.Time
UserLimit int
UserMax int
Version int
}
