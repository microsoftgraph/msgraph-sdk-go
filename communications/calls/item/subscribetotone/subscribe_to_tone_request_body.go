package subscribetotone

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SubscribeToToneRequestBody struct {
    additionalData map[string]interface{};
    clientContext *string;
}
func NewSubscribeToToneRequestBody()(*SubscribeToToneRequestBody) {
    m := &SubscribeToToneRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SubscribeToToneRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SubscribeToToneRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
func (m *SubscribeToToneRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clientContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientContext(val)
        return nil
    }
    return res
}
func (m *SubscribeToToneRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *SubscribeToToneRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
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
func (m *SubscribeToToneRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SubscribeToToneRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
