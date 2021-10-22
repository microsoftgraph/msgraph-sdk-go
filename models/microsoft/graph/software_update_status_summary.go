package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SoftwareUpdateStatusSummary struct {
    Entity
    compliantDeviceCount *int32;
    compliantUserCount *int32;
    conflictDeviceCount *int32;
    conflictUserCount *int32;
    displayName *string;
    errorDeviceCount *int32;
    errorUserCount *int32;
    nonCompliantDeviceCount *int32;
    nonCompliantUserCount *int32;
    notApplicableDeviceCount *int32;
    notApplicableUserCount *int32;
    remediatedDeviceCount *int32;
    remediatedUserCount *int32;
    unknownDeviceCount *int32;
    unknownUserCount *int32;
}
func NewSoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummary) {
    m := &SoftwareUpdateStatusSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SoftwareUpdateStatusSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantUserCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetConflictUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictUserCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *SoftwareUpdateStatusSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetNonCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantUserCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetRemediatedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedUserCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetUnknownUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownUserCount
    }
}
func (m *SoftwareUpdateStatusSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantDeviceCount(val)
        return nil
    }
    res["compliantUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantUserCount(val)
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
    res["conflictUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConflictUserCount(val)
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
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorDeviceCount(val)
        return nil
    }
    res["errorUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorUserCount(val)
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
    res["nonCompliantUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNonCompliantUserCount(val)
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
    res["notApplicableUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotApplicableUserCount(val)
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
    res["remediatedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediatedUserCount(val)
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
    res["unknownUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownUserCount(val)
        return nil
    }
    return res
}
func (m *SoftwareUpdateStatusSummary) IsNil()(bool) {
    return m == nil
}
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
func (m *SoftwareUpdateStatusSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
func (m *SoftwareUpdateStatusSummary) SetCompliantUserCount(value *int32)() {
    m.compliantUserCount = value
}
func (m *SoftwareUpdateStatusSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
func (m *SoftwareUpdateStatusSummary) SetConflictUserCount(value *int32)() {
    m.conflictUserCount = value
}
func (m *SoftwareUpdateStatusSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *SoftwareUpdateStatusSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
func (m *SoftwareUpdateStatusSummary) SetErrorUserCount(value *int32)() {
    m.errorUserCount = value
}
func (m *SoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
func (m *SoftwareUpdateStatusSummary) SetNonCompliantUserCount(value *int32)() {
    m.nonCompliantUserCount = value
}
func (m *SoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
func (m *SoftwareUpdateStatusSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
func (m *SoftwareUpdateStatusSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
func (m *SoftwareUpdateStatusSummary) SetRemediatedUserCount(value *int32)() {
    m.remediatedUserCount = value
}
func (m *SoftwareUpdateStatusSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
func (m *SoftwareUpdateStatusSummary) SetUnknownUserCount(value *int32)() {
    m.unknownUserCount = value
}
