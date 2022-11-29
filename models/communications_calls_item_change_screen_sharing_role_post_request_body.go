package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallsItemChangeScreenSharingRolePostRequestBody provides operations to call the changeScreenSharingRole method.
type CommunicationsCallsItemChangeScreenSharingRolePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The role property
    role *ScreenSharingRole
}
// NewCommunicationsCallsItemChangeScreenSharingRolePostRequestBody instantiates a new CommunicationsCallsItemChangeScreenSharingRolePostRequestBody and sets the default values.
func NewCommunicationsCallsItemChangeScreenSharingRolePostRequestBody()(*CommunicationsCallsItemChangeScreenSharingRolePostRequestBody) {
    m := &CommunicationsCallsItemChangeScreenSharingRolePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommunicationsCallsItemChangeScreenSharingRolePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsCallsItemChangeScreenSharingRolePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsCallsItemChangeScreenSharingRolePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsCallsItemChangeScreenSharingRolePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsCallsItemChangeScreenSharingRolePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["role"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseScreenSharingRole , m.SetRole)
    return res
}
// GetRole gets the role property value. The role property
func (m *CommunicationsCallsItemChangeScreenSharingRolePostRequestBody) GetRole()(*ScreenSharingRole) {
    return m.role
}
// Serialize serializes information the current object
func (m *CommunicationsCallsItemChangeScreenSharingRolePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
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
func (m *CommunicationsCallsItemChangeScreenSharingRolePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRole sets the role property value. The role property
func (m *CommunicationsCallsItemChangeScreenSharingRolePostRequestBody) SetRole(value *ScreenSharingRole)() {
    m.role = value
}
