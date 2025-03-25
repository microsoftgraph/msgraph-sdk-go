package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetAllSitesgetResponseable instead.
type GetAllSitesResponse struct {
    GetAllSitesgetResponse
}
// NewGetAllSitesResponse instantiates a new GetAllSitesResponse and sets the default values.
func NewGetAllSitesResponse()(*GetAllSitesResponse) {
    m := &GetAllSitesResponse{
        GetAllSitesgetResponse: *NewGetAllSitesgetResponse(),
    }
    return m
}
// CreateGetAllSitesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetAllSitesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAllSitesResponse(), nil
}
// Deprecated: This class is obsolete. Use GetAllSitesgetResponseable instead.
type GetAllSitesResponseable interface {
    GetAllSitesgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
