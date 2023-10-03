package organization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSetMobileDeviceManagementAuthorityResponse 
// Deprecated: This class is obsolete. Use setMobileDeviceManagementAuthorityPostResponse instead.
type ItemSetMobileDeviceManagementAuthorityResponse struct {
    ItemSetMobileDeviceManagementAuthorityPostResponse
}
// NewItemSetMobileDeviceManagementAuthorityResponse instantiates a new ItemSetMobileDeviceManagementAuthorityResponse and sets the default values.
func NewItemSetMobileDeviceManagementAuthorityResponse()(*ItemSetMobileDeviceManagementAuthorityResponse) {
    m := &ItemSetMobileDeviceManagementAuthorityResponse{
        ItemSetMobileDeviceManagementAuthorityPostResponse: *NewItemSetMobileDeviceManagementAuthorityPostResponse(),
    }
    return m
}
// CreateItemSetMobileDeviceManagementAuthorityResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSetMobileDeviceManagementAuthorityResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSetMobileDeviceManagementAuthorityResponse(), nil
}
// ItemSetMobileDeviceManagementAuthorityResponseable 
// Deprecated: This class is obsolete. Use setMobileDeviceManagementAuthorityPostResponse instead.
type ItemSetMobileDeviceManagementAuthorityResponseable interface {
    ItemSetMobileDeviceManagementAuthorityPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
