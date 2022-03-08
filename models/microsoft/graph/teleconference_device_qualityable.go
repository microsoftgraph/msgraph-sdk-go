package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeleconferenceDeviceQualityable 
type TeleconferenceDeviceQualityable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCallChainId()(*string)
    GetCloudServiceDeploymentEnvironment()(*string)
    GetCloudServiceDeploymentId()(*string)
    GetCloudServiceInstanceName()(*string)
    GetCloudServiceName()(*string)
    GetDeviceDescription()(*string)
    GetDeviceName()(*string)
    GetMediaLegId()(*string)
    GetMediaQualityList()([]TeleconferenceDeviceMediaQualityable)
    GetParticipantId()(*string)
    SetCallChainId(value *string)()
    SetCloudServiceDeploymentEnvironment(value *string)()
    SetCloudServiceDeploymentId(value *string)()
    SetCloudServiceInstanceName(value *string)()
    SetCloudServiceName(value *string)()
    SetDeviceDescription(value *string)()
    SetDeviceName(value *string)()
    SetMediaLegId(value *string)()
    SetMediaQualityList(value []TeleconferenceDeviceMediaQualityable)()
    SetParticipantId(value *string)()
}
