package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCompliancePolicy 
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
// NewDeviceCompliancePolicy instantiates a new deviceCompliancePolicy and sets the default values.
func NewDeviceCompliancePolicy()(*DeviceCompliancePolicy) {
    m := &DeviceCompliancePolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The collection of assignments for this compliance policy.
func (m *DeviceCompliancePolicy) GetAssignments()([]DeviceCompliancePolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *DeviceCompliancePolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceCompliancePolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Compliance Setting State Device Summary
func (m *DeviceCompliancePolicy) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
// GetDeviceStatuses gets the deviceStatuses property value. List of DeviceComplianceDeviceStatus.
func (m *DeviceCompliancePolicy) GetDeviceStatuses()([]DeviceComplianceDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// GetDeviceStatusOverview gets the deviceStatusOverview property value. Device compliance devices status overview
func (m *DeviceCompliancePolicy) GetDeviceStatusOverview()(*DeviceComplianceDeviceOverview) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatusOverview
    }
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceCompliancePolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceCompliancePolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetScheduledActionsForRule gets the scheduledActionsForRule property value. The list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceCompliancePolicy) GetScheduledActionsForRule()([]DeviceComplianceScheduledActionForRule) {
    if m == nil {
        return nil
    } else {
        return m.scheduledActionsForRule
    }
}
// GetUserStatuses gets the userStatuses property value. List of DeviceComplianceUserStatus.
func (m *DeviceCompliancePolicy) GetUserStatuses()([]DeviceComplianceUserStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// GetUserStatusOverview gets the userStatusOverview property value. Device compliance users status overview
func (m *DeviceCompliancePolicy) GetUserStatusOverview()(*DeviceComplianceUserOverview) {
    if m == nil {
        return nil
    } else {
        return m.userStatusOverview
    }
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *DeviceCompliancePolicy) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicyAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceCompliancePolicyAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceSettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingStateDeviceSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingStateDeviceSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*SettingStateDeviceSummary))
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceDeviceStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceDeviceStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceDeviceStatus))
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
    res["deviceStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceDeviceOverview() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStatusOverview(val.(*DeviceComplianceDeviceOverview))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["scheduledActionsForRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScheduledActionForRule() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScheduledActionForRule, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceScheduledActionForRule))
            }
            m.SetScheduledActionsForRule(res)
        }
        return nil
    }
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceUserStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceUserStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceUserStatus))
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["userStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceUserOverview() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStatusOverview(val.(*DeviceComplianceUserOverview))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *DeviceCompliancePolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAssignments sets the assignments property value. The collection of assignments for this compliance policy.
func (m *DeviceCompliancePolicy) SetAssignments(value []DeviceCompliancePolicyAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *DeviceCompliancePolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceCompliancePolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Compliance Setting State Device Summary
func (m *DeviceCompliancePolicy) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummary)() {
    if m != nil {
        m.deviceSettingStateSummaries = value
    }
}
// SetDeviceStatuses sets the deviceStatuses property value. List of DeviceComplianceDeviceStatus.
func (m *DeviceCompliancePolicy) SetDeviceStatuses(value []DeviceComplianceDeviceStatus)() {
    if m != nil {
        m.deviceStatuses = value
    }
}
// SetDeviceStatusOverview sets the deviceStatusOverview property value. Device compliance devices status overview
func (m *DeviceCompliancePolicy) SetDeviceStatusOverview(value *DeviceComplianceDeviceOverview)() {
    if m != nil {
        m.deviceStatusOverview = value
    }
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceCompliancePolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceCompliancePolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetScheduledActionsForRule sets the scheduledActionsForRule property value. The list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceCompliancePolicy) SetScheduledActionsForRule(value []DeviceComplianceScheduledActionForRule)() {
    if m != nil {
        m.scheduledActionsForRule = value
    }
}
// SetUserStatuses sets the userStatuses property value. List of DeviceComplianceUserStatus.
func (m *DeviceCompliancePolicy) SetUserStatuses(value []DeviceComplianceUserStatus)() {
    if m != nil {
        m.userStatuses = value
    }
}
// SetUserStatusOverview sets the userStatusOverview property value. Device compliance users status overview
func (m *DeviceCompliancePolicy) SetUserStatusOverview(value *DeviceComplianceUserOverview)() {
    if m != nil {
        m.userStatusOverview = value
    }
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *DeviceCompliancePolicy) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
