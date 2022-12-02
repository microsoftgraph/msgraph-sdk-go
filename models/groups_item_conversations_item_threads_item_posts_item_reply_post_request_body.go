package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody provides operations to call the reply method.
type GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Post property
    post Postable
}
// NewGroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody instantiates a new GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody and sets the default values.
func NewGroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody()(*GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody) {
    m := &GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["post"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPost(val.(Postable))
        }
        return nil
    }
    return res
}
// GetPost gets the post property value. The Post property
func (m *GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody) GetPost()(Postable) {
    return m.post
}
// Serialize serializes information the current object
func (m *GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("post", m.GetPost())
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
func (m *GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPost sets the post property value. The Post property
func (m *GroupsItemConversationsItemThreadsItemPostsItemReplyPostRequestBody) SetPost(value Postable)() {
    m.post = value
}
