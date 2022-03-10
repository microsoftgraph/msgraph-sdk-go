package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedInsight provides operations to manage the collection of drive entities.
type SharedInsight struct {
    Entity
    // Details about the shared item. Read only.
    lastShared SharingDetailable;
    // 
    lastSharedMethod Entityable;
    // Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
    resource Entityable;
    // Reference properties of the shared document, such as the url and type of the document. Read-only
    resourceReference ResourceReferenceable;
    // Properties that you can use to visualize the document in your experience. Read-only
    resourceVisualization ResourceVisualizationable;
    // 
    sharingHistory []SharingDetailable;
}
// NewSharedInsight instantiates a new sharedInsight and sets the default values.
func NewSharedInsight()(*SharedInsight) {
    m := &SharedInsight{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSharedInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedInsightFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSharedInsight(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedInsight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastShared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastShared(val.(SharingDetailable))
        }
        return nil
    }
    res["lastSharedMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSharedMethod(val.(Entityable))
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(Entityable))
        }
        return nil
    }
    res["resourceReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceReference(val.(ResourceReferenceable))
        }
        return nil
    }
    res["resourceVisualization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceVisualizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceVisualization(val.(ResourceVisualizationable))
        }
        return nil
    }
    res["sharingHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharingDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharingDetailable, len(val))
            for i, v := range val {
                res[i] = v.(SharingDetailable)
            }
            m.SetSharingHistory(res)
        }
        return nil
    }
    return res
}
// GetLastShared gets the lastShared property value. Details about the shared item. Read only.
func (m *SharedInsight) GetLastShared()(SharingDetailable) {
    if m == nil {
        return nil
    } else {
        return m.lastShared
    }
}
// GetLastSharedMethod gets the lastSharedMethod property value. 
func (m *SharedInsight) GetLastSharedMethod()(Entityable) {
    if m == nil {
        return nil
    } else {
        return m.lastSharedMethod
    }
}
// GetResource gets the resource property value. Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *SharedInsight) GetResource()(Entityable) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceReference gets the resourceReference property value. Reference properties of the shared document, such as the url and type of the document. Read-only
func (m *SharedInsight) GetResourceReference()(ResourceReferenceable) {
    if m == nil {
        return nil
    } else {
        return m.resourceReference
    }
}
// GetResourceVisualization gets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *SharedInsight) GetResourceVisualization()(ResourceVisualizationable) {
    if m == nil {
        return nil
    } else {
        return m.resourceVisualization
    }
}
// GetSharingHistory gets the sharingHistory property value. 
func (m *SharedInsight) GetSharingHistory()([]SharingDetailable) {
    if m == nil {
        return nil
    } else {
        return m.sharingHistory
    }
}
func (m *SharedInsight) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetSharingHistory() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSharingHistory()))
        for i, v := range m.GetSharingHistory() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sharingHistory", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastShared sets the lastShared property value. Details about the shared item. Read only.
func (m *SharedInsight) SetLastShared(value SharingDetailable)() {
    if m != nil {
        m.lastShared = value
    }
}
// SetLastSharedMethod sets the lastSharedMethod property value. 
func (m *SharedInsight) SetLastSharedMethod(value Entityable)() {
    if m != nil {
        m.lastSharedMethod = value
    }
}
// SetResource sets the resource property value. Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *SharedInsight) SetResource(value Entityable)() {
    if m != nil {
        m.resource = value
    }
}
// SetResourceReference sets the resourceReference property value. Reference properties of the shared document, such as the url and type of the document. Read-only
func (m *SharedInsight) SetResourceReference(value ResourceReferenceable)() {
    if m != nil {
        m.resourceReference = value
    }
}
// SetResourceVisualization sets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *SharedInsight) SetResourceVisualization(value ResourceVisualizationable)() {
    if m != nil {
        m.resourceVisualization = value
    }
}
// SetSharingHistory sets the sharingHistory property value. 
func (m *SharedInsight) SetSharingHistory(value []SharingDetailable)() {
    if m != nil {
        m.sharingHistory = value
    }
}
