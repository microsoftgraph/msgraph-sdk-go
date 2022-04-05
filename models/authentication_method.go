package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethod 
type AuthenticationMethod struct {
    Entity
}
// NewAuthenticationMethod instantiates a new authenticationMethod and sets the default values.
func NewAuthenticationMethod()(*AuthenticationMethod) {
    m := &AuthenticationMethod{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethod(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
