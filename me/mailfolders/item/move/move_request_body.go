package move

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MoveRequestBody struct {
    additionalData map[string]interface{};
    destinationId *string;
}
func NewMoveRequestBody()(*MoveRequestBody) {
    m := &MoveRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MoveRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MoveRequestBody) GetDestinationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationId
    }
}
func (m *MoveRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["destinationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationId(val)
        return nil
    }
    return res
}
func (m *MoveRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *MoveRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationId", m.GetDestinationId())
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
func (m *MoveRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MoveRequestBody) SetDestinationId(value *string)() {
    m.destinationId = value
}
