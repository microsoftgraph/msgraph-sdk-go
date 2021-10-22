package merge

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MergeRequestBody struct {
    across *bool;
    additionalData map[string]interface{};
}
func NewMergeRequestBody()(*MergeRequestBody) {
    m := &MergeRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MergeRequestBody) GetAcross()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.across
    }
}
func (m *MergeRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MergeRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["across"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAcross(val)
        return nil
    }
    return res
}
func (m *MergeRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *MergeRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("across", m.GetAcross())
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
func (m *MergeRequestBody) SetAcross(value *bool)() {
    m.across = value
}
func (m *MergeRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
