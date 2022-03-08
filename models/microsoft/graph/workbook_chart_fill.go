package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartFill provides operations to manage the drive singleton.
type WorkbookChartFill struct {
    Entity
}
// NewWorkbookChartFill instantiates a new workbookChartFill and sets the default values.
func NewWorkbookChartFill()(*WorkbookChartFill) {
    m := &WorkbookChartFill{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartFillFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartFillFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartFill(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartFill) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *WorkbookChartFill) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartFill) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
