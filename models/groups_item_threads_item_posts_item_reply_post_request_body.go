package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemThreadsItemPostsItemReplyPostRequestBody provides operations to call the reply method.
type GroupsItemThreadsItemPostsItemReplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Post property
    post Postable
}
// NewGroupsItemThreadsItemPostsItemReplyPostRequestBody instantiates a new GroupsItemThreadsItemPostsItemReplyPostRequestBody and sets the default values.
func NewGroupsItemThreadsItemPostsItemReplyPostRequestBody()(*GroupsItemThreadsItemPostsItemReplyPostRequestBody) {
    m := &GroupsItemThreadsItemPostsItemReplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemThreadsItemPostsItemReplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemThreadsItemPostsItemReplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemThreadsItemPostsItemReplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemThreadsItemPostsItemReplyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemThreadsItemPostsItemReplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *GroupsItemThreadsItemPostsItemReplyPostRequestBody) GetPost()(Postable) {
    return m.post
}
// Serialize serializes information the current object
func (m *GroupsItemThreadsItemPostsItemReplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GroupsItemThreadsItemPostsItemReplyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPost sets the post property value. The Post property
func (m *GroupsItemThreadsItemPostsItemReplyPostRequestBody) SetPost(value Postable)() {
    m.post = value
}
