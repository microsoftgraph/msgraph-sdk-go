package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PermissionGrantPolicyable 
type PermissionGrantPolicyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    PolicyBaseable
    GetExcludes()([]PermissionGrantConditionSetable)
    GetIncludes()([]PermissionGrantConditionSetable)
    SetExcludes(value []PermissionGrantConditionSetable)()
    SetIncludes(value []PermissionGrantConditionSetable)()
}
