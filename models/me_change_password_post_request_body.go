package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeChangePasswordPostRequestBody provides operations to call the changePassword method.
type MeChangePasswordPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The currentPassword property
    currentPassword *string
    // The newPassword property
    newPassword *string
}
// NewMeChangePasswordPostRequestBody instantiates a new MeChangePasswordPostRequestBody and sets the default values.
func NewMeChangePasswordPostRequestBody()(*MeChangePasswordPostRequestBody) {
    m := &MeChangePasswordPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeChangePasswordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeChangePasswordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeChangePasswordPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeChangePasswordPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCurrentPassword gets the currentPassword property value. The currentPassword property
func (m *MeChangePasswordPostRequestBody) GetCurrentPassword()(*string) {
    return m.currentPassword
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeChangePasswordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPassword(val)
        }
        return nil
    }
    res["newPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewPassword(val)
        }
        return nil
    }
    return res
}
// GetNewPassword gets the newPassword property value. The newPassword property
func (m *MeChangePasswordPostRequestBody) GetNewPassword()(*string) {
    return m.newPassword
}
// Serialize serializes information the current object
func (m *MeChangePasswordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("currentPassword", m.GetCurrentPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newPassword", m.GetNewPassword())
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
func (m *MeChangePasswordPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCurrentPassword sets the currentPassword property value. The currentPassword property
func (m *MeChangePasswordPostRequestBody) SetCurrentPassword(value *string)() {
    m.currentPassword = value
}
// SetNewPassword sets the newPassword property value. The newPassword property
func (m *MeChangePasswordPostRequestBody) SetNewPassword(value *string)() {
    m.newPassword = value
}
