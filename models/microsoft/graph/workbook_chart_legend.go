package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChartLegend struct {
    Entity
    format *WorkbookChartLegendFormat;
    overlay *bool;
    position *string;
    visible *bool;
}
func NewWorkbookChartLegend()(*WorkbookChartLegend) {
    m := &WorkbookChartLegend{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChartLegend) GetFormat()(*WorkbookChartLegendFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *WorkbookChartLegend) GetOverlay()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overlay
    }
}
func (m *WorkbookChartLegend) GetPosition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
func (m *WorkbookChartLegend) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
func (m *WorkbookChartLegend) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartLegendFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartLegendFormat))
        return nil
    }
    res["overlay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOverlay(val)
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPosition(val)
        return nil
    }
    res["visible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetVisible(val)
        return nil
    }
    return res
}
func (m *WorkbookChartLegend) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookChartLegend) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("overlay", m.GetOverlay())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("visible", m.GetVisible())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookChartLegend) SetFormat(value *WorkbookChartLegendFormat)() {
    m.format = value
}
func (m *WorkbookChartLegend) SetOverlay(value *bool)() {
    m.overlay = value
}
func (m *WorkbookChartLegend) SetPosition(value *string)() {
    m.position = value
}
func (m *WorkbookChartLegend) SetVisible(value *bool)() {
    m.visible = value
}
