package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEndpoint struct {
    Entity
}
// NewVirtualEndpoint instantiates a new VirtualEndpoint and sets the default values.
func NewVirtualEndpoint()(*VirtualEndpoint) {
    m := &VirtualEndpoint{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpoint(), nil
}
// GetAuditEvents gets the auditEvents property value. The auditEvents property
// returns a []CloudPcAuditEventable when successful
func (m *VirtualEndpoint) GetAuditEvents()([]CloudPcAuditEventable) {
    val, err := m.GetBackingStore().Get("auditEvents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcAuditEventable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["auditEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcAuditEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcAuditEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcAuditEventable)
                }
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["provisioningPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcProvisioningPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcProvisioningPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcProvisioningPolicyable)
                }
            }
            m.SetProvisioningPolicies(res)
        }
        return nil
    }
    res["userSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcUserSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcUserSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcUserSettingable)
                }
            }
            m.SetUserSettings(res)
        }
        return nil
    }
    return res
}
// GetProvisioningPolicies gets the provisioningPolicies property value. The provisioningPolicies property
// returns a []CloudPcProvisioningPolicyable when successful
func (m *VirtualEndpoint) GetProvisioningPolicies()([]CloudPcProvisioningPolicyable) {
    val, err := m.GetBackingStore().Get("provisioningPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcProvisioningPolicyable)
    }
    return nil
}
// GetUserSettings gets the userSettings property value. The userSettings property
// returns a []CloudPcUserSettingable when successful
func (m *VirtualEndpoint) GetUserSettings()([]CloudPcUserSettingable) {
    val, err := m.GetBackingStore().Get("userSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcUserSettingable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuditEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisioningPolicies()))
        for i, v := range m.GetProvisioningPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("provisioningPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSettings()))
        for i, v := range m.GetUserSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuditEvents sets the auditEvents property value. The auditEvents property
func (m *VirtualEndpoint) SetAuditEvents(value []CloudPcAuditEventable)() {
    err := m.GetBackingStore().Set("auditEvents", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningPolicies sets the provisioningPolicies property value. The provisioningPolicies property
func (m *VirtualEndpoint) SetProvisioningPolicies(value []CloudPcProvisioningPolicyable)() {
    err := m.GetBackingStore().Set("provisioningPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetUserSettings sets the userSettings property value. The userSettings property
func (m *VirtualEndpoint) SetUserSettings(value []CloudPcUserSettingable)() {
    err := m.GetBackingStore().Set("userSettings", value)
    if err != nil {
        panic(err)
    }
}
type VirtualEndpointable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuditEvents()([]CloudPcAuditEventable)
    GetProvisioningPolicies()([]CloudPcProvisioningPolicyable)
    GetUserSettings()([]CloudPcUserSettingable)
    SetAuditEvents(value []CloudPcAuditEventable)()
    SetProvisioningPolicies(value []CloudPcProvisioningPolicyable)()
    SetUserSettings(value []CloudPcUserSettingable)()
}
