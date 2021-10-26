package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceCompliancePolicy struct {
    Entity
    // The collection of assignments for this compliance policy.
    assignments []DeviceCompliancePolicyAssignment;
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Admin provided description of the Device Configuration.
    description *string;
    // Compliance Setting State Device Summary
    deviceSettingStateSummaries []SettingStateDeviceSummary;
    // List of DeviceComplianceDeviceStatus.
    deviceStatuses []DeviceComplianceDeviceStatus;
    // Device compliance devices status overview
    deviceStatusOverview *DeviceComplianceDeviceOverview;
    // Admin provided name of the device configuration.
    displayName *string;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
    scheduledActionsForRule []DeviceComplianceScheduledActionForRule;
    // List of DeviceComplianceUserStatus.
    userStatuses []DeviceComplianceUserStatus;
    // Device compliance users status overview
    userStatusOverview *DeviceComplianceUserOverview;
    // Version of the device configuration.
    version *int32;
}
// Instantiates a new deviceCompliancePolicy and sets the default values.
func NewDeviceCompliancePolicy()(*DeviceCompliancePolicy) {
    m := &DeviceCompliancePolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The collection of assignments for this compliance policy.
func (m *DeviceCompliancePolicy) GetAssignments()([]DeviceCompliancePolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. DateTime the object was created.
func (m *DeviceCompliancePolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceCompliancePolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceSettingStateSummaries property value. Compliance Setting State Device Summary
func (m *DeviceCompliancePolicy) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
// Gets the deviceStatuses property value. List of DeviceComplianceDeviceStatus.
func (m *DeviceCompliancePolicy) GetDeviceStatuses()([]DeviceComplianceDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// Gets the deviceStatusOverview property value. Device compliance devices status overview
func (m *DeviceCompliancePolicy) GetDeviceStatusOverview()(*DeviceComplianceDeviceOverview) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatusOverview
    }
}
// Gets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceCompliancePolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceCompliancePolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the scheduledActionsForRule property value. The list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceCompliancePolicy) GetScheduledActionsForRule()([]DeviceComplianceScheduledActionForRule) {
    if m == nil {
        return nil
    } else {
        return m.scheduledActionsForRule
    }
}
// Gets the userStatuses property value. List of DeviceComplianceUserStatus.
func (m *DeviceCompliancePolicy) GetUserStatuses()([]DeviceComplianceUserStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// Gets the userStatusOverview property value. Device compliance users status overview
func (m *DeviceCompliancePolicy) GetUserStatusOverview()(*DeviceComplianceUserOverview) {
    if m == nil {
        return nil
    } else {
        return m.userStatusOverview
    }
}
// Gets the version property value. Version of the device configuration.
func (m *DeviceCompliancePolicy) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceCompliancePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicyAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceCompliancePolicyAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCompliancePolicyAssignment))
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
    res["deviceSettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingStateDeviceSummary() })
        if err != nil {
            return err
        }
        res := make([]SettingStateDeviceSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*SettingStateDeviceSummary))
        }
        m.SetDeviceSettingStateSummaries(res)
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceDeviceStatus() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceDeviceStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceDeviceStatus))
        }
        m.SetDeviceStatuses(res)
        return nil
    }
    res["deviceStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceDeviceOverview() })
        if err != nil {
            return err
        }
        m.SetDeviceStatusOverview(val.(*DeviceComplianceDeviceOverview))
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
    res["scheduledActionsForRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScheduledActionForRule() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceScheduledActionForRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceScheduledActionForRule))
        }
        m.SetScheduledActionsForRule(res)
        return nil
    }
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceUserStatus() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceUserStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceUserStatus))
        }
        m.SetUserStatuses(res)
        return nil
    }
    res["userStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceUserOverview() })
        if err != nil {
            return err
        }
        m.SetUserStatusOverview(val.(*DeviceComplianceUserOverview))
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
func (m *DeviceCompliancePolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceCompliancePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
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
        err = writer.WriteObjectValue("deviceStatusOverview", m.GetDeviceStatusOverview())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScheduledActionsForRule()))
        for i, v := range m.GetScheduledActionsForRule() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("scheduledActionsForRule", cast)
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
        err = writer.WriteObjectValue("userStatusOverview", m.GetUserStatusOverview())
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
// Sets the assignments property value. The collection of assignments for this compliance policy.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *DeviceCompliancePolicy) SetAssignments(value []DeviceCompliancePolicyAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. DateTime the object was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceCompliancePolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Admin provided description of the Device Configuration.
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceCompliancePolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceSettingStateSummaries property value. Compliance Setting State Device Summary
// Parameters:
//  - value : Value to set for the deviceSettingStateSummaries property.
func (m *DeviceCompliancePolicy) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummary)() {
    m.deviceSettingStateSummaries = value
}
// Sets the deviceStatuses property value. List of DeviceComplianceDeviceStatus.
// Parameters:
//  - value : Value to set for the deviceStatuses property.
func (m *DeviceCompliancePolicy) SetDeviceStatuses(value []DeviceComplianceDeviceStatus)() {
    m.deviceStatuses = value
}
// Sets the deviceStatusOverview property value. Device compliance devices status overview
// Parameters:
//  - value : Value to set for the deviceStatusOverview property.
func (m *DeviceCompliancePolicy) SetDeviceStatusOverview(value *DeviceComplianceDeviceOverview)() {
    m.deviceStatusOverview = value
}
// Sets the displayName property value. Admin provided name of the device configuration.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceCompliancePolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. DateTime the object was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceCompliancePolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the scheduledActionsForRule property value. The list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
// Parameters:
//  - value : Value to set for the scheduledActionsForRule property.
func (m *DeviceCompliancePolicy) SetScheduledActionsForRule(value []DeviceComplianceScheduledActionForRule)() {
    m.scheduledActionsForRule = value
}
// Sets the userStatuses property value. List of DeviceComplianceUserStatus.
// Parameters:
//  - value : Value to set for the userStatuses property.
func (m *DeviceCompliancePolicy) SetUserStatuses(value []DeviceComplianceUserStatus)() {
    m.userStatuses = value
}
// Sets the userStatusOverview property value. Device compliance users status overview
// Parameters:
//  - value : Value to set for the userStatusOverview property.
func (m *DeviceCompliancePolicy) SetUserStatusOverview(value *DeviceComplianceUserOverview)() {
    m.userStatusOverview = value
}
// Sets the version property value. Version of the device configuration.
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceCompliancePolicy) SetVersion(value *int32)() {
    m.version = value
}
