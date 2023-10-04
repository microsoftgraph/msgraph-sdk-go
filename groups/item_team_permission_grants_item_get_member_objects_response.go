package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamPermissionGrantsItemGetMemberObjectsResponse 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemTeamPermissionGrantsItemGetMemberObjectsResponse struct {
    ItemTeamPermissionGrantsItemGetMemberObjectsPostResponse
}
// NewItemTeamPermissionGrantsItemGetMemberObjectsResponse instantiates a new ItemTeamPermissionGrantsItemGetMemberObjectsResponse and sets the default values.
func NewItemTeamPermissionGrantsItemGetMemberObjectsResponse()(*ItemTeamPermissionGrantsItemGetMemberObjectsResponse) {
    m := &ItemTeamPermissionGrantsItemGetMemberObjectsResponse{
        ItemTeamPermissionGrantsItemGetMemberObjectsPostResponse: *NewItemTeamPermissionGrantsItemGetMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemTeamPermissionGrantsItemGetMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamPermissionGrantsItemGetMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPermissionGrantsItemGetMemberObjectsResponse(), nil
}
// ItemTeamPermissionGrantsItemGetMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemTeamPermissionGrantsItemGetMemberObjectsResponseable interface {
    ItemTeamPermissionGrantsItemGetMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
