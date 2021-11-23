package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// workbookChartSeriesFormat 
type WorkbookChartSeriesFormat struct {
    Entity
    // Represents the fill format of a chart series, which includes background formating information. Read-only.
    fill *WorkbookChartFill;
    // Represents line formatting. Read-only.
    line *WorkbookChartLineFormat;
}
// NewWorkbookChartSeriesFormat instantiates a new workbookChartSeriesFormat and sets the default values.
func NewWorkbookChartSeriesFormat()(*WorkbookChartSeriesFormat) {
    m := &WorkbookChartSeriesFormat{
        Entity: *NewEntity(),
    }
    return m
}
// GetFill gets the fill property value. Represents the fill format of a chart series, which includes background formating information. Read-only.
func (m *WorkbookChartSeriesFormat) GetFill()(*WorkbookChartFill) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
// GetLine gets the line property value. Represents line formatting. Read-only.
func (m *WorkbookChartSeriesFormat) GetLine()(*WorkbookChartLineFormat) {
    if m == nil {
        return nil
    } else {
        return m.line
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartSeriesFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *WorkbookChartSeriesFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartSeriesFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("line", m.GetLine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFill sets the fill property value. Represents the fill format of a chart series, which includes background formating information. Read-only.
func (m *WorkbookChartSeriesFormat) SetFill(value *WorkbookChartFill)() {
    m.fill = value
}
// SetLine sets the line property value. Represents line formatting. Read-only.
func (m *WorkbookChartSeriesFormat) SetLine(value *WorkbookChartLineFormat)() {
    m.line = value
}
