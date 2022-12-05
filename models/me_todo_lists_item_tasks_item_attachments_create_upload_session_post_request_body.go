package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody provides operations to call the createUploadSession method.
type MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The attachmentInfo property
    attachmentInfo AttachmentInfoable
}
// NewMeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewMeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody()(*MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttachmentInfo gets the attachmentInfo property value. The attachmentInfo property
func (m *MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentInfo()(AttachmentInfoable) {
    return m.attachmentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttachmentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentInfo(val.(AttachmentInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachmentInfo", m.GetAttachmentInfo())
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
func (m *MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttachmentInfo sets the attachmentInfo property value. The attachmentInfo property
func (m *MeTodoListsItemTasksItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentInfo(value AttachmentInfoable)() {
    m.attachmentInfo = value
}
