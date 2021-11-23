package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UsedInsight 
type UsedInsight struct {
    Entity
    // Information about when the item was last viewed or modified by the user. Read only.
    lastUsed *UsageDetails;
    // Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
    resource *Entity;
    // Reference properties of the used document, such as the url and type of the document. Read-only
    resourceReference *ResourceReference;
    // Properties that you can use to visualize the document in your experience. Read-only
    resourceVisualization *ResourceVisualization;
}
// NewUsedInsight instantiates a new usedInsight and sets the default values.
func NewUsedInsight()(*UsedInsight) {
    m := &UsedInsight{
        Entity: *NewEntity(),
    }
    return m
}
// GetLastUsed gets the lastUsed property value. Information about when the item was last viewed or modified by the user. Read only.
func (m *UsedInsight) GetLastUsed()(*UsageDetails) {
    if m == nil {
        return nil
    } else {
        return m.lastUsed
    }
}
// GetResource gets the resource property value. Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *UsedInsight) GetResource()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceReference gets the resourceReference property value. Reference properties of the used document, such as the url and type of the document. Read-only
func (m *UsedInsight) GetResourceReference()(*ResourceReference) {
    if m == nil {
        return nil
    } else {
        return m.resourceReference
    }
}
// GetResourceVisualization gets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *UsedInsight) GetResourceVisualization()(*ResourceVisualization) {
    if m == nil {
        return nil
    } else {
        return m.resourceVisualization
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsedInsight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUsageDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val.(*UsageDetails))
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
    return res
}
func (m *UsedInsight) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UsedInsight) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("lastUsed", m.GetLastUsed())
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
    return nil
}
// SetLastUsed sets the lastUsed property value. Information about when the item was last viewed or modified by the user. Read only.
func (m *UsedInsight) SetLastUsed(value *UsageDetails)() {
    m.lastUsed = value
}
// SetResource sets the resource property value. Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *UsedInsight) SetResource(value *Entity)() {
    m.resource = value
}
// SetResourceReference sets the resourceReference property value. Reference properties of the used document, such as the url and type of the document. Read-only
func (m *UsedInsight) SetResourceReference(value *ResourceReference)() {
    m.resourceReference = value
}
// SetResourceVisualization sets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *UsedInsight) SetResourceVisualization(value *ResourceVisualization)() {
    m.resourceVisualization = value
}
