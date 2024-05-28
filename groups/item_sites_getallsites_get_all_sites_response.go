package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesGetallsitesGetAllSitesGetResponseable instead.
type ItemSitesGetallsitesGetAllSitesResponse struct {
    ItemSitesGetallsitesGetAllSitesGetResponse
}
// NewItemSitesGetallsitesGetAllSitesResponse instantiates a new ItemSitesGetallsitesGetAllSitesResponse and sets the default values.
func NewItemSitesGetallsitesGetAllSitesResponse()(*ItemSitesGetallsitesGetAllSitesResponse) {
    m := &ItemSitesGetallsitesGetAllSitesResponse{
        ItemSitesGetallsitesGetAllSitesGetResponse: *NewItemSitesGetallsitesGetAllSitesGetResponse(),
    }
    return m
}
// CreateItemSitesGetallsitesGetAllSitesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesGetallsitesGetAllSitesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesGetallsitesGetAllSitesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesGetallsitesGetAllSitesGetResponseable instead.
type ItemSitesGetallsitesGetAllSitesResponseable interface {
    ItemSitesGetallsitesGetAllSitesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
