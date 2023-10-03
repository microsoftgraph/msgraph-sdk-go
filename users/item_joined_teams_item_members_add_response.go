package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemJoinedTeamsItemMembersAddResponse struct {
    ItemJoinedTeamsItemMembersAddPostResponse
}
// NewItemJoinedTeamsItemMembersAddResponse instantiates a new ItemJoinedTeamsItemMembersAddResponse and sets the default values.
func NewItemJoinedTeamsItemMembersAddResponse()(*ItemJoinedTeamsItemMembersAddResponse) {
    m := &ItemJoinedTeamsItemMembersAddResponse{
        ItemJoinedTeamsItemMembersAddPostResponse: *NewItemJoinedTeamsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemMembersAddResponse(), nil
}
// ItemJoinedTeamsItemMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemJoinedTeamsItemMembersAddResponseable interface {
    ItemJoinedTeamsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
