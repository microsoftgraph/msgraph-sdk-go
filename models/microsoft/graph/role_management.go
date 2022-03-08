package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleManagement provides operations to manage the roleManagement singleton.
type RoleManagement struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Read-only. Nullable.
    directory RbacApplicationable;
    // Container for all entitlement management resources in Azure AD identity governance.
    entitlementManagement RbacApplicationable;
}
// NewRoleManagement instantiates a new RoleManagement and sets the default values.
func NewRoleManagement()(*RoleManagement) {
    m := &RoleManagement{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleManagementFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRoleManagement(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleManagement) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDirectory gets the directory property value. Read-only. Nullable.
func (m *RoleManagement) GetDirectory()(RbacApplicationable) {
    if m == nil {
        return nil
    } else {
        return m.directory
    }
}
// GetEntitlementManagement gets the entitlementManagement property value. Container for all entitlement management resources in Azure AD identity governance.
func (m *RoleManagement) GetEntitlementManagement()(RbacApplicationable) {
    if m == nil {
        return nil
    } else {
        return m.entitlementManagement
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["directory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRbacApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(RbacApplicationable))
        }
        return nil
    }
    res["entitlementManagement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRbacApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitlementManagement(val.(RbacApplicationable))
        }
        return nil
    }
    return res
}
func (m *RoleManagement) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleManagement) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDirectory sets the directory property value. Read-only. Nullable.
func (m *RoleManagement) SetDirectory(value RbacApplicationable)() {
    if m != nil {
        m.directory = value
    }
}
// SetEntitlementManagement sets the entitlementManagement property value. Container for all entitlement management resources in Azure AD identity governance.
func (m *RoleManagement) SetEntitlementManagement(value RbacApplicationable)() {
    if m != nil {
        m.entitlementManagement = value
    }
}
