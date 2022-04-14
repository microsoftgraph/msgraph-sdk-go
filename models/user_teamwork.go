package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserTeamwork 
type UserTeamwork struct {
    Entity
    // The apps installed in the personal scope of this user.
    installedApps []UserScopeTeamsAppInstallationable
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
// GetFieldDeserializers the deserialization information for the current model
func (m *UserTeamwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["installedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserScopeTeamsAppInstallationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserScopeTeamsAppInstallationable, len(val))
            for i, v := range val {
                res[i] = v.(UserScopeTeamsAppInstallationable)
            }
            m.SetInstalledApps(res)
        }
        return nil
    }
    return res
}
// GetInstalledApps gets the installedApps property value. The apps installed in the personal scope of this user.
func (m *UserTeamwork) GetInstalledApps()([]UserScopeTeamsAppInstallationable) {
    if m == nil {
        return nil
    } else {
        return m.installedApps
    }
}
// Serialize serializes information the current object
func (m *UserTeamwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInstalledApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstalledApps()))
        for i, v := range m.GetInstalledApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("installedApps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInstalledApps sets the installedApps property value. The apps installed in the personal scope of this user.
func (m *UserTeamwork) SetInstalledApps(value []UserScopeTeamsAppInstallationable)() {
    if m != nil {
        m.installedApps = value
    }
}
