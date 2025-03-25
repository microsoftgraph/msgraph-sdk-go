package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemPrimaryChannelAllMembersAddpostResponseable instead.
type ItemJoinedTeamsItemPrimaryChannelAllMembersAddResponse struct {
    ItemJoinedTeamsItemPrimaryChannelAllMembersAddpostResponse
}
// NewItemJoinedTeamsItemPrimaryChannelAllMembersAddResponse instantiates a new ItemJoinedTeamsItemPrimaryChannelAllMembersAddResponse and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelAllMembersAddResponse()(*ItemJoinedTeamsItemPrimaryChannelAllMembersAddResponse) {
    m := &ItemJoinedTeamsItemPrimaryChannelAllMembersAddResponse{
        ItemJoinedTeamsItemPrimaryChannelAllMembersAddpostResponse: *NewItemJoinedTeamsItemPrimaryChannelAllMembersAddpostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemPrimaryChannelAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedTeamsItemPrimaryChannelAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemPrimaryChannelAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemPrimaryChannelAllMembersAddpostResponseable instead.
type ItemJoinedTeamsItemPrimaryChannelAllMembersAddResponseable interface {
    ItemJoinedTeamsItemPrimaryChannelAllMembersAddpostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
