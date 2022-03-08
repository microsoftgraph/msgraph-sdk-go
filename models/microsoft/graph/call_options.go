package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CallOptions provides operations to manage the cloudCommunications singleton.
type CallOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    hideBotAfterEscalation *bool;
}
// NewCallOptions instantiates a new callOptions and sets the default values.
func NewCallOptions()(*CallOptions) {
    m := &CallOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCallOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallOptionsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCallOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallOptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hideBotAfterEscalation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideBotAfterEscalation(val)
        }
        return nil
    }
    return res
}
// GetHideBotAfterEscalation gets the hideBotAfterEscalation property value. 
func (m *CallOptions) GetHideBotAfterEscalation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideBotAfterEscalation
    }
}
func (m *CallOptions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CallOptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hideBotAfterEscalation", m.GetHideBotAfterEscalation())
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
func (m *CallOptions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHideBotAfterEscalation sets the hideBotAfterEscalation property value. 
func (m *CallOptions) SetHideBotAfterEscalation(value *bool)() {
    if m != nil {
        m.hideBotAfterEscalation = value
    }
}
