package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody 
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The securityEnabledOnly property
    securityEnabledOnly *bool
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody instantiates a new AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody) {
    m := &AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["securityEnabledOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabledOnly(val)
        }
        return nil
    }
    return res
}
// GetSecurityEnabledOnly gets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody) GetSecurityEnabledOnly()(*bool) {
    return m.securityEnabledOnly
}
// Serialize serializes information the current object
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("securityEnabledOnly", m.GetSecurityEnabledOnly())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSecurityEnabledOnly sets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBody) SetSecurityEnabledOnly(value *bool)() {
    m.securityEnabledOnly = value
}
