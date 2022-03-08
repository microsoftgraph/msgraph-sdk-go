package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtectionAppable 
type WindowsInformationProtectionAppable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDenied()(*bool)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetProductName()(*string)
    GetPublisherName()(*string)
    SetDenied(value *bool)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetProductName(value *string)()
    SetPublisherName(value *string)()
}
