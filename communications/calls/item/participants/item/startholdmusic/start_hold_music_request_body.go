package startholdmusic

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// StartHoldMusicRequestBody provides operations to call the startHoldMusic method.
type StartHoldMusicRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The clientContext property
    clientContext *string;
    // The customPrompt property
    customPrompt iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable;
}
// NewStartHoldMusicRequestBody instantiates a new startHoldMusicRequestBody and sets the default values.
func NewStartHoldMusicRequestBody()(*StartHoldMusicRequestBody) {
    m := &StartHoldMusicRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateStartHoldMusicRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStartHoldMusicRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStartHoldMusicRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StartHoldMusicRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *StartHoldMusicRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
// GetCustomPrompt gets the customPrompt property value. The customPrompt property
func (m *StartHoldMusicRequestBody) GetCustomPrompt()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable) {
    if m == nil {
        return nil
    } else {
        return m.customPrompt
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StartHoldMusicRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientContext"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["customPrompt"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePromptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrompt(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *StartHoldMusicRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *StartHoldMusicRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *StartHoldMusicRequestBody) SetClientContext(value *string)() {
    if m != nil {
        m.clientContext = value
    }
}
// SetCustomPrompt sets the customPrompt property value. The customPrompt property
func (m *StartHoldMusicRequestBody) SetCustomPrompt(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable)() {
    if m != nil {
        m.customPrompt = value
    }
}
