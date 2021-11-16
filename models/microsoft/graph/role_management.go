package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RoleManagement struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Read-only. Nullable.
    directory *RbacApplication;
    // The RbacApplication for Entitlement Management
    entitlementManagement *RbacApplication;
}
// Instantiates a new RoleManagement and sets the default values.
func NewRoleManagement()(*RoleManagement) {
    m := &RoleManagement{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleManagement) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the directory property value. Read-only. Nullable.
func (m *RoleManagement) GetDirectory()(*RbacApplication) {
    if m == nil {
        return nil
    } else {
        return m.directory
    }
}
// Gets the entitlementManagement property value. The RbacApplication for Entitlement Management
func (m *RoleManagement) GetEntitlementManagement()(*RbacApplication) {
    if m == nil {
        return nil
    } else {
        return m.entitlementManagement
    }
}
// The deserialization information for the current model
func (m *RoleManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["directory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRbacApplication() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(*RbacApplication))
        }
        return nil
    }
    res["entitlementManagement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRbacApplication() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitlementManagement(val.(*RbacApplication))
        }
        return nil
    }
    return res
}
func (m *RoleManagement) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RoleManagement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("entitlementManagement", m.GetEntitlementManagement())
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
func (m *RoleManagement) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the directory property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the directory property.
func (m *RoleManagement) SetDirectory(value *RbacApplication)() {
    m.directory = value
}
// Sets the entitlementManagement property value. The RbacApplication for Entitlement Management
// Parameters:
//  - value : Value to set for the entitlementManagement property.
func (m *RoleManagement) SetEntitlementManagement(value *RbacApplication)() {
    m.entitlementManagement = value
}
