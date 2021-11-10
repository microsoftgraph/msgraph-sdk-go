package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SharedInsight struct {
    Entity
    // Details about the shared item. Read only.
    lastShared *SharingDetail;
    // 
    lastSharedMethod *Entity;
    // Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
    resource *Entity;
    // Reference properties of the shared document, such as the url and type of the document. Read-only
    resourceReference *ResourceReference;
    // Properties that you can use to visualize the document in your experience. Read-only
    resourceVisualization *ResourceVisualization;
    // 
    sharingHistory []SharingDetail;
}
// Instantiates a new sharedInsight and sets the default values.
func NewSharedInsight()(*SharedInsight) {
    m := &SharedInsight{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the lastShared property value. Details about the shared item. Read only.
func (m *SharedInsight) GetLastShared()(*SharingDetail) {
    if m == nil {
        return nil
    } else {
        return m.lastShared
    }
}
// Gets the lastSharedMethod property value. 
func (m *SharedInsight) GetLastSharedMethod()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m.lastSharedMethod
    }
}
// Gets the resource property value. Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *SharedInsight) GetResource()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// Gets the resourceReference property value. Reference properties of the shared document, such as the url and type of the document. Read-only
func (m *SharedInsight) GetResourceReference()(*ResourceReference) {
    if m == nil {
        return nil
    } else {
        return m.resourceReference
    }
}
// Gets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *SharedInsight) GetResourceVisualization()(*ResourceVisualization) {
    if m == nil {
        return nil
    } else {
        return m.resourceVisualization
    }
}
// Gets the sharingHistory property value. 
func (m *SharedInsight) GetSharingHistory()([]SharingDetail) {
    if m == nil {
        return nil
    } else {
        return m.sharingHistory
    }
}
// The deserialization information for the current model
func (m *SharedInsight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastShared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharingDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastShared(val.(*SharingDetail))
        }
        return nil
    }
    res["lastSharedMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSharedMethod(val.(*Entity))
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(*Entity))
        }
        return nil
    }
    res["resourceReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceReference() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceReference(val.(*ResourceReference))
        }
        return nil
    }
    res["resourceVisualization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceVisualization() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceVisualization(val.(*ResourceVisualization))
        }
        return nil
    }
    res["sharingHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharingDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharingDetail, len(val))
            for i, v := range val {
                res[i] = *(v.(*SharingDetail))
            }
            m.SetSharingHistory(res)
        }
        return nil
    }
    return res
}
func (m *SharedInsight) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the lastShared property value. Details about the shared item. Read only.
// Parameters:
//  - value : Value to set for the lastShared property.
func (m *SharedInsight) SetLastShared(value *SharingDetail)() {
    m.lastShared = value
}
// Sets the lastSharedMethod property value. 
// Parameters:
//  - value : Value to set for the lastSharedMethod property.
func (m *SharedInsight) SetLastSharedMethod(value *Entity)() {
    m.lastSharedMethod = value
}
// Sets the resource property value. Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
// Parameters:
//  - value : Value to set for the resource property.
func (m *SharedInsight) SetResource(value *Entity)() {
    m.resource = value
}
// Sets the resourceReference property value. Reference properties of the shared document, such as the url and type of the document. Read-only
// Parameters:
//  - value : Value to set for the resourceReference property.
func (m *SharedInsight) SetResourceReference(value *ResourceReference)() {
    m.resourceReference = value
}
// Sets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
// Parameters:
//  - value : Value to set for the resourceVisualization property.
func (m *SharedInsight) SetResourceVisualization(value *ResourceVisualization)() {
    m.resourceVisualization = value
}
// Sets the sharingHistory property value. 
// Parameters:
//  - value : Value to set for the sharingHistory property.
func (m *SharedInsight) SetSharingHistory(value []SharingDetail)() {
    m.sharingHistory = value
}
