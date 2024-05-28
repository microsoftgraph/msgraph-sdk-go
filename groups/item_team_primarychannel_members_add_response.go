package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamPrimarychannelMembersAddPostResponseable instead.
type ItemTeamPrimarychannelMembersAddResponse struct {
    ItemTeamPrimarychannelMembersAddPostResponse
}
// NewItemTeamPrimarychannelMembersAddResponse instantiates a new ItemTeamPrimarychannelMembersAddResponse and sets the default values.
func NewItemTeamPrimarychannelMembersAddResponse()(*ItemTeamPrimarychannelMembersAddResponse) {
    m := &ItemTeamPrimarychannelMembersAddResponse{
        ItemTeamPrimarychannelMembersAddPostResponse: *NewItemTeamPrimarychannelMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamPrimarychannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamPrimarychannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimarychannelMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamPrimarychannelMembersAddPostResponseable instead.
type ItemTeamPrimarychannelMembersAddResponseable interface {
    ItemTeamPrimarychannelMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
