package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResourceReference provides operations to manage the collection of drive entities.
type ResourceReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The item's unique identifier.
    id *string;
    // A string value that can be used to classify the item, such as 'microsoft.graph.driveItem'
    type_escaped *string;
    // A URL leading to the referenced item.
    webUrl *string;
}
// NewResourceReference instantiates a new resourceReference and sets the default values.
func NewResourceReference()(*ResourceReference) {
    m := &ResourceReference{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateResourceReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceReferenceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewResourceReference(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceReference) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The item's unique identifier.
func (m *ResourceReference) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetType gets the type property value. A string value that can be used to classify the item, such as 'microsoft.graph.driveItem'
func (m *ResourceReference) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetWebUrl gets the webUrl property value. A URL leading to the referenced item.
func (m *ResourceReference) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *ResourceReference) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResourceReference) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceReference) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. The item's unique identifier.
func (m *ResourceReference) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetType sets the type property value. A string value that can be used to classify the item, such as 'microsoft.graph.driveItem'
func (m *ResourceReference) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetWebUrl sets the webUrl property value. A URL leading to the referenced item.
func (m *ResourceReference) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
