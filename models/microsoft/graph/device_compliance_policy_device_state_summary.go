package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCompliancePolicyDeviceStateSummary provides operations to manage the deviceManagement singleton.
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
// NewDeviceCompliancePolicyDeviceStateSummary instantiates a new deviceCompliancePolicyDeviceStateSummary and sets the default values.
func NewDeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummary) {
    m := &DeviceCompliancePolicyDeviceStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceCompliancePolicyDeviceStateSummary(), nil
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. Number of compliant devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// GetConfigManagerCount gets the configManagerCount property value. Number of devices that have compliance managed by System Center Configuration Manager
func (m *DeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configManagerCount
    }
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. Number of conflict devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Number of error devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetInGracePeriodCount gets the inGracePeriodCount property value. Number of devices that are in grace period
func (m *DeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inGracePeriodCount
    }
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. Number of remediated devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Number of unknown devices
func (m *DeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
func (m *DeviceCompliancePolicyDeviceStateSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCompliantDeviceCount sets the compliantDeviceCount property value. Number of compliant devices
func (m *DeviceCompliancePolicyDeviceStateSummary) SetCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.compliantDeviceCount = value
    }
}
// SetConfigManagerCount sets the configManagerCount property value. Number of devices that have compliance managed by System Center Configuration Manager
func (m *DeviceCompliancePolicyDeviceStateSummary) SetConfigManagerCount(value *int32)() {
    if m != nil {
        m.configManagerCount = value
    }
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. Number of conflict devices
func (m *DeviceCompliancePolicyDeviceStateSummary) SetConflictDeviceCount(value *int32)() {
    if m != nil {
        m.conflictDeviceCount = value
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Number of error devices
func (m *DeviceCompliancePolicyDeviceStateSummary) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetInGracePeriodCount sets the inGracePeriodCount property value. Number of devices that are in grace period
func (m *DeviceCompliancePolicyDeviceStateSummary) SetInGracePeriodCount(value *int32)() {
    if m != nil {
        m.inGracePeriodCount = value
    }
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *DeviceCompliancePolicyDeviceStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.nonCompliantDeviceCount = value
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *DeviceCompliancePolicyDeviceStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    if m != nil {
        m.notApplicableDeviceCount = value
    }
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. Number of remediated devices
func (m *DeviceCompliancePolicyDeviceStateSummary) SetRemediatedDeviceCount(value *int32)() {
    if m != nil {
        m.remediatedDeviceCount = value
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Number of unknown devices
func (m *DeviceCompliancePolicyDeviceStateSummary) SetUnknownDeviceCount(value *int32)() {
    if m != nil {
        m.unknownDeviceCount = value
    }
}
