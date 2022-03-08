package startholdmusic

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// StartHoldMusicRequestBody provides operations to call the startHoldMusic method.
type StartHoldMusicRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    clientContext *string;
    // 
    customPrompt i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Promptable;
}
// NewStartHoldMusicRequestBody instantiates a new startHoldMusicRequestBody and sets the default values.
func NewStartHoldMusicRequestBody()(*StartHoldMusicRequestBody) {
    m := &StartHoldMusicRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateStartHoldMusicRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStartHoldMusicRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
// GetClientContext gets the clientContext property value. 
func (m *StartHoldMusicRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
// GetCustomPrompt gets the customPrompt property value. 
func (m *StartHoldMusicRequestBody) GetCustomPrompt()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Promptable) {
    if m == nil {
        return nil
    } else {
        return m.customPrompt
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StartHoldMusicRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clientContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["customPrompt"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreatePromptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrompt(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Promptable))
        }
        return nil
    }
    return res
}
func (m *StartHoldMusicRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *StartHoldMusicRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetClientContext sets the clientContext property value. 
func (m *StartHoldMusicRequestBody) SetClientContext(value *string)() {
    if m != nil {
        m.clientContext = value
    }
}
// SetCustomPrompt sets the customPrompt property value. 
func (m *StartHoldMusicRequestBody) SetCustomPrompt(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Promptable)() {
    if m != nil {
        m.customPrompt = value
    }
}
