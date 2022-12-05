package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeJoinedTeamsItemChannelsGetAllMessagesResponse provides operations to call the getAllMessages method.
type MeJoinedTeamsItemChannelsGetAllMessagesResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ChatMessageable
}
// NewMeJoinedTeamsItemChannelsGetAllMessagesResponse instantiates a new MeJoinedTeamsItemChannelsGetAllMessagesResponse and sets the default values.
func NewMeJoinedTeamsItemChannelsGetAllMessagesResponse()(*MeJoinedTeamsItemChannelsGetAllMessagesResponse) {
    m := &MeJoinedTeamsItemChannelsGetAllMessagesResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMeJoinedTeamsItemChannelsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeJoinedTeamsItemChannelsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeJoinedTeamsItemChannelsGetAllMessagesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeJoinedTeamsItemChannelsGetAllMessagesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MeJoinedTeamsItemChannelsGetAllMessagesResponse) GetValue()([]ChatMessageable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeJoinedTeamsItemChannelsGetAllMessagesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MeJoinedTeamsItemChannelsGetAllMessagesResponse) SetValue(value []ChatMessageable)() {
    m.value = value
}
