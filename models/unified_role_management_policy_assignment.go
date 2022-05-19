package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyAssignment provides operations to manage the policyRoot singleton.
type UnifiedRoleManagementPolicyAssignment struct {
    Entity
    // The policy for the assignment.
    policy UnifiedRoleManagementPolicyable
    // The id of the policy.
    policyId *string
    // The id of the role definition where the policy applies. If not specified, the policy applies to all roles.
    roleDefinitionId *string
    // The id of the scope where the policy is assigned. E.g. '/', groupId, etc.
    scopeId *string
    // The type of the scope where the policy is assigned. One of Directory, DirectoryRole, Group.
    scopeType *string
}
// NewUnifiedRoleManagementPolicyAssignment instantiates a new unifiedRoleManagementPolicyAssignment and sets the default values.
func NewUnifiedRoleManagementPolicyAssignment()(*UnifiedRoleManagementPolicyAssignment) {
    m := &UnifiedRoleManagementPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(UnifiedRoleManagementPolicyable))
        }
        return nil
    }
    res["policyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
    res["scopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeId(val)
        }
        return nil
    }
    res["scopeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeType(val)
        }
        return nil
    }
    return res
}
// GetPolicy gets the policy property value. The policy for the assignment.
func (m *UnifiedRoleManagementPolicyAssignment) GetPolicy()(UnifiedRoleManagementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// GetPolicyId gets the policyId property value. The id of the policy.
func (m *UnifiedRoleManagementPolicyAssignment) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// GetRoleDefinitionId gets the roleDefinitionId property value. The id of the role definition where the policy applies. If not specified, the policy applies to all roles.
func (m *UnifiedRoleManagementPolicyAssignment) GetRoleDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionId
    }
}
// GetScopeId gets the scopeId property value. The id of the scope where the policy is assigned. E.g. '/', groupId, etc.
func (m *UnifiedRoleManagementPolicyAssignment) GetScopeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeId
    }
}
// GetScopeType gets the scopeType property value. The type of the scope where the policy is assigned. One of Directory, DirectoryRole, Group.
func (m *UnifiedRoleManagementPolicyAssignment) GetScopeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeType
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeId", m.GetScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeType", m.GetScopeType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPolicy sets the policy property value. The policy for the assignment.
func (m *UnifiedRoleManagementPolicyAssignment) SetPolicy(value UnifiedRoleManagementPolicyable)() {
    if m != nil {
        m.policy = value
    }
}
// SetPolicyId sets the policyId property value. The id of the policy.
func (m *UnifiedRoleManagementPolicyAssignment) SetPolicyId(value *string)() {
    if m != nil {
        m.policyId = value
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. The id of the role definition where the policy applies. If not specified, the policy applies to all roles.
func (m *UnifiedRoleManagementPolicyAssignment) SetRoleDefinitionId(value *string)() {
    if m != nil {
        m.roleDefinitionId = value
    }
}
// SetScopeId sets the scopeId property value. The id of the scope where the policy is assigned. E.g. '/', groupId, etc.
func (m *UnifiedRoleManagementPolicyAssignment) SetScopeId(value *string)() {
    if m != nil {
        m.scopeId = value
    }
}
// SetScopeType sets the scopeType property value. The type of the scope where the policy is assigned. One of Directory, DirectoryRole, Group.
func (m *UnifiedRoleManagementPolicyAssignment) SetScopeType(value *string)() {
    if m != nil {
        m.scopeType = value
    }
}
