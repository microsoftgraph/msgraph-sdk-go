package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CalendarPermissionable 
type CalendarPermissionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowedRoles()([]CalendarRoleType)
    GetEmailAddress()(EmailAddressable)
    GetIsInsideOrganization()(*bool)
    GetIsRemovable()(*bool)
    GetRole()(*CalendarRoleType)
    SetAllowedRoles(value []CalendarRoleType)()
    SetEmailAddress(value EmailAddressable)()
    SetIsInsideOrganization(value *bool)()
    SetIsRemovable(value *bool)()
    SetRole(value *CalendarRoleType)()
}
