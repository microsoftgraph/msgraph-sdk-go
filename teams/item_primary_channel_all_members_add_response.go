package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimaryChannelAllMembersAddpostResponseable instead.
type ItemPrimaryChannelAllMembersAddResponse struct {
    ItemPrimaryChannelAllMembersAddpostResponse
}
// NewItemPrimaryChannelAllMembersAddResponse instantiates a new ItemPrimaryChannelAllMembersAddResponse and sets the default values.
func NewItemPrimaryChannelAllMembersAddResponse()(*ItemPrimaryChannelAllMembersAddResponse) {
    m := &ItemPrimaryChannelAllMembersAddResponse{
        ItemPrimaryChannelAllMembersAddpostResponse: *NewItemPrimaryChannelAllMembersAddpostResponse(),
    }
    return m
}
// CreateItemPrimaryChannelAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimaryChannelAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimaryChannelAllMembersAddpostResponseable instead.
type ItemPrimaryChannelAllMembersAddResponseable interface {
    ItemPrimaryChannelAllMembersAddpostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
