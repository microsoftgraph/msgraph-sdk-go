package applycellcolorfilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApplyCellColorFilterRequestBody struct {
    additionalData map[string]interface{};
    color *string;
}
func NewApplyCellColorFilterRequestBody()(*ApplyCellColorFilterRequestBody) {
    m := &ApplyCellColorFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApplyCellColorFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApplyCellColorFilterRequestBody) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
func (m *ApplyCellColorFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *ApplyCellColorFilterRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ApplyCellColorFilterRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *ApplyCellColorFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApplyCellColorFilterRequestBody) SetColor(value *string)() {
    m.color = value
}
