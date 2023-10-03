package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetManagedDevicesWithAppFailuresResponse 
// Deprecated: This class is obsolete. Use getManagedDevicesWithAppFailuresGetResponse instead.
type ItemGetManagedDevicesWithAppFailuresResponse struct {
    ItemGetManagedDevicesWithAppFailuresGetResponse
}
// NewItemGetManagedDevicesWithAppFailuresResponse instantiates a new ItemGetManagedDevicesWithAppFailuresResponse and sets the default values.
func NewItemGetManagedDevicesWithAppFailuresResponse()(*ItemGetManagedDevicesWithAppFailuresResponse) {
    m := &ItemGetManagedDevicesWithAppFailuresResponse{
        ItemGetManagedDevicesWithAppFailuresGetResponse: *NewItemGetManagedDevicesWithAppFailuresGetResponse(),
    }
    return m
}
// CreateItemGetManagedDevicesWithAppFailuresResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetManagedDevicesWithAppFailuresResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetManagedDevicesWithAppFailuresResponse(), nil
}
// ItemGetManagedDevicesWithAppFailuresResponseable 
// Deprecated: This class is obsolete. Use getManagedDevicesWithAppFailuresGetResponse instead.
type ItemGetManagedDevicesWithAppFailuresResponseable interface {
    ItemGetManagedDevicesWithAppFailuresGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
