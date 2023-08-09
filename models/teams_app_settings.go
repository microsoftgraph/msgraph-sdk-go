package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppSettings 
type TeamsAppSettings struct {
    Entity
}
// NewTeamsAppSettings instantiates a new teamsAppSettings and sets the default values.
func NewTeamsAppSettings()(*TeamsAppSettings) {
    m := &TeamsAppSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppSettings(), nil
}
// GetAllowUserRequestsForAppAccess gets the allowUserRequestsForAppAccess property value. The allowUserRequestsForAppAccess property
func (m *TeamsAppSettings) GetAllowUserRequestsForAppAccess()(*bool) {
    val, err := m.GetBackingStore().Get("allowUserRequestsForAppAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowUserRequestsForAppAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUserRequestsForAppAccess(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeamsAppSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowUserRequestsForAppAccess", m.GetAllowUserRequestsForAppAccess())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowUserRequestsForAppAccess sets the allowUserRequestsForAppAccess property value. The allowUserRequestsForAppAccess property
func (m *TeamsAppSettings) SetAllowUserRequestsForAppAccess(value *bool)() {
    err := m.GetBackingStore().Set("allowUserRequestsForAppAccess", value)
    if err != nil {
        panic(err)
    }
}
// TeamsAppSettingsable 
type TeamsAppSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowUserRequestsForAppAccess()(*bool)
    SetAllowUserRequestsForAppAccess(value *bool)()
}
