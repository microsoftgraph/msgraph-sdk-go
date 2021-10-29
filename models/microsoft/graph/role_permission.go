package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RolePermission struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Resource Actions each containing a set of allowed and not allowed permissions.
    resourceActions []ResourceAction;
}
// Instantiates a new rolePermission and sets the default values.
func NewRolePermission()(*RolePermission) {
    m := &RolePermission{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RolePermission) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) GetResourceActions()([]ResourceAction) {
    if m == nil {
        return nil
    } else {
        return m.resourceActions
    }
}
// The deserialization information for the current model
func (m *RolePermission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceAction() })
        if err != nil {
            return err
        }
        res := make([]ResourceAction, len(val))
        for i, v := range val {
            res[i] = *(v.(*ResourceAction))
        }
        m.SetResourceActions(res)
        return nil
    }
    return res
}
func (m *RolePermission) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RolePermission) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceActions()))
        for i, v := range m.GetResourceActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("resourceActions", cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RolePermission) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
// Parameters:
//  - value : Value to set for the resourceActions property.
func (m *RolePermission) SetResourceActions(value []ResourceAction)() {
    m.resourceActions = value
}
