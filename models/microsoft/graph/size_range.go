package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SizeRange struct {
    additionalData map[string]interface{};
    maximumSize *int32;
    minimumSize *int32;
}
func NewSizeRange()(*SizeRange) {
    m := &SizeRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SizeRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SizeRange) GetMaximumSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumSize
    }
}
func (m *SizeRange) GetMinimumSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumSize
    }
}
func (m *SizeRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maximumSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaximumSize(val)
        return nil
    }
    res["minimumSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMinimumSize(val)
        return nil
    }
    return res
}
func (m *SizeRange) IsNil()(bool) {
    return m == nil
}
func (m *SizeRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maximumSize", m.GetMaximumSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minimumSize", m.GetMinimumSize())
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
func (m *SizeRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SizeRange) SetMaximumSize(value *int32)() {
    m.maximumSize = value
}
func (m *SizeRange) SetMinimumSize(value *int32)() {
    m.minimumSize = value
}
