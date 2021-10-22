package cancelmediaprocessing

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CancelMediaProcessingRequestBody struct {
    additionalData map[string]interface{};
    clientContext *string;
}
func NewCancelMediaProcessingRequestBody()(*CancelMediaProcessingRequestBody) {
    m := &CancelMediaProcessingRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CancelMediaProcessingRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CancelMediaProcessingRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
func (m *CancelMediaProcessingRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *CancelMediaProcessingRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *CancelMediaProcessingRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *CancelMediaProcessingRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CancelMediaProcessingRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
