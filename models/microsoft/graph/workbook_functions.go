package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookFunctions provides operations to manage the collection of drive entities.
type WorkbookFunctions struct {
    Entity
}
// NewWorkbookFunctions instantiates a new workbookFunctions and sets the default values.
func NewWorkbookFunctions()(*WorkbookFunctions) {
    m := &WorkbookFunctions{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookFunctionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFunctionsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookFunctions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFunctions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *WorkbookFunctions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookFunctions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
