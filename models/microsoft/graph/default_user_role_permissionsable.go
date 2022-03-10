package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DefaultUserRolePermissionsable 
type DefaultUserRolePermissionsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowedToCreateApps()(*bool)
    GetAllowedToCreateSecurityGroups()(*bool)
    GetAllowedToReadOtherUsers()(*bool)
    GetPermissionGrantPoliciesAssigned()([]string)
    SetAllowedToCreateApps(value *bool)()
    SetAllowedToCreateSecurityGroups(value *bool)()
    SetAllowedToReadOtherUsers(value *bool)()
    SetPermissionGrantPoliciesAssigned(value []string)()
}
