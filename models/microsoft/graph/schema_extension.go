package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SchemaExtension 
type SchemaExtension struct {
    Entity
    // Description for the schema extension. Supports $filter (eq).
    description *string;
    // The appId of the application that is the owner of the schema extension. This property can be supplied on creation, to set the owner.  If not supplied, then the calling application's appId will be set as the owner. In either case, the signed-in user must be the owner of the application. So, for example, if creating a new schema extension definition using Graph Explorer, you must supply the owner property. Once set, this property is read-only and cannot be changed. Supports $filter (eq).
    owner *string;
    // The collection of property names and types that make up the schema extension definition.
    properties []ExtensionSchemaProperty;
    // The lifecycle state of the schema extension. Possible states are InDevelopment, Available, and Deprecated. Automatically set to InDevelopment on creation. Schema extensions provides more information on the possible state transitions and behaviors. Supports $filter (eq).
    status *string;
    // Set of Microsoft Graph types (that can support extensions) that the schema extension can be applied to. Select from contact, device, event, group, message, organization, post, or user.
    targetTypes []string;
}
// NewSchemaExtension instantiates a new schemaExtension and sets the default values.
func NewSchemaExtension()(*SchemaExtension) {
    m := &SchemaExtension{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Description for the schema extension. Supports $filter (eq).
func (m *SchemaExtension) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetOwner gets the owner property value. The appId of the application that is the owner of the schema extension. This property can be supplied on creation, to set the owner.  If not supplied, then the calling application's appId will be set as the owner. In either case, the signed-in user must be the owner of the application. So, for example, if creating a new schema extension definition using Graph Explorer, you must supply the owner property. Once set, this property is read-only and cannot be changed. Supports $filter (eq).
func (m *SchemaExtension) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetProperties gets the properties property value. The collection of property names and types that make up the schema extension definition.
func (m *SchemaExtension) GetProperties()([]ExtensionSchemaProperty) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
// GetStatus gets the status property value. The lifecycle state of the schema extension. Possible states are InDevelopment, Available, and Deprecated. Automatically set to InDevelopment on creation. Schema extensions provides more information on the possible state transitions and behaviors. Supports $filter (eq).
func (m *SchemaExtension) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTargetTypes gets the targetTypes property value. Set of Microsoft Graph types (that can support extensions) that the schema extension can be applied to. Select from contact, device, event, group, message, organization, post, or user.
func (m *SchemaExtension) GetTargetTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SchemaExtension) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtensionSchemaProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExtensionSchemaProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExtensionSchemaProperty))
            }
            m.SetProperties(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["targetTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTargetTypes(res)
        }
        return nil
    }
    return res
}
func (m *SchemaExtension) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SchemaExtension) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    if m.GetTargetTypes() != nil {
        err = writer.WriteCollectionOfStringValues("targetTypes", m.GetTargetTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description for the schema extension. Supports $filter (eq).
func (m *SchemaExtension) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetOwner sets the owner property value. The appId of the application that is the owner of the schema extension. This property can be supplied on creation, to set the owner.  If not supplied, then the calling application's appId will be set as the owner. In either case, the signed-in user must be the owner of the application. So, for example, if creating a new schema extension definition using Graph Explorer, you must supply the owner property. Once set, this property is read-only and cannot be changed. Supports $filter (eq).
func (m *SchemaExtension) SetOwner(value *string)() {
    if m != nil {
        m.owner = value
    }
}
// SetProperties sets the properties property value. The collection of property names and types that make up the schema extension definition.
func (m *SchemaExtension) SetProperties(value []ExtensionSchemaProperty)() {
    if m != nil {
        m.properties = value
    }
}
// SetStatus sets the status property value. The lifecycle state of the schema extension. Possible states are InDevelopment, Available, and Deprecated. Automatically set to InDevelopment on creation. Schema extensions provides more information on the possible state transitions and behaviors. Supports $filter (eq).
func (m *SchemaExtension) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTargetTypes sets the targetTypes property value. Set of Microsoft Graph types (that can support extensions) that the schema extension can be applied to. Select from contact, device, event, group, message, organization, post, or user.
func (m *SchemaExtension) SetTargetTypes(value []string)() {
    if m != nil {
        m.targetTypes = value
    }
}
