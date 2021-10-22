package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookFilter struct {
    Entity
    criteria *WorkbookFilterCriteria;
}
func NewWorkbookFilter()(*WorkbookFilter) {
    m := &WorkbookFilter{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookFilter) GetCriteria()(*WorkbookFilterCriteria) {
    if m == nil {
        return nil
    } else {
        return m.criteria
    }
}
func (m *WorkbookFilter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["criteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookFilterCriteria() })
        if err != nil {
            return err
        }
        m.SetCriteria(val.(*WorkbookFilterCriteria))
        return nil
    }
    return res
}
func (m *WorkbookFilter) IsNil()(bool) {
    return m == nil
}
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
func (m *WorkbookFilter) SetCriteria(value *WorkbookFilterCriteria)() {
    m.criteria = value
}
