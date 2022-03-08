package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GeoCoordinatesable 
type GeoCoordinatesable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAltitude()(*float64)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    SetAltitude(value *float64)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
}
