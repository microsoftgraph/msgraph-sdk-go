package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamPermissionGrantsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamPermissionGrantsDeltaResponse struct {
    ItemTeamPermissionGrantsDeltaGetResponse
}
// NewItemTeamPermissionGrantsDeltaResponse instantiates a new ItemTeamPermissionGrantsDeltaResponse and sets the default values.
func NewItemTeamPermissionGrantsDeltaResponse()(*ItemTeamPermissionGrantsDeltaResponse) {
    m := &ItemTeamPermissionGrantsDeltaResponse{
        ItemTeamPermissionGrantsDeltaGetResponse: *NewItemTeamPermissionGrantsDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamPermissionGrantsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamPermissionGrantsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPermissionGrantsDeltaResponse(), nil
}
// ItemTeamPermissionGrantsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamPermissionGrantsDeltaResponseable interface {
    ItemTeamPermissionGrantsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
