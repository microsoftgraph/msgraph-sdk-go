package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCompliancePolicyState provides operations to manage the deviceManagement singleton.
type DeviceCompliancePolicyState struct {
    Entity
    // The name of the policy for this policyBase
    displayName *string;
    // Platform type that the policy applies to
    platformType *PolicyPlatformType;
    // Count of how many setting a policy holds
    settingCount *int32;
    // 
    settingStates []DeviceCompliancePolicySettingStateable;
    // The compliance state of the policy
    state *ComplianceStatus;
    // The version of the policy
    version *int32;
}
// NewDeviceCompliancePolicyState instantiates a new deviceCompliancePolicyState and sets the default values.
func NewDeviceCompliancePolicyState()(*DeviceCompliancePolicyState) {
    m := &DeviceCompliancePolicyState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceCompliancePolicyStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceCompliancePolicyState(), nil
}
// GetDisplayName gets the displayName property value. The name of the policy for this policyBase
func (m *DeviceCompliancePolicyState) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicyState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    res["settingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCount(val)
        }
        return nil
    }
    res["settingStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicySettingStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicySettingStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCompliancePolicySettingStateable)
            }
            m.SetSettingStates(res)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetPlatformType gets the platformType property value. Platform type that the policy applies to
func (m *DeviceCompliancePolicyState) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// GetSettingCount gets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceCompliancePolicyState) GetSettingCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.settingCount
    }
}
// GetSettingStates gets the settingStates property value. 
func (m *DeviceCompliancePolicyState) GetSettingStates()([]DeviceCompliancePolicySettingStateable) {
    if m == nil {
        return nil
    } else {
        return m.settingStates
    }
}
// GetState gets the state property value. The compliance state of the policy
func (m *DeviceCompliancePolicyState) GetState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetVersion gets the version property value. The version of the policy
func (m *DeviceCompliancePolicyState) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *DeviceCompliancePolicyState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetPlatformType()).String()
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
    if m.GetSettingStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettingStates()))
        for i, v := range m.GetSettingStates() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settingStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
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
// SetDisplayName sets the displayName property value. The name of the policy for this policyBase
func (m *DeviceCompliancePolicyState) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPlatformType sets the platformType property value. Platform type that the policy applies to
func (m *DeviceCompliancePolicyState) SetPlatformType(value *PolicyPlatformType)() {
    if m != nil {
        m.platformType = value
    }
}
// SetSettingCount sets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceCompliancePolicyState) SetSettingCount(value *int32)() {
    if m != nil {
        m.settingCount = value
    }
}
// SetSettingStates sets the settingStates property value. 
func (m *DeviceCompliancePolicyState) SetSettingStates(value []DeviceCompliancePolicySettingStateable)() {
    if m != nil {
        m.settingStates = value
    }
}
// SetState sets the state property value. The compliance state of the policy
func (m *DeviceCompliancePolicyState) SetState(value *ComplianceStatus)() {
    if m != nil {
        m.state = value
    }
}
// SetVersion sets the version property value. The version of the policy
func (m *DeviceCompliancePolicyState) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
