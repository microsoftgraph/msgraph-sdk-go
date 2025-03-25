package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsItemMembersAddpostResponseable instead.
type ItemChannelsItemMembersAddResponse struct {
    ItemChannelsItemMembersAddpostResponse
}
// NewItemChannelsItemMembersAddResponse instantiates a new ItemChannelsItemMembersAddResponse and sets the default values.
func NewItemChannelsItemMembersAddResponse()(*ItemChannelsItemMembersAddResponse) {
    m := &ItemChannelsItemMembersAddResponse{
        ItemChannelsItemMembersAddpostResponse: *NewItemChannelsItemMembersAddpostResponse(),
    }
    return m
}
// CreateItemChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsItemMembersAddpostResponseable instead.
type ItemChannelsItemMembersAddResponseable interface {
    ItemChannelsItemMembersAddpostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
