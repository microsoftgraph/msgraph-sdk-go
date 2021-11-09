package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookFilter struct {
    Entity
    // The currently applied filter on the given column. Read-only.
    criteria *WorkbookFilterCriteria;
}
// Instantiates a new workbookFilter and sets the default values.
func NewWorkbookFilter()(*WorkbookFilter) {
    m := &WorkbookFilter{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the criteria property value. The currently applied filter on the given column. Read-only.
func (m *WorkbookFilter) GetCriteria()(*WorkbookFilterCriteria) {
    if m == nil {
        return nil
    } else {
        return m.criteria
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the criteria property value. The currently applied filter on the given column. Read-only.
// Parameters:
//  - value : Value to set for the criteria property.
func (m *WorkbookFilter) SetCriteria(value *WorkbookFilterCriteria)() {
    m.criteria = value
}
