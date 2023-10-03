package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPrimaryChannelMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemPrimaryChannelMembersAddResponse struct {
    ItemPrimaryChannelMembersAddPostResponse
}
// NewItemPrimaryChannelMembersAddResponse instantiates a new ItemPrimaryChannelMembersAddResponse and sets the default values.
func NewItemPrimaryChannelMembersAddResponse()(*ItemPrimaryChannelMembersAddResponse) {
    m := &ItemPrimaryChannelMembersAddResponse{
        ItemPrimaryChannelMembersAddPostResponse: *NewItemPrimaryChannelMembersAddPostResponse(),
    }
    return m
}
// CreateItemPrimaryChannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPrimaryChannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelMembersAddResponse(), nil
}
// ItemPrimaryChannelMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemPrimaryChannelMembersAddResponseable interface {
    ItemPrimaryChannelMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
