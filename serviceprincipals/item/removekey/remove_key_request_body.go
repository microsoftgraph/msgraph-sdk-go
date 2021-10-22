package removekey

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RemoveKeyRequestBody struct {
    additionalData map[string]interface{};
    keyId *string;
    proof *string;
}
func NewRemoveKeyRequestBody()(*RemoveKeyRequestBody) {
    m := &RemoveKeyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RemoveKeyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RemoveKeyRequestBody) GetKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyId
    }
}
func (m *RemoveKeyRequestBody) GetProof()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proof
    }
}
func (m *RemoveKeyRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKeyId(val)
        return nil
    }
    res["proof"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProof(val)
        return nil
    }
    return res
}
func (m *RemoveKeyRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *RemoveKeyRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("keyId", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proof", m.GetProof())
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
func (m *RemoveKeyRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RemoveKeyRequestBody) SetKeyId(value *string)() {
    m.keyId = value
}
func (m *RemoveKeyRequestBody) SetProof(value *string)() {
    m.proof = value
}
