package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeJoinedTeamsItemPrimaryChannelMembersAddResponse provides operations to call the add method.
type MeJoinedTeamsItemPrimaryChannelMembersAddResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ActionResultPartable
}
// NewMeJoinedTeamsItemPrimaryChannelMembersAddResponse instantiates a new MeJoinedTeamsItemPrimaryChannelMembersAddResponse and sets the default values.
func NewMeJoinedTeamsItemPrimaryChannelMembersAddResponse()(*MeJoinedTeamsItemPrimaryChannelMembersAddResponse) {
    m := &MeJoinedTeamsItemPrimaryChannelMembersAddResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMeJoinedTeamsItemPrimaryChannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeJoinedTeamsItemPrimaryChannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeJoinedTeamsItemPrimaryChannelMembersAddResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeJoinedTeamsItemPrimaryChannelMembersAddResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateActionResultPartFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *MeJoinedTeamsItemPrimaryChannelMembersAddResponse) GetValue()([]ActionResultPartable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeJoinedTeamsItemPrimaryChannelMembersAddResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MeJoinedTeamsItemPrimaryChannelMembersAddResponse) SetValue(value []ActionResultPartable)() {
    m.value = value
}
