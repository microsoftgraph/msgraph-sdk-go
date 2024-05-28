package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedteamsItemChannelsItemMembersAddPostResponseable instead.
type ItemJoinedteamsItemChannelsItemMembersAddResponse struct {
    ItemJoinedteamsItemChannelsItemMembersAddPostResponse
}
// NewItemJoinedteamsItemChannelsItemMembersAddResponse instantiates a new ItemJoinedteamsItemChannelsItemMembersAddResponse and sets the default values.
func NewItemJoinedteamsItemChannelsItemMembersAddResponse()(*ItemJoinedteamsItemChannelsItemMembersAddResponse) {
    m := &ItemJoinedteamsItemChannelsItemMembersAddResponse{
        ItemJoinedteamsItemChannelsItemMembersAddPostResponse: *NewItemJoinedteamsItemChannelsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemJoinedteamsItemChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedteamsItemChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedteamsItemChannelsItemMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedteamsItemChannelsItemMembersAddPostResponseable instead.
type ItemJoinedteamsItemChannelsItemMembersAddResponseable interface {
    ItemJoinedteamsItemChannelsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
