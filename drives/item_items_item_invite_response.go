package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemInviteResponse 
// Deprecated: This class is obsolete. Use invitePostResponse instead.
type ItemItemsItemInviteResponse struct {
    ItemItemsItemInvitePostResponse
}
// NewItemItemsItemInviteResponse instantiates a new ItemItemsItemInviteResponse and sets the default values.
func NewItemItemsItemInviteResponse()(*ItemItemsItemInviteResponse) {
    m := &ItemItemsItemInviteResponse{
        ItemItemsItemInvitePostResponse: *NewItemItemsItemInvitePostResponse(),
    }
    return m
}
// CreateItemItemsItemInviteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemInviteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemInviteResponse(), nil
}
// ItemItemsItemInviteResponseable 
// Deprecated: This class is obsolete. Use invitePostResponse instead.
type ItemItemsItemInviteResponseable interface {
    ItemItemsItemInvitePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
