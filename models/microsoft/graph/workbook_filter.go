package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookFilter 
type WorkbookFilter struct {
    Entity
    // The currently applied filter on the given column. Read-only.
    criteria WorkbookFilterCriteriaable;
}
// NewWorkbookFilter instantiates a new workbookFilter and sets the default values.
func NewWorkbookFilter()(*WorkbookFilter) {
    m := &WorkbookFilter{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFilterFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookFilter(), nil
}
// GetCriteria gets the criteria property value. The currently applied filter on the given column. Read-only.
func (m *WorkbookFilter) GetCriteria()(WorkbookFilterCriteriaable) {
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
        val, err := n.GetObjectValue(CreateWorkbookFilterCriteriaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val.(WorkbookFilterCriteriaable))
        }
        return nil
    }
    return res
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
func (m *WorkbookFilter) SetCriteria(value WorkbookFilterCriteriaable)() {
    if m != nil {
        m.criteria = value
    }
}
