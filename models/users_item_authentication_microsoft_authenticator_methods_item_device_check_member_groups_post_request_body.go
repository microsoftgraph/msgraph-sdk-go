package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody provides operations to call the checkMemberGroups method.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groupIds property
    groupIds []string
}
// NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody instantiates a new UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody and sets the default values.
func NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) {
    m := &UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupIds(res)
        }
        return nil
    }
    return res
}
// GetGroupIds gets the groupIds property value. The groupIds property
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) GetGroupIds()([]string) {
    return m.groupIds
}
// Serialize serializes information the current object
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGroupIds sets the groupIds property value. The groupIds property
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBody) SetGroupIds(value []string)() {
    m.groupIds = value
}
