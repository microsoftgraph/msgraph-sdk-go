package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamPermissionGrantsItemCheckMemberGroupsResponse 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemTeamPermissionGrantsItemCheckMemberGroupsResponse struct {
    ItemTeamPermissionGrantsItemCheckMemberGroupsPostResponse
}
// NewItemTeamPermissionGrantsItemCheckMemberGroupsResponse instantiates a new ItemTeamPermissionGrantsItemCheckMemberGroupsResponse and sets the default values.
func NewItemTeamPermissionGrantsItemCheckMemberGroupsResponse()(*ItemTeamPermissionGrantsItemCheckMemberGroupsResponse) {
    m := &ItemTeamPermissionGrantsItemCheckMemberGroupsResponse{
        ItemTeamPermissionGrantsItemCheckMemberGroupsPostResponse: *NewItemTeamPermissionGrantsItemCheckMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemTeamPermissionGrantsItemCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamPermissionGrantsItemCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPermissionGrantsItemCheckMemberGroupsResponse(), nil
}
// ItemTeamPermissionGrantsItemCheckMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemTeamPermissionGrantsItemCheckMemberGroupsResponseable interface {
    ItemTeamPermissionGrantsItemCheckMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
