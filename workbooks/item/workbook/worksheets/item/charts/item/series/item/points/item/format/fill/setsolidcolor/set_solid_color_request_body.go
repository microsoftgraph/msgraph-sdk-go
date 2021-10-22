package setsolidcolor

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SetSolidColorRequestBody struct {
    additionalData map[string]interface{};
    color *string;
}
func NewSetSolidColorRequestBody()(*SetSolidColorRequestBody) {
    m := &SetSolidColorRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SetSolidColorRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SetSolidColorRequestBody) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
func (m *SetSolidColorRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColor(val)
        return nil
    }
    return res
}
func (m *SetSolidColorRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *SetSolidColorRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
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
func (m *SetSolidColorRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SetSolidColorRequestBody) SetColor(value *string)() {
    m.color = value
}
