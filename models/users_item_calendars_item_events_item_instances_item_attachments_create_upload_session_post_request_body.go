package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody provides operations to call the createUploadSession method.
type UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The AttachmentItem property
    attachmentItem AttachmentItemable
}
// NewUsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewUsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody()(*UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttachmentItem gets the attachmentItem property value. The AttachmentItem property
func (m *UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentItem()(AttachmentItemable) {
    return m.attachmentItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttachmentItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentItem(val.(AttachmentItemable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttachmentItem sets the attachmentItem property value. The AttachmentItem property
func (m *UsersItemCalendarsItemEventsItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentItem(value AttachmentItemable)() {
    m.attachmentItem = value
}
