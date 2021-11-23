package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityProviderBase 
type IdentityProviderBase struct {
    Entity
    // The display name of the identity provider.
    displayName *string;
}
// NewIdentityProviderBase instantiates a new identityProviderBase and sets the default values.
func NewIdentityProviderBase()(*IdentityProviderBase) {
    m := &IdentityProviderBase{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The display name of the identity provider.
func (m *IdentityProviderBase) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityProviderBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
func (m *IdentityProviderBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IdentityProviderBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the identity provider.
func (m *IdentityProviderBase) SetDisplayName(value *string)() {
    m.displayName = value
}
