package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// StsPolicyable 
type StsPolicyable interface {
    PolicyBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppliesTo()([]DirectoryObjectable)
    GetDefinition()([]string)
    GetIsOrganizationDefault()(*bool)
    SetAppliesTo(value []DirectoryObjectable)()
    SetDefinition(value []string)()
    SetIsOrganizationDefault(value *bool)()
}
