package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// samlSingleSignOnSettings 
type SamlSingleSignOnSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The relative URI the service provider would redirect to after completion of the single sign-on flow.
    relayState *string;
}
// NewSamlSingleSignOnSettings instantiates a new samlSingleSignOnSettings and sets the default values.
func NewSamlSingleSignOnSettings()(*SamlSingleSignOnSettings) {
    m := &SamlSingleSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SamlSingleSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetRelayState gets the relayState property value. The relative URI the service provider would redirect to after completion of the single sign-on flow.
func (m *SamlSingleSignOnSettings) GetRelayState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relayState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SamlSingleSignOnSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["relayState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelayState(val)
        }
        return nil
    }
    return res
}
func (m *SamlSingleSignOnSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SamlSingleSignOnSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("relayState", m.GetRelayState())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SamlSingleSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRelayState sets the relayState property value. The relative URI the service provider would redirect to after completion of the single sign-on flow.
func (m *SamlSingleSignOnSettings) SetRelayState(value *string)() {
    m.relayState = value
}
