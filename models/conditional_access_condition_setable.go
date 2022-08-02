package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessConditionSetable 
type ConditionalAccessConditionSetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplications()(ConditionalAccessApplicationsable)
    GetClientApplications()(ConditionalAccessClientApplicationsable)
    GetClientAppTypes()([]string)
    GetDevices()(ConditionalAccessDevicesable)
    GetLocations()(ConditionalAccessLocationsable)
    GetOdataType()(*string)
    GetPlatforms()(ConditionalAccessPlatformsable)
    GetServicePrincipalRiskLevels()([]string)
    GetSignInRiskLevels()([]string)
    GetUserRiskLevels()([]string)
    GetUsers()(ConditionalAccessUsersable)
    SetApplications(value ConditionalAccessApplicationsable)()
    SetClientApplications(value ConditionalAccessClientApplicationsable)()
    SetClientAppTypes(value []string)()
    SetDevices(value ConditionalAccessDevicesable)()
    SetLocations(value ConditionalAccessLocationsable)()
    SetOdataType(value *string)()
    SetPlatforms(value ConditionalAccessPlatformsable)()
    SetServicePrincipalRiskLevels(value []string)()
    SetSignInRiskLevels(value []string)()
    SetUserRiskLevels(value []string)()
    SetUsers(value ConditionalAccessUsersable)()
}
