package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemChatsItemMembersAddResponse struct {
    ItemChatsItemMembersAddPostResponse
}
// NewItemChatsItemMembersAddResponse instantiates a new ItemChatsItemMembersAddResponse and sets the default values.
func NewItemChatsItemMembersAddResponse()(*ItemChatsItemMembersAddResponse) {
    m := &ItemChatsItemMembersAddResponse{
        ItemChatsItemMembersAddPostResponse: *NewItemChatsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemChatsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemMembersAddResponse(), nil
}
// ItemChatsItemMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemChatsItemMembersAddResponseable interface {
    ItemChatsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
