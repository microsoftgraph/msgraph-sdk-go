package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedteamsItemMembersAddPostResponseable instead.
type ItemJoinedteamsItemMembersAddResponse struct {
    ItemJoinedteamsItemMembersAddPostResponse
}
// NewItemJoinedteamsItemMembersAddResponse instantiates a new ItemJoinedteamsItemMembersAddResponse and sets the default values.
func NewItemJoinedteamsItemMembersAddResponse()(*ItemJoinedteamsItemMembersAddResponse) {
    m := &ItemJoinedteamsItemMembersAddResponse{
        ItemJoinedteamsItemMembersAddPostResponse: *NewItemJoinedteamsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemJoinedteamsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedteamsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedteamsItemMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedteamsItemMembersAddPostResponseable instead.
type ItemJoinedteamsItemMembersAddResponseable interface {
    ItemJoinedteamsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
