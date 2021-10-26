package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CalculatedColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // For dateTime output types, the format of the value. Must be one of dateOnly or dateTime.
    format *string;
    // The formula used to compute the value for this column.
    formula *string;
    // The output type used to format values in this column. Must be one of boolean, currency, dateTime, number, or text.
    outputType *string;
}
// Instantiates a new calculatedColumn and sets the default values.
func NewCalculatedColumn()(*CalculatedColumn) {
    m := &CalculatedColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CalculatedColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the format property value. For dateTime output types, the format of the value. Must be one of dateOnly or dateTime.
func (m *CalculatedColumn) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the formula property value. The formula used to compute the value for this column.
func (m *CalculatedColumn) GetFormula()(*string) {
    if m == nil {
        return nil
    } else {
        return m.formula
    }
}
// Gets the outputType property value. The output type used to format values in this column. Must be one of boolean, currency, dateTime, number, or text.
func (m *CalculatedColumn) GetOutputType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputType
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CalculatedColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the format property value. For dateTime output types, the format of the value. Must be one of dateOnly or dateTime.
// Parameters:
//  - value : Value to set for the format property.
func (m *CalculatedColumn) SetFormat(value *string)() {
    m.format = value
}
// Sets the formula property value. The formula used to compute the value for this column.
// Parameters:
//  - value : Value to set for the formula property.
func (m *CalculatedColumn) SetFormula(value *string)() {
    m.formula = value
}
// Sets the outputType property value. The output type used to format values in this column. Must be one of boolean, currency, dateTime, number, or text.
// Parameters:
//  - value : Value to set for the outputType property.
func (m *CalculatedColumn) SetOutputType(value *string)() {
    m.outputType = value
}
