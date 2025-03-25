package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsItemAllMembersRemovepostResponseable instead.
type ItemChannelsItemAllMembersRemoveResponse struct {
    ItemChannelsItemAllMembersRemovepostResponse
}
// NewItemChannelsItemAllMembersRemoveResponse instantiates a new ItemChannelsItemAllMembersRemoveResponse and sets the default values.
func NewItemChannelsItemAllMembersRemoveResponse()(*ItemChannelsItemAllMembersRemoveResponse) {
    m := &ItemChannelsItemAllMembersRemoveResponse{
        ItemChannelsItemAllMembersRemovepostResponse: *NewItemChannelsItemAllMembersRemovepostResponse(),
    }
    return m
}
// CreateItemChannelsItemAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsItemAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsItemAllMembersRemovepostResponseable instead.
type ItemChannelsItemAllMembersRemoveResponseable interface {
    ItemChannelsItemAllMembersRemovepostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
