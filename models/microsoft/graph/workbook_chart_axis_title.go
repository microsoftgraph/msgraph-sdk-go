package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartAxisTitle struct {
    Entity
    // Represents the formatting of chart axis title. Read-only.
    format *WorkbookChartAxisTitleFormat;
    // Represents the axis title.
    text *string;
    // A boolean that specifies the visibility of an axis title.
    visible *bool;
}
// Instantiates a new workbookChartAxisTitle and sets the default values.
func NewWorkbookChartAxisTitle()(*WorkbookChartAxisTitle) {
    m := &WorkbookChartAxisTitle{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the format property value. Represents the formatting of chart axis title. Read-only.
func (m *WorkbookChartAxisTitle) GetFormat()(*WorkbookChartAxisTitleFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the text property value. Represents the axis title.
func (m *WorkbookChartAxisTitle) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// Gets the visible property value. A boolean that specifies the visibility of an axis title.
func (m *WorkbookChartAxisTitle) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
// The deserialization information for the current model
func (m *WorkbookChartAxisTitle) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxisTitleFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartAxisTitleFormat))
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
func (m *WorkbookChartAxisTitle) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartAxisTitle) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the format property value. Represents the formatting of chart axis title. Read-only.
// Parameters:
//  - value : Value to set for the format property.
func (m *WorkbookChartAxisTitle) SetFormat(value *WorkbookChartAxisTitleFormat)() {
    m.format = value
}
// Sets the text property value. Represents the axis title.
// Parameters:
//  - value : Value to set for the text property.
func (m *WorkbookChartAxisTitle) SetText(value *string)() {
    m.text = value
}
// Sets the visible property value. A boolean that specifies the visibility of an axis title.
// Parameters:
//  - value : Value to set for the visible property.
func (m *WorkbookChartAxisTitle) SetVisible(value *bool)() {
    m.visible = value
}
