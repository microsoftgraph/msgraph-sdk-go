package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeviceconfigurationsItemAssignPostResponseable instead.
type DeviceconfigurationsItemAssignResponse struct {
    DeviceconfigurationsItemAssignPostResponse
}
// NewDeviceconfigurationsItemAssignResponse instantiates a new DeviceconfigurationsItemAssignResponse and sets the default values.
func NewDeviceconfigurationsItemAssignResponse()(*DeviceconfigurationsItemAssignResponse) {
    m := &DeviceconfigurationsItemAssignResponse{
        DeviceconfigurationsItemAssignPostResponse: *NewDeviceconfigurationsItemAssignPostResponse(),
    }
    return m
}
// CreateDeviceconfigurationsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceconfigurationsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceconfigurationsItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use DeviceconfigurationsItemAssignPostResponseable instead.
type DeviceconfigurationsItemAssignResponseable interface {
    DeviceconfigurationsItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
