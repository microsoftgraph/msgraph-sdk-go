package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Teamwork struct {
    Entity
}
// NewTeamwork instantiates a new Teamwork and sets the default values.
func NewTeamwork()(*Teamwork) {
    m := &Teamwork{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamwork(), nil
}
// GetDeletedTeams gets the deletedTeams property value. The deleted team.
// returns a []DeletedTeamable when successful
func (m *Teamwork) GetDeletedTeams()([]DeletedTeamable) {
    val, err := m.GetBackingStore().Get("deletedTeams")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeletedTeamable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Teamwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedTeams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeletedTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeletedTeamable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeletedTeamable)
                }
            }
            m.SetDeletedTeams(res)
        }
        return nil
    }
    res["teamsAppSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppSettings(val.(TeamsAppSettingsable))
        }
        return nil
    }
    res["workforceIntegrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkforceIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkforceIntegrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkforceIntegrationable)
                }
            }
            m.SetWorkforceIntegrations(res)
        }
        return nil
    }
    return res
}
// GetTeamsAppSettings gets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
// returns a TeamsAppSettingsable when successful
func (m *Teamwork) GetTeamsAppSettings()(TeamsAppSettingsable) {
    val, err := m.GetBackingStore().Get("teamsAppSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamsAppSettingsable)
    }
    return nil
}
// GetWorkforceIntegrations gets the workforceIntegrations property value. The workforceIntegrations property
// returns a []WorkforceIntegrationable when successful
func (m *Teamwork) GetWorkforceIntegrations()([]WorkforceIntegrationable) {
    val, err := m.GetBackingStore().Get("workforceIntegrations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkforceIntegrationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Teamwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeletedTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedTeams()))
        for i, v := range m.GetDeletedTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deletedTeams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsAppSettings", m.GetTeamsAppSettings())
        if err != nil {
            return err
        }
    }
    if m.GetWorkforceIntegrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkforceIntegrations()))
        for i, v := range m.GetWorkforceIntegrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("workforceIntegrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedTeams sets the deletedTeams property value. The deleted team.
func (m *Teamwork) SetDeletedTeams(value []DeletedTeamable)() {
    err := m.GetBackingStore().Set("deletedTeams", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsAppSettings sets the teamsAppSettings property value. Represents tenant-wide settings for all Teams apps in the tenant.
func (m *Teamwork) SetTeamsAppSettings(value TeamsAppSettingsable)() {
    err := m.GetBackingStore().Set("teamsAppSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkforceIntegrations sets the workforceIntegrations property value. The workforceIntegrations property
func (m *Teamwork) SetWorkforceIntegrations(value []WorkforceIntegrationable)() {
    err := m.GetBackingStore().Set("workforceIntegrations", value)
    if err != nil {
        panic(err)
    }
}
type Teamworkable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeletedTeams()([]DeletedTeamable)
    GetTeamsAppSettings()(TeamsAppSettingsable)
    GetWorkforceIntegrations()([]WorkforceIntegrationable)
    SetDeletedTeams(value []DeletedTeamable)()
    SetTeamsAppSettings(value TeamsAppSettingsable)()
    SetWorkforceIntegrations(value []WorkforceIntegrationable)()
}
