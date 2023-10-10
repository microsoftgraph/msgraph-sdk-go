package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeletedItemsGetAvailableExtensionPropertiesResponse 
// Deprecated: This class is obsolete. Use getAvailableExtensionPropertiesPostResponse instead.
type DeletedItemsGetAvailableExtensionPropertiesResponse struct {
    DeletedItemsGetAvailableExtensionPropertiesPostResponse
}
// NewDeletedItemsGetAvailableExtensionPropertiesResponse instantiates a new DeletedItemsGetAvailableExtensionPropertiesResponse and sets the default values.
func NewDeletedItemsGetAvailableExtensionPropertiesResponse()(*DeletedItemsGetAvailableExtensionPropertiesResponse) {
    m := &DeletedItemsGetAvailableExtensionPropertiesResponse{
        DeletedItemsGetAvailableExtensionPropertiesPostResponse: *NewDeletedItemsGetAvailableExtensionPropertiesPostResponse(),
    }
    return m
}
// CreateDeletedItemsGetAvailableExtensionPropertiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletedItemsGetAvailableExtensionPropertiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedItemsGetAvailableExtensionPropertiesResponse(), nil
}
// DeletedItemsGetAvailableExtensionPropertiesResponseable 
// Deprecated: This class is obsolete. Use getAvailableExtensionPropertiesPostResponse instead.
type DeletedItemsGetAvailableExtensionPropertiesResponseable interface {
    DeletedItemsGetAvailableExtensionPropertiesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
