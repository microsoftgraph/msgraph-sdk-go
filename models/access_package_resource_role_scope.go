package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceRoleScope 
type AccessPackageResourceRoleScope struct {
    Entity
}
// NewAccessPackageResourceRoleScope instantiates a new accessPackageResourceRoleScope and sets the default values.
func NewAccessPackageResourceRoleScope()(*AccessPackageResourceRoleScope) {
    m := &AccessPackageResourceRoleScope{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageResourceRoleScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceRoleScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceRoleScope(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *AccessPackageResourceRoleScope) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceRoleScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(AccessPackageResourceRoleable))
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(AccessPackageResourceScopeable))
        }
        return nil
    }
    return res
}
// GetRole gets the role property value. The role property
func (m *AccessPackageResourceRoleScope) GetRole()(AccessPackageResourceRoleable) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageResourceRoleable)
    }
    return nil
}
// GetScope gets the scope property value. The scope property
func (m *AccessPackageResourceRoleScope) GetScope()(AccessPackageResourceScopeable) {
    val, err := m.GetBackingStore().Get("scope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageResourceScopeable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageResourceRoleScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *AccessPackageResourceRoleScope) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. The role property
func (m *AccessPackageResourceRoleScope) SetRole(value AccessPackageResourceRoleable)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// SetScope sets the scope property value. The scope property
func (m *AccessPackageResourceRoleScope) SetScope(value AccessPackageResourceScopeable)() {
    err := m.GetBackingStore().Set("scope", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageResourceRoleScopeable 
type AccessPackageResourceRoleScopeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRole()(AccessPackageResourceRoleable)
    GetScope()(AccessPackageResourceScopeable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRole(value AccessPackageResourceRoleable)()
    SetScope(value AccessPackageResourceScopeable)()
}
