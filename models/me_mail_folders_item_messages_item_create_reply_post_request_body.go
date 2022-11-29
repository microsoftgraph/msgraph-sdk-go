package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeMailFoldersItemMessagesItemCreateReplyPostRequestBody provides operations to call the createReply method.
type MeMailFoldersItemMessagesItemCreateReplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Comment property
    comment *string
    // The Message property
    message Messageable
}
// NewMeMailFoldersItemMessagesItemCreateReplyPostRequestBody instantiates a new MeMailFoldersItemMessagesItemCreateReplyPostRequestBody and sets the default values.
func NewMeMailFoldersItemMessagesItemCreateReplyPostRequestBody()(*MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) {
    m := &MeMailFoldersItemMessagesItemCreateReplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeMailFoldersItemMessagesItemCreateReplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeMailFoldersItemMessagesItemCreateReplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeMailFoldersItemMessagesItemCreateReplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetComment)
    res["message"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMessageFromDiscriminatorValue , m.SetMessage)
    return res
}
// GetMessage gets the message property value. The Message property
func (m *MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) GetMessage()(Messageable) {
    return m.message
}
// Serialize serializes information the current object
func (m *MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
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
func (m *MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetMessage sets the message property value. The Message property
func (m *MeMailFoldersItemMessagesItemCreateReplyPostRequestBody) SetMessage(value Messageable)() {
    m.message = value
}
