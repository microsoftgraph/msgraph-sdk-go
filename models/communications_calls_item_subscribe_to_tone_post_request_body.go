package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallsItemSubscribeToTonePostRequestBody provides operations to call the subscribeToTone method.
type CommunicationsCallsItemSubscribeToTonePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The clientContext property
    clientContext *string
}
// NewCommunicationsCallsItemSubscribeToTonePostRequestBody instantiates a new CommunicationsCallsItemSubscribeToTonePostRequestBody and sets the default values.
func NewCommunicationsCallsItemSubscribeToTonePostRequestBody()(*CommunicationsCallsItemSubscribeToTonePostRequestBody) {
    m := &CommunicationsCallsItemSubscribeToTonePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommunicationsCallsItemSubscribeToTonePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsCallsItemSubscribeToTonePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsCallsItemSubscribeToTonePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsCallsItemSubscribeToTonePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemSubscribeToTonePostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsCallsItemSubscribeToTonePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CommunicationsCallsItemSubscribeToTonePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
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
func (m *CommunicationsCallsItemSubscribeToTonePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemSubscribeToTonePostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
