package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemPermissionGrantsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemJoinedTeamsItemPermissionGrantsDeltaResponse struct {
    ItemJoinedTeamsItemPermissionGrantsDeltaGetResponse
}
// NewItemJoinedTeamsItemPermissionGrantsDeltaResponse instantiates a new ItemJoinedTeamsItemPermissionGrantsDeltaResponse and sets the default values.
func NewItemJoinedTeamsItemPermissionGrantsDeltaResponse()(*ItemJoinedTeamsItemPermissionGrantsDeltaResponse) {
    m := &ItemJoinedTeamsItemPermissionGrantsDeltaResponse{
        ItemJoinedTeamsItemPermissionGrantsDeltaGetResponse: *NewItemJoinedTeamsItemPermissionGrantsDeltaGetResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemPermissionGrantsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemPermissionGrantsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemPermissionGrantsDeltaResponse(), nil
}
// ItemJoinedTeamsItemPermissionGrantsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemJoinedTeamsItemPermissionGrantsDeltaResponseable interface {
    ItemJoinedTeamsItemPermissionGrantsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
