package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleDefinitionable 
type UnifiedRoleDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInheritsPermissionsFrom()([]UnifiedRoleDefinitionable)
    GetIsBuiltIn()(*bool)
    GetIsEnabled()(*bool)
    GetResourceScopes()([]string)
    GetRolePermissions()([]UnifiedRolePermissionable)
    GetTemplateId()(*string)
    GetVersion()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInheritsPermissionsFrom(value []UnifiedRoleDefinitionable)()
    SetIsBuiltIn(value *bool)()
    SetIsEnabled(value *bool)()
    SetResourceScopes(value []string)()
    SetRolePermissions(value []UnifiedRolePermissionable)()
    SetTemplateId(value *string)()
    SetVersion(value *string)()
}
