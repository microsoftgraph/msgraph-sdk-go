package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new softwareUpdateStatusSummary and sets the default values.
func NewSoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummary) {
    m := &SoftwareUpdateStatusSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the compliantDeviceCount property value. Number of compliant devices.
func (m *SoftwareUpdateStatusSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// Gets the compliantUserCount property value. Number of compliant users.
func (m *SoftwareUpdateStatusSummary) GetCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantUserCount
    }
}
// Gets the conflictDeviceCount property value. Number of conflict devices.
func (m *SoftwareUpdateStatusSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// Gets the conflictUserCount property value. Number of conflict users.
func (m *SoftwareUpdateStatusSummary) GetConflictUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictUserCount
    }
}
// Gets the displayName property value. The name of the policy.
func (m *SoftwareUpdateStatusSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the errorDeviceCount property value. Number of devices had error.
func (m *SoftwareUpdateStatusSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// Gets the errorUserCount property value. Number of users had error.
func (m *SoftwareUpdateStatusSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
// Gets the nonCompliantDeviceCount property value. Number of non compliant devices.
func (m *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// Gets the nonCompliantUserCount property value. Number of non compliant users.
func (m *SoftwareUpdateStatusSummary) GetNonCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantUserCount
    }
}
// Gets the notApplicableDeviceCount property value. Number of not applicable devices.
func (m *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// Gets the notApplicableUserCount property value. Number of not applicable users.
func (m *SoftwareUpdateStatusSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
// Gets the remediatedDeviceCount property value. Number of remediated devices.
func (m *SoftwareUpdateStatusSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// Gets the remediatedUserCount property value. Number of remediated users.
func (m *SoftwareUpdateStatusSummary) GetRemediatedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedUserCount
    }
}
// Gets the unknownDeviceCount property value. Number of unknown devices.
func (m *SoftwareUpdateStatusSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// Gets the unknownUserCount property value. Number of unknown users.
func (m *SoftwareUpdateStatusSummary) GetUnknownUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownUserCount
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the compliantDeviceCount property value. Number of compliant devices.
// Parameters:
//  - value : Value to set for the compliantDeviceCount property.
func (m *SoftwareUpdateStatusSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// Sets the compliantUserCount property value. Number of compliant users.
// Parameters:
//  - value : Value to set for the compliantUserCount property.
func (m *SoftwareUpdateStatusSummary) SetCompliantUserCount(value *int32)() {
    m.compliantUserCount = value
}
// Sets the conflictDeviceCount property value. Number of conflict devices.
// Parameters:
//  - value : Value to set for the conflictDeviceCount property.
func (m *SoftwareUpdateStatusSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
// Sets the conflictUserCount property value. Number of conflict users.
// Parameters:
//  - value : Value to set for the conflictUserCount property.
func (m *SoftwareUpdateStatusSummary) SetConflictUserCount(value *int32)() {
    m.conflictUserCount = value
}
// Sets the displayName property value. The name of the policy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SoftwareUpdateStatusSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the errorDeviceCount property value. Number of devices had error.
// Parameters:
//  - value : Value to set for the errorDeviceCount property.
func (m *SoftwareUpdateStatusSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// Sets the errorUserCount property value. Number of users had error.
// Parameters:
//  - value : Value to set for the errorUserCount property.
func (m *SoftwareUpdateStatusSummary) SetErrorUserCount(value *int32)() {
    m.errorUserCount = value
}
// Sets the nonCompliantDeviceCount property value. Number of non compliant devices.
// Parameters:
//  - value : Value to set for the nonCompliantDeviceCount property.
func (m *SoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// Sets the nonCompliantUserCount property value. Number of non compliant users.
// Parameters:
//  - value : Value to set for the nonCompliantUserCount property.
func (m *SoftwareUpdateStatusSummary) SetNonCompliantUserCount(value *int32)() {
    m.nonCompliantUserCount = value
}
// Sets the notApplicableDeviceCount property value. Number of not applicable devices.
// Parameters:
//  - value : Value to set for the notApplicableDeviceCount property.
func (m *SoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// Sets the notApplicableUserCount property value. Number of not applicable users.
// Parameters:
//  - value : Value to set for the notApplicableUserCount property.
func (m *SoftwareUpdateStatusSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
// Sets the remediatedDeviceCount property value. Number of remediated devices.
// Parameters:
//  - value : Value to set for the remediatedDeviceCount property.
func (m *SoftwareUpdateStatusSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// Sets the remediatedUserCount property value. Number of remediated users.
// Parameters:
//  - value : Value to set for the remediatedUserCount property.
func (m *SoftwareUpdateStatusSummary) SetRemediatedUserCount(value *int32)() {
    m.remediatedUserCount = value
}
// Sets the unknownDeviceCount property value. Number of unknown devices.
// Parameters:
//  - value : Value to set for the unknownDeviceCount property.
func (m *SoftwareUpdateStatusSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
// Sets the unknownUserCount property value. Number of unknown users.
// Parameters:
//  - value : Value to set for the unknownUserCount property.
func (m *SoftwareUpdateStatusSummary) SetUnknownUserCount(value *int32)() {
    m.unknownUserCount = value
}
