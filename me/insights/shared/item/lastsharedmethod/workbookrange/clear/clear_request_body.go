package clear

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ClearRequestBody struct {
    additionalData map[string]interface{};
    applyTo *string;
}
func NewClearRequestBody()(*ClearRequestBody) {
    m := &ClearRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ClearRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ClearRequestBody) GetApplyTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applyTo
    }
}
func (m *ClearRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applyTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplyTo(val)
        return nil
    }
    return res
}
func (m *ClearRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ClearRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applyTo", m.GetApplyTo())
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
func (m *ClearRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ClearRequestBody) SetApplyTo(value *string)() {
    m.applyTo = value
}
