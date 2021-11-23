package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookRangeSort 
type WorkbookRangeSort struct {
    Entity
}
// NewWorkbookRangeSort instantiates a new workbookRangeSort and sets the default values.
func NewWorkbookRangeSort()(*WorkbookRangeSort) {
    m := &WorkbookRangeSort{
        Entity: *NewEntity(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRangeSort) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *WorkbookRangeSort) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookRangeSort) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
