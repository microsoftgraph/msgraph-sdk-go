package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RestrictedSignIn 
type RestrictedSignIn struct {
    SignIn
    // The targetTenantId property
    targetTenantId *string
}
// NewRestrictedSignIn instantiates a new RestrictedSignIn and sets the default values.
func NewRestrictedSignIn()(*RestrictedSignIn) {
    m := &RestrictedSignIn{
        SignIn: *NewSignIn(),
    }
    odataTypeValue := "#microsoft.graph.restrictedSignIn";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRestrictedSignInFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRestrictedSignInFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRestrictedSignIn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RestrictedSignIn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SignIn.GetFieldDeserializers()
    res["targetTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetTenantId(val)
        }
        return nil
    }
    return res
}
// GetTargetTenantId gets the targetTenantId property value. The targetTenantId property
func (m *RestrictedSignIn) GetTargetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetTenantId
    }
}
// Serialize serializes information the current object
func (m *RestrictedSignIn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SignIn.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetTenantId", m.GetTargetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetTenantId sets the targetTenantId property value. The targetTenantId property
func (m *RestrictedSignIn) SetTargetTenantId(value *string)() {
    if m != nil {
        m.targetTenantId = value
    }
}
