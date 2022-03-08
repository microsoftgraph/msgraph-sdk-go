package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LicenseUnitsDetailable 
type LicenseUnitsDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetEnabled()(*int32)
    GetSuspended()(*int32)
    GetWarning()(*int32)
    SetEnabled(value *int32)()
    SetSuspended(value *int32)()
    SetWarning(value *int32)()
}
