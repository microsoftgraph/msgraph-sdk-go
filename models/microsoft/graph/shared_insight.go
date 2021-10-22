package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SharedInsight struct {
    Entity
    lastShared *SharingDetail;
    lastSharedMethod *Entity;
    resource *Entity;
    resourceReference *ResourceReference;
    resourceVisualization *ResourceVisualization;
    sharingHistory []SharingDetail;
}
func NewSharedInsight()(*SharedInsight) {
    m := &SharedInsight{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SharedInsight) GetLastShared()(*SharingDetail) {
    if m == nil {
        return nil
    } else {
        return m.lastShared
    }
}
func (m *SharedInsight) GetLastSharedMethod()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m.lastSharedMethod
    }
}
func (m *SharedInsight) GetResource()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *SharedInsight) GetResourceReference()(*ResourceReference) {
    if m == nil {
        return nil
    } else {
        return m.resourceReference
    }
}
func (m *SharedInsight) GetResourceVisualization()(*ResourceVisualization) {
    if m == nil {
        return nil
    } else {
        return m.resourceVisualization
    }
}
func (m *SharedInsight) GetSharingHistory()([]SharingDetail) {
    if m == nil {
        return nil
    } else {
        return m.sharingHistory
    }
}
func (m *SharedInsight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastShared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharingDetail() })
        if err != nil {
            return err
        }
        m.SetLastShared(val.(*SharingDetail))
        return nil
    }
    res["lastSharedMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntity() })
        if err != nil {
            return err
        }
        m.SetLastSharedMethod(val.(*Entity))
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntity() })
        if err != nil {
            return err
        }
        m.SetResource(val.(*Entity))
        return nil
    }
    res["resourceReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceReference() })
        if err != nil {
            return err
        }
        m.SetResourceReference(val.(*ResourceReference))
        return nil
    }
    res["resourceVisualization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceVisualization() })
        if err != nil {
            return err
        }
        m.SetResourceVisualization(val.(*ResourceVisualization))
        return nil
    }
    res["sharingHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharingDetail() })
        if err != nil {
            return err
        }
        res := make([]SharingDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*SharingDetail))
        }
        m.SetSharingHistory(res)
        return nil
    }
    return res
}
func (m *SharedInsight) IsNil()(bool) {
    return m == nil
}
func (m *SharedInsight) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("lastShared", m.GetLastShared())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastSharedMethod", m.GetLastSharedMethod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourceReference", m.GetResourceReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourceVisualization", m.GetResourceVisualization())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSharingHistory()))
        for i, v := range m.GetSharingHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sharingHistory", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SharedInsight) SetLastShared(value *SharingDetail)() {
    m.lastShared = value
}
func (m *SharedInsight) SetLastSharedMethod(value *Entity)() {
    m.lastSharedMethod = value
}
func (m *SharedInsight) SetResource(value *Entity)() {
    m.resource = value
}
func (m *SharedInsight) SetResourceReference(value *ResourceReference)() {
    m.resourceReference = value
}
func (m *SharedInsight) SetResourceVisualization(value *ResourceVisualization)() {
    m.resourceVisualization = value
}
func (m *SharedInsight) SetSharingHistory(value []SharingDetail)() {
    m.sharingHistory = value
}
