package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthorizationPolicyable 
type AuthorizationPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyBaseable
    GetAllowedToSignUpEmailBasedSubscriptions()(*bool)
    GetAllowedToUseSSPR()(*bool)
    GetAllowEmailVerifiedUsersToJoinOrganization()(*bool)
    GetAllowInvitesFrom()(*AllowInvitesFrom)
    GetBlockMsolPowerShell()(*bool)
    GetDefaultUserRolePermissions()(DefaultUserRolePermissionsable)
    GetGuestUserRoleId()(*UUID)
    SetAllowedToSignUpEmailBasedSubscriptions(value *bool)()
    SetAllowedToUseSSPR(value *bool)()
    SetAllowEmailVerifiedUsersToJoinOrganization(value *bool)()
    SetAllowInvitesFrom(value *AllowInvitesFrom)()
    SetBlockMsolPowerShell(value *bool)()
    SetDefaultUserRolePermissions(value DefaultUserRolePermissionsable)()
    SetGuestUserRoleId(value *UUID)()
}
