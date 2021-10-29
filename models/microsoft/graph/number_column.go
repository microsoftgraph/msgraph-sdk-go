package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type NumberColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // How many decimal places to display. See below for information about the possible values.
    decimalPlaces *string;
    // How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
    displayAs *string;
    // The maximum permitted value.
    maximum *float64;
    // The minimum permitted value.
    minimum *float64;
}
// Instantiates a new numberColumn and sets the default values.
func NewNumberColumn()(*NumberColumn) {
    m := &NumberColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NumberColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the decimalPlaces property value. How many decimal places to display. See below for information about the possible values.
func (m *NumberColumn) GetDecimalPlaces()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decimalPlaces
    }
}
// Gets the displayAs property value. How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
func (m *NumberColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
// Gets the maximum property value. The maximum permitted value.
func (m *NumberColumn) GetMaximum()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximum
    }
}
// Gets the minimum property value. The minimum permitted value.
func (m *NumberColumn) GetMinimum()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.minimum
    }
}
// The deserialization information for the current model
func (m *NumberColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decimalPlaces"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDecimalPlaces(val)
        return nil
    }
    res["displayAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayAs(val)
        return nil
    }
    res["maximum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetMaximum(val)
        return nil
    }
    res["minimum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetMinimum(val)
        return nil
    }
    return res
}
func (m *NumberColumn) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *NumberColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("decimalPlaces", m.GetDecimalPlaces())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("maximum", m.GetMaximum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("minimum", m.GetMinimum())
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
func (m *NumberColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the decimalPlaces property value. How many decimal places to display. See below for information about the possible values.
// Parameters:
//  - value : Value to set for the decimalPlaces property.
func (m *NumberColumn) SetDecimalPlaces(value *string)() {
    m.decimalPlaces = value
}
// Sets the displayAs property value. How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
// Parameters:
//  - value : Value to set for the displayAs property.
func (m *NumberColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
// Sets the maximum property value. The maximum permitted value.
// Parameters:
//  - value : Value to set for the maximum property.
func (m *NumberColumn) SetMaximum(value *float64)() {
    m.maximum = value
}
// Sets the minimum property value. The minimum permitted value.
// Parameters:
//  - value : Value to set for the minimum property.
func (m *NumberColumn) SetMinimum(value *float64)() {
    m.minimum = value
}
