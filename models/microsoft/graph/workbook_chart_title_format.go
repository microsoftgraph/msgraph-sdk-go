package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartTitleFormat struct {
    Entity
    // Represents the fill format of an object, which includes background formatting information. Read-only.
    fill *WorkbookChartFill;
    // Represents the font attributes (font name, font size, color, etc.) for the current object. Read-only.
    font *WorkbookChartFont;
}
// Instantiates a new workbookChartTitleFormat and sets the default values.
func NewWorkbookChartTitleFormat()(*WorkbookChartTitleFormat) {
    m := &WorkbookChartTitleFormat{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the fill property value. Represents the fill format of an object, which includes background formatting information. Read-only.
func (m *WorkbookChartTitleFormat) GetFill()(*WorkbookChartFill) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
// Gets the font property value. Represents the font attributes (font name, font size, color, etc.) for the current object. Read-only.
func (m *WorkbookChartTitleFormat) GetFont()(*WorkbookChartFont) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
// The deserialization information for the current model
func (m *WorkbookChartTitleFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartFill() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(*WorkbookChartFill))
        }
        return nil
    }
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
    return res
}
func (m *WorkbookChartTitleFormat) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartTitleFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fill", m.GetFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the fill property value. Represents the fill format of an object, which includes background formatting information. Read-only.
// Parameters:
//  - value : Value to set for the fill property.
func (m *WorkbookChartTitleFormat) SetFill(value *WorkbookChartFill)() {
    m.fill = value
}
// Sets the font property value. Represents the font attributes (font name, font size, color, etc.) for the current object. Read-only.
// Parameters:
//  - value : Value to set for the font property.
func (m *WorkbookChartTitleFormat) SetFont(value *WorkbookChartFont)() {
    m.font = value
}
