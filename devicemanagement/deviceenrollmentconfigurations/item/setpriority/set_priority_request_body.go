package setpriority

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SetPriorityRequestBody struct {
    additionalData map[string]interface{};
    priority *int32;
}
func NewSetPriorityRequestBody()(*SetPriorityRequestBody) {
    m := &SetPriorityRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SetPriorityRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SetPriorityRequestBody) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
func (m *SetPriorityRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPriority(val)
        return nil
    }
    return res
}
func (m *SetPriorityRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *SetPriorityRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
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
func (m *SetPriorityRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SetPriorityRequestBody) SetPriority(value *int32)() {
    m.priority = value
}
