package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleDefinitionable 
type RoleDefinitionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsBuiltIn()(*bool)
    GetRoleAssignments()([]RoleAssignmentable)
    GetRolePermissions()([]RolePermissionable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsBuiltIn(value *bool)()
    SetRoleAssignments(value []RoleAssignmentable)()
    SetRolePermissions(value []RolePermissionable)()
}
