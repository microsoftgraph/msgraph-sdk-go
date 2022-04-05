package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleManagement 
type RoleManagement struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Read-only. Nullable.
    directory RbacApplicationable;
    // Container for roles and assignments for entitlement management resources.
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
func CreateRoleManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
// GetEntitlementManagement gets the entitlementManagement property value. Container for roles and assignments for entitlement management resources.
func (m *RoleManagement) GetEntitlementManagement()(RbacApplicationable) {
    if m == nil {
        return nil
    } else {
        return m.entitlementManagement
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleManagement) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["directory"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRbacApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(RbacApplicationable))
        }
        return nil
    }
    res["entitlementManagement"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *RoleManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetEntitlementManagement sets the entitlementManagement property value. Container for roles and assignments for entitlement management resources.
func (m *RoleManagement) SetEntitlementManagement(value RbacApplicationable)() {
    if m != nil {
        m.entitlementManagement = value
    }
}
