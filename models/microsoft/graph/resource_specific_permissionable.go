package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResourceSpecificPermissionable 
type ResourceSpecificPermissionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetId()(*string)
    GetIsEnabled()(*bool)
    GetValue()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetIsEnabled(value *bool)()
    SetValue(value *string)()
}
