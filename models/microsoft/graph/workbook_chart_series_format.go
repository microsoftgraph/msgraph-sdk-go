package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartSeriesFormat provides operations to manage the collection of drive entities.
type WorkbookChartSeriesFormat struct {
    Entity
    // Represents the fill format of a chart series, which includes background formating information. Read-only.
    fill WorkbookChartFillable;
    // Represents line formatting. Read-only.
    line WorkbookChartLineFormatable;
}
// NewWorkbookChartSeriesFormat instantiates a new workbookChartSeriesFormat and sets the default values.
func NewWorkbookChartSeriesFormat()(*WorkbookChartSeriesFormat) {
    m := &WorkbookChartSeriesFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartSeriesFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartSeriesFormatFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartSeriesFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartSeriesFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFillFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(WorkbookChartFillable))
        }
        return nil
    }
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
// GetFill gets the fill property value. Represents the fill format of a chart series, which includes background formating information. Read-only.
func (m *WorkbookChartSeriesFormat) GetFill()(WorkbookChartFillable) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
// GetLine gets the line property value. Represents line formatting. Read-only.
func (m *WorkbookChartSeriesFormat) GetLine()(WorkbookChartLineFormatable) {
    if m == nil {
        return nil
    } else {
        return m.line
    }
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
func (m *WorkbookChartSeriesFormat) SetFill(value WorkbookChartFillable)() {
    if m != nil {
        m.fill = value
    }
}
// SetLine sets the line property value. Represents line formatting. Read-only.
func (m *WorkbookChartSeriesFormat) SetLine(value WorkbookChartLineFormatable)() {
    if m != nil {
        m.line = value
    }
}
