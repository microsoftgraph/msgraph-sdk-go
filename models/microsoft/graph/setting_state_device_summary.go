package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SettingStateDeviceSummary 
type SettingStateDeviceSummary struct {
    Entity
    // Device Compliant count for the setting
    compliantDeviceCount *int32;
    // Device conflict error count for the setting
    conflictDeviceCount *int32;
    // Device error count for the setting
    errorDeviceCount *int32;
    // Name of the InstancePath for the setting
    instancePath *string;
    // Device NonCompliant count for the setting
    nonCompliantDeviceCount *int32;
    // Device Not Applicable count for the setting
    notApplicableDeviceCount *int32;
    // Device Compliant count for the setting
    remediatedDeviceCount *int32;
    // Name of the setting
    settingName *string;
    // Device Unkown count for the setting
    unknownDeviceCount *int32;
}
// NewSettingStateDeviceSummary instantiates a new settingStateDeviceSummary and sets the default values.
func NewSettingStateDeviceSummary()(*SettingStateDeviceSummary) {
    m := &SettingStateDeviceSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. Device Compliant count for the setting
func (m *SettingStateDeviceSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. Device conflict error count for the setting
func (m *SettingStateDeviceSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Device error count for the setting
func (m *SettingStateDeviceSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetInstancePath gets the instancePath property value. Name of the InstancePath for the setting
func (m *SettingStateDeviceSummary) GetInstancePath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.instancePath
    }
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. Device NonCompliant count for the setting
func (m *SettingStateDeviceSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Device Not Applicable count for the setting
func (m *SettingStateDeviceSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. Device Compliant count for the setting
func (m *SettingStateDeviceSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// GetSettingName gets the settingName property value. Name of the setting
func (m *SettingStateDeviceSummary) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Device Unkown count for the setting
func (m *SettingStateDeviceSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SettingStateDeviceSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantDeviceCount(val)
        }
        return nil
    }
    res["conflictDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictDeviceCount(val)
        }
        return nil
    }
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["instancePath"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstancePath(val)
        }
        return nil
    }
    res["nonCompliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantDeviceCount(val)
        }
        return nil
    }
    res["notApplicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["remediatedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedDeviceCount(val)
        }
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["unknownDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *SettingStateDeviceSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SettingStateDeviceSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("compliantDeviceCount", m.GetCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("conflictDeviceCount", m.GetConflictDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("instancePath", m.GetInstancePath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("nonCompliantDeviceCount", m.GetNonCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableDeviceCount", m.GetNotApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. Device Compliant count for the setting
func (m *SettingStateDeviceSummary) SetCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.compliantDeviceCount = value
    }
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. Device conflict error count for the setting
func (m *SettingStateDeviceSummary) SetConflictDeviceCount(value *int32)() {
    if m != nil {
        m.conflictDeviceCount = value
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Device error count for the setting
func (m *SettingStateDeviceSummary) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetInstancePath sets the instancePath property value. Name of the InstancePath for the setting
func (m *SettingStateDeviceSummary) SetInstancePath(value *string)() {
    if m != nil {
        m.instancePath = value
    }
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. Device NonCompliant count for the setting
func (m *SettingStateDeviceSummary) SetNonCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.nonCompliantDeviceCount = value
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Device Not Applicable count for the setting
func (m *SettingStateDeviceSummary) SetNotApplicableDeviceCount(value *int32)() {
    if m != nil {
        m.notApplicableDeviceCount = value
    }
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. Device Compliant count for the setting
func (m *SettingStateDeviceSummary) SetRemediatedDeviceCount(value *int32)() {
    if m != nil {
        m.remediatedDeviceCount = value
    }
}
// SetSettingName sets the settingName property value. Name of the setting
func (m *SettingStateDeviceSummary) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Device Unkown count for the setting
func (m *SettingStateDeviceSummary) SetUnknownDeviceCount(value *int32)() {
    if m != nil {
        m.unknownDeviceCount = value
    }
}
