package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartAxis struct {
    Entity
    // Represents the formatting of a chart object, which includes line and font formatting. Read-only.
    format *WorkbookChartAxisFormat;
    // Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
    majorGridlines *WorkbookChartGridlines;
    // Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number.
    majorUnit *Json;
    // Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
    maximum *Json;
    // Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
    minimum *Json;
    // Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
    minorGridlines *WorkbookChartGridlines;
    // Represents the interval between two minor tick marks. 'Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number.
    minorUnit *Json;
    // Represents the axis title. Read-only.
    title *WorkbookChartAxisTitle;
}
// Instantiates a new workbookChartAxis and sets the default values.
func NewWorkbookChartAxis()(*WorkbookChartAxis) {
    m := &WorkbookChartAxis{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the format property value. Represents the formatting of a chart object, which includes line and font formatting. Read-only.
func (m *WorkbookChartAxis) GetFormat()(*WorkbookChartAxisFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the majorGridlines property value. Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) GetMajorGridlines()(*WorkbookChartGridlines) {
    if m == nil {
        return nil
    } else {
        return m.majorGridlines
    }
}
// Gets the majorUnit property value. Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number.
func (m *WorkbookChartAxis) GetMajorUnit()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.majorUnit
    }
}
// Gets the maximum property value. Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
func (m *WorkbookChartAxis) GetMaximum()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.maximum
    }
}
// Gets the minimum property value. Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
func (m *WorkbookChartAxis) GetMinimum()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.minimum
    }
}
// Gets the minorGridlines property value. Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) GetMinorGridlines()(*WorkbookChartGridlines) {
    if m == nil {
        return nil
    } else {
        return m.minorGridlines
    }
}
// Gets the minorUnit property value. Represents the interval between two minor tick marks. 'Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number.
func (m *WorkbookChartAxis) GetMinorUnit()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.minorUnit
    }
}
// Gets the title property value. Represents the axis title. Read-only.
func (m *WorkbookChartAxis) GetTitle()(*WorkbookChartAxisTitle) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// The deserialization information for the current model
func (m *WorkbookChartAxis) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxisFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*WorkbookChartAxisFormat))
        }
        return nil
    }
    res["majorGridlines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartGridlines() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorGridlines(val.(*WorkbookChartGridlines))
        }
        return nil
    }
    res["majorUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorUnit(val.(*Json))
        }
        return nil
    }
    res["maximum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximum(val.(*Json))
        }
        return nil
    }
    res["minimum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimum(val.(*Json))
        }
        return nil
    }
    res["minorGridlines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartGridlines() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorGridlines(val.(*WorkbookChartGridlines))
        }
        return nil
    }
    res["minorUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorUnit(val.(*Json))
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxisTitle() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(*WorkbookChartAxisTitle))
        }
        return nil
    }
    return res
}
func (m *WorkbookChartAxis) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartAxis) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("majorGridlines", m.GetMajorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("majorUnit", m.GetMajorUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("maximum", m.GetMaximum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimum", m.GetMinimum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorGridlines", m.GetMinorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorUnit", m.GetMinorUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the format property value. Represents the formatting of a chart object, which includes line and font formatting. Read-only.
// Parameters:
//  - value : Value to set for the format property.
func (m *WorkbookChartAxis) SetFormat(value *WorkbookChartAxisFormat)() {
    m.format = value
}
// Sets the majorGridlines property value. Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
// Parameters:
//  - value : Value to set for the majorGridlines property.
func (m *WorkbookChartAxis) SetMajorGridlines(value *WorkbookChartGridlines)() {
    m.majorGridlines = value
}
// Sets the majorUnit property value. Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number.
// Parameters:
//  - value : Value to set for the majorUnit property.
func (m *WorkbookChartAxis) SetMajorUnit(value *Json)() {
    m.majorUnit = value
}
// Sets the maximum property value. Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
// Parameters:
//  - value : Value to set for the maximum property.
func (m *WorkbookChartAxis) SetMaximum(value *Json)() {
    m.maximum = value
}
// Sets the minimum property value. Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number.
// Parameters:
//  - value : Value to set for the minimum property.
func (m *WorkbookChartAxis) SetMinimum(value *Json)() {
    m.minimum = value
}
// Sets the minorGridlines property value. Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
// Parameters:
//  - value : Value to set for the minorGridlines property.
func (m *WorkbookChartAxis) SetMinorGridlines(value *WorkbookChartGridlines)() {
    m.minorGridlines = value
}
// Sets the minorUnit property value. Represents the interval between two minor tick marks. 'Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number.
// Parameters:
//  - value : Value to set for the minorUnit property.
func (m *WorkbookChartAxis) SetMinorUnit(value *Json)() {
    m.minorUnit = value
}
// Sets the title property value. Represents the axis title. Read-only.
// Parameters:
//  - value : Value to set for the title property.
func (m *WorkbookChartAxis) SetTitle(value *WorkbookChartAxisTitle)() {
    m.title = value
}
