package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartSeries 
type WorkbookChartSeries struct {
    Entity
    // Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
    format *WorkbookChartSeriesFormat;
    // Represents the name of a series in a chart.
    name *string;
    // Represents a collection of all points in the series. Read-only.
    points []WorkbookChartPoint;
}
// NewWorkbookChartSeries instantiates a new workbookChartSeries and sets the default values.
func NewWorkbookChartSeries()(*WorkbookChartSeries) {
    m := &WorkbookChartSeries{
        Entity: *NewEntity(),
    }
    return m
}
// GetFormat gets the format property value. Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
func (m *WorkbookChartSeries) GetFormat()(*WorkbookChartSeriesFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetName gets the name property value. Represents the name of a series in a chart.
func (m *WorkbookChartSeries) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPoints gets the points property value. Represents a collection of all points in the series. Read-only.
func (m *WorkbookChartSeries) GetPoints()([]WorkbookChartPoint) {
    if m == nil {
        return nil
    } else {
        return m.points
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartSeries) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartSeriesFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*WorkbookChartSeriesFormat))
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["points"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartPoint() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookChartPoint, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookChartPoint))
            }
            m.SetPoints(res)
        }
        return nil
    }
    return res
}
func (m *WorkbookChartSeries) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetFormat sets the format property value. Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
func (m *WorkbookChartSeries) SetFormat(value *WorkbookChartSeriesFormat)() {
    m.format = value
}
// SetName sets the name property value. Represents the name of a series in a chart.
func (m *WorkbookChartSeries) SetName(value *string)() {
    m.name = value
}
// SetPoints sets the points property value. Represents a collection of all points in the series. Read-only.
func (m *WorkbookChartSeries) SetPoints(value []WorkbookChartPoint)() {
    m.points = value
}
