package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Locationable 
type Locationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddress()(PhysicalAddressable)
    GetCoordinates()(OutlookGeoCoordinatesable)
    GetDisplayName()(*string)
    GetLocationEmailAddress()(*string)
    GetLocationType()(*LocationType)
    GetLocationUri()(*string)
    GetUniqueId()(*string)
    GetUniqueIdType()(*LocationUniqueIdType)
    SetAddress(value PhysicalAddressable)()
    SetCoordinates(value OutlookGeoCoordinatesable)()
    SetDisplayName(value *string)()
    SetLocationEmailAddress(value *string)()
    SetLocationType(value *LocationType)()
    SetLocationUri(value *string)()
    SetUniqueId(value *string)()
    SetUniqueIdType(value *LocationUniqueIdType)()
}
