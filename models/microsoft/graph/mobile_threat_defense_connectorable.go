package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileThreatDefenseConnectorable 
type MobileThreatDefenseConnectorable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAndroidDeviceBlockedOnMissingPartnerData()(*bool)
    GetAndroidEnabled()(*bool)
    GetIosDeviceBlockedOnMissingPartnerData()(*bool)
    GetIosEnabled()(*bool)
    GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPartnerState()(*MobileThreatPartnerTenantState)
    GetPartnerUnresponsivenessThresholdInDays()(*int32)
    GetPartnerUnsupportedOsVersionBlocked()(*bool)
    SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)()
    SetAndroidEnabled(value *bool)()
    SetIosDeviceBlockedOnMissingPartnerData(value *bool)()
    SetIosEnabled(value *bool)()
    SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPartnerState(value *MobileThreatPartnerTenantState)()
    SetPartnerUnresponsivenessThresholdInDays(value *int32)()
    SetPartnerUnsupportedOsVersionBlocked(value *bool)()
}
