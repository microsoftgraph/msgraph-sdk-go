package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CallMediaState struct {
    additionalData map[string]interface{};
    audio *MediaState;
}
func NewCallMediaState()(*CallMediaState) {
    m := &CallMediaState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CallMediaState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CallMediaState) GetAudio()(*MediaState) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
func (m *CallMediaState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["audio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMediaState)
        if err != nil {
            return err
        }
        cast := val.(MediaState)
        m.SetAudio(&cast)
        return nil
    }
    return res
}
func (m *CallMediaState) IsNil()(bool) {
    return m == nil
}
func (m *CallMediaState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAudio() != nil {
        cast := m.GetAudio().String()
        err := writer.WriteStringValue("audio", &cast)
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
func (m *CallMediaState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CallMediaState) SetAudio(value *MediaState)() {
    m.audio = value
}
