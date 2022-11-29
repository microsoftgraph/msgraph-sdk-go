package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemChatsItemMarkChatUnreadForUserPostRequestBody provides operations to call the markChatUnreadForUser method.
type UsersItemChatsItemMarkChatUnreadForUserPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The lastMessageReadDateTime property
    lastMessageReadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user property
    user TeamworkUserIdentityable
}
// NewUsersItemChatsItemMarkChatUnreadForUserPostRequestBody instantiates a new UsersItemChatsItemMarkChatUnreadForUserPostRequestBody and sets the default values.
func NewUsersItemChatsItemMarkChatUnreadForUserPostRequestBody()(*UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) {
    m := &UsersItemChatsItemMarkChatUnreadForUserPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemChatsItemMarkChatUnreadForUserPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemChatsItemMarkChatUnreadForUserPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemChatsItemMarkChatUnreadForUserPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lastMessageReadDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastMessageReadDateTime)
    res["user"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkUserIdentityFromDiscriminatorValue , m.SetUser)
    return res
}
// GetLastMessageReadDateTime gets the lastMessageReadDateTime property value. The lastMessageReadDateTime property
func (m *UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) GetLastMessageReadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastMessageReadDateTime
}
// GetUser gets the user property value. The user property
func (m *UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) GetUser()(TeamworkUserIdentityable) {
    return m.user
}
// Serialize serializes information the current object
func (m *UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastMessageReadDateTime", m.GetLastMessageReadDateTime())
        if err != nil {
            return err
        }
    }
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
func (m *UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLastMessageReadDateTime sets the lastMessageReadDateTime property value. The lastMessageReadDateTime property
func (m *UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) SetLastMessageReadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastMessageReadDateTime = value
}
// SetUser sets the user property value. The user property
func (m *UsersItemChatsItemMarkChatUnreadForUserPostRequestBody) SetUser(value TeamworkUserIdentityable)() {
    m.user = value
}
