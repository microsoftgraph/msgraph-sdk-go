package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TokenLifetimePolicy provides operations to manage the collection of application entities.
type TokenLifetimePolicy struct {
    StsPolicy
}
// NewTokenLifetimePolicy instantiates a new tokenLifetimePolicy and sets the default values.
func NewTokenLifetimePolicy()(*TokenLifetimePolicy) {
    m := &TokenLifetimePolicy{
        StsPolicy: *NewStsPolicy(),
    }
    return m
}
// CreateTokenLifetimePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTokenLifetimePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTokenLifetimePolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TokenLifetimePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StsPolicy.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TokenLifetimePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StsPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
