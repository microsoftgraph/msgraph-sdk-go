package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamRenamedEventMessageDetail 
type TeamRenamedEventMessageDetail struct {
    EventMessageDetail
    // Initiator of the event.
    initiator IdentitySetable
    // The updated name of the team.
    teamDisplayName *string
    // Unique identifier of the team.
    teamId *string
}
// NewTeamRenamedEventMessageDetail instantiates a new TeamRenamedEventMessageDetail and sets the default values.
func NewTeamRenamedEventMessageDetail()(*TeamRenamedEventMessageDetail) {
    m := &TeamRenamedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    return m
}
// CreateTeamRenamedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamRenamedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamRenamedEventMessageDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamRenamedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["initiator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
        }
        return nil
    }
    res["teamDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamDisplayName(val)
        }
        return nil
    }
    res["teamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *TeamRenamedEventMessageDetail) GetInitiator()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.initiator
    }
}
// GetTeamDisplayName gets the teamDisplayName property value. The updated name of the team.
func (m *TeamRenamedEventMessageDetail) GetTeamDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamDisplayName
    }
}
// GetTeamId gets the teamId property value. Unique identifier of the team.
func (m *TeamRenamedEventMessageDetail) GetTeamId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamId
    }
}
// Serialize serializes information the current object
func (m *TeamRenamedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamDisplayName", m.GetTeamDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *TeamRenamedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    if m != nil {
        m.initiator = value
    }
}
// SetTeamDisplayName sets the teamDisplayName property value. The updated name of the team.
func (m *TeamRenamedEventMessageDetail) SetTeamDisplayName(value *string)() {
    if m != nil {
        m.teamDisplayName = value
    }
}
// SetTeamId sets the teamId property value. Unique identifier of the team.
func (m *TeamRenamedEventMessageDetail) SetTeamId(value *string)() {
    if m != nil {
        m.teamId = value
    }
}
