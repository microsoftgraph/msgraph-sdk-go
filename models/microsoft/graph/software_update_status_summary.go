package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SoftwareUpdateStatusSummary 
type SoftwareUpdateStatusSummary struct {
    Entity
    // Number of compliant devices.
    compliantDeviceCount *int32;
    // Number of compliant users.
    compliantUserCount *int32;
    // Number of conflict devices.
    conflictDeviceCount *int32;
    // Number of conflict users.
    conflictUserCount *int32;
    // The name of the policy.
    displayName *string;
    // Number of devices had error.
    errorDeviceCount *int32;
    // Number of users had error.
    errorUserCount *int32;
    // Number of non compliant devices.
    nonCompliantDeviceCount *int32;
    // Number of non compliant users.
    nonCompliantUserCount *int32;
    // Number of not applicable devices.
    notApplicableDeviceCount *int32;
    // Number of not applicable users.
    notApplicableUserCount *int32;
    // Number of remediated devices.
    remediatedDeviceCount *int32;
    // Number of remediated users.
    remediatedUserCount *int32;
    // Number of unknown devices.
    unknownDeviceCount *int32;
    // Number of unknown users.
    unknownUserCount *int32;
}
// NewSoftwareUpdateStatusSummary instantiates a new softwareUpdateStatusSummary and sets the default values.
func NewSoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummary) {
    m := &SoftwareUpdateStatusSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. Number of compliant devices.
func (m *SoftwareUpdateStatusSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// GetCompliantUserCount gets the compliantUserCount property value. Number of compliant users.
func (m *SoftwareUpdateStatusSummary) GetCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantUserCount
    }
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. Number of conflict devices.
func (m *SoftwareUpdateStatusSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// GetConflictUserCount gets the conflictUserCount property value. Number of conflict users.
func (m *SoftwareUpdateStatusSummary) GetConflictUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictUserCount
    }
}
// GetDisplayName gets the displayName property value. The name of the policy.
func (m *SoftwareUpdateStatusSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Number of devices had error.
func (m *SoftwareUpdateStatusSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetErrorUserCount gets the errorUserCount property value. Number of users had error.
func (m *SoftwareUpdateStatusSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. Number of non compliant devices.
func (m *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// GetNonCompliantUserCount gets the nonCompliantUserCount property value. Number of non compliant users.
func (m *SoftwareUpdateStatusSummary) GetNonCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantUserCount
    }
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of not applicable devices.
func (m *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// GetNotApplicableUserCount gets the notApplicableUserCount property value. Number of not applicable users.
func (m *SoftwareUpdateStatusSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. Number of remediated devices.
func (m *SoftwareUpdateStatusSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// GetRemediatedUserCount gets the remediatedUserCount property value. Number of remediated users.
func (m *SoftwareUpdateStatusSummary) GetRemediatedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedUserCount
    }
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Number of unknown devices.
func (m *SoftwareUpdateStatusSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// GetUnknownUserCount gets the unknownUserCount property value. Number of unknown users.
func (m *SoftwareUpdateStatusSummary) GetUnknownUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownUserCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SoftwareUpdateStatusSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["compliantUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantUserCount(val)
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
    res["conflictUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictUserCount(val)
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
    res["errorUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorUserCount(val)
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
    res["nonCompliantUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantUserCount(val)
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
    res["notApplicableUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableUserCount(val)
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
    res["remediatedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedUserCount(val)
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
    res["unknownUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownUserCount(val)
        }
        return nil
    }
    return res
}
func (m *SoftwareUpdateStatusSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SoftwareUpdateStatusSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("compliantUserCount", m.GetCompliantUserCount())
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
        err = writer.WriteInt32Value("conflictUserCount", m.GetConflictUserCount())
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
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorUserCount", m.GetErrorUserCount())
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
        err = writer.WriteInt32Value("nonCompliantUserCount", m.GetNonCompliantUserCount())
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
        err = writer.WriteInt32Value("notApplicableUserCount", m.GetNotApplicableUserCount())
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
        err = writer.WriteInt32Value("remediatedUserCount", m.GetRemediatedUserCount())
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
    {
        err = writer.WriteInt32Value("unknownUserCount", m.GetUnknownUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. Number of compliant devices.
func (m *SoftwareUpdateStatusSummary) SetCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.compliantDeviceCount = value
    }
}
// SetCompliantUserCount sets the compliantUserCount property value. Number of compliant users.
func (m *SoftwareUpdateStatusSummary) SetCompliantUserCount(value *int32)() {
    if m != nil {
        m.compliantUserCount = value
    }
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. Number of conflict devices.
func (m *SoftwareUpdateStatusSummary) SetConflictDeviceCount(value *int32)() {
    if m != nil {
        m.conflictDeviceCount = value
    }
}
// SetConflictUserCount sets the conflictUserCount property value. Number of conflict users.
func (m *SoftwareUpdateStatusSummary) SetConflictUserCount(value *int32)() {
    if m != nil {
        m.conflictUserCount = value
    }
}
// SetDisplayName sets the displayName property value. The name of the policy.
func (m *SoftwareUpdateStatusSummary) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Number of devices had error.
func (m *SoftwareUpdateStatusSummary) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetErrorUserCount sets the errorUserCount property value. Number of users had error.
func (m *SoftwareUpdateStatusSummary) SetErrorUserCount(value *int32)() {
    if m != nil {
        m.errorUserCount = value
    }
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. Number of non compliant devices.
func (m *SoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.nonCompliantDeviceCount = value
    }
}
// SetNonCompliantUserCount sets the nonCompliantUserCount property value. Number of non compliant users.
func (m *SoftwareUpdateStatusSummary) SetNonCompliantUserCount(value *int32)() {
    if m != nil {
        m.nonCompliantUserCount = value
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of not applicable devices.
func (m *SoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(value *int32)() {
    if m != nil {
        m.notApplicableDeviceCount = value
    }
}
// SetNotApplicableUserCount sets the notApplicableUserCount property value. Number of not applicable users.
func (m *SoftwareUpdateStatusSummary) SetNotApplicableUserCount(value *int32)() {
    if m != nil {
        m.notApplicableUserCount = value
    }
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. Number of remediated devices.
func (m *SoftwareUpdateStatusSummary) SetRemediatedDeviceCount(value *int32)() {
    if m != nil {
        m.remediatedDeviceCount = value
    }
}
// SetRemediatedUserCount sets the remediatedUserCount property value. Number of remediated users.
func (m *SoftwareUpdateStatusSummary) SetRemediatedUserCount(value *int32)() {
    if m != nil {
        m.remediatedUserCount = value
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Number of unknown devices.
func (m *SoftwareUpdateStatusSummary) SetUnknownDeviceCount(value *int32)() {
    if m != nil {
        m.unknownDeviceCount = value
    }
}
// SetUnknownUserCount sets the unknownUserCount property value. Number of unknown users.
func (m *SoftwareUpdateStatusSummary) SetUnknownUserCount(value *int32)() {
    if m != nil {
        m.unknownUserCount = value
    }
}
