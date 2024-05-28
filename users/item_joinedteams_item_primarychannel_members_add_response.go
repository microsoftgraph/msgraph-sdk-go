package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedteamsItemPrimarychannelMembersAddPostResponseable instead.
type ItemJoinedteamsItemPrimarychannelMembersAddResponse struct {
    ItemJoinedteamsItemPrimarychannelMembersAddPostResponse
}
// NewItemJoinedteamsItemPrimarychannelMembersAddResponse instantiates a new ItemJoinedteamsItemPrimarychannelMembersAddResponse and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMembersAddResponse()(*ItemJoinedteamsItemPrimarychannelMembersAddResponse) {
    m := &ItemJoinedteamsItemPrimarychannelMembersAddResponse{
        ItemJoinedteamsItemPrimarychannelMembersAddPostResponse: *NewItemJoinedteamsItemPrimarychannelMembersAddPostResponse(),
    }
    return m
}
// CreateItemJoinedteamsItemPrimarychannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedteamsItemPrimarychannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedteamsItemPrimarychannelMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedteamsItemPrimarychannelMembersAddPostResponseable instead.
type ItemJoinedteamsItemPrimarychannelMembersAddResponseable interface {
    ItemJoinedteamsItemPrimarychannelMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
