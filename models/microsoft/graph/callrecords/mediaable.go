package callrecords

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Mediaable 
type Mediaable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCalleeDevice()(DeviceInfoable)
    GetCalleeNetwork()(NetworkInfoable)
    GetCallerDevice()(DeviceInfoable)
    GetCallerNetwork()(NetworkInfoable)
    GetLabel()(*string)
    GetStreams()([]MediaStreamable)
    SetCalleeDevice(value DeviceInfoable)()
    SetCalleeNetwork(value NetworkInfoable)()
    SetCallerDevice(value DeviceInfoable)()
    SetCallerNetwork(value NetworkInfoable)()
    SetLabel(value *string)()
    SetStreams(value []MediaStreamable)()
}
