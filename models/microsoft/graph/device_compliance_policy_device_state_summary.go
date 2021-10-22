package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceCompliancePolicyDeviceStateSummary struct {
    Entity
    compliantDeviceCount *int32;
    configManagerCount *int32;
    conflictDeviceCount *int32;
    errorDeviceCount *int32;
    inGracePeriodCount *int32;
    nonCompliantDeviceCount *int32;
    notApplicableDeviceCount *int32;
    remediatedDeviceCount *int32;
    unknownDeviceCount *int32;
}
func NewDeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummary) {
    m := &DeviceCompliancePolicyDeviceStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configManagerCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inGracePeriodCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantDeviceCount(val)
        return nil
    }
    res["configManagerCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfigManagerCount(val)
        return nil
    }
    res["conflictDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConflictDeviceCount(val)
        return nil
    }
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorDeviceCount(val)
        return nil
    }
    res["inGracePeriodCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInGracePeriodCount(val)
        return nil
    }
    res["nonCompliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNonCompliantDeviceCount(val)
        return nil
    }
    res["notApplicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotApplicableDeviceCount(val)
        return nil
    }
    res["remediatedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediatedDeviceCount(val)
        return nil
    }
    res["unknownDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownDeviceCount(val)
        return nil
    }
    return res
}
func (m *DeviceCompliancePolicyDeviceStateSummary) IsNil()(bool) {
    return m == nil
}
func (m *DeviceCompliancePolicyDeviceStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("configManagerCount", m.GetConfigManagerCount())
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
        err = writer.WriteInt32Value("inGracePeriodCount", m.GetInGracePeriodCount())
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
        err = writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetConfigManagerCount(value *int32)() {
    m.configManagerCount = value
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetInGracePeriodCount(value *int32)() {
    m.inGracePeriodCount = value
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
func (m *DeviceCompliancePolicyDeviceStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
