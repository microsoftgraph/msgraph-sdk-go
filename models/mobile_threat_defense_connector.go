package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileThreatDefenseConnector 
type MobileThreatDefenseConnector struct {
    Entity
    // For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    androidDeviceBlockedOnMissingPartnerData *bool
    // For Android, set whether data from the data sync partner should be used during compliance evaluations
    androidEnabled *bool
    // For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    iosDeviceBlockedOnMissingPartnerData *bool
    // For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
    iosEnabled *bool
    // DateTime of last Heartbeat recieved from the Data Sync Partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Partner state of this tenant.
    partnerState *MobileThreatPartnerTenantState
    // Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
    partnerUnresponsivenessThresholdInDays *int32
    // Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
    partnerUnsupportedOsVersionBlocked *bool
}
// NewMobileThreatDefenseConnector instantiates a new MobileThreatDefenseConnector and sets the default values.
func NewMobileThreatDefenseConnector()(*MobileThreatDefenseConnector) {
    m := &MobileThreatDefenseConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileThreatDefenseConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileThreatDefenseConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileThreatDefenseConnector(), nil
}
// GetAndroidDeviceBlockedOnMissingPartnerData gets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceBlockedOnMissingPartnerData
    }
}
// GetAndroidEnabled gets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetAndroidEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileThreatDefenseConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["androidEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidEnabled(val)
        }
        return nil
    }
    res["iosDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["iosEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosEnabled(val)
        }
        return nil
    }
    res["lastHeartbeatDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHeartbeatDateTime(val)
        }
        return nil
    }
    res["partnerState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileThreatPartnerTenantState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerState(val.(*MobileThreatPartnerTenantState))
        }
        return nil
    }
    res["partnerUnresponsivenessThresholdInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerUnresponsivenessThresholdInDays(val)
        }
        return nil
    }
    res["partnerUnsupportedOsVersionBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerUnsupportedOsVersionBlocked(val)
        }
        return nil
    }
    return res
}
// GetIosDeviceBlockedOnMissingPartnerData gets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosDeviceBlockedOnMissingPartnerData
    }
}
// GetIosEnabled gets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetIosEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosEnabled
    }
}
// GetLastHeartbeatDateTime gets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
// GetPartnerState gets the partnerState property value. Partner state of this tenant.
func (m *MobileThreatDefenseConnector) GetPartnerState()(*MobileThreatPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
// GetPartnerUnresponsivenessThresholdInDays gets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.partnerUnresponsivenessThresholdInDays
    }
}
// GetPartnerUnsupportedOsVersionBlocked gets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.partnerUnsupportedOsVersionBlocked
    }
}
// Serialize serializes information the current object
func (m *MobileThreatDefenseConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("androidDeviceBlockedOnMissingPartnerData", m.GetAndroidDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidEnabled", m.GetAndroidEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosDeviceBlockedOnMissingPartnerData", m.GetIosDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosEnabled", m.GetIosEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastHeartbeatDateTime", m.GetLastHeartbeatDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerState() != nil {
        cast := (*m.GetPartnerState()).String()
        err = writer.WriteStringValue("partnerState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("partnerUnresponsivenessThresholdInDays", m.GetPartnerUnresponsivenessThresholdInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("partnerUnsupportedOsVersionBlocked", m.GetPartnerUnsupportedOsVersionBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidDeviceBlockedOnMissingPartnerData sets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)() {
    if m != nil {
        m.androidDeviceBlockedOnMissingPartnerData = value
    }
}
// SetAndroidEnabled sets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetAndroidEnabled(value *bool)() {
    if m != nil {
        m.androidEnabled = value
    }
}
// SetIosDeviceBlockedOnMissingPartnerData sets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(value *bool)() {
    if m != nil {
        m.iosDeviceBlockedOnMissingPartnerData = value
    }
}
// SetIosEnabled sets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetIosEnabled(value *bool)() {
    if m != nil {
        m.iosEnabled = value
    }
}
// SetLastHeartbeatDateTime sets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastHeartbeatDateTime = value
    }
}
// SetPartnerState sets the partnerState property value. Partner state of this tenant.
func (m *MobileThreatDefenseConnector) SetPartnerState(value *MobileThreatPartnerTenantState)() {
    if m != nil {
        m.partnerState = value
    }
}
// SetPartnerUnresponsivenessThresholdInDays sets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(value *int32)() {
    if m != nil {
        m.partnerUnresponsivenessThresholdInDays = value
    }
}
// SetPartnerUnsupportedOsVersionBlocked sets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(value *bool)() {
    if m != nil {
        m.partnerUnsupportedOsVersionBlocked = value
    }
}
