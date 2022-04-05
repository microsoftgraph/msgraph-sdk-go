package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RbacApplication 
type RbacApplication struct {
    Entity
    // Resource to grant access to users or groups.
    roleAssignments []UnifiedRoleAssignmentable;
    // Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
    roleDefinitions []UnifiedRoleDefinitionable;
}
// NewRbacApplication instantiates a new rbacApplication and sets the default values.
func NewRbacApplication()(*RbacApplication) {
    m := &RbacApplication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRbacApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRbacApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRbacApplication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RbacApplication) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["roleAssignments"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleAssignmentable)
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleDefinitionable)
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    return res
}
// GetRoleAssignments gets the roleAssignments property value. Resource to grant access to users or groups.
func (m *RbacApplication) GetRoleAssignments()([]UnifiedRoleAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRoleDefinitions gets the roleDefinitions property value. Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
func (m *RbacApplication) GetRoleDefinitions()([]UnifiedRoleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// Serialize serializes information the current object
func (m *RbacApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRoleAssignments sets the roleAssignments property value. Resource to grant access to users or groups.
func (m *RbacApplication) SetRoleAssignments(value []UnifiedRoleAssignmentable)() {
    if m != nil {
        m.roleAssignments = value
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
func (m *RbacApplication) SetRoleDefinitions(value []UnifiedRoleDefinitionable)() {
    if m != nil {
        m.roleDefinitions = value
    }
}
