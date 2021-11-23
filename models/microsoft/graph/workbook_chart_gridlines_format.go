package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// workbookChartGridlinesFormat 
type WorkbookChartGridlinesFormat struct {
    Entity
    // Represents chart line formatting. Read-only.
    line *WorkbookChartLineFormat;
}
// NewWorkbookChartGridlinesFormat instantiates a new workbookChartGridlinesFormat and sets the default values.
func NewWorkbookChartGridlinesFormat()(*WorkbookChartGridlinesFormat) {
    m := &WorkbookChartGridlinesFormat{
        Entity: *NewEntity(),
    }
    return m
}
// GetLine gets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartGridlinesFormat) GetLine()(*WorkbookChartLineFormat) {
    if m == nil {
        return nil
    } else {
        return m.line
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartGridlinesFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
func (m *WorkbookChartGridlinesFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartGridlinesFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("line", m.GetLine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLine sets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartGridlinesFormat) SetLine(value *WorkbookChartLineFormat)() {
    m.line = value
}
