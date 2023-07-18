package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyApprovalRule 
type UnifiedRoleManagementPolicyApprovalRule struct {
    UnifiedRoleManagementPolicyRule
}
// NewUnifiedRoleManagementPolicyApprovalRule instantiates a new unifiedRoleManagementPolicyApprovalRule and sets the default values.
func NewUnifiedRoleManagementPolicyApprovalRule()(*UnifiedRoleManagementPolicyApprovalRule) {
    m := &UnifiedRoleManagementPolicyApprovalRule{
        UnifiedRoleManagementPolicyRule: *NewUnifiedRoleManagementPolicyRule(),
    }
    odataTypeValue := "#microsoft.graph.unifiedRoleManagementPolicyApprovalRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUnifiedRoleManagementPolicyApprovalRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyApprovalRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyApprovalRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyApprovalRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementPolicyRule.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val.(ApprovalSettingsable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UnifiedRoleManagementPolicyApprovalRule) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSetting gets the setting property value. The settings for approval of the role assignment.
func (m *UnifiedRoleManagementPolicyApprovalRule) GetSetting()(ApprovalSettingsable) {
    val, err := m.GetBackingStore().Get("setting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApprovalSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyApprovalRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementPolicyRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UnifiedRoleManagementPolicyApprovalRule) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSetting sets the setting property value. The settings for approval of the role assignment.
func (m *UnifiedRoleManagementPolicyApprovalRule) SetSetting(value ApprovalSettingsable)() {
    err := m.GetBackingStore().Set("setting", value)
    if err != nil {
        panic(err)
    }
}
// UnifiedRoleManagementPolicyApprovalRuleable 
type UnifiedRoleManagementPolicyApprovalRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementPolicyRuleable
    GetOdataType()(*string)
    GetSetting()(ApprovalSettingsable)
    SetOdataType(value *string)()
    SetSetting(value ApprovalSettingsable)()
}
