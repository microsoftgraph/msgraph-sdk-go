package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRolePermission 
type UnifiedRolePermission struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Set of tasks that can be performed on a resource. Required.
    allowedResourceActions []string;
    // Optional constraints that must be met for the permission to be effective.
    condition *string;
    // Set of tasks that may not be performed on a resource. Not yet supported.
    excludedResourceActions []string;
}
// NewUnifiedRolePermission instantiates a new unifiedRolePermission and sets the default values.
func NewUnifiedRolePermission()(*UnifiedRolePermission) {
    m := &UnifiedRolePermission{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnifiedRolePermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRolePermissionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUnifiedRolePermission(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRolePermission) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedResourceActions gets the allowedResourceActions property value. Set of tasks that can be performed on a resource. Required.
func (m *UnifiedRolePermission) GetAllowedResourceActions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.allowedResourceActions
    }
}
// GetCondition gets the condition property value. Optional constraints that must be met for the permission to be effective.
func (m *UnifiedRolePermission) GetCondition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.condition
    }
}
// GetExcludedResourceActions gets the excludedResourceActions property value. Set of tasks that may not be performed on a resource. Not yet supported.
func (m *UnifiedRolePermission) GetExcludedResourceActions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludedResourceActions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRolePermission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedResourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedResourceActions(res)
        }
        return nil
    }
    res["condition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val)
        }
        return nil
    }
    res["excludedResourceActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludedResourceActions(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UnifiedRolePermission) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAllowedResourceActions() != nil {
        err := writer.WriteCollectionOfStringValues("allowedResourceActions", m.GetAllowedResourceActions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("condition", m.GetCondition())
        if err != nil {
            return err
        }
    }
    if m.GetExcludedResourceActions() != nil {
        err := writer.WriteCollectionOfStringValues("excludedResourceActions", m.GetExcludedResourceActions())
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
func (m *UnifiedRolePermission) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowedResourceActions sets the allowedResourceActions property value. Set of tasks that can be performed on a resource. Required.
func (m *UnifiedRolePermission) SetAllowedResourceActions(value []string)() {
    if m != nil {
        m.allowedResourceActions = value
    }
}
// SetCondition sets the condition property value. Optional constraints that must be met for the permission to be effective.
func (m *UnifiedRolePermission) SetCondition(value *string)() {
    if m != nil {
        m.condition = value
    }
}
// SetExcludedResourceActions sets the excludedResourceActions property value. Set of tasks that may not be performed on a resource. Not yet supported.
func (m *UnifiedRolePermission) SetExcludedResourceActions(value []string)() {
    if m != nil {
        m.excludedResourceActions = value
    }
}
