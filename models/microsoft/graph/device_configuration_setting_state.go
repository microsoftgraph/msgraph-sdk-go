package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceConfigurationSettingState struct {
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
// Instantiates a new deviceConfigurationSettingState and sets the default values.
func NewDeviceConfigurationSettingState()(*DeviceConfigurationSettingState) {
    m := &DeviceConfigurationSettingState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurationSettingState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the currentValue property value. Current value of setting on device
func (m *DeviceConfigurationSettingState) GetCurrentValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currentValue
    }
}
// Gets the errorCode property value. Error code for the setting
func (m *DeviceConfigurationSettingState) GetErrorCode()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the errorDescription property value. Error description
func (m *DeviceConfigurationSettingState) GetErrorDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorDescription
    }
}
// Gets the instanceDisplayName property value. Name of setting instance that is being reported.
func (m *DeviceConfigurationSettingState) GetInstanceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.instanceDisplayName
    }
}
// Gets the setting property value. The setting that is being reported
func (m *DeviceConfigurationSettingState) GetSetting()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setting
    }
}
// Gets the settingName property value. Localized/user friendly setting name that is being reported
func (m *DeviceConfigurationSettingState) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// Gets the sources property value. Contributing policies
func (m *DeviceConfigurationSettingState) GetSources()([]SettingSource) {
    if m == nil {
        return nil
    } else {
        return m.sources
    }
}
// Gets the state property value. The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *DeviceConfigurationSettingState) GetState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the userEmail property value. UserEmail
func (m *DeviceConfigurationSettingState) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
// Gets the userId property value. UserId
func (m *DeviceConfigurationSettingState) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userName property value. UserName
func (m *DeviceConfigurationSettingState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// Gets the userPrincipalName property value. UserPrincipalName.
func (m *DeviceConfigurationSettingState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *DeviceConfigurationSettingState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["currentValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCurrentValue(val)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["errorDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorDescription(val)
        return nil
    }
    res["instanceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInstanceDisplayName(val)
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
    res["sources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingSource() })
        if err != nil {
            return err
        }
        res := make([]SettingSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*SettingSource))
        }
        m.SetSources(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        cast := val.(ComplianceStatus)
        m.SetState(&cast)
        return nil
    }
    res["userEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserEmail(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *DeviceConfigurationSettingState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceConfigurationSettingState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceConfigurationSettingState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the currentValue property value. Current value of setting on device
// Parameters:
//  - value : Value to set for the currentValue property.
func (m *DeviceConfigurationSettingState) SetCurrentValue(value *string)() {
    m.currentValue = value
}
// Sets the errorCode property value. Error code for the setting
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *DeviceConfigurationSettingState) SetErrorCode(value *int64)() {
    m.errorCode = value
}
// Sets the errorDescription property value. Error description
// Parameters:
//  - value : Value to set for the errorDescription property.
func (m *DeviceConfigurationSettingState) SetErrorDescription(value *string)() {
    m.errorDescription = value
}
// Sets the instanceDisplayName property value. Name of setting instance that is being reported.
// Parameters:
//  - value : Value to set for the instanceDisplayName property.
func (m *DeviceConfigurationSettingState) SetInstanceDisplayName(value *string)() {
    m.instanceDisplayName = value
}
// Sets the setting property value. The setting that is being reported
// Parameters:
//  - value : Value to set for the setting property.
func (m *DeviceConfigurationSettingState) SetSetting(value *string)() {
    m.setting = value
}
// Sets the settingName property value. Localized/user friendly setting name that is being reported
// Parameters:
//  - value : Value to set for the settingName property.
func (m *DeviceConfigurationSettingState) SetSettingName(value *string)() {
    m.settingName = value
}
// Sets the sources property value. Contributing policies
// Parameters:
//  - value : Value to set for the sources property.
func (m *DeviceConfigurationSettingState) SetSources(value []SettingSource)() {
    m.sources = value
}
// Sets the state property value. The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
// Parameters:
//  - value : Value to set for the state property.
func (m *DeviceConfigurationSettingState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// Sets the userEmail property value. UserEmail
// Parameters:
//  - value : Value to set for the userEmail property.
func (m *DeviceConfigurationSettingState) SetUserEmail(value *string)() {
    m.userEmail = value
}
// Sets the userId property value. UserId
// Parameters:
//  - value : Value to set for the userId property.
func (m *DeviceConfigurationSettingState) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userName property value. UserName
// Parameters:
//  - value : Value to set for the userName property.
func (m *DeviceConfigurationSettingState) SetUserName(value *string)() {
    m.userName = value
}
// Sets the userPrincipalName property value. UserPrincipalName.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *DeviceConfigurationSettingState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
