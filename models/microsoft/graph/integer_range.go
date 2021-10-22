package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IntegerRange struct {
    additionalData map[string]interface{};
    end *int64;
    start *int64;
}
func NewIntegerRange()(*IntegerRange) {
    m := &IntegerRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *IntegerRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *IntegerRange) GetEnd()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
func (m *IntegerRange) GetStart()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
func (m *IntegerRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetEnd(val)
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetStart(val)
        return nil
    }
    return res
}
func (m *IntegerRange) IsNil()(bool) {
    return m == nil
}
func (m *IntegerRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("start", m.GetStart())
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
func (m *IntegerRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *IntegerRange) SetEnd(value *int64)() {
    m.end = value
}
func (m *IntegerRange) SetStart(value *int64)() {
    m.start = value
}
