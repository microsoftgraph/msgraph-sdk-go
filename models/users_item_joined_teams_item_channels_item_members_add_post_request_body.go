package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody provides operations to call the add method.
type UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The values property
    values []ConversationMemberable
}
// NewUsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody instantiates a new UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody and sets the default values.
func NewUsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody()(*UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody) {
    m := &UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["values"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateConversationMemberFromDiscriminatorValue , m.SetValues)
    return res
}
// GetValues gets the values property value. The values property
func (m *UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody) GetValues()([]ConversationMemberable) {
    return m.values
}
// Serialize serializes information the current object
func (m *UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValues() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValues())
        err := writer.WriteCollectionOfObjectValues("values", cast)
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
func (m *UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValues sets the values property value. The values property
func (m *UsersItemJoinedTeamsItemChannelsItemMembersAddPostRequestBody) SetValues(value []ConversationMemberable)() {
    m.values = value
}
