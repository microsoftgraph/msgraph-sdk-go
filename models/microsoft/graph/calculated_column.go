package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CalculatedColumn struct {
    additionalData map[string]interface{};
    format *string;
    formula *string;
    outputType *string;
}
func NewCalculatedColumn()(*CalculatedColumn) {
    m := &CalculatedColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CalculatedColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CalculatedColumn) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *CalculatedColumn) GetFormula()(*string) {
    if m == nil {
        return nil
    } else {
        return m.formula
    }
}
func (m *CalculatedColumn) GetOutputType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputType
    }
}
func (m *CalculatedColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFormat(val)
        return nil
    }
    res["formula"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFormula(val)
        return nil
    }
    res["outputType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOutputType(val)
        return nil
    }
    return res
}
func (m *CalculatedColumn) IsNil()(bool) {
    return m == nil
}
func (m *CalculatedColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputType", m.GetOutputType())
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
func (m *CalculatedColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CalculatedColumn) SetFormat(value *string)() {
    m.format = value
}
func (m *CalculatedColumn) SetFormula(value *string)() {
    m.formula = value
}
func (m *CalculatedColumn) SetOutputType(value *string)() {
    m.outputType = value
}
