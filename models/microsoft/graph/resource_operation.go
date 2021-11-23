package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// resourceOperation 
type ResourceOperation struct {
    Entity
    // Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
    actionName *string;
    // Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
    description *string;
    // Name of the Resource this operation is performed on.
    resourceName *string;
}
// NewResourceOperation instantiates a new resourceOperation and sets the default values.
func NewResourceOperation()(*ResourceOperation) {
    m := &ResourceOperation{
        Entity: *NewEntity(),
    }
    return m
}
// GetActionName gets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) GetActionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionName
    }
}
// GetDescription gets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetResourceName gets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) GetResourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val)
        }
        return nil
    }
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
    res["resourceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    return res
}
func (m *ResourceOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResourceOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionName sets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) SetActionName(value *string)() {
    m.actionName = value
}
// SetDescription sets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) SetDescription(value *string)() {
    m.description = value
}
// SetResourceName sets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) SetResourceName(value *string)() {
    m.resourceName = value
}
