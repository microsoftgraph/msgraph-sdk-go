package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemMembersAddResponse struct {
    ItemMembersAddPostResponse
}
// NewItemMembersAddResponse instantiates a new ItemMembersAddResponse and sets the default values.
func NewItemMembersAddResponse()(*ItemMembersAddResponse) {
    m := &ItemMembersAddResponse{
        ItemMembersAddPostResponse: *NewItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMembersAddResponse(), nil
}
// ItemMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemMembersAddResponseable interface {
    ItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
