package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// NumberColumn 
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
// NewNumberColumn instantiates a new numberColumn and sets the default values.
func NewNumberColumn()(*NumberColumn) {
    m := &NumberColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateNumberColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNumberColumnFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewNumberColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NumberColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDecimalPlaces gets the decimalPlaces property value. How many decimal places to display. See below for information about the possible values.
func (m *NumberColumn) GetDecimalPlaces()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decimalPlaces
    }
}
// GetDisplayAs gets the displayAs property value. How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
func (m *NumberColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NumberColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decimalPlaces"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecimalPlaces(val)
        }
        return nil
    }
    res["displayAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAs(val)
        }
        return nil
    }
    res["maximum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximum(val)
        }
        return nil
    }
    res["minimum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimum(val)
        }
        return nil
    }
    return res
}
// GetMaximum gets the maximum property value. The maximum permitted value.
func (m *NumberColumn) GetMaximum()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximum
    }
}
// GetMinimum gets the minimum property value. The minimum permitted value.
func (m *NumberColumn) GetMinimum()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.minimum
    }
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NumberColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDecimalPlaces sets the decimalPlaces property value. How many decimal places to display. See below for information about the possible values.
func (m *NumberColumn) SetDecimalPlaces(value *string)() {
    if m != nil {
        m.decimalPlaces = value
    }
}
// SetDisplayAs sets the displayAs property value. How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
func (m *NumberColumn) SetDisplayAs(value *string)() {
    if m != nil {
        m.displayAs = value
    }
}
// SetMaximum sets the maximum property value. The maximum permitted value.
func (m *NumberColumn) SetMaximum(value *float64)() {
    if m != nil {
        m.maximum = value
    }
}
// SetMinimum sets the minimum property value. The minimum permitted value.
func (m *NumberColumn) SetMinimum(value *float64)() {
    if m != nil {
        m.minimum = value
    }
}
