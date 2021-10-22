package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExternalItem struct {
    Entity
    acl []Acl;
    content *ExternalItemContent;
    properties *Properties;
}
func NewExternalItem()(*ExternalItem) {
    m := &ExternalItem{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ExternalItem) GetAcl()([]Acl) {
    if m == nil {
        return nil
    } else {
        return m.acl
    }
}
func (m *ExternalItem) GetContent()(*ExternalItemContent) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *ExternalItem) GetProperties()(*Properties) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
func (m *ExternalItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAcl() })
        if err != nil {
            return err
        }
        res := make([]Acl, len(val))
        for i, v := range val {
            res[i] = *(v.(*Acl))
        }
        m.SetAcl(res)
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalItemContent() })
        if err != nil {
            return err
        }
        m.SetContent(val.(*ExternalItemContent))
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProperties() })
        if err != nil {
            return err
        }
        m.SetProperties(val.(*Properties))
        return nil
    }
    return res
}
func (m *ExternalItem) IsNil()(bool) {
    return m == nil
}
func (m *ExternalItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAcl()))
        for i, v := range m.GetAcl() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("acl", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("properties", m.GetProperties())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ExternalItem) SetAcl(value []Acl)() {
    m.acl = value
}
func (m *ExternalItem) SetContent(value *ExternalItemContent)() {
    m.content = value
}
func (m *ExternalItem) SetProperties(value *Properties)() {
    m.properties = value
}
