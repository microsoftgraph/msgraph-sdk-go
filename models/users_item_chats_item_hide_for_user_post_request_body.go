package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemChatsItemHideForUserPostRequestBody provides operations to call the hideForUser method.
type UsersItemChatsItemHideForUserPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The user property
    user TeamworkUserIdentityable
}
// NewUsersItemChatsItemHideForUserPostRequestBody instantiates a new UsersItemChatsItemHideForUserPostRequestBody and sets the default values.
func NewUsersItemChatsItemHideForUserPostRequestBody()(*UsersItemChatsItemHideForUserPostRequestBody) {
    m := &UsersItemChatsItemHideForUserPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemChatsItemHideForUserPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemChatsItemHideForUserPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemChatsItemHideForUserPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemChatsItemHideForUserPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemChatsItemHideForUserPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(TeamworkUserIdentityable))
        }
        return nil
    }
    return res
}
// GetUser gets the user property value. The user property
func (m *UsersItemChatsItemHideForUserPostRequestBody) GetUser()(TeamworkUserIdentityable) {
    return m.user
}
// Serialize serializes information the current object
func (m *UsersItemChatsItemHideForUserPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *UsersItemChatsItemHideForUserPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetUser sets the user property value. The user property
func (m *UsersItemChatsItemHideForUserPostRequestBody) SetUser(value TeamworkUserIdentityable)() {
    m.user = value
}
