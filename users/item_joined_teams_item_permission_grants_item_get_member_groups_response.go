package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponse 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponse struct {
    ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsPostResponse
}
// NewItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponse instantiates a new ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponse and sets the default values.
func NewItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponse()(*ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponse) {
    m := &ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponse{
        ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsPostResponse: *NewItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponse(), nil
}
// ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsResponseable interface {
    ItemJoinedTeamsItemPermissionGrantsItemGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
