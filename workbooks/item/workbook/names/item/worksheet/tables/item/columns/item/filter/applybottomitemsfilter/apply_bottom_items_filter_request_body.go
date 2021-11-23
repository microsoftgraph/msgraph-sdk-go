package applybottomitemsfilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// applyBottomItemsFilterRequestBody 
type ApplyBottomItemsFilterRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    count *int32;
}
// NewApplyBottomItemsFilterRequestBody instantiates a new applyBottomItemsFilterRequestBody and sets the default values.
func NewApplyBottomItemsFilterRequestBody()(*ApplyBottomItemsFilterRequestBody) {
    m := &ApplyBottomItemsFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyBottomItemsFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCount gets the count property value. 
func (m *ApplyBottomItemsFilterRequestBody) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyBottomItemsFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["count"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    return res
}
func (m *ApplyBottomItemsFilterRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApplyBottomItemsFilterRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
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
func (m *ApplyBottomItemsFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCount sets the count property value. 
func (m *ApplyBottomItemsFilterRequestBody) SetCount(value *int32)() {
    m.count = value
}
