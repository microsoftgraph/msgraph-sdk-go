package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemPrimaryChannelMembersRemovepostResponseable instead.
type ItemJoinedTeamsItemPrimaryChannelMembersRemoveResponse struct {
    ItemJoinedTeamsItemPrimaryChannelMembersRemovepostResponse
}
// NewItemJoinedTeamsItemPrimaryChannelMembersRemoveResponse instantiates a new ItemJoinedTeamsItemPrimaryChannelMembersRemoveResponse and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelMembersRemoveResponse()(*ItemJoinedTeamsItemPrimaryChannelMembersRemoveResponse) {
    m := &ItemJoinedTeamsItemPrimaryChannelMembersRemoveResponse{
        ItemJoinedTeamsItemPrimaryChannelMembersRemovepostResponse: *NewItemJoinedTeamsItemPrimaryChannelMembersRemovepostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemPrimaryChannelMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedTeamsItemPrimaryChannelMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemPrimaryChannelMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemPrimaryChannelMembersRemovepostResponseable instead.
type ItemJoinedTeamsItemPrimaryChannelMembersRemoveResponseable interface {
    ItemJoinedTeamsItemPrimaryChannelMembersRemovepostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
