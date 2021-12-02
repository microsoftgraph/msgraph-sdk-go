package removekey

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RemoveKeyRequestBody 
type RemoveKeyRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    keyId *string;
    // 
    proof *string;
}
// NewRemoveKeyRequestBody instantiates a new removeKeyRequestBody and sets the default values.
func NewRemoveKeyRequestBody()(*RemoveKeyRequestBody) {
    m := &RemoveKeyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoveKeyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetKeyId gets the keyId property value. 
func (m *RemoveKeyRequestBody) GetKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyId
    }
}
// GetProof gets the proof property value. 
func (m *RemoveKeyRequestBody) GetProof()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proof
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoveKeyRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["proof"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProof(val)
        }
        return nil
    }
    return res
}
func (m *RemoveKeyRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoveKeyRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeyId sets the keyId property value. 
func (m *RemoveKeyRequestBody) SetKeyId(value *string)() {
    if m != nil {
        m.keyId = value
    }
}
// SetProof sets the proof property value. 
func (m *RemoveKeyRequestBody) SetProof(value *string)() {
    if m != nil {
        m.proof = value
    }
}
