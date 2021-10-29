package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartAxisTitleFormat struct {
    Entity
    // Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
    font *WorkbookChartFont;
}
// Instantiates a new workbookChartAxisTitleFormat and sets the default values.
func NewWorkbookChartAxisTitleFormat()(*WorkbookChartAxisTitleFormat) {
    m := &WorkbookChartAxisTitleFormat{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the font property value. Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
func (m *WorkbookChartAxisTitleFormat) GetFont()(*WorkbookChartFont) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
// The deserialization information for the current model
func (m *WorkbookChartAxisTitleFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["font"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartFont() })
        if err != nil {
            return err
        }
        m.SetFont(val.(*WorkbookChartFont))
        return nil
    }
    return res
}
func (m *WorkbookChartAxisTitleFormat) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartAxisTitleFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
// Sets the font property value. Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
// Parameters:
//  - value : Value to set for the font property.
func (m *WorkbookChartAxisTitleFormat) SetFont(value *WorkbookChartFont)() {
    m.font = value
}
