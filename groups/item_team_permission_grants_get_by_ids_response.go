package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamPermissionGrantsGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemTeamPermissionGrantsGetByIdsResponse struct {
    ItemTeamPermissionGrantsGetByIdsPostResponse
}
// NewItemTeamPermissionGrantsGetByIdsResponse instantiates a new ItemTeamPermissionGrantsGetByIdsResponse and sets the default values.
func NewItemTeamPermissionGrantsGetByIdsResponse()(*ItemTeamPermissionGrantsGetByIdsResponse) {
    m := &ItemTeamPermissionGrantsGetByIdsResponse{
        ItemTeamPermissionGrantsGetByIdsPostResponse: *NewItemTeamPermissionGrantsGetByIdsPostResponse(),
    }
    return m
}
// CreateItemTeamPermissionGrantsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamPermissionGrantsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPermissionGrantsGetByIdsResponse(), nil
}
// ItemTeamPermissionGrantsGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemTeamPermissionGrantsGetByIdsResponseable interface {
    ItemTeamPermissionGrantsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
