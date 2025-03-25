package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use IdentityProvidersAvailableProviderTypesgetResponseable instead.
type IdentityProvidersAvailableProviderTypesResponse struct {
    IdentityProvidersAvailableProviderTypesgetResponse
}
// NewIdentityProvidersAvailableProviderTypesResponse instantiates a new IdentityProvidersAvailableProviderTypesResponse and sets the default values.
func NewIdentityProvidersAvailableProviderTypesResponse()(*IdentityProvidersAvailableProviderTypesResponse) {
    m := &IdentityProvidersAvailableProviderTypesResponse{
        IdentityProvidersAvailableProviderTypesgetResponse: *NewIdentityProvidersAvailableProviderTypesgetResponse(),
    }
    return m
}
// CreateIdentityProvidersAvailableProviderTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentityProvidersAvailableProviderTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityProvidersAvailableProviderTypesResponse(), nil
}
// Deprecated: This class is obsolete. Use IdentityProvidersAvailableProviderTypesgetResponseable instead.
type IdentityProvidersAvailableProviderTypesResponseable interface {
    IdentityProvidersAvailableProviderTypesgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
