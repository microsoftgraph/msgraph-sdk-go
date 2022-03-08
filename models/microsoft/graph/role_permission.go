package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RolePermission provides operations to manage the deviceManagement singleton.
type RolePermission struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Resource Actions each containing a set of allowed and not allowed permissions.
    resourceActions []ResourceActionable;
}
// NewRolePermission instantiates a new rolePermission and sets the default values.
func NewRolePermission()(*RolePermission) {
    m := &RolePermission{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRolePermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRolePermissionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRolePermission(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RolePermission) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RolePermission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceActionable, len(val))
            for i, v := range val {
                res[i] = v.(ResourceActionable)
            }
            m.SetResourceActions(res)
        }
        return nil
    }
    return res
}
// GetResourceActions gets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) GetResourceActions()([]ResourceActionable) {
    if m == nil {
        return nil
    } else {
        return m.resourceActions
    }
}
func (m *RolePermission) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RolePermission) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetResourceActions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceActions()))
        for i, v := range m.GetResourceActions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RolePermission) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResourceActions sets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) SetResourceActions(value []ResourceActionable)() {
    if m != nil {
        m.resourceActions = value
    }
}
