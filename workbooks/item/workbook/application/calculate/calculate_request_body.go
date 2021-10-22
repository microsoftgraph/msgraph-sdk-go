package calculate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CalculateRequestBody struct {
    additionalData map[string]interface{};
    calculationType *string;
}
func NewCalculateRequestBody()(*CalculateRequestBody) {
    m := &CalculateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CalculateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CalculateRequestBody) GetCalculationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.calculationType
    }
}
func (m *CalculateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calculationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCalculationType(val)
        return nil
    }
    return res
}
func (m *CalculateRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *CalculateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("calculationType", m.GetCalculationType())
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
func (m *CalculateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CalculateRequestBody) SetCalculationType(value *string)() {
    m.calculationType = value
}
