package structs

type AccountMembership struct {
Account *string //owner,admin,account_manager,billing_manager
Reviewer bool
Team *string //team_manager,member
}
