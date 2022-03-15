package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartLineFormat provides operations to manage the collection of drive entities.
type WorkbookChartLineFormat struct {
    Entity
    // HTML color code representing the color of lines in the chart.
    color *string;
}
// NewWorkbookChartLineFormat instantiates a new workbookChartLineFormat and sets the default values.
func NewWorkbookChartLineFormat()(*WorkbookChartLineFormat) {
    m := &WorkbookChartLineFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartLineFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartLineFormatFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartLineFormat(), nil
}
// GetColor gets the color property value. HTML color code representing the color of lines in the chart.
func (m *WorkbookChartLineFormat) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetColor sets the color property value. HTML color code representing the color of lines in the chart.
func (m *WorkbookChartLineFormat) SetColor(value *string)() {
    if m != nil {
        m.color = value
    }
}
