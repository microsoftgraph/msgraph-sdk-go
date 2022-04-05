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
    GetClientAppTypes()([]ConditionalAccessClientApp)
    GetDevices()(ConditionalAccessDevicesable)
    GetLocations()(ConditionalAccessLocationsable)
    GetPlatforms()(ConditionalAccessPlatformsable)
    GetSignInRiskLevels()([]RiskLevel)
    GetUserRiskLevels()([]RiskLevel)
    GetUsers()(ConditionalAccessUsersable)
    SetApplications(value ConditionalAccessApplicationsable)()
    SetClientApplications(value ConditionalAccessClientApplicationsable)()
    SetClientAppTypes(value []ConditionalAccessClientApp)()
    SetDevices(value ConditionalAccessDevicesable)()
    SetLocations(value ConditionalAccessLocationsable)()
    SetPlatforms(value ConditionalAccessPlatformsable)()
    SetSignInRiskLevels(value []RiskLevel)()
    SetUserRiskLevels(value []RiskLevel)()
    SetUsers(value ConditionalAccessUsersable)()
}
