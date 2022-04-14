package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyRule 
type UnifiedRoleManagementPolicyRule struct {
    Entity
    // The target for the policy rule.
    target UnifiedRoleManagementPolicyRuleTargetable
}
// NewUnifiedRoleManagementPolicyRule instantiates a new unifiedRoleManagementPolicyRule and sets the default values.
func NewUnifiedRoleManagementPolicyRule()(*UnifiedRoleManagementPolicyRule) {
    m := &UnifiedRoleManagementPolicyRule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleManagementPolicyRuleTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(UnifiedRoleManagementPolicyRuleTargetable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The target for the policy rule.
func (m *UnifiedRoleManagementPolicyRule) GetTarget()(UnifiedRoleManagementPolicyRuleTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTarget sets the target property value. The target for the policy rule.
func (m *UnifiedRoleManagementPolicyRule) SetTarget(value UnifiedRoleManagementPolicyRuleTargetable)() {
    if m != nil {
        m.target = value
    }
}
