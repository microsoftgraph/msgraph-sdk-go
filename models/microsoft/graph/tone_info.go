package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ToneInfo 
type ToneInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // An incremental identifier used for ordering DTMF events.
    sequenceId *int64;
    // Possible values are: tone0, tone1, tone2, tone3, tone4, tone5, tone6, tone7, tone8, tone9, star, pound, a, b, c, d, flash.
    tone *Tone;
}
// NewToneInfo instantiates a new toneInfo and sets the default values.
func NewToneInfo()(*ToneInfo) {
    m := &ToneInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ToneInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetSequenceId gets the sequenceId property value. An incremental identifier used for ordering DTMF events.
func (m *ToneInfo) GetSequenceId()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sequenceId
    }
}
// GetTone gets the tone property value. Possible values are: tone0, tone1, tone2, tone3, tone4, tone5, tone6, tone7, tone8, tone9, star, pound, a, b, c, d, flash.
func (m *ToneInfo) GetTone()(*Tone) {
    if m == nil {
        return nil
    } else {
        return m.tone
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ToneInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["sequenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequenceId(val)
        }
        return nil
    }
    res["tone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTone)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(Tone)
            m.SetTone(&cast)
        }
        return nil
    }
    return res
}
func (m *ToneInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ToneInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSequenceId sets the sequenceId property value. An incremental identifier used for ordering DTMF events.
func (m *ToneInfo) SetSequenceId(value *int64)() {
    if m != nil {
        m.sequenceId = value
    }
}
// SetTone sets the tone property value. Possible values are: tone0, tone1, tone2, tone3, tone4, tone5, tone6, tone7, tone8, tone9, star, pound, a, b, c, d, flash.
func (m *ToneInfo) SetTone(value *Tone)() {
    if m != nil {
        m.tone = value
    }
}
