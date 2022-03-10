package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppRoleable 
type AppRoleable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowedMemberTypes()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetId()(*string)
    GetIsEnabled()(*bool)
    GetOrigin()(*string)
    GetValue()(*string)
    SetAllowedMemberTypes(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetIsEnabled(value *bool)()
    SetOrigin(value *string)()
    SetValue(value *string)()
}
