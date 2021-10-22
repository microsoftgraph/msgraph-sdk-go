package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SamlSingleSignOnSettings struct {
    additionalData map[string]interface{};
    relayState *string;
}
func NewSamlSingleSignOnSettings()(*SamlSingleSignOnSettings) {
    m := &SamlSingleSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SamlSingleSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SamlSingleSignOnSettings) GetRelayState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relayState
    }
}
func (m *SamlSingleSignOnSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["relayState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRelayState(val)
        return nil
    }
    return res
}
func (m *SamlSingleSignOnSettings) IsNil()(bool) {
    return m == nil
}
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
func (m *SamlSingleSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SamlSingleSignOnSettings) SetRelayState(value *string)() {
    m.relayState = value
}
