package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeSendMailPostRequestBody provides operations to call the sendMail method.
type MeSendMailPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Message property
    message Messageable
    // The SaveToSentItems property
    saveToSentItems *bool
}
// NewMeSendMailPostRequestBody instantiates a new MeSendMailPostRequestBody and sets the default values.
func NewMeSendMailPostRequestBody()(*MeSendMailPostRequestBody) {
    m := &MeSendMailPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeSendMailPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeSendMailPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeSendMailPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeSendMailPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeSendMailPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(Messageable))
        }
        return nil
    }
    res["saveToSentItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaveToSentItems(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The Message property
func (m *MeSendMailPostRequestBody) GetMessage()(Messageable) {
    return m.message
}
// GetSaveToSentItems gets the saveToSentItems property value. The SaveToSentItems property
func (m *MeSendMailPostRequestBody) GetSaveToSentItems()(*bool) {
    return m.saveToSentItems
}
// Serialize serializes information the current object
func (m *MeSendMailPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("saveToSentItems", m.GetSaveToSentItems())
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
func (m *MeSendMailPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMessage sets the message property value. The Message property
func (m *MeSendMailPostRequestBody) SetMessage(value Messageable)() {
    m.message = value
}
// SetSaveToSentItems sets the saveToSentItems property value. The SaveToSentItems property
func (m *MeSendMailPostRequestBody) SetSaveToSentItems(value *bool)() {
    m.saveToSentItems = value
}
