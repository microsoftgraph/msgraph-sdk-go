package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChatsItemMembersRemovepostResponseable instead.
type ItemChatsItemMembersRemoveResponse struct {
    ItemChatsItemMembersRemovepostResponse
}
// NewItemChatsItemMembersRemoveResponse instantiates a new ItemChatsItemMembersRemoveResponse and sets the default values.
func NewItemChatsItemMembersRemoveResponse()(*ItemChatsItemMembersRemoveResponse) {
    m := &ItemChatsItemMembersRemoveResponse{
        ItemChatsItemMembersRemovepostResponse: *NewItemChatsItemMembersRemovepostResponse(),
    }
    return m
}
// CreateItemChatsItemMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChatsItemMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChatsItemMembersRemovepostResponseable instead.
type ItemChatsItemMembersRemoveResponseable interface {
    ItemChatsItemMembersRemovepostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
