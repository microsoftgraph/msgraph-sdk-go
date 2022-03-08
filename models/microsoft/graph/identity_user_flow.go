package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityUserFlow provides operations to manage the identityContainer singleton.
type IdentityUserFlow struct {
    Entity
    // 
    userFlowType *UserFlowType;
    // 
    userFlowTypeVersion *float32;
}
// NewIdentityUserFlow instantiates a new identityUserFlow and sets the default values.
func NewIdentityUserFlow()(*IdentityUserFlow) {
    m := &IdentityUserFlow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIdentityUserFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityUserFlowFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewIdentityUserFlow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityUserFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["userFlowType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserFlowType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserFlowType(val.(*UserFlowType))
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
// GetUserFlowType gets the userFlowType property value. 
func (m *IdentityUserFlow) GetUserFlowType()(*UserFlowType) {
    if m == nil {
        return nil
    } else {
        return m.userFlowType
    }
}
// GetUserFlowTypeVersion gets the userFlowTypeVersion property value. 
func (m *IdentityUserFlow) GetUserFlowTypeVersion()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.userFlowTypeVersion
    }
}
func (m *IdentityUserFlow) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IdentityUserFlow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserFlowType() != nil {
        cast := (*m.GetUserFlowType()).String()
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
// SetUserFlowType sets the userFlowType property value. 
func (m *IdentityUserFlow) SetUserFlowType(value *UserFlowType)() {
    if m != nil {
        m.userFlowType = value
    }
}
// SetUserFlowTypeVersion sets the userFlowTypeVersion property value. 
func (m *IdentityUserFlow) SetUserFlowTypeVersion(value *float32)() {
    if m != nil {
        m.userFlowTypeVersion = value
    }
}
