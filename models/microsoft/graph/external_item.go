package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExternalItem 
type ExternalItem struct {
    Entity
    // An array of access control entries. Each entry specifies the access granted to a user or group. Required.
    acl []Acl;
    // A plain-text  representation of the contents of the item. The text in this property is full-text indexed. Optional.
    content *ExternalItemContent;
    // A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
    properties *Properties;
}
// NewExternalItem instantiates a new externalItem and sets the default values.
func NewExternalItem()(*ExternalItem) {
    m := &ExternalItem{
        Entity: *NewEntity(),
    }
    return m
}
// GetAcl gets the acl property value. An array of access control entries. Each entry specifies the access granted to a user or group. Required.
func (m *ExternalItem) GetAcl()([]Acl) {
    if m == nil {
        return nil
    } else {
        return m.acl
    }
}
// GetContent gets the content property value. A plain-text  representation of the contents of the item. The text in this property is full-text indexed. Optional.
func (m *ExternalItem) GetContent()(*ExternalItemContent) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetProperties gets the properties property value. A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
func (m *ExternalItem) GetProperties()(*Properties) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAcl() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Acl, len(val))
            for i, v := range val {
                res[i] = *(v.(*Acl))
            }
            m.SetAcl(res)
        }
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalItemContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(*ExternalItemContent))
        }
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProperties() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperties(val.(*Properties))
        }
        return nil
    }
    return res
}
func (m *ExternalItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAcl sets the acl property value. An array of access control entries. Each entry specifies the access granted to a user or group. Required.
func (m *ExternalItem) SetAcl(value []Acl)() {
    if m != nil {
        m.acl = value
    }
}
// SetContent sets the content property value. A plain-text  representation of the contents of the item. The text in this property is full-text indexed. Optional.
func (m *ExternalItem) SetContent(value *ExternalItemContent)() {
    if m != nil {
        m.content = value
    }
}
// SetProperties sets the properties property value. A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
func (m *ExternalItem) SetProperties(value *Properties)() {
    if m != nil {
        m.properties = value
    }
}
