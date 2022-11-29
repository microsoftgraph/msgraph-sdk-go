package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody provides operations to call the checkMemberGroups method.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groupIds property
    groupIds []string
}
// NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody instantiates a new MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody and sets the default values.
func NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) {
    m := &MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetGroupIds)
    return res
}
// GetGroupIds gets the groupIds property value. The groupIds property
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) GetGroupIds()([]string) {
    return m.groupIds
}
// Serialize serializes information the current object
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroupIds() != nil {
        err := writer.WriteCollectionOfStringValues("groupIds", m.GetGroupIds())
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
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGroupIds sets the groupIds property value. The groupIds property
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) SetGroupIds(value []string)() {
    m.groupIds = value
}
