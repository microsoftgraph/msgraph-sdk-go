package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// workbookChartPointFormat 
type WorkbookChartPointFormat struct {
    Entity
    // Represents the fill format of a chart, which includes background formating information. Read-only.
    fill *WorkbookChartFill;
}
// NewWorkbookChartPointFormat instantiates a new workbookChartPointFormat and sets the default values.
func NewWorkbookChartPointFormat()(*WorkbookChartPointFormat) {
    m := &WorkbookChartPointFormat{
        Entity: *NewEntity(),
    }
    return m
}
// GetFill gets the fill property value. Represents the fill format of a chart, which includes background formating information. Read-only.
func (m *WorkbookChartPointFormat) GetFill()(*WorkbookChartFill) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartPointFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
func (m *WorkbookChartPointFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartPointFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
// SetFill sets the fill property value. Represents the fill format of a chart, which includes background formating information. Read-only.
func (m *WorkbookChartPointFormat) SetFill(value *WorkbookChartFill)() {
    m.fill = value
}
