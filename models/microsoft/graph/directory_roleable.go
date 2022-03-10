package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DirectoryRoleable 
type DirectoryRoleable interface {
    DirectoryObjectable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetMembers()([]DirectoryObjectable)
    GetRoleTemplateId()(*string)
    GetScopedMembers()([]ScopedRoleMembershipable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetMembers(value []DirectoryObjectable)()
    SetRoleTemplateId(value *string)()
    SetScopedMembers(value []ScopedRoleMembershipable)()
}
