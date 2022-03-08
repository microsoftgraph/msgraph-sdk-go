package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PhysicalOfficeAddressable 
type PhysicalOfficeAddressable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCity()(*string)
    GetCountryOrRegion()(*string)
    GetOfficeLocation()(*string)
    GetPostalCode()(*string)
    GetState()(*string)
    GetStreet()(*string)
    SetCity(value *string)()
    SetCountryOrRegion(value *string)()
    SetOfficeLocation(value *string)()
    SetPostalCode(value *string)()
    SetState(value *string)()
    SetStreet(value *string)()
}
