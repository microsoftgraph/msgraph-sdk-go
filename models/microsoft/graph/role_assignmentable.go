package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleAssignmentable 
type RoleAssignmentable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetResourceScopes()([]string)
    GetRoleDefinition()(RoleDefinitionable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetResourceScopes(value []string)()
    SetRoleDefinition(value RoleDefinitionable)()
}
