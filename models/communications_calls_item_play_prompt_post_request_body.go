package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallsItemPlayPromptPostRequestBody provides operations to call the playPrompt method.
type CommunicationsCallsItemPlayPromptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The clientContext property
    clientContext *string
    // The prompts property
    prompts []Promptable
}
// NewCommunicationsCallsItemPlayPromptPostRequestBody instantiates a new CommunicationsCallsItemPlayPromptPostRequestBody and sets the default values.
func NewCommunicationsCallsItemPlayPromptPostRequestBody()(*CommunicationsCallsItemPlayPromptPostRequestBody) {
    m := &CommunicationsCallsItemPlayPromptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommunicationsCallsItemPlayPromptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsCallsItemPlayPromptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsCallsItemPlayPromptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsCallsItemPlayPromptPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemPlayPromptPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsCallsItemPlayPromptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["prompts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePromptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Promptable, len(val))
            for i, v := range val {
                res[i] = v.(Promptable)
            }
            m.SetPrompts(res)
        }
        return nil
    }
    return res
}
// GetPrompts gets the prompts property value. The prompts property
func (m *CommunicationsCallsItemPlayPromptPostRequestBody) GetPrompts()([]Promptable) {
    return m.prompts
}
// Serialize serializes information the current object
func (m *CommunicationsCallsItemPlayPromptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    if m.GetPrompts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrompts()))
        for i, v := range m.GetPrompts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("prompts", cast)
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
func (m *CommunicationsCallsItemPlayPromptPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemPlayPromptPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetPrompts sets the prompts property value. The prompts property
func (m *CommunicationsCallsItemPlayPromptPostRequestBody) SetPrompts(value []Promptable)() {
    m.prompts = value
}
