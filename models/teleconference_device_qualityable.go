package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeleconferenceDeviceQualityable 
type TeleconferenceDeviceQualityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallChainId()(*UUID)
    GetCloudServiceDeploymentEnvironment()(*string)
    GetCloudServiceDeploymentId()(*string)
    GetCloudServiceInstanceName()(*string)
    GetCloudServiceName()(*string)
    GetDeviceDescription()(*string)
    GetDeviceName()(*string)
    GetMediaLegId()(*UUID)
    GetMediaQualityList()([]TeleconferenceDeviceMediaQualityable)
    GetOdataType()(*string)
    GetParticipantId()(*UUID)
    SetCallChainId(value *UUID)()
    SetCloudServiceDeploymentEnvironment(value *string)()
    SetCloudServiceDeploymentId(value *string)()
    SetCloudServiceInstanceName(value *string)()
    SetCloudServiceName(value *string)()
    SetDeviceDescription(value *string)()
    SetDeviceName(value *string)()
    SetMediaLegId(value *UUID)()
    SetMediaQualityList(value []TeleconferenceDeviceMediaQualityable)()
    SetOdataType(value *string)()
    SetParticipantId(value *UUID)()
}
