package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartDataLabelFormat struct {
    Entity
    // Represents the fill format of the current chart data label. Read-only.
    fill *WorkbookChartFill;
    // Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
    font *WorkbookChartFont;
}
// Instantiates a new workbookChartDataLabelFormat and sets the default values.
func NewWorkbookChartDataLabelFormat()(*WorkbookChartDataLabelFormat) {
    m := &WorkbookChartDataLabelFormat{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the fill property value. Represents the fill format of the current chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) GetFill()(*WorkbookChartFill) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
// Gets the font property value. Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) GetFont()(*WorkbookChartFont) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
// The deserialization information for the current model
func (m *WorkbookChartDataLabelFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartFill() })
        if err != nil {
            return err
        }
        m.SetFill(val.(*WorkbookChartFill))
        return nil
    }
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
func (m *WorkbookChartDataLabelFormat) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartDataLabelFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the fill property value. Represents the fill format of the current chart data label. Read-only.
// Parameters:
//  - value : Value to set for the fill property.
func (m *WorkbookChartDataLabelFormat) SetFill(value *WorkbookChartFill)() {
    m.fill = value
}
// Sets the font property value. Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
// Parameters:
//  - value : Value to set for the font property.
func (m *WorkbookChartDataLabelFormat) SetFont(value *WorkbookChartFont)() {
    m.font = value
}
