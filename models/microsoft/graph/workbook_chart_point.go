package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartPoint provides operations to manage the collection of drive entities.
type WorkbookChartPoint struct {
    Entity
    // Encapsulates the format properties chart point. Read-only.
    format WorkbookChartPointFormatable;
    // Returns the value of a chart point. Read-only.
    value Jsonable;
}
// NewWorkbookChartPoint instantiates a new workbookChartPoint and sets the default values.
func NewWorkbookChartPoint()(*WorkbookChartPoint) {
    m := &WorkbookChartPoint{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartPointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartPointFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartPoint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartPoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartPointFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartPointFormatable))
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Encapsulates the format properties chart point. Read-only.
func (m *WorkbookChartPoint) GetFormat()(WorkbookChartPointFormatable) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetValue gets the value property value. Returns the value of a chart point. Read-only.
func (m *WorkbookChartPoint) GetValue()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *WorkbookChartPoint) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartPoint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Encapsulates the format properties chart point. Read-only.
func (m *WorkbookChartPoint) SetFormat(value WorkbookChartPointFormatable)() {
    if m != nil {
        m.format = value
    }
}
// SetValue sets the value property value. Returns the value of a chart point. Read-only.
func (m *WorkbookChartPoint) SetValue(value Jsonable)() {
    if m != nil {
        m.value = value
    }
}
