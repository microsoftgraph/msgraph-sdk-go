package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ToneInfo struct {
    additionalData map[string]interface{};
    sequenceId *int64;
    tone *Tone;
}
func NewToneInfo()(*ToneInfo) {
    m := &ToneInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ToneInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ToneInfo) GetSequenceId()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sequenceId
    }
}
func (m *ToneInfo) GetTone()(*Tone) {
    if m == nil {
        return nil
    } else {
        return m.tone
    }
}
func (m *ToneInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["sequenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSequenceId(val)
        return nil
    }
    res["tone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTone)
        if err != nil {
            return err
        }
        cast := val.(Tone)
        m.SetTone(&cast)
        return nil
    }
    return res
}
func (m *ToneInfo) IsNil()(bool) {
    return m == nil
}
func (m *ToneInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("sequenceId", m.GetSequenceId())
        if err != nil {
            return err
        }
    }
    if m.GetTone() != nil {
        cast := m.GetTone().String()
        err := writer.WriteStringValue("tone", &cast)
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
func (m *ToneInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ToneInfo) SetSequenceId(value *int64)() {
    m.sequenceId = value
}
func (m *ToneInfo) SetTone(value *Tone)() {
    m.tone = value
}
