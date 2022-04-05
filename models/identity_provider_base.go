package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
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
// CreateIdentityProviderBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityProviderBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityProviderBase(), nil
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
func (m *IdentityProviderBase) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *IdentityProviderBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m != nil {
        m.displayName = value
    }
}
