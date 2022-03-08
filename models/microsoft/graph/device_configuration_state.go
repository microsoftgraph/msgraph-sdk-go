package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceConfigurationState provides operations to manage the drive singleton.
type DeviceConfigurationState struct {
    Entity
    // The name of the policy for this policyBase
    displayName *string;
    // Platform type that the policy applies to
    platformType *PolicyPlatformType;
    // Count of how many setting a policy holds
    settingCount *int32;
    // 
    settingStates []DeviceConfigurationSettingStateable;
    // The compliance state of the policy
    state *ComplianceStatus;
    // The version of the policy
    version *int32;
}
// NewDeviceConfigurationState instantiates a new deviceConfigurationState and sets the default values.
func NewDeviceConfigurationState()(*DeviceConfigurationState) {
    m := &DeviceConfigurationState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceConfigurationState(), nil
}
// GetDisplayName gets the displayName property value. The name of the policy for this policyBase
func (m *DeviceConfigurationState) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationSettingStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationSettingStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationSettingStateable)
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
func (m *DeviceConfigurationState) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// GetSettingCount gets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceConfigurationState) GetSettingCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.settingCount
    }
}
// GetSettingStates gets the settingStates property value. 
func (m *DeviceConfigurationState) GetSettingStates()([]DeviceConfigurationSettingStateable) {
    if m == nil {
        return nil
    } else {
        return m.settingStates
    }
}
// GetState gets the state property value. The compliance state of the policy
func (m *DeviceConfigurationState) GetState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetVersion gets the version property value. The version of the policy
func (m *DeviceConfigurationState) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *DeviceConfigurationState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceConfigurationState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *DeviceConfigurationState) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPlatformType sets the platformType property value. Platform type that the policy applies to
func (m *DeviceConfigurationState) SetPlatformType(value *PolicyPlatformType)() {
    if m != nil {
        m.platformType = value
    }
}
// SetSettingCount sets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceConfigurationState) SetSettingCount(value *int32)() {
    if m != nil {
        m.settingCount = value
    }
}
// SetSettingStates sets the settingStates property value. 
func (m *DeviceConfigurationState) SetSettingStates(value []DeviceConfigurationSettingStateable)() {
    if m != nil {
        m.settingStates = value
    }
}
// SetState sets the state property value. The compliance state of the policy
func (m *DeviceConfigurationState) SetState(value *ComplianceStatus)() {
    if m != nil {
        m.state = value
    }
}
// SetVersion sets the version property value. The version of the policy
func (m *DeviceConfigurationState) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
