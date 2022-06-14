package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateForBusinessConfigurationable 
type WindowsUpdateForBusinessConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutomaticUpdateMode()(*AutomaticUpdateMode)
    GetBusinessReadyUpdatesOnly()(*WindowsUpdateType)
    GetDeliveryOptimizationMode()(*WindowsDeliveryOptimizationMode)
    GetDriversExcluded()(*bool)
    GetFeatureUpdatesDeferralPeriodInDays()(*int32)
    GetFeatureUpdatesPaused()(*bool)
    GetFeatureUpdatesPauseExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInstallationSchedule()(WindowsUpdateInstallScheduleTypeable)
    GetMicrosoftUpdateServiceAllowed()(*bool)
    GetPrereleaseFeatures()(*PrereleaseFeatures)
    GetQualityUpdatesDeferralPeriodInDays()(*int32)
    GetQualityUpdatesPaused()(*bool)
    GetQualityUpdatesPauseExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAutomaticUpdateMode(value *AutomaticUpdateMode)()
    SetBusinessReadyUpdatesOnly(value *WindowsUpdateType)()
    SetDeliveryOptimizationMode(value *WindowsDeliveryOptimizationMode)()
    SetDriversExcluded(value *bool)()
    SetFeatureUpdatesDeferralPeriodInDays(value *int32)()
    SetFeatureUpdatesPaused(value *bool)()
    SetFeatureUpdatesPauseExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInstallationSchedule(value WindowsUpdateInstallScheduleTypeable)()
    SetMicrosoftUpdateServiceAllowed(value *bool)()
    SetPrereleaseFeatures(value *PrereleaseFeatures)()
    SetQualityUpdatesDeferralPeriodInDays(value *int32)()
    SetQualityUpdatesPaused(value *bool)()
    SetQualityUpdatesPauseExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
