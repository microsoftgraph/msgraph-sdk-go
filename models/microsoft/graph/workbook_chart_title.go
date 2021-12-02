package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartTitle 
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
// NewWorkbookChartTitle instantiates a new workbookChartTitle and sets the default values.
func NewWorkbookChartTitle()(*WorkbookChartTitle) {
    m := &WorkbookChartTitle{
        Entity: *NewEntity(),
    }
    return m
}
// GetFormat gets the format property value. Represents the formatting of a chart title, which includes fill and font formatting. Read-only.
func (m *WorkbookChartTitle) GetFormat()(*WorkbookChartTitleFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetOverlay gets the overlay property value. Boolean value representing if the chart title will overlay the chart or not.
func (m *WorkbookChartTitle) GetOverlay()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overlay
    }
}
// GetText gets the text property value. Represents the title text of a chart.
func (m *WorkbookChartTitle) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetVisible gets the visible property value. A boolean value the represents the visibility of a chart title object.
func (m *WorkbookChartTitle) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartTitle) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartTitleFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*WorkbookChartTitleFormat))
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
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
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
func (m *WorkbookChartTitle) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetFormat sets the format property value. Represents the formatting of a chart title, which includes fill and font formatting. Read-only.
func (m *WorkbookChartTitle) SetFormat(value *WorkbookChartTitleFormat)() {
    if m != nil {
        m.format = value
    }
}
// SetOverlay sets the overlay property value. Boolean value representing if the chart title will overlay the chart or not.
func (m *WorkbookChartTitle) SetOverlay(value *bool)() {
    if m != nil {
        m.overlay = value
    }
}
// SetText sets the text property value. Represents the title text of a chart.
func (m *WorkbookChartTitle) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
// SetVisible sets the visible property value. A boolean value the represents the visibility of a chart title object.
func (m *WorkbookChartTitle) SetVisible(value *bool)() {
    if m != nil {
        m.visible = value
    }
}
