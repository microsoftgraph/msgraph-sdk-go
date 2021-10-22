package forcedelete

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ForceDeleteRequestBody struct {
    additionalData map[string]interface{};
    disableUserAccounts *bool;
}
func NewForceDeleteRequestBody()(*ForceDeleteRequestBody) {
    m := &ForceDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ForceDeleteRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ForceDeleteRequestBody) GetDisableUserAccounts()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableUserAccounts
    }
}
func (m *ForceDeleteRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["disableUserAccounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDisableUserAccounts(val)
        return nil
    }
    return res
}
func (m *ForceDeleteRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ForceDeleteRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("disableUserAccounts", m.GetDisableUserAccounts())
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
func (m *ForceDeleteRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ForceDeleteRequestBody) SetDisableUserAccounts(value *bool)() {
    m.disableUserAccounts = value
}
