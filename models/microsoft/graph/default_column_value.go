package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DefaultColumnValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The formula used to compute the default value for this column.
    formula *string;
    // The direct value to use as the default value for this column.
    value *string;
}
// Instantiates a new defaultColumnValue and sets the default values.
func NewDefaultColumnValue()(*DefaultColumnValue) {
    m := &DefaultColumnValue{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultColumnValue) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the formula property value. The formula used to compute the default value for this column.
func (m *DefaultColumnValue) GetFormula()(*string) {
    if m == nil {
        return nil
    } else {
        return m.formula
    }
}
// Gets the value property value. The direct value to use as the default value for this column.
func (m *DefaultColumnValue) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *DefaultColumnValue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["formula"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFormula(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *DefaultColumnValue) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DefaultColumnValue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *DefaultColumnValue) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the formula property value. The formula used to compute the default value for this column.
// Parameters:
//  - value : Value to set for the formula property.
func (m *DefaultColumnValue) SetFormula(value *string)() {
    m.formula = value
}
// Sets the value property value. The direct value to use as the default value for this column.
// Parameters:
//  - value : Value to set for the value property.
func (m *DefaultColumnValue) SetValue(value *string)() {
    m.value = value
}
