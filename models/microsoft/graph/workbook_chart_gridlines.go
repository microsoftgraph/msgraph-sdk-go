package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartGridlines struct {
    Entity
    // Represents the formatting of chart gridlines. Read-only.
    format *WorkbookChartGridlinesFormat;
    // Boolean value representing if the axis gridlines are visible or not.
    visible *bool;
}
// Instantiates a new workbookChartGridlines and sets the default values.
func NewWorkbookChartGridlines()(*WorkbookChartGridlines) {
    m := &WorkbookChartGridlines{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the format property value. Represents the formatting of chart gridlines. Read-only.
func (m *WorkbookChartGridlines) GetFormat()(*WorkbookChartGridlinesFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the visible property value. Boolean value representing if the axis gridlines are visible or not.
func (m *WorkbookChartGridlines) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
// The deserialization information for the current model
func (m *WorkbookChartGridlines) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartGridlinesFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*WorkbookChartGridlinesFormat))
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
func (m *WorkbookChartGridlines) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartGridlines) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("visible", m.GetVisible())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the format property value. Represents the formatting of chart gridlines. Read-only.
// Parameters:
//  - value : Value to set for the format property.
func (m *WorkbookChartGridlines) SetFormat(value *WorkbookChartGridlinesFormat)() {
    m.format = value
}
// Sets the visible property value. Boolean value representing if the axis gridlines are visible or not.
// Parameters:
//  - value : Value to set for the visible property.
func (m *WorkbookChartGridlines) SetVisible(value *bool)() {
    m.visible = value
}
