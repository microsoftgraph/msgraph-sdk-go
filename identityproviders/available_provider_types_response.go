package identityproviders

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AvailableProviderTypesResponse 
// Deprecated: This class is obsolete. Use availableProviderTypesGetResponse instead.
type AvailableProviderTypesResponse struct {
    AvailableProviderTypesGetResponse
}
// NewAvailableProviderTypesResponse instantiates a new AvailableProviderTypesResponse and sets the default values.
func NewAvailableProviderTypesResponse()(*AvailableProviderTypesResponse) {
    m := &AvailableProviderTypesResponse{
        AvailableProviderTypesGetResponse: *NewAvailableProviderTypesGetResponse(),
    }
    return m
}
// CreateAvailableProviderTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAvailableProviderTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAvailableProviderTypesResponse(), nil
}
// AvailableProviderTypesResponseable 
// Deprecated: This class is obsolete. Use availableProviderTypesGetResponse instead.
type AvailableProviderTypesResponseable interface {
    AvailableProviderTypesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
