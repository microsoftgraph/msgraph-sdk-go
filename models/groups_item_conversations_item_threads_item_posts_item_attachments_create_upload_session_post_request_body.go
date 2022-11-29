package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody provides operations to call the createUploadSession method.
type GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The AttachmentItem property
    attachmentItem AttachmentItemable
}
// NewGroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewGroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody()(*GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttachmentItem gets the attachmentItem property value. The AttachmentItem property
func (m *GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentItem()(AttachmentItemable) {
    return m.attachmentItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentItem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAttachmentItemFromDiscriminatorValue , m.SetAttachmentItem)
    return res
}
// Serialize serializes information the current object
func (m *GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachmentItem", m.GetAttachmentItem())
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
func (m *GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttachmentItem sets the attachmentItem property value. The AttachmentItem property
func (m *GroupsItemConversationsItemThreadsItemPostsItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentItem(value AttachmentItemable)() {
    m.attachmentItem = value
}
