package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserFlowLanguageConfigurationable 
type UserFlowLanguageConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDefaultPages()([]UserFlowLanguagePageable)
    GetDisplayName()(*string)
    GetIsEnabled()(*bool)
    GetOverridesPages()([]UserFlowLanguagePageable)
    SetDefaultPages(value []UserFlowLanguagePageable)()
    SetDisplayName(value *string)()
    SetIsEnabled(value *bool)()
    SetOverridesPages(value []UserFlowLanguagePageable)()
}
