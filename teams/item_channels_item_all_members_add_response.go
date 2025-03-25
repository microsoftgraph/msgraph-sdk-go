package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsItemAllMembersAddpostResponseable instead.
type ItemChannelsItemAllMembersAddResponse struct {
    ItemChannelsItemAllMembersAddpostResponse
}
// NewItemChannelsItemAllMembersAddResponse instantiates a new ItemChannelsItemAllMembersAddResponse and sets the default values.
func NewItemChannelsItemAllMembersAddResponse()(*ItemChannelsItemAllMembersAddResponse) {
    m := &ItemChannelsItemAllMembersAddResponse{
        ItemChannelsItemAllMembersAddpostResponse: *NewItemChannelsItemAllMembersAddpostResponse(),
    }
    return m
}
// CreateItemChannelsItemAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsItemAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsItemAllMembersAddpostResponseable instead.
type ItemChannelsItemAllMembersAddResponseable interface {
    ItemChannelsItemAllMembersAddpostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
