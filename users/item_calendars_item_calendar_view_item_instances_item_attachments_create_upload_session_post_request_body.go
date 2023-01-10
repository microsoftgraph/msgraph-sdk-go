package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody 
type ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The AttachmentItem property
    attachmentItem iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable
}
// NewItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody()(*ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttachmentItem gets the attachmentItem property value. The AttachmentItem property
func (m *ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentItem()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable) {
    return m.attachmentItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttachmentItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentItem(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttachmentItem sets the attachmentItem property value. The AttachmentItem property
func (m *ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentItem(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable)() {
    m.attachmentItem = value
}
