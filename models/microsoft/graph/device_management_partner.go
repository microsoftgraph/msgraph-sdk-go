package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementPartner struct {
    Entity
    // Partner display name
    displayName *string;
    // Whether device management partner is configured or not
    isConfigured *bool;
    // Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Partner App type. Possible values are: unknown, singleTenantApp, multiTenantApp.
    partnerAppType *DeviceManagementPartnerAppType;
    // Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
    partnerState *DeviceManagementPartnerTenantState;
    // Partner Single tenant App id
    singleTenantAppId *string;
    // DateTime in UTC when PartnerDevices will be marked as NonCompliant
    whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // DateTime in UTC when PartnerDevices will be removed
    whenPartnerDevicesWillBeRemovedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new deviceManagementPartner and sets the default values.
func NewDeviceManagementPartner()(*DeviceManagementPartner) {
    m := &DeviceManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. Partner display name
func (m *DeviceManagementPartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isConfigured property value. Whether device management partner is configured or not
func (m *DeviceManagementPartner) GetIsConfigured()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isConfigured
    }
}
// Gets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
func (m *DeviceManagementPartner) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
// Gets the partnerAppType property value. Partner App type. Possible values are: unknown, singleTenantApp, multiTenantApp.
func (m *DeviceManagementPartner) GetPartnerAppType()(*DeviceManagementPartnerAppType) {
    if m == nil {
        return nil
    } else {
        return m.partnerAppType
    }
}
// Gets the partnerState property value. Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
func (m *DeviceManagementPartner) GetPartnerState()(*DeviceManagementPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
// Gets the singleTenantAppId property value. Partner Single tenant App id
func (m *DeviceManagementPartner) GetSingleTenantAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.singleTenantAppId
    }
}
// Gets the whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime property value. DateTime in UTC when PartnerDevices will be marked as NonCompliant
func (m *DeviceManagementPartner) GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime
    }
}
// Gets the whenPartnerDevicesWillBeRemovedDateTime property value. DateTime in UTC when PartnerDevices will be removed
func (m *DeviceManagementPartner) GetWhenPartnerDevicesWillBeRemovedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.whenPartnerDevicesWillBeRemovedDateTime
    }
}
// The deserialization information for the current model
func (m *DeviceManagementPartner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isConfigured"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsConfigured(val)
        return nil
    }
    res["lastHeartbeatDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastHeartbeatDateTime(val)
        return nil
    }
    res["partnerAppType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerAppType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementPartnerAppType)
        m.SetPartnerAppType(&cast)
        return nil
    }
    res["partnerState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerTenantState)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementPartnerTenantState)
        m.SetPartnerState(&cast)
        return nil
    }
    res["singleTenantAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSingleTenantAppId(val)
        return nil
    }
    res["whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime(val)
        return nil
    }
    res["whenPartnerDevicesWillBeRemovedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetWhenPartnerDevicesWillBeRemovedDateTime(val)
        return nil
    }
    return res
}
func (m *DeviceManagementPartner) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementPartner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isConfigured", m.GetIsConfigured())
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
    if m.GetPartnerAppType() != nil {
        cast := m.GetPartnerAppType().String()
        err = writer.WriteStringValue("partnerAppType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPartnerState() != nil {
        cast := m.GetPartnerState().String()
        err = writer.WriteStringValue("partnerState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("singleTenantAppId", m.GetSingleTenantAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime", m.GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("whenPartnerDevicesWillBeRemovedDateTime", m.GetWhenPartnerDevicesWillBeRemovedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. Partner display name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementPartner) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isConfigured property value. Whether device management partner is configured or not
// Parameters:
//  - value : Value to set for the isConfigured property.
func (m *DeviceManagementPartner) SetIsConfigured(value *bool)() {
    m.isConfigured = value
}
// Sets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
// Parameters:
//  - value : Value to set for the lastHeartbeatDateTime property.
func (m *DeviceManagementPartner) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
// Sets the partnerAppType property value. Partner App type. Possible values are: unknown, singleTenantApp, multiTenantApp.
// Parameters:
//  - value : Value to set for the partnerAppType property.
func (m *DeviceManagementPartner) SetPartnerAppType(value *DeviceManagementPartnerAppType)() {
    m.partnerAppType = value
}
// Sets the partnerState property value. Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
// Parameters:
//  - value : Value to set for the partnerState property.
func (m *DeviceManagementPartner) SetPartnerState(value *DeviceManagementPartnerTenantState)() {
    m.partnerState = value
}
// Sets the singleTenantAppId property value. Partner Single tenant App id
// Parameters:
//  - value : Value to set for the singleTenantAppId property.
func (m *DeviceManagementPartner) SetSingleTenantAppId(value *string)() {
    m.singleTenantAppId = value
}
// Sets the whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime property value. DateTime in UTC when PartnerDevices will be marked as NonCompliant
// Parameters:
//  - value : Value to set for the whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime property.
func (m *DeviceManagementPartner) SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime = value
}
// Sets the whenPartnerDevicesWillBeRemovedDateTime property value. DateTime in UTC when PartnerDevices will be removed
// Parameters:
//  - value : Value to set for the whenPartnerDevicesWillBeRemovedDateTime property.
func (m *DeviceManagementPartner) SetWhenPartnerDevicesWillBeRemovedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.whenPartnerDevicesWillBeRemovedDateTime = value
}
