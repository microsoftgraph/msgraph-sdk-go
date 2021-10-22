package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChartAxes struct {
    Entity
    categoryAxis *WorkbookChartAxis;
    seriesAxis *WorkbookChartAxis;
    valueAxis *WorkbookChartAxis;
}
func NewWorkbookChartAxes()(*WorkbookChartAxes) {
    m := &WorkbookChartAxes{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChartAxes) GetCategoryAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.categoryAxis
    }
}
func (m *WorkbookChartAxes) GetSeriesAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.seriesAxis
    }
}
func (m *WorkbookChartAxes) GetValueAxis()(*WorkbookChartAxis) {
    if m == nil {
        return nil
    } else {
        return m.valueAxis
    }
}
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
func (m *WorkbookChartAxes) SetCategoryAxis(value *WorkbookChartAxis)() {
    m.categoryAxis = value
}
func (m *WorkbookChartAxes) SetSeriesAxis(value *WorkbookChartAxis)() {
    m.seriesAxis = value
}
func (m *WorkbookChartAxes) SetValueAxis(value *WorkbookChartAxis)() {
    m.valueAxis = value
}
