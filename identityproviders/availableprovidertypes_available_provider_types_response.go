package identityproviders

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AvailableprovidertypesAvailableProviderTypesGetResponseable instead.
type AvailableprovidertypesAvailableProviderTypesResponse struct {
    AvailableprovidertypesAvailableProviderTypesGetResponse
}
// NewAvailableprovidertypesAvailableProviderTypesResponse instantiates a new AvailableprovidertypesAvailableProviderTypesResponse and sets the default values.
func NewAvailableprovidertypesAvailableProviderTypesResponse()(*AvailableprovidertypesAvailableProviderTypesResponse) {
    m := &AvailableprovidertypesAvailableProviderTypesResponse{
        AvailableprovidertypesAvailableProviderTypesGetResponse: *NewAvailableprovidertypesAvailableProviderTypesGetResponse(),
    }
    return m
}
// CreateAvailableprovidertypesAvailableProviderTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAvailableprovidertypesAvailableProviderTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAvailableprovidertypesAvailableProviderTypesResponse(), nil
}
// Deprecated: This class is obsolete. Use AvailableprovidertypesAvailableProviderTypesGetResponseable instead.
type AvailableprovidertypesAvailableProviderTypesResponseable interface {
    AvailableprovidertypesAvailableProviderTypesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
