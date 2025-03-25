package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimaryChannelAllMembersRemovepostResponseable instead.
type ItemPrimaryChannelAllMembersRemoveResponse struct {
    ItemPrimaryChannelAllMembersRemovepostResponse
}
// NewItemPrimaryChannelAllMembersRemoveResponse instantiates a new ItemPrimaryChannelAllMembersRemoveResponse and sets the default values.
func NewItemPrimaryChannelAllMembersRemoveResponse()(*ItemPrimaryChannelAllMembersRemoveResponse) {
    m := &ItemPrimaryChannelAllMembersRemoveResponse{
        ItemPrimaryChannelAllMembersRemovepostResponse: *NewItemPrimaryChannelAllMembersRemovepostResponse(),
    }
    return m
}
// CreateItemPrimaryChannelAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimaryChannelAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimaryChannelAllMembersRemovepostResponseable instead.
type ItemPrimaryChannelAllMembersRemoveResponseable interface {
    ItemPrimaryChannelAllMembersRemovepostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
