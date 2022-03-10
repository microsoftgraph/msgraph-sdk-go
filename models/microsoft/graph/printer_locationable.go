package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterLocationable 
type PrinterLocationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAltitudeInMeters()(*int32)
    GetBuilding()(*string)
    GetCity()(*string)
    GetCountryOrRegion()(*string)
    GetFloor()(*string)
    GetFloorDescription()(*string)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    GetOrganization()([]string)
    GetPostalCode()(*string)
    GetRoomDescription()(*string)
    GetRoomName()(*string)
    GetSite()(*string)
    GetStateOrProvince()(*string)
    GetStreetAddress()(*string)
    GetSubdivision()([]string)
    GetSubunit()([]string)
    SetAltitudeInMeters(value *int32)()
    SetBuilding(value *string)()
    SetCity(value *string)()
    SetCountryOrRegion(value *string)()
    SetFloor(value *string)()
    SetFloorDescription(value *string)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
    SetOrganization(value []string)()
    SetPostalCode(value *string)()
    SetRoomDescription(value *string)()
    SetRoomName(value *string)()
    SetSite(value *string)()
    SetStateOrProvince(value *string)()
    SetStreetAddress(value *string)()
    SetSubdivision(value []string)()
    SetSubunit(value []string)()
}
