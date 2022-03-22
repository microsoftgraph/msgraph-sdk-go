package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Identity 
type Identity struct {
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity
    // The type of identity. Possible values are: user or group for Azure AD identities and externalgroup for groups in an external system.
    type_escaped *IdentityType;
}
// NewIdentity instantiates a new identity and sets the default values.
func NewIdentity()(*Identity) {
    m := &Identity{
        Entity: *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEntity(),
    }
    return m
}
// CreateIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Identity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*IdentityType))
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type of identity. Possible values are: user or group for Azure AD identities and externalgroup for groups in an external system.
func (m *Identity) GetType()(*IdentityType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *Identity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetType sets the type property value. The type of identity. Possible values are: user or group for Azure AD identities and externalgroup for groups in an external system.
func (m *Identity) SetType(value *IdentityType)() {
    if m != nil {
        m.type_escaped = value
    }
}
