package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FederationConfigurationsAvailableProviderTypesResponse 
// Deprecated: This class is obsolete. Use availableProviderTypesGetResponse instead.
type FederationConfigurationsAvailableProviderTypesResponse struct {
    FederationConfigurationsAvailableProviderTypesGetResponse
}
// NewFederationConfigurationsAvailableProviderTypesResponse instantiates a new FederationConfigurationsAvailableProviderTypesResponse and sets the default values.
func NewFederationConfigurationsAvailableProviderTypesResponse()(*FederationConfigurationsAvailableProviderTypesResponse) {
    m := &FederationConfigurationsAvailableProviderTypesResponse{
        FederationConfigurationsAvailableProviderTypesGetResponse: *NewFederationConfigurationsAvailableProviderTypesGetResponse(),
    }
    return m
}
// CreateFederationConfigurationsAvailableProviderTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFederationConfigurationsAvailableProviderTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFederationConfigurationsAvailableProviderTypesResponse(), nil
}
// FederationConfigurationsAvailableProviderTypesResponseable 
// Deprecated: This class is obsolete. Use availableProviderTypesGetResponse instead.
type FederationConfigurationsAvailableProviderTypesResponseable interface {
    FederationConfigurationsAvailableProviderTypesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
