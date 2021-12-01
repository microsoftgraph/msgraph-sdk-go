package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxes 
type WorkbookChartAxes struct {
    Entity
    // Represents the category axis in a chart. Read-only.
    categoryAxis *WorkbookChartAxis;
    // Represents the series axis of a 3-dimensional chart. Read-only.
    seriesAxis *WorkbookChartAxis;
    // Represents the value axis in an axis. Read-only.
    valueAxis *WorkbookChartAxis;
}
// NewWorkbookChartAxes instantiates a new workbookChartAxes and sets the default values.
func NewWorkbookChartAxes()(*WorkbookChartAxes) {
    m := &WorkbookChartAxes{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategoryAxis gets the categoryAxis property value. Represents the category axis in a chart. Read-only.
func (m *WorkbookChartAxes) GetCategoryAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.categoryAxis
    }
}
// GetSeriesAxis gets the seriesAxis property value. Represents the series axis of a 3-dimensional chart. Read-only.
func (m *WorkbookChartAxes) GetSeriesAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.seriesAxis
    }
}
// GetValueAxis gets the valueAxis property value. Represents the value axis in an axis. Read-only.
func (m *WorkbookChartAxes) GetValueAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.valueAxis
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxis() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryAxis(val.(*WorkbookChartAxis))
        }
        return nil
    }
    res["seriesAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxis() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesAxis(val.(*WorkbookChartAxis))
        }
        return nil
    }
    res["valueAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxis() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueAxis(val.(*WorkbookChartAxis))
        }
        return nil
    }
    return res
}
func (m *WorkbookChartAxes) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCategoryAxis sets the categoryAxis property value. Represents the category axis in a chart. Read-only.
func (m *WorkbookChartAxes) SetCategoryAxis(value *WorkbookChartAxis)() {
    if m != nil {
        m.categoryAxis = value
    }
}
// SetSeriesAxis sets the seriesAxis property value. Represents the series axis of a 3-dimensional chart. Read-only.
func (m *WorkbookChartAxes) SetSeriesAxis(value *WorkbookChartAxis)() {
    if m != nil {
        m.seriesAxis = value
    }
}
// SetValueAxis sets the valueAxis property value. Represents the value axis in an axis. Read-only.
func (m *WorkbookChartAxes) SetValueAxis(value *WorkbookChartAxis)() {
    if m != nil {
        m.valueAxis = value
    }
}
