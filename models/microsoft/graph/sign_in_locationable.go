package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SignInLocationable 
type SignInLocationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCity()(*string)
    GetCountryOrRegion()(*string)
    GetGeoCoordinates()(GeoCoordinatesable)
    GetState()(*string)
    SetCity(value *string)()
    SetCountryOrRegion(value *string)()
    SetGeoCoordinates(value GeoCoordinatesable)()
    SetState(value *string)()
}
