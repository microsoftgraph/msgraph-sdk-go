package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookFilter 
type WorkbookFilter struct {
    Entity
    // The currently applied filter on the given column. Read-only.
    criteria *WorkbookFilterCriteria;
}
// NewWorkbookFilter instantiates a new workbookFilter and sets the default values.
func NewWorkbookFilter()(*WorkbookFilter) {
    m := &WorkbookFilter{
        Entity: *NewEntity(),
    }
    return m
}
// GetCriteria gets the criteria property value. The currently applied filter on the given column. Read-only.
func (m *WorkbookFilter) GetCriteria()(*WorkbookFilterCriteria) {
    if m == nil {
        return nil
    } else {
        return m.criteria
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFilter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["criteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookFilterCriteria() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val.(*WorkbookFilterCriteria))
        }
        return nil
    }
    return res
}
func (m *WorkbookFilter) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookFilter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("criteria", m.GetCriteria())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCriteria sets the criteria property value. The currently applied filter on the given column. Read-only.
func (m *WorkbookFilter) SetCriteria(value *WorkbookFilterCriteria)() {
    m.criteria = value
}
