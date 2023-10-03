package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemChannelsItemMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemJoinedTeamsItemChannelsItemMembersAddResponse struct {
    ItemJoinedTeamsItemChannelsItemMembersAddPostResponse
}
// NewItemJoinedTeamsItemChannelsItemMembersAddResponse instantiates a new ItemJoinedTeamsItemChannelsItemMembersAddResponse and sets the default values.
func NewItemJoinedTeamsItemChannelsItemMembersAddResponse()(*ItemJoinedTeamsItemChannelsItemMembersAddResponse) {
    m := &ItemJoinedTeamsItemChannelsItemMembersAddResponse{
        ItemJoinedTeamsItemChannelsItemMembersAddPostResponse: *NewItemJoinedTeamsItemChannelsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemChannelsItemMembersAddResponse(), nil
}
// ItemJoinedTeamsItemChannelsItemMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemJoinedTeamsItemChannelsItemMembersAddResponseable interface {
    ItemJoinedTeamsItemChannelsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
