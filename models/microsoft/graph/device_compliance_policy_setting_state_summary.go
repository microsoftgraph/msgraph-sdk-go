package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCompliancePolicySettingStateSummary 
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
    // Setting platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
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
// NewDeviceCompliancePolicySettingStateSummary instantiates a new deviceCompliancePolicySettingStateSummary and sets the default values.
func NewDeviceCompliancePolicySettingStateSummary()(*DeviceCompliancePolicySettingStateSummary) {
    m := &DeviceCompliancePolicySettingStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. Number of compliant devices
func (m *DeviceCompliancePolicySettingStateSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. Number of conflict devices
func (m *DeviceCompliancePolicySettingStateSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// GetDeviceComplianceSettingStates gets the deviceComplianceSettingStates property value. Not yet documented
func (m *DeviceCompliancePolicySettingStateSummary) GetDeviceComplianceSettingStates()([]DeviceComplianceSettingState) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceSettingStates
    }
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Number of error devices
func (m *DeviceCompliancePolicySettingStateSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *DeviceCompliancePolicySettingStateSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *DeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// GetPlatformType gets the platformType property value. Setting platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
func (m *DeviceCompliancePolicySettingStateSummary) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. Number of remediated devices
func (m *DeviceCompliancePolicySettingStateSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// GetSetting gets the setting property value. The setting class name and property name.
func (m *DeviceCompliancePolicySettingStateSummary) GetSetting()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setting
    }
}
// GetSettingName gets the settingName property value. Name of the setting.
func (m *DeviceCompliancePolicySettingStateSummary) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Number of unknown devices
func (m *DeviceCompliancePolicySettingStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicySettingStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["deviceComplianceSettingStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceSettingState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceSettingState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceSettingState))
            }
            m.SetDeviceComplianceSettingStates(res)
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
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PolicyPlatformType)
            m.SetPlatformType(&cast)
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
    res["setting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val)
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
func (m *DeviceCompliancePolicySettingStateSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCompliantDeviceCount sets the compliantDeviceCount property value. Number of compliant devices
func (m *DeviceCompliancePolicySettingStateSummary) SetCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.compliantDeviceCount = value
    }
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. Number of conflict devices
func (m *DeviceCompliancePolicySettingStateSummary) SetConflictDeviceCount(value *int32)() {
    if m != nil {
        m.conflictDeviceCount = value
    }
}
// SetDeviceComplianceSettingStates sets the deviceComplianceSettingStates property value. Not yet documented
func (m *DeviceCompliancePolicySettingStateSummary) SetDeviceComplianceSettingStates(value []DeviceComplianceSettingState)() {
    if m != nil {
        m.deviceComplianceSettingStates = value
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Number of error devices
func (m *DeviceCompliancePolicySettingStateSummary) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *DeviceCompliancePolicySettingStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.nonCompliantDeviceCount = value
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *DeviceCompliancePolicySettingStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    if m != nil {
        m.notApplicableDeviceCount = value
    }
}
// SetPlatformType sets the platformType property value. Setting platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
func (m *DeviceCompliancePolicySettingStateSummary) SetPlatformType(value *PolicyPlatformType)() {
    if m != nil {
        m.platformType = value
    }
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. Number of remediated devices
func (m *DeviceCompliancePolicySettingStateSummary) SetRemediatedDeviceCount(value *int32)() {
    if m != nil {
        m.remediatedDeviceCount = value
    }
}
// SetSetting sets the setting property value. The setting class name and property name.
func (m *DeviceCompliancePolicySettingStateSummary) SetSetting(value *string)() {
    if m != nil {
        m.setting = value
    }
}
// SetSettingName sets the settingName property value. Name of the setting.
func (m *DeviceCompliancePolicySettingStateSummary) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Number of unknown devices
func (m *DeviceCompliancePolicySettingStateSummary) SetUnknownDeviceCount(value *int32)() {
    if m != nil {
        m.unknownDeviceCount = value
    }
}
