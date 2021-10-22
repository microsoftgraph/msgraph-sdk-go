package removegroup

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RemoveGroupRequestBody struct {
    additionalData map[string]interface{};
    groupId *string;
}
func NewRemoveGroupRequestBody()(*RemoveGroupRequestBody) {
    m := &RemoveGroupRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RemoveGroupRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RemoveGroupRequestBody) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
func (m *RemoveGroupRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["groupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupId(val)
        return nil
    }
    return res
}
func (m *RemoveGroupRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *RemoveGroupRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
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
func (m *RemoveGroupRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RemoveGroupRequestBody) SetGroupId(value *string)() {
    m.groupId = value
}
