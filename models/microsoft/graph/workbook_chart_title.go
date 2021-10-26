package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartTitle struct {
    Entity
    // Represents the formatting of a chart title, which includes fill and font formatting. Read-only.
    format *WorkbookChartTitleFormat;
    // Boolean value representing if the chart title will overlay the chart or not.
    overlay *bool;
    // Represents the title text of a chart.
    text *string;
    // A boolean value the represents the visibility of a chart title object.
    visible *bool;
}
// Instantiates a new workbookChartTitle and sets the default values.
func NewWorkbookChartTitle()(*WorkbookChartTitle) {
    m := &WorkbookChartTitle{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the format property value. Represents the formatting of a chart title, which includes fill and font formatting. Read-only.
func (m *WorkbookChartTitle) GetFormat()(*WorkbookChartTitleFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the overlay property value. Boolean value representing if the chart title will overlay the chart or not.
func (m *WorkbookChartTitle) GetOverlay()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overlay
    }
}
// Gets the text property value. Represents the title text of a chart.
func (m *WorkbookChartTitle) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// Gets the visible property value. A boolean value the represents the visibility of a chart title object.
func (m *WorkbookChartTitle) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
// The deserialization information for the current model
func (m *WorkbookChartTitle) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartTitleFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartTitleFormat))
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
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetText(val)
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
func (m *WorkbookChartTitle) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartTitle) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("text", m.GetText())
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
// Sets the format property value. Represents the formatting of a chart title, which includes fill and font formatting. Read-only.
// Parameters:
//  - value : Value to set for the format property.
func (m *WorkbookChartTitle) SetFormat(value *WorkbookChartTitleFormat)() {
    m.format = value
}
// Sets the overlay property value. Boolean value representing if the chart title will overlay the chart or not.
// Parameters:
//  - value : Value to set for the overlay property.
func (m *WorkbookChartTitle) SetOverlay(value *bool)() {
    m.overlay = value
}
// Sets the text property value. Represents the title text of a chart.
// Parameters:
//  - value : Value to set for the text property.
func (m *WorkbookChartTitle) SetText(value *string)() {
    m.text = value
}
// Sets the visible property value. A boolean value the represents the visibility of a chart title object.
// Parameters:
//  - value : Value to set for the visible property.
func (m *WorkbookChartTitle) SetVisible(value *bool)() {
    m.visible = value
}
