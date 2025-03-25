package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsItemMembersAddpostResponseable instead.
type ItemTeamChannelsItemMembersAddResponse struct {
    ItemTeamChannelsItemMembersAddpostResponse
}
// NewItemTeamChannelsItemMembersAddResponse instantiates a new ItemTeamChannelsItemMembersAddResponse and sets the default values.
func NewItemTeamChannelsItemMembersAddResponse()(*ItemTeamChannelsItemMembersAddResponse) {
    m := &ItemTeamChannelsItemMembersAddResponse{
        ItemTeamChannelsItemMembersAddpostResponse: *NewItemTeamChannelsItemMembersAddpostResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsItemMembersAddpostResponseable instead.
type ItemTeamChannelsItemMembersAddResponseable interface {
    ItemTeamChannelsItemMembersAddpostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
