package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChannelsItemMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemChannelsItemMembersAddResponse struct {
    ItemChannelsItemMembersAddPostResponse
}
// NewItemChannelsItemMembersAddResponse instantiates a new ItemChannelsItemMembersAddResponse and sets the default values.
func NewItemChannelsItemMembersAddResponse()(*ItemChannelsItemMembersAddResponse) {
    m := &ItemChannelsItemMembersAddResponse{
        ItemChannelsItemMembersAddPostResponse: *NewItemChannelsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemMembersAddResponse(), nil
}
// ItemChannelsItemMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemChannelsItemMembersAddResponseable interface {
    ItemChannelsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
