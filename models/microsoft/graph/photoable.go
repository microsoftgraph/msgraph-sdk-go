package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Photoable 
type Photoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCameraMake()(*string)
    GetCameraModel()(*string)
    GetExposureDenominator()(*float64)
    GetExposureNumerator()(*float64)
    GetFNumber()(*float64)
    GetFocalLength()(*float64)
    GetIso()(*int32)
    GetOrientation()(*int32)
    GetTakenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCameraMake(value *string)()
    SetCameraModel(value *string)()
    SetExposureDenominator(value *float64)()
    SetExposureNumerator(value *float64)()
    SetFNumber(value *float64)()
    SetFocalLength(value *float64)()
    SetIso(value *int32)()
    SetOrientation(value *int32)()
    SetTakenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
