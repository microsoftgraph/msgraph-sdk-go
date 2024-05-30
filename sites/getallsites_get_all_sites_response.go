package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetallsitesGetAllSitesGetResponseable instead.
type GetallsitesGetAllSitesResponse struct {
    GetallsitesGetAllSitesGetResponse
}
// NewGetallsitesGetAllSitesResponse instantiates a new GetallsitesGetAllSitesResponse and sets the default values.
func NewGetallsitesGetAllSitesResponse()(*GetallsitesGetAllSitesResponse) {
    m := &GetallsitesGetAllSitesResponse{
        GetallsitesGetAllSitesGetResponse: *NewGetallsitesGetAllSitesGetResponse(),
    }
    return m
}
// CreateGetallsitesGetAllSitesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetallsitesGetAllSitesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetallsitesGetAllSitesResponse(), nil
}
// Deprecated: This class is obsolete. Use GetallsitesGetAllSitesGetResponseable instead.
type GetallsitesGetAllSitesResponseable interface {
    GetallsitesGetAllSitesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
