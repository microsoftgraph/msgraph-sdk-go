package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemPermissionGrantsGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemJoinedTeamsItemPermissionGrantsGetByIdsResponse struct {
    ItemJoinedTeamsItemPermissionGrantsGetByIdsPostResponse
}
// NewItemJoinedTeamsItemPermissionGrantsGetByIdsResponse instantiates a new ItemJoinedTeamsItemPermissionGrantsGetByIdsResponse and sets the default values.
func NewItemJoinedTeamsItemPermissionGrantsGetByIdsResponse()(*ItemJoinedTeamsItemPermissionGrantsGetByIdsResponse) {
    m := &ItemJoinedTeamsItemPermissionGrantsGetByIdsResponse{
        ItemJoinedTeamsItemPermissionGrantsGetByIdsPostResponse: *NewItemJoinedTeamsItemPermissionGrantsGetByIdsPostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemPermissionGrantsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemPermissionGrantsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemPermissionGrantsGetByIdsResponse(), nil
}
// ItemJoinedTeamsItemPermissionGrantsGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemJoinedTeamsItemPermissionGrantsGetByIdsResponseable interface {
    ItemJoinedTeamsItemPermissionGrantsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
