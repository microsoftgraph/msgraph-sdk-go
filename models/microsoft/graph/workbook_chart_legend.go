package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new workbookChartLegend and sets the default values.
func NewWorkbookChartLegend()(*WorkbookChartLegend) {
    m := &WorkbookChartLegend{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the format property value. Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
func (m *WorkbookChartLegend) GetFormat()(*WorkbookChartLegendFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the overlay property value. Boolean value for whether the chart legend should overlap with the main body of the chart.
func (m *WorkbookChartLegend) GetOverlay()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overlay
    }
}
// Gets the position property value. Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
func (m *WorkbookChartLegend) GetPosition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// Gets the visible property value. A boolean value the represents the visibility of a ChartLegend object.
func (m *WorkbookChartLegend) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the format property value. Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
// Parameters:
//  - value : Value to set for the format property.
func (m *WorkbookChartLegend) SetFormat(value *WorkbookChartLegendFormat)() {
    m.format = value
}
// Sets the overlay property value. Boolean value for whether the chart legend should overlap with the main body of the chart.
// Parameters:
//  - value : Value to set for the overlay property.
func (m *WorkbookChartLegend) SetOverlay(value *bool)() {
    m.overlay = value
}
// Sets the position property value. Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
// Parameters:
//  - value : Value to set for the position property.
func (m *WorkbookChartLegend) SetPosition(value *string)() {
    m.position = value
}
// Sets the visible property value. A boolean value the represents the visibility of a ChartLegend object.
// Parameters:
//  - value : Value to set for the visible property.
func (m *WorkbookChartLegend) SetVisible(value *bool)() {
    m.visible = value
}
