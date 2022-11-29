package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody provides operations to call the startHoldMusic method.
type CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The clientContext property
    clientContext *string
    // The customPrompt property
    customPrompt Promptable
}
// NewCommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody instantiates a new CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody and sets the default values.
func NewCommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody()(*CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) {
    m := &CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetCustomPrompt gets the customPrompt property value. The customPrompt property
func (m *CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) GetCustomPrompt()(Promptable) {
    return m.customPrompt
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientContext"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetClientContext)
    res["customPrompt"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePromptFromDiscriminatorValue , m.SetCustomPrompt)
    return res
}
// Serialize serializes information the current object
func (m *CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("customPrompt", m.GetCustomPrompt())
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
func (m *CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetCustomPrompt sets the customPrompt property value. The customPrompt property
func (m *CommunicationsCallsItemParticipantsItemStartHoldMusicPostRequestBody) SetCustomPrompt(value Promptable)() {
    m.customPrompt = value
}
