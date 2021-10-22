package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Phone struct {
    additionalData map[string]interface{};
    language *string;
    number *string;
    region *string;
    type_escpaped *PhoneType;
}
func NewPhone()(*Phone) {
    m := &Phone{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Phone) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Phone) GetLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
func (m *Phone) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
func (m *Phone) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
func (m *Phone) GetType_escpaped()(*PhoneType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *Phone) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["language"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLanguage(val)
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNumber(val)
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRegion(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePhoneType)
        if err != nil {
            return err
        }
        cast := val.(PhoneType)
        m.SetType_escpaped(&cast)
        return nil
    }
    return res
}
func (m *Phone) IsNil()(bool) {
    return m == nil
}
func (m *Phone) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err := writer.WriteStringValue("type_escpaped", &cast)
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
func (m *Phone) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Phone) SetLanguage(value *string)() {
    m.language = value
}
func (m *Phone) SetNumber(value *string)() {
    m.number = value
}
func (m *Phone) SetRegion(value *string)() {
    m.region = value
}
func (m *Phone) SetType_escpaped(value *PhoneType)() {
    m.type_escpaped = value
}
