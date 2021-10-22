package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChartPoint struct {
    Entity
    format *WorkbookChartPointFormat;
    value *Json;
}
func NewWorkbookChartPoint()(*WorkbookChartPoint) {
    m := &WorkbookChartPoint{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChartPoint) GetFormat()(*WorkbookChartPointFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *WorkbookChartPoint) GetValue()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *WorkbookChartPoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartPointFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartPointFormat))
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetValue(val.(*Json))
        return nil
    }
    return res
}
func (m *WorkbookChartPoint) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookChartPoint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookChartPoint) SetFormat(value *WorkbookChartPointFormat)() {
    m.format = value
}
func (m *WorkbookChartPoint) SetValue(value *Json)() {
    m.value = value
}
