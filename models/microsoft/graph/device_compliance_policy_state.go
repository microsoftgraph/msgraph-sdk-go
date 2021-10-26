package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceCompliancePolicyState struct {
    Entity
    // The name of the policy for this policyBase
    displayName *string;
    // Platform type that the policy applies to
    platformType *PolicyPlatformType;
    // Count of how many setting a policy holds
    settingCount *int32;
    // 
    settingStates []DeviceCompliancePolicySettingState;
    // The compliance state of the policy
    state *ComplianceStatus;
    // The version of the policy
    version *int32;
}
// Instantiates a new deviceCompliancePolicyState and sets the default values.
func NewDeviceCompliancePolicyState()(*DeviceCompliancePolicyState) {
    m := &DeviceCompliancePolicyState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The name of the policy for this policyBase
func (m *DeviceCompliancePolicyState) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the platformType property value. Platform type that the policy applies to
func (m *DeviceCompliancePolicyState) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// Gets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceCompliancePolicyState) GetSettingCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.settingCount
    }
}
// Gets the settingStates property value. 
func (m *DeviceCompliancePolicyState) GetSettingStates()([]DeviceCompliancePolicySettingState) {
    if m == nil {
        return nil
    } else {
        return m.settingStates
    }
}
// Gets the state property value. The compliance state of the policy
func (m *DeviceCompliancePolicyState) GetState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the version property value. The version of the policy
func (m *DeviceCompliancePolicyState) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceCompliancePolicyState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
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
    res["settingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSettingCount(val)
        return nil
    }
    res["settingStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicySettingState() })
        if err != nil {
            return err
        }
        res := make([]DeviceCompliancePolicySettingState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCompliancePolicySettingState))
        }
        m.SetSettingStates(res)
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
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *DeviceCompliancePolicyState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceCompliancePolicyState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteInt32Value("settingCount", m.GetSettingCount())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettingStates()))
        for i, v := range m.GetSettingStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settingStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The name of the policy for this policyBase
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceCompliancePolicyState) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the platformType property value. Platform type that the policy applies to
// Parameters:
//  - value : Value to set for the platformType property.
func (m *DeviceCompliancePolicyState) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// Sets the settingCount property value. Count of how many setting a policy holds
// Parameters:
//  - value : Value to set for the settingCount property.
func (m *DeviceCompliancePolicyState) SetSettingCount(value *int32)() {
    m.settingCount = value
}
// Sets the settingStates property value. 
// Parameters:
//  - value : Value to set for the settingStates property.
func (m *DeviceCompliancePolicyState) SetSettingStates(value []DeviceCompliancePolicySettingState)() {
    m.settingStates = value
}
// Sets the state property value. The compliance state of the policy
// Parameters:
//  - value : Value to set for the state property.
func (m *DeviceCompliancePolicyState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// Sets the version property value. The version of the policy
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceCompliancePolicyState) SetVersion(value *int32)() {
    m.version = value
}
