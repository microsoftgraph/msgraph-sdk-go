package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AudioConferencing struct {
    additionalData map[string]interface{};
    conferenceId *string;
    dialinUrl *string;
    tollFreeNumber *string;
    tollNumber *string;
}
func NewAudioConferencing()(*AudioConferencing) {
    m := &AudioConferencing{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AudioConferencing) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AudioConferencing) GetConferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conferenceId
    }
}
func (m *AudioConferencing) GetDialinUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dialinUrl
    }
}
func (m *AudioConferencing) GetTollFreeNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tollFreeNumber
    }
}
func (m *AudioConferencing) GetTollNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tollNumber
    }
}
func (m *AudioConferencing) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["conferenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConferenceId(val)
        return nil
    }
    res["dialinUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDialinUrl(val)
        return nil
    }
    res["tollFreeNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTollFreeNumber(val)
        return nil
    }
    res["tollNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTollNumber(val)
        return nil
    }
    return res
}
func (m *AudioConferencing) IsNil()(bool) {
    return m == nil
}
func (m *AudioConferencing) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("conferenceId", m.GetConferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dialinUrl", m.GetDialinUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tollFreeNumber", m.GetTollFreeNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tollNumber", m.GetTollNumber())
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
func (m *AudioConferencing) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AudioConferencing) SetConferenceId(value *string)() {
    m.conferenceId = value
}
func (m *AudioConferencing) SetDialinUrl(value *string)() {
    m.dialinUrl = value
}
func (m *AudioConferencing) SetTollFreeNumber(value *string)() {
    m.tollFreeNumber = value
}
func (m *AudioConferencing) SetTollNumber(value *string)() {
    m.tollNumber = value
}
