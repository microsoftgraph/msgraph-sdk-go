package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxisFormat 
type WorkbookChartAxisFormat struct {
    Entity
    // Represents the font attributes (font name, font size, color, etc.) for a chart axis element. Read-only.
    font *WorkbookChartFont;
    // Represents chart line formatting. Read-only.
    line *WorkbookChartLineFormat;
}
// NewWorkbookChartAxisFormat instantiates a new workbookChartAxisFormat and sets the default values.
func NewWorkbookChartAxisFormat()(*WorkbookChartAxisFormat) {
    m := &WorkbookChartAxisFormat{
        Entity: *NewEntity(),
    }
    return m
}
// GetFont gets the font property value. Represents the font attributes (font name, font size, color, etc.) for a chart axis element. Read-only.
func (m *WorkbookChartAxisFormat) GetFont()(*WorkbookChartFont) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
// GetLine gets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartAxisFormat) GetLine()(*WorkbookChartLineFormat) {
    if m == nil {
        return nil
    } else {
        return m.line
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxisFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["font"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartFont() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFont(val.(*WorkbookChartFont))
        }
        return nil
    }
    res["line"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartLineFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val.(*WorkbookChartLineFormat))
        }
        return nil
    }
    return res
}
func (m *WorkbookChartAxisFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartAxisFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("line", m.GetLine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFont sets the font property value. Represents the font attributes (font name, font size, color, etc.) for a chart axis element. Read-only.
func (m *WorkbookChartAxisFormat) SetFont(value *WorkbookChartFont)() {
    m.font = value
}
// SetLine sets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartAxisFormat) SetLine(value *WorkbookChartLineFormat)() {
    m.line = value
}
