package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPermissionGrantsGetAvailableExtensionPropertiesResponse 
// Deprecated: This class is obsolete. Use getAvailableExtensionPropertiesPostResponse instead.
type ItemPermissionGrantsGetAvailableExtensionPropertiesResponse struct {
    ItemPermissionGrantsGetAvailableExtensionPropertiesPostResponse
}
// NewItemPermissionGrantsGetAvailableExtensionPropertiesResponse instantiates a new ItemPermissionGrantsGetAvailableExtensionPropertiesResponse and sets the default values.
func NewItemPermissionGrantsGetAvailableExtensionPropertiesResponse()(*ItemPermissionGrantsGetAvailableExtensionPropertiesResponse) {
    m := &ItemPermissionGrantsGetAvailableExtensionPropertiesResponse{
        ItemPermissionGrantsGetAvailableExtensionPropertiesPostResponse: *NewItemPermissionGrantsGetAvailableExtensionPropertiesPostResponse(),
    }
    return m
}
// CreateItemPermissionGrantsGetAvailableExtensionPropertiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPermissionGrantsGetAvailableExtensionPropertiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPermissionGrantsGetAvailableExtensionPropertiesResponse(), nil
}
// ItemPermissionGrantsGetAvailableExtensionPropertiesResponseable 
// Deprecated: This class is obsolete. Use getAvailableExtensionPropertiesPostResponse instead.
type ItemPermissionGrantsGetAvailableExtensionPropertiesResponseable interface {
    ItemPermissionGrantsGetAvailableExtensionPropertiesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
