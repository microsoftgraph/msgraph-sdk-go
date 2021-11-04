package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedDeviceMobileAppConfiguration struct {
    Entity
    // The list of group assignemenets for app configration.
    assignments []ManagedDeviceMobileAppConfigurationAssignment;
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Admin provided description of the Device Configuration.
    description *string;
    // List of ManagedDeviceMobileAppConfigurationDeviceStatus.
    deviceStatuses []ManagedDeviceMobileAppConfigurationDeviceStatus;
    // App configuration device status summary.
    deviceStatusSummary *ManagedDeviceMobileAppConfigurationDeviceSummary;
    // Admin provided name of the device configuration.
    displayName *string;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // the associated app.
    targetedMobileApps []string;
    // List of ManagedDeviceMobileAppConfigurationUserStatus.
    userStatuses []ManagedDeviceMobileAppConfigurationUserStatus;
    // App configuration user status summary.
    userStatusSummary *ManagedDeviceMobileAppConfigurationUserSummary;
    // Version of the device configuration.
    version *int32;
}
// Instantiates a new managedDeviceMobileAppConfiguration and sets the default values.
func NewManagedDeviceMobileAppConfiguration()(*ManagedDeviceMobileAppConfiguration) {
    m := &ManagedDeviceMobileAppConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignemenets for app configration.
func (m *ManagedDeviceMobileAppConfiguration) GetAssignments()([]ManagedDeviceMobileAppConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. DateTime the object was created.
func (m *ManagedDeviceMobileAppConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Admin provided description of the Device Configuration.
func (m *ManagedDeviceMobileAppConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceStatuses property value. List of ManagedDeviceMobileAppConfigurationDeviceStatus.
func (m *ManagedDeviceMobileAppConfiguration) GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// Gets the deviceStatusSummary property value. App configuration device status summary.
func (m *ManagedDeviceMobileAppConfiguration) GetDeviceStatusSummary()(*ManagedDeviceMobileAppConfigurationDeviceSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatusSummary
    }
}
// Gets the displayName property value. Admin provided name of the device configuration.
func (m *ManagedDeviceMobileAppConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *ManagedDeviceMobileAppConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the targetedMobileApps property value. the associated app.
func (m *ManagedDeviceMobileAppConfiguration) GetTargetedMobileApps()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetedMobileApps
    }
}
// Gets the userStatuses property value. List of ManagedDeviceMobileAppConfigurationUserStatus.
func (m *ManagedDeviceMobileAppConfiguration) GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// Gets the userStatusSummary property value. App configuration user status summary.
func (m *ManagedDeviceMobileAppConfiguration) GetUserStatusSummary()(*ManagedDeviceMobileAppConfigurationUserSummary) {
    if m == nil {
        return nil
    } else {
        return m.userStatusSummary
    }
}
// Gets the version property value. Version of the device configuration.
func (m *ManagedDeviceMobileAppConfiguration) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *ManagedDeviceMobileAppConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationAssignment() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceMobileAppConfigurationAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceMobileAppConfigurationAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationDeviceStatus() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceMobileAppConfigurationDeviceStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceMobileAppConfigurationDeviceStatus))
        }
        m.SetDeviceStatuses(res)
        return nil
    }
    res["deviceStatusSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationDeviceSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceStatusSummary(val.(*ManagedDeviceMobileAppConfigurationDeviceSummary))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["targetedMobileApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetTargetedMobileApps(res)
        return nil
    }
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationUserStatus() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceMobileAppConfigurationUserStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceMobileAppConfigurationUserStatus))
        }
        m.SetUserStatuses(res)
        return nil
    }
    res["userStatusSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationUserSummary() })
        if err != nil {
            return err
        }
        m.SetUserStatusSummary(val.(*ManagedDeviceMobileAppConfigurationUserSummary))
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *ManagedDeviceMobileAppConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedDeviceMobileAppConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStatusSummary", m.GetDeviceStatusSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("targetedMobileApps", m.GetTargetedMobileApps())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStatusSummary", m.GetUserStatusSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignemenets for app configration.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *ManagedDeviceMobileAppConfiguration) SetAssignments(value []ManagedDeviceMobileAppConfigurationAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. DateTime the object was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ManagedDeviceMobileAppConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Admin provided description of the Device Configuration.
// Parameters:
//  - value : Value to set for the description property.
func (m *ManagedDeviceMobileAppConfiguration) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceStatuses property value. List of ManagedDeviceMobileAppConfigurationDeviceStatus.
// Parameters:
//  - value : Value to set for the deviceStatuses property.
func (m *ManagedDeviceMobileAppConfiguration) SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatus)() {
    m.deviceStatuses = value
}
// Sets the deviceStatusSummary property value. App configuration device status summary.
// Parameters:
//  - value : Value to set for the deviceStatusSummary property.
func (m *ManagedDeviceMobileAppConfiguration) SetDeviceStatusSummary(value *ManagedDeviceMobileAppConfigurationDeviceSummary)() {
    m.deviceStatusSummary = value
}
// Sets the displayName property value. Admin provided name of the device configuration.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagedDeviceMobileAppConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. DateTime the object was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *ManagedDeviceMobileAppConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the targetedMobileApps property value. the associated app.
// Parameters:
//  - value : Value to set for the targetedMobileApps property.
func (m *ManagedDeviceMobileAppConfiguration) SetTargetedMobileApps(value []string)() {
    m.targetedMobileApps = value
}
// Sets the userStatuses property value. List of ManagedDeviceMobileAppConfigurationUserStatus.
// Parameters:
//  - value : Value to set for the userStatuses property.
func (m *ManagedDeviceMobileAppConfiguration) SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatus)() {
    m.userStatuses = value
}
// Sets the userStatusSummary property value. App configuration user status summary.
// Parameters:
//  - value : Value to set for the userStatusSummary property.
func (m *ManagedDeviceMobileAppConfiguration) SetUserStatusSummary(value *ManagedDeviceMobileAppConfigurationUserSummary)() {
    m.userStatusSummary = value
}
// Sets the version property value. Version of the device configuration.
// Parameters:
//  - value : Value to set for the version property.
func (m *ManagedDeviceMobileAppConfiguration) SetVersion(value *int32)() {
    m.version = value
}
