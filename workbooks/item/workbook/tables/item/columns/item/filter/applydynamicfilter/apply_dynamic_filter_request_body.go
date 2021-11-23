package applydynamicfilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplyDynamicFilterRequestBody 
type ApplyDynamicFilterRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    criteria *string;
}
// NewApplyDynamicFilterRequestBody instantiates a new applyDynamicFilterRequestBody and sets the default values.
func NewApplyDynamicFilterRequestBody()(*ApplyDynamicFilterRequestBody) {
    m := &ApplyDynamicFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyDynamicFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCriteria gets the criteria property value. 
func (m *ApplyDynamicFilterRequestBody) GetCriteria()(*string) {
    if m == nil {
        return nil
    } else {
        return m.criteria
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyDynamicFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["criteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val)
        }
        return nil
    }
    return res
}
func (m *ApplyDynamicFilterRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApplyDynamicFilterRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("criteria", m.GetCriteria())
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
func (m *ApplyDynamicFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCriteria sets the criteria property value. 
func (m *ApplyDynamicFilterRequestBody) SetCriteria(value *string)() {
    m.criteria = value
}
