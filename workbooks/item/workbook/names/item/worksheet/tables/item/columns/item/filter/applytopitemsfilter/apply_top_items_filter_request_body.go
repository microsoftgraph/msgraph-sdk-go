package applytopitemsfilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApplyTopItemsFilterRequestBody struct {
    additionalData map[string]interface{};
    count *int32;
}
func NewApplyTopItemsFilterRequestBody()(*ApplyTopItemsFilterRequestBody) {
    m := &ApplyTopItemsFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApplyTopItemsFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApplyTopItemsFilterRequestBody) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
func (m *ApplyTopItemsFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["count"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCount(val)
        return nil
    }
    return res
}
func (m *ApplyTopItemsFilterRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ApplyTopItemsFilterRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *ApplyTopItemsFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApplyTopItemsFilterRequestBody) SetCount(value *int32)() {
    m.count = value
}
