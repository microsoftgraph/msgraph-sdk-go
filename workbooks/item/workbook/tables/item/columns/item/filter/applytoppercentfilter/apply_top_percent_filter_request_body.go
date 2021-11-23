package applytoppercentfilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplyTopPercentFilterRequestBody 
type ApplyTopPercentFilterRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    percent *int32;
}
// NewApplyTopPercentFilterRequestBody instantiates a new applyTopPercentFilterRequestBody and sets the default values.
func NewApplyTopPercentFilterRequestBody()(*ApplyTopPercentFilterRequestBody) {
    m := &ApplyTopPercentFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyTopPercentFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPercent gets the percent property value. 
func (m *ApplyTopPercentFilterRequestBody) GetPercent()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.percent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyTopPercentFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["percent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercent(val)
        }
        return nil
    }
    return res
}
func (m *ApplyTopPercentFilterRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApplyTopPercentFilterRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("percent", m.GetPercent())
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
func (m *ApplyTopPercentFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPercent sets the percent property value. 
func (m *ApplyTopPercentFilterRequestBody) SetPercent(value *int32)() {
    m.percent = value
}
