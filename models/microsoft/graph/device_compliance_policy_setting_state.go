package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCompliancePolicySettingState 
type DeviceCompliancePolicySettingState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Current value of setting on device
    currentValue *string;
    // Error code for the setting
    errorCode *int64;
    // Error description
    errorDescription *string;
    // Name of setting instance that is being reported.
    instanceDisplayName *string;
    // The setting that is being reported
    setting *string;
    // Localized/user friendly setting name that is being reported
    settingName *string;
    // Contributing policies
    sources []SettingSource;
    // The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
    state *ComplianceStatus;
    // UserEmail
    userEmail *string;
    // UserId
    userId *string;
    // UserName
    userName *string;
    // UserPrincipalName.
    userPrincipalName *string;
}
// NewDeviceCompliancePolicySettingState instantiates a new deviceCompliancePolicySettingState and sets the default values.
func NewDeviceCompliancePolicySettingState()(*DeviceCompliancePolicySettingState) {
    m := &DeviceCompliancePolicySettingState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePolicySettingState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCurrentValue gets the currentValue property value. Current value of setting on device
func (m *DeviceCompliancePolicySettingState) GetCurrentValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currentValue
    }
}
// GetErrorCode gets the errorCode property value. Error code for the setting
func (m *DeviceCompliancePolicySettingState) GetErrorCode()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetErrorDescription gets the errorDescription property value. Error description
func (m *DeviceCompliancePolicySettingState) GetErrorDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorDescription
    }
}
// GetInstanceDisplayName gets the instanceDisplayName property value. Name of setting instance that is being reported.
func (m *DeviceCompliancePolicySettingState) GetInstanceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.instanceDisplayName
    }
}
// GetSetting gets the setting property value. The setting that is being reported
func (m *DeviceCompliancePolicySettingState) GetSetting()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setting
    }
}
// GetSettingName gets the settingName property value. Localized/user friendly setting name that is being reported
func (m *DeviceCompliancePolicySettingState) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// GetSources gets the sources property value. Contributing policies
func (m *DeviceCompliancePolicySettingState) GetSources()([]SettingSource) {
    if m == nil {
        return nil
    } else {
        return m.sources
    }
}
// GetState gets the state property value. The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *DeviceCompliancePolicySettingState) GetState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUserEmail gets the userEmail property value. UserEmail
func (m *DeviceCompliancePolicySettingState) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
// GetUserId gets the userId property value. UserId
func (m *DeviceCompliancePolicySettingState) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserName gets the userName property value. UserName
func (m *DeviceCompliancePolicySettingState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. UserPrincipalName.
func (m *DeviceCompliancePolicySettingState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicySettingState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["currentValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentValue(val)
        }
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["errorDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDescription(val)
        }
        return nil
    }
    res["instanceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceDisplayName(val)
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
    res["sources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingSource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingSource, len(val))
            for i, v := range val {
                res[i] = *(v.(*SettingSource))
            }
            m.SetSources(res)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ComplianceStatus)
            m.SetState(&cast)
        }
        return nil
    }
    res["userEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *DeviceCompliancePolicySettingState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicySettingState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("currentValue", m.GetCurrentValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorDescription", m.GetErrorDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instanceDisplayName", m.GetInstanceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSources()))
        for i, v := range m.GetSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("sources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userEmail", m.GetUserEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePolicySettingState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCurrentValue sets the currentValue property value. Current value of setting on device
func (m *DeviceCompliancePolicySettingState) SetCurrentValue(value *string)() {
    if m != nil {
        m.currentValue = value
    }
}
// SetErrorCode sets the errorCode property value. Error code for the setting
func (m *DeviceCompliancePolicySettingState) SetErrorCode(value *int64)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetErrorDescription sets the errorDescription property value. Error description
func (m *DeviceCompliancePolicySettingState) SetErrorDescription(value *string)() {
    if m != nil {
        m.errorDescription = value
    }
}
// SetInstanceDisplayName sets the instanceDisplayName property value. Name of setting instance that is being reported.
func (m *DeviceCompliancePolicySettingState) SetInstanceDisplayName(value *string)() {
    if m != nil {
        m.instanceDisplayName = value
    }
}
// SetSetting sets the setting property value. The setting that is being reported
func (m *DeviceCompliancePolicySettingState) SetSetting(value *string)() {
    if m != nil {
        m.setting = value
    }
}
// SetSettingName sets the settingName property value. Localized/user friendly setting name that is being reported
func (m *DeviceCompliancePolicySettingState) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
// SetSources sets the sources property value. Contributing policies
func (m *DeviceCompliancePolicySettingState) SetSources(value []SettingSource)() {
    if m != nil {
        m.sources = value
    }
}
// SetState sets the state property value. The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *DeviceCompliancePolicySettingState) SetState(value *ComplianceStatus)() {
    if m != nil {
        m.state = value
    }
}
// SetUserEmail sets the userEmail property value. UserEmail
func (m *DeviceCompliancePolicySettingState) SetUserEmail(value *string)() {
    if m != nil {
        m.userEmail = value
    }
}
// SetUserId sets the userId property value. UserId
func (m *DeviceCompliancePolicySettingState) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserName sets the userName property value. UserName
func (m *DeviceCompliancePolicySettingState) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. UserPrincipalName.
func (m *DeviceCompliancePolicySettingState) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
