package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamPermissionGrantsItemCheckMemberObjectsResponse 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type ItemTeamPermissionGrantsItemCheckMemberObjectsResponse struct {
    ItemTeamPermissionGrantsItemCheckMemberObjectsPostResponse
}
// NewItemTeamPermissionGrantsItemCheckMemberObjectsResponse instantiates a new ItemTeamPermissionGrantsItemCheckMemberObjectsResponse and sets the default values.
func NewItemTeamPermissionGrantsItemCheckMemberObjectsResponse()(*ItemTeamPermissionGrantsItemCheckMemberObjectsResponse) {
    m := &ItemTeamPermissionGrantsItemCheckMemberObjectsResponse{
        ItemTeamPermissionGrantsItemCheckMemberObjectsPostResponse: *NewItemTeamPermissionGrantsItemCheckMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemTeamPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPermissionGrantsItemCheckMemberObjectsResponse(), nil
}
// ItemTeamPermissionGrantsItemCheckMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type ItemTeamPermissionGrantsItemCheckMemberObjectsResponseable interface {
    ItemTeamPermissionGrantsItemCheckMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
