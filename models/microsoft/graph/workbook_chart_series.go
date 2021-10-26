package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartSeries struct {
    Entity
    // Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
    format *WorkbookChartSeriesFormat;
    // Represents the name of a series in a chart.
    name *string;
    // Represents a collection of all points in the series. Read-only.
    points []WorkbookChartPoint;
}
// Instantiates a new workbookChartSeries and sets the default values.
func NewWorkbookChartSeries()(*WorkbookChartSeries) {
    m := &WorkbookChartSeries{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the format property value. Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
func (m *WorkbookChartSeries) GetFormat()(*WorkbookChartSeriesFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the name property value. Represents the name of a series in a chart.
func (m *WorkbookChartSeries) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the points property value. Represents a collection of all points in the series. Read-only.
func (m *WorkbookChartSeries) GetPoints()([]WorkbookChartPoint) {
    if m == nil {
        return nil
    } else {
        return m.points
    }
}
// The deserialization information for the current model
func (m *WorkbookChartSeries) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartSeriesFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartSeriesFormat))
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["points"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartPoint() })
        if err != nil {
            return err
        }
        res := make([]WorkbookChartPoint, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookChartPoint))
        }
        m.SetPoints(res)
        return nil
    }
    return res
}
func (m *WorkbookChartSeries) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartSeries) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPoints()))
        for i, v := range m.GetPoints() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("points", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the format property value. Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
// Parameters:
//  - value : Value to set for the format property.
func (m *WorkbookChartSeries) SetFormat(value *WorkbookChartSeriesFormat)() {
    m.format = value
}
// Sets the name property value. Represents the name of a series in a chart.
// Parameters:
//  - value : Value to set for the name property.
func (m *WorkbookChartSeries) SetName(value *string)() {
    m.name = value
}
// Sets the points property value. Represents a collection of all points in the series. Read-only.
// Parameters:
//  - value : Value to set for the points property.
func (m *WorkbookChartSeries) SetPoints(value []WorkbookChartPoint)() {
    m.points = value
}
