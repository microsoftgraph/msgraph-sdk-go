package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserTeamwork 
type UserTeamwork struct {
    Entity
}
// NewUserTeamwork instantiates a new userTeamwork and sets the default values.
func NewUserTeamwork()(*UserTeamwork) {
    m := &UserTeamwork{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserTeamworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserTeamworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserTeamwork(), nil
}
// GetAssociatedTeams gets the associatedTeams property value. The list of associatedTeamInfo objects that a user is associated with.
func (m *UserTeamwork) GetAssociatedTeams()([]AssociatedTeamInfoable) {
    val, err := m.GetBackingStore().Get("associatedTeams")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AssociatedTeamInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserTeamwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["associatedTeams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssociatedTeamInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssociatedTeamInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AssociatedTeamInfoable)
                }
            }
            m.SetAssociatedTeams(res)
        }
        return nil
    }
    res["installedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserScopeTeamsAppInstallationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserScopeTeamsAppInstallationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserScopeTeamsAppInstallationable)
                }
            }
            m.SetInstalledApps(res)
        }
        return nil
    }
    return res
}
// GetInstalledApps gets the installedApps property value. The apps installed in the personal scope of this user.
func (m *UserTeamwork) GetInstalledApps()([]UserScopeTeamsAppInstallationable) {
    val, err := m.GetBackingStore().Get("installedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserScopeTeamsAppInstallationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserTeamwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociatedTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociatedTeams()))
        for i, v := range m.GetAssociatedTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("associatedTeams", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInstalledApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstalledApps()))
        for i, v := range m.GetInstalledApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("installedApps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssociatedTeams sets the associatedTeams property value. The list of associatedTeamInfo objects that a user is associated with.
func (m *UserTeamwork) SetAssociatedTeams(value []AssociatedTeamInfoable)() {
    err := m.GetBackingStore().Set("associatedTeams", value)
    if err != nil {
        panic(err)
    }
}
// SetInstalledApps sets the installedApps property value. The apps installed in the personal scope of this user.
func (m *UserTeamwork) SetInstalledApps(value []UserScopeTeamsAppInstallationable)() {
    err := m.GetBackingStore().Set("installedApps", value)
    if err != nil {
        panic(err)
    }
}
// UserTeamworkable 
type UserTeamworkable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssociatedTeams()([]AssociatedTeamInfoable)
    GetInstalledApps()([]UserScopeTeamsAppInstallationable)
    SetAssociatedTeams(value []AssociatedTeamInfoable)()
    SetInstalledApps(value []UserScopeTeamsAppInstallationable)()
}
