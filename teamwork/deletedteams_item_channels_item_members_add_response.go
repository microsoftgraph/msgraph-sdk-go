package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedteamsItemChannelsItemMembersAddPostResponseable instead.
type DeletedteamsItemChannelsItemMembersAddResponse struct {
    DeletedteamsItemChannelsItemMembersAddPostResponse
}
// NewDeletedteamsItemChannelsItemMembersAddResponse instantiates a new DeletedteamsItemChannelsItemMembersAddResponse and sets the default values.
func NewDeletedteamsItemChannelsItemMembersAddResponse()(*DeletedteamsItemChannelsItemMembersAddResponse) {
    m := &DeletedteamsItemChannelsItemMembersAddResponse{
        DeletedteamsItemChannelsItemMembersAddPostResponse: *NewDeletedteamsItemChannelsItemMembersAddPostResponse(),
    }
    return m
}
// CreateDeletedteamsItemChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedteamsItemChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedteamsItemChannelsItemMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedteamsItemChannelsItemMembersAddPostResponseable instead.
type DeletedteamsItemChannelsItemMembersAddResponseable interface {
    DeletedteamsItemChannelsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
