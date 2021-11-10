package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookChartLineFormat struct {
    Entity
    // HTML color code representing the color of lines in the chart.
    color *string;
}
// Instantiates a new workbookChartLineFormat and sets the default values.
func NewWorkbookChartLineFormat()(*WorkbookChartLineFormat) {
    m := &WorkbookChartLineFormat{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the color property value. HTML color code representing the color of lines in the chart.
func (m *WorkbookChartLineFormat) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// The deserialization information for the current model
func (m *WorkbookChartLineFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookChartLineFormat) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookChartLineFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the color property value. HTML color code representing the color of lines in the chart.
// Parameters:
//  - value : Value to set for the color property.
func (m *WorkbookChartLineFormat) SetColor(value *string)() {
    m.color = value
}
