package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceCompliancePolicyDeviceStateSummary struct {
    Entity
    // Number of compliant devices
    compliantDeviceCount *int32;
    // Number of devices that have compliance managed by System Center Configuration Manager
    configManagerCount *int32;
    // Number of conflict devices
    conflictDeviceCount *int32;
    // Number of error devices
    errorDeviceCount *int32;
    // Number of devices that are in grace period
    inGracePeriodCount *int32;
    // Number of NonCompliant devices
    nonCompliantDeviceCount *int32;
    // Number of not applicable devices
    notApplicableDeviceCount *int32;
    // Number of remediated devices
    remediatedDeviceCount *int32;
    // Number of unknown devices
    unknownDeviceCount *int32;
}
// Instantiates a new deviceCompliancePolicyDeviceStateSummary and sets the default values.
func NewDeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummary) {
    m := &DeviceCompliancePolicyDeviceStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the compliantDeviceCount property value. Number of compliant devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// Gets the configManagerCount property value. Number of devices that have compliance managed by System Center Configuration Manager
func (m *DeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configManagerCount
    }
}
// Gets the conflictDeviceCount property value. Number of conflict devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// Gets the errorDeviceCount property value. Number of error devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// Gets the inGracePeriodCount property value. Number of devices that are in grace period
func (m *DeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inGracePeriodCount
    }
}
// Gets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// Gets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// Gets the remediatedDeviceCount property value. Number of remediated devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// Gets the unknownDeviceCount property value. Number of unknown devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// The deserialization information for the current model
func (m *DeviceCompliancePolicyDeviceStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["configManagerCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigManagerCount(val)
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
    res["inGracePeriodCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInGracePeriodCount(val)
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
func (m *DeviceCompliancePolicyDeviceStateSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the compliantDeviceCount property value. Number of compliant devices
// Parameters:
//  - value : Value to set for the compliantDeviceCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// Sets the configManagerCount property value. Number of devices that have compliance managed by System Center Configuration Manager
// Parameters:
//  - value : Value to set for the configManagerCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetConfigManagerCount(value *int32)() {
    m.configManagerCount = value
}
// Sets the conflictDeviceCount property value. Number of conflict devices
// Parameters:
//  - value : Value to set for the conflictDeviceCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
// Sets the errorDeviceCount property value. Number of error devices
// Parameters:
//  - value : Value to set for the errorDeviceCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// Sets the inGracePeriodCount property value. Number of devices that are in grace period
// Parameters:
//  - value : Value to set for the inGracePeriodCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetInGracePeriodCount(value *int32)() {
    m.inGracePeriodCount = value
}
// Sets the nonCompliantDeviceCount property value. Number of NonCompliant devices
// Parameters:
//  - value : Value to set for the nonCompliantDeviceCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// Sets the notApplicableDeviceCount property value. Number of not applicable devices
// Parameters:
//  - value : Value to set for the notApplicableDeviceCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// Sets the remediatedDeviceCount property value. Number of remediated devices
// Parameters:
//  - value : Value to set for the remediatedDeviceCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// Sets the unknownDeviceCount property value. Number of unknown devices
// Parameters:
//  - value : Value to set for the unknownDeviceCount property.
func (m *DeviceCompliancePolicyDeviceStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
