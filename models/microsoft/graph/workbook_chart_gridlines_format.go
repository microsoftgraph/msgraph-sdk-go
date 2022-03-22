package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartGridlinesFormat 
type WorkbookChartGridlinesFormat struct {
    Entity
    // Represents chart line formatting. Read-only.
    line WorkbookChartLineFormatable;
}
// NewWorkbookChartGridlinesFormat instantiates a new workbookChartGridlinesFormat and sets the default values.
func NewWorkbookChartGridlinesFormat()(*WorkbookChartGridlinesFormat) {
    m := &WorkbookChartGridlinesFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartGridlinesFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartGridlinesFormatFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartGridlinesFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartGridlinesFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["line"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartLineFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val.(WorkbookChartLineFormatable))
        }
        return nil
    }
    return res
}
// GetLine gets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartGridlinesFormat) GetLine()(WorkbookChartLineFormatable) {
    if m == nil {
        return nil
    } else {
        return m.line
    }
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
func (m *WorkbookChartGridlinesFormat) SetLine(value WorkbookChartLineFormatable)() {
    if m != nil {
        m.line = value
    }
}
