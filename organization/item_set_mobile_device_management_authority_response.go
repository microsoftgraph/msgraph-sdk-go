package organization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSetMobileDeviceManagementAuthoritypostResponseable instead.
type ItemSetMobileDeviceManagementAuthorityResponse struct {
    ItemSetMobileDeviceManagementAuthoritypostResponse
}
// NewItemSetMobileDeviceManagementAuthorityResponse instantiates a new ItemSetMobileDeviceManagementAuthorityResponse and sets the default values.
func NewItemSetMobileDeviceManagementAuthorityResponse()(*ItemSetMobileDeviceManagementAuthorityResponse) {
    m := &ItemSetMobileDeviceManagementAuthorityResponse{
        ItemSetMobileDeviceManagementAuthoritypostResponse: *NewItemSetMobileDeviceManagementAuthoritypostResponse(),
    }
    return m
}
// CreateItemSetMobileDeviceManagementAuthorityResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSetMobileDeviceManagementAuthorityResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSetMobileDeviceManagementAuthorityResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSetMobileDeviceManagementAuthoritypostResponseable instead.
type ItemSetMobileDeviceManagementAuthorityResponseable interface {
    ItemSetMobileDeviceManagementAuthoritypostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
