package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimarychannelMembersAddPostResponseable instead.
type ItemPrimarychannelMembersAddResponse struct {
    ItemPrimarychannelMembersAddPostResponse
}
// NewItemPrimarychannelMembersAddResponse instantiates a new ItemPrimarychannelMembersAddResponse and sets the default values.
func NewItemPrimarychannelMembersAddResponse()(*ItemPrimarychannelMembersAddResponse) {
    m := &ItemPrimarychannelMembersAddResponse{
        ItemPrimarychannelMembersAddPostResponse: *NewItemPrimarychannelMembersAddPostResponse(),
    }
    return m
}
// CreateItemPrimarychannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimarychannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimarychannelMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimarychannelMembersAddPostResponseable instead.
type ItemPrimarychannelMembersAddResponseable interface {
    ItemPrimarychannelMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
