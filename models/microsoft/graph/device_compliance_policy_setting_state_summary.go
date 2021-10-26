package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceCompliancePolicySettingStateSummary struct {
    Entity
    // Number of compliant devices
    compliantDeviceCount *int32;
    // Number of conflict devices
    conflictDeviceCount *int32;
    // Not yet documented
    deviceComplianceSettingStates []DeviceComplianceSettingState;
    // Number of error devices
    errorDeviceCount *int32;
    // Number of NonCompliant devices
    nonCompliantDeviceCount *int32;
    // Number of not applicable devices
    notApplicableDeviceCount *int32;
    // Setting platform. Possible values are: android, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, all.
    platformType *PolicyPlatformType;
    // Number of remediated devices
    remediatedDeviceCount *int32;
    // The setting class name and property name.
    setting *string;
    // Name of the setting.
    settingName *string;
    // Number of unknown devices
    unknownDeviceCount *int32;
}
// Instantiates a new deviceCompliancePolicySettingStateSummary and sets the default values.
func NewDeviceCompliancePolicySettingStateSummary()(*DeviceCompliancePolicySettingStateSummary) {
    m := &DeviceCompliancePolicySettingStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the compliantDeviceCount property value. Number of compliant devices
func (m *DeviceCompliancePolicySettingStateSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// Gets the conflictDeviceCount property value. Number of conflict devices
func (m *DeviceCompliancePolicySettingStateSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// Gets the deviceComplianceSettingStates property value. Not yet documented
func (m *DeviceCompliancePolicySettingStateSummary) GetDeviceComplianceSettingStates()([]DeviceComplianceSettingState) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceSettingStates
    }
}
// Gets the errorDeviceCount property value. Number of error devices
func (m *DeviceCompliancePolicySettingStateSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// Gets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *DeviceCompliancePolicySettingStateSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// Gets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *DeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// Gets the platformType property value. Setting platform. Possible values are: android, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, all.
func (m *DeviceCompliancePolicySettingStateSummary) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// Gets the remediatedDeviceCount property value. Number of remediated devices
func (m *DeviceCompliancePolicySettingStateSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// Gets the setting property value. The setting class name and property name.
func (m *DeviceCompliancePolicySettingStateSummary) GetSetting()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setting
    }
}
// Gets the settingName property value. Name of the setting.
func (m *DeviceCompliancePolicySettingStateSummary) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// Gets the unknownDeviceCount property value. Number of unknown devices
func (m *DeviceCompliancePolicySettingStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// The deserialization information for the current model
func (m *DeviceCompliancePolicySettingStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantDeviceCount(val)
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
    res["deviceComplianceSettingStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceSettingState() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceSettingState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceSettingState))
        }
        m.SetDeviceComplianceSettingStates(res)
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
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        cast := val.(PolicyPlatformType)
        m.SetPlatformType(&cast)
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
    res["setting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSetting(val)
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingName(val)
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
func (m *DeviceCompliancePolicySettingStateSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceCompliancePolicySettingStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceComplianceSettingStates()))
        for i, v := range m.GetDeviceComplianceSettingStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceComplianceSettingStates", cast)
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
    if m.GetPlatformType() != nil {
        cast := m.GetPlatformType().String()
        err = writer.WriteStringValue("platformType", &cast)
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
        err = writer.WriteStringValue("setting", m.GetSetting())
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
// Sets the compliantDeviceCount property value. Number of compliant devices
// Parameters:
//  - value : Value to set for the compliantDeviceCount property.
func (m *DeviceCompliancePolicySettingStateSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// Sets the conflictDeviceCount property value. Number of conflict devices
// Parameters:
//  - value : Value to set for the conflictDeviceCount property.
func (m *DeviceCompliancePolicySettingStateSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
// Sets the deviceComplianceSettingStates property value. Not yet documented
// Parameters:
//  - value : Value to set for the deviceComplianceSettingStates property.
func (m *DeviceCompliancePolicySettingStateSummary) SetDeviceComplianceSettingStates(value []DeviceComplianceSettingState)() {
    m.deviceComplianceSettingStates = value
}
// Sets the errorDeviceCount property value. Number of error devices
// Parameters:
//  - value : Value to set for the errorDeviceCount property.
func (m *DeviceCompliancePolicySettingStateSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// Sets the nonCompliantDeviceCount property value. Number of NonCompliant devices
// Parameters:
//  - value : Value to set for the nonCompliantDeviceCount property.
func (m *DeviceCompliancePolicySettingStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// Sets the notApplicableDeviceCount property value. Number of not applicable devices
// Parameters:
//  - value : Value to set for the notApplicableDeviceCount property.
func (m *DeviceCompliancePolicySettingStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// Sets the platformType property value. Setting platform. Possible values are: android, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, all.
// Parameters:
//  - value : Value to set for the platformType property.
func (m *DeviceCompliancePolicySettingStateSummary) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// Sets the remediatedDeviceCount property value. Number of remediated devices
// Parameters:
//  - value : Value to set for the remediatedDeviceCount property.
func (m *DeviceCompliancePolicySettingStateSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// Sets the setting property value. The setting class name and property name.
// Parameters:
//  - value : Value to set for the setting property.
func (m *DeviceCompliancePolicySettingStateSummary) SetSetting(value *string)() {
    m.setting = value
}
// Sets the settingName property value. Name of the setting.
// Parameters:
//  - value : Value to set for the settingName property.
func (m *DeviceCompliancePolicySettingStateSummary) SetSettingName(value *string)() {
    m.settingName = value
}
// Sets the unknownDeviceCount property value. Number of unknown devices
// Parameters:
//  - value : Value to set for the unknownDeviceCount property.
func (m *DeviceCompliancePolicySettingStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
