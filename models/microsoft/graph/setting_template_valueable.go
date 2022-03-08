package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SettingTemplateValueable 
type SettingTemplateValueable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDefaultValue()(*string)
    GetDescription()(*string)
    GetName()(*string)
    GetType()(*string)
    SetDefaultValue(value *string)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetType(value *string)()
}
