package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookTableColumn struct {
    Entity
    filter *WorkbookFilter;
    index *int32;
    name *string;
    values *Json;
}
func NewWorkbookTableColumn()(*WorkbookTableColumn) {
    m := &WorkbookTableColumn{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookTableColumn) GetFilter()(*WorkbookFilter) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
func (m *WorkbookTableColumn) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
func (m *WorkbookTableColumn) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *WorkbookTableColumn) GetValues()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *WorkbookTableColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookFilter() })
        if err != nil {
            return err
        }
        m.SetFilter(val.(*WorkbookFilter))
        return nil
    }
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIndex(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetValues(val.(*Json))
        return nil
    }
    return res
}
func (m *WorkbookTableColumn) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookTableColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookTableColumn) SetFilter(value *WorkbookFilter)() {
    m.filter = value
}
func (m *WorkbookTableColumn) SetIndex(value *int32)() {
    m.index = value
}
func (m *WorkbookTableColumn) SetName(value *string)() {
    m.name = value
}
func (m *WorkbookTableColumn) SetValues(value *Json)() {
    m.values = value
}
