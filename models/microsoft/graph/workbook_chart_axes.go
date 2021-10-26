package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartAxes struct {
    Entity
    // Represents the category axis in a chart. Read-only.
    categoryAxis *WorkbookChartAxis;
    // Represents the series axis of a 3-dimensional chart. Read-only.
    seriesAxis *WorkbookChartAxis;
    // Represents the value axis in an axis. Read-only.
    valueAxis *WorkbookChartAxis;
}
// Instantiates a new workbookChartAxes and sets the default values.
func NewWorkbookChartAxes()(*WorkbookChartAxes) {
    m := &WorkbookChartAxes{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the categoryAxis property value. Represents the category axis in a chart. Read-only.
func (m *WorkbookChartAxes) GetCategoryAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.categoryAxis
    }
}
// Gets the seriesAxis property value. Represents the series axis of a 3-dimensional chart. Read-only.
func (m *WorkbookChartAxes) GetSeriesAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.seriesAxis
    }
}
// Gets the valueAxis property value. Represents the value axis in an axis. Read-only.
func (m *WorkbookChartAxes) GetValueAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.valueAxis
    }
}
// The deserialization information for the current model
func (m *WorkbookChartAxes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxis() })
        if err != nil {
            return err
        }
        m.SetCategoryAxis(val.(*WorkbookChartAxis))
        return nil
    }
    res["seriesAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxis() })
        if err != nil {
            return err
        }
        m.SetSeriesAxis(val.(*WorkbookChartAxis))
        return nil
    }
    res["valueAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxis() })
        if err != nil {
            return err
        }
        m.SetValueAxis(val.(*WorkbookChartAxis))
        return nil
    }
    return res
}
func (m *WorkbookChartAxes) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartAxes) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("categoryAxis", m.GetCategoryAxis())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("seriesAxis", m.GetSeriesAxis())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("valueAxis", m.GetValueAxis())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the categoryAxis property value. Represents the category axis in a chart. Read-only.
// Parameters:
//  - value : Value to set for the categoryAxis property.
func (m *WorkbookChartAxes) SetCategoryAxis(value *WorkbookChartAxis)() {
    m.categoryAxis = value
}
// Sets the seriesAxis property value. Represents the series axis of a 3-dimensional chart. Read-only.
// Parameters:
//  - value : Value to set for the seriesAxis property.
func (m *WorkbookChartAxes) SetSeriesAxis(value *WorkbookChartAxis)() {
    m.seriesAxis = value
}
// Sets the valueAxis property value. Represents the value axis in an axis. Read-only.
// Parameters:
//  - value : Value to set for the valueAxis property.
func (m *WorkbookChartAxes) SetValueAxis(value *WorkbookChartAxis)() {
    m.valueAxis = value
}
