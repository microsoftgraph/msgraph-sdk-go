package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemChannelsItemAllMembersRemovepostResponseable instead.
type ItemJoinedTeamsItemChannelsItemAllMembersRemoveResponse struct {
    ItemJoinedTeamsItemChannelsItemAllMembersRemovepostResponse
}
// NewItemJoinedTeamsItemChannelsItemAllMembersRemoveResponse instantiates a new ItemJoinedTeamsItemChannelsItemAllMembersRemoveResponse and sets the default values.
func NewItemJoinedTeamsItemChannelsItemAllMembersRemoveResponse()(*ItemJoinedTeamsItemChannelsItemAllMembersRemoveResponse) {
    m := &ItemJoinedTeamsItemChannelsItemAllMembersRemoveResponse{
        ItemJoinedTeamsItemChannelsItemAllMembersRemovepostResponse: *NewItemJoinedTeamsItemChannelsItemAllMembersRemovepostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemChannelsItemAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedTeamsItemChannelsItemAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemChannelsItemAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemChannelsItemAllMembersRemovepostResponseable instead.
type ItemJoinedTeamsItemChannelsItemAllMembersRemoveResponseable interface {
    ItemJoinedTeamsItemChannelsItemAllMembersRemovepostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
