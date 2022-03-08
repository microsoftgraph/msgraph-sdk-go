package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingsable 
type DeviceManagementSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceComplianceCheckinThresholdDays()(*int32)
    GetIsScheduledActionEnabled()(*bool)
    GetSecureByDefault()(*bool)
    SetDeviceComplianceCheckinThresholdDays(value *int32)()
    SetIsScheduledActionEnabled(value *bool)()
    SetSecureByDefault(value *bool)()
}
