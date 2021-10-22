package delete

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeleteRequestBody struct {
    additionalData map[string]interface{};
    shift *string;
}
func NewDeleteRequestBody()(*DeleteRequestBody) {
    m := &DeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeleteRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeleteRequestBody) GetShift()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shift
    }
}
func (m *DeleteRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["shift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShift(val)
        return nil
    }
    return res
}
func (m *DeleteRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *DeleteRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("shift", m.GetShift())
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
func (m *DeleteRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeleteRequestBody) SetShift(value *string)() {
    m.shift = value
}
