package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesConditionalAccessSettingsable 
type OnPremisesConditionalAccessSettingsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEnabled()(*bool)
    GetExcludedGroups()([]string)
    GetIncludedGroups()([]string)
    GetOverrideDefaultRule()(*bool)
    SetEnabled(value *bool)()
    SetExcludedGroups(value []string)()
    SetIncludedGroups(value []string)()
    SetOverrideDefaultRule(value *bool)()
}
