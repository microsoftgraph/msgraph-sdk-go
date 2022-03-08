package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceDetailable 
type DeviceDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetBrowser()(*string)
    GetDeviceId()(*string)
    GetDisplayName()(*string)
    GetIsCompliant()(*bool)
    GetIsManaged()(*bool)
    GetOperatingSystem()(*string)
    GetTrustType()(*string)
    SetBrowser(value *string)()
    SetDeviceId(value *string)()
    SetDisplayName(value *string)()
    SetIsCompliant(value *bool)()
    SetIsManaged(value *bool)()
    SetOperatingSystem(value *string)()
    SetTrustType(value *string)()
}
