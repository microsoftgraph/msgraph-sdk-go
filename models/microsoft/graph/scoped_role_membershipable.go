package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ScopedRoleMembershipable 
type ScopedRoleMembershipable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAdministrativeUnitId()(*string)
    GetRoleId()(*string)
    GetRoleMemberInfo()(Identityable)
    SetAdministrativeUnitId(value *string)()
    SetRoleId(value *string)()
    SetRoleMemberInfo(value Identityable)()
}
