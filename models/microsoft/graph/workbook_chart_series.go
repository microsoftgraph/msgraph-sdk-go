package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartSeries 
type WorkbookChartSeries struct {
    Entity
    // Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
    format WorkbookChartSeriesFormatable;
    // Represents the name of a series in a chart.
    name *string;
    // Represents a collection of all points in the series. Read-only.
    points []WorkbookChartPointable;
}
// NewWorkbookChartSeries instantiates a new workbookChartSeries and sets the default values.
func NewWorkbookChartSeries()(*WorkbookChartSeries) {
    m := &WorkbookChartSeries{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartSeriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartSeriesFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartSeries(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartSeries) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartSeriesFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartSeriesFormatable))
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
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookChartPointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookChartPointable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookChartPointable)
            }
            m.SetPoints(res)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
func (m *WorkbookChartSeries) GetFormat()(WorkbookChartSeriesFormatable) {
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
func (m *WorkbookChartSeries) GetPoints()([]WorkbookChartPointable) {
    if m == nil {
        return nil
    } else {
        return m.points
    }
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
    if m.GetPoints() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPoints()))
        for i, v := range m.GetPoints() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("points", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Represents the formatting of a chart series, which includes fill and line formatting. Read-only.
func (m *WorkbookChartSeries) SetFormat(value WorkbookChartSeriesFormatable)() {
    if m != nil {
        m.format = value
    }
}
// SetName sets the name property value. Represents the name of a series in a chart.
func (m *WorkbookChartSeries) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPoints sets the points property value. Represents a collection of all points in the series. Read-only.
func (m *WorkbookChartSeries) SetPoints(value []WorkbookChartPointable)() {
    if m != nil {
        m.points = value
    }
}
