package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartLegend 
type WorkbookChartLegend struct {
    Entity
    // Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
    format *WorkbookChartLegendFormat;
    // Boolean value for whether the chart legend should overlap with the main body of the chart.
    overlay *bool;
    // Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
    position *string;
    // A boolean value the represents the visibility of a ChartLegend object.
    visible *bool;
}
// NewWorkbookChartLegend instantiates a new workbookChartLegend and sets the default values.
func NewWorkbookChartLegend()(*WorkbookChartLegend) {
    m := &WorkbookChartLegend{
        Entity: *NewEntity(),
    }
    return m
}
// GetFormat gets the format property value. Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
func (m *WorkbookChartLegend) GetFormat()(*WorkbookChartLegendFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetOverlay gets the overlay property value. Boolean value for whether the chart legend should overlap with the main body of the chart.
func (m *WorkbookChartLegend) GetOverlay()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overlay
    }
}
// GetPosition gets the position property value. Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
func (m *WorkbookChartLegend) GetPosition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// GetVisible gets the visible property value. A boolean value the represents the visibility of a ChartLegend object.
func (m *WorkbookChartLegend) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartLegend) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartLegendFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*WorkbookChartLegendFormat))
        }
        return nil
    }
    res["overlay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverlay(val)
        }
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    res["visible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisible(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookChartLegend) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetFormat sets the format property value. Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
func (m *WorkbookChartLegend) SetFormat(value *WorkbookChartLegendFormat)() {
    m.format = value
}
// SetOverlay sets the overlay property value. Boolean value for whether the chart legend should overlap with the main body of the chart.
func (m *WorkbookChartLegend) SetOverlay(value *bool)() {
    m.overlay = value
}
// SetPosition sets the position property value. Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
func (m *WorkbookChartLegend) SetPosition(value *string)() {
    m.position = value
}
// SetVisible sets the visible property value. A boolean value the represents the visibility of a ChartLegend object.
func (m *WorkbookChartLegend) SetVisible(value *bool)() {
    m.visible = value
}
