package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdministrativeUnitsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type AdministrativeUnitsDeltaResponse struct {
    AdministrativeUnitsDeltaGetResponse
}
// NewAdministrativeUnitsDeltaResponse instantiates a new AdministrativeUnitsDeltaResponse and sets the default values.
func NewAdministrativeUnitsDeltaResponse()(*AdministrativeUnitsDeltaResponse) {
    m := &AdministrativeUnitsDeltaResponse{
        AdministrativeUnitsDeltaGetResponse: *NewAdministrativeUnitsDeltaGetResponse(),
    }
    return m
}
// CreateAdministrativeUnitsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdministrativeUnitsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdministrativeUnitsDeltaResponse(), nil
}
// AdministrativeUnitsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type AdministrativeUnitsDeltaResponseable interface {
    AdministrativeUnitsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
