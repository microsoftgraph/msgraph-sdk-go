package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemChannelsItemAllMembersAddpostResponseable instead.
type ItemJoinedTeamsItemChannelsItemAllMembersAddResponse struct {
    ItemJoinedTeamsItemChannelsItemAllMembersAddpostResponse
}
// NewItemJoinedTeamsItemChannelsItemAllMembersAddResponse instantiates a new ItemJoinedTeamsItemChannelsItemAllMembersAddResponse and sets the default values.
func NewItemJoinedTeamsItemChannelsItemAllMembersAddResponse()(*ItemJoinedTeamsItemChannelsItemAllMembersAddResponse) {
    m := &ItemJoinedTeamsItemChannelsItemAllMembersAddResponse{
        ItemJoinedTeamsItemChannelsItemAllMembersAddpostResponse: *NewItemJoinedTeamsItemChannelsItemAllMembersAddpostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemChannelsItemAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedTeamsItemChannelsItemAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemChannelsItemAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemChannelsItemAllMembersAddpostResponseable instead.
type ItemJoinedTeamsItemChannelsItemAllMembersAddResponseable interface {
    ItemJoinedTeamsItemChannelsItemAllMembersAddpostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
