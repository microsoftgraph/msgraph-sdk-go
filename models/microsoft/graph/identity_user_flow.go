package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IdentityUserFlow struct {
    Entity
    // 
    userFlowType *UserFlowType;
    // 
    userFlowTypeVersion *float32;
}
// Instantiates a new identityUserFlow and sets the default values.
func NewIdentityUserFlow()(*IdentityUserFlow) {
    m := &IdentityUserFlow{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the userFlowType property value. 
func (m *IdentityUserFlow) GetUserFlowType()(*UserFlowType) {
    if m == nil {
        return nil
    } else {
        return m.userFlowType
    }
}
// Gets the userFlowTypeVersion property value. 
func (m *IdentityUserFlow) GetUserFlowTypeVersion()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.userFlowTypeVersion
    }
}
// The deserialization information for the current model
func (m *IdentityUserFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["userFlowType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserFlowType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(UserFlowType)
            m.SetUserFlowType(&cast)
        }
        return nil
    }
    res["userFlowTypeVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserFlowTypeVersion(val)
        }
        return nil
    }
    return res
}
func (m *IdentityUserFlow) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IdentityUserFlow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserFlowType() != nil {
        cast := m.GetUserFlowType().String()
        err = writer.WriteStringValue("userFlowType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat32Value("userFlowTypeVersion", m.GetUserFlowTypeVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the userFlowType property value. 
// Parameters:
//  - value : Value to set for the userFlowType property.
func (m *IdentityUserFlow) SetUserFlowType(value *UserFlowType)() {
    m.userFlowType = value
}
// Sets the userFlowTypeVersion property value. 
// Parameters:
//  - value : Value to set for the userFlowTypeVersion property.
func (m *IdentityUserFlow) SetUserFlowTypeVersion(value *float32)() {
    m.userFlowTypeVersion = value
}
