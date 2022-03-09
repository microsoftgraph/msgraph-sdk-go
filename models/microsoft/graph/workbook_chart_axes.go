package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxes provides operations to manage the collection of drive entities.
type WorkbookChartAxes struct {
    Entity
    // Represents the category axis in a chart. Read-only.
    categoryAxis WorkbookChartAxisable;
    // Represents the series axis of a 3-dimensional chart. Read-only.
    seriesAxis WorkbookChartAxisable;
    // Represents the value axis in an axis. Read-only.
    valueAxis WorkbookChartAxisable;
}
// NewWorkbookChartAxes instantiates a new workbookChartAxes and sets the default values.
func NewWorkbookChartAxes()(*WorkbookChartAxes) {
    m := &WorkbookChartAxes{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxesFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartAxes(), nil
}
// GetCategoryAxis gets the categoryAxis property value. Represents the category axis in a chart. Read-only.
func (m *WorkbookChartAxes) GetCategoryAxis()(WorkbookChartAxisable) {
    if m == nil {
        return nil
    } else {
        return m.categoryAxis
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryAxis(val.(WorkbookChartAxisable))
        }
        return nil
    }
    res["seriesAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesAxis(val.(WorkbookChartAxisable))
        }
        return nil
    }
    res["valueAxis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueAxis(val.(WorkbookChartAxisable))
        }
        return nil
    }
    return res
}
// GetSeriesAxis gets the seriesAxis property value. Represents the series axis of a 3-dimensional chart. Read-only.
func (m *WorkbookChartAxes) GetSeriesAxis()(WorkbookChartAxisable) {
    if m == nil {
        return nil
    } else {
        return m.seriesAxis
    }
}
// GetValueAxis gets the valueAxis property value. Represents the value axis in an axis. Read-only.
func (m *WorkbookChartAxes) GetValueAxis()(WorkbookChartAxisable) {
    if m == nil {
        return nil
    } else {
        return m.valueAxis
    }
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
func (m *WorkbookChartAxes) SetCategoryAxis(value WorkbookChartAxisable)() {
    if m != nil {
        m.categoryAxis = value
    }
}
// SetSeriesAxis sets the seriesAxis property value. Represents the series axis of a 3-dimensional chart. Read-only.
func (m *WorkbookChartAxes) SetSeriesAxis(value WorkbookChartAxisable)() {
    if m != nil {
        m.seriesAxis = value
    }
}
// SetValueAxis sets the valueAxis property value. Represents the value axis in an axis. Read-only.
func (m *WorkbookChartAxes) SetValueAxis(value WorkbookChartAxisable)() {
    if m != nil {
        m.valueAxis = value
    }
}
