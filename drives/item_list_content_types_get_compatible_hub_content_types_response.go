package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListContentTypesGetCompatibleHubContentTypesgetResponseable instead.
type ItemListContentTypesGetCompatibleHubContentTypesResponse struct {
    ItemListContentTypesGetCompatibleHubContentTypesgetResponse
}
// NewItemListContentTypesGetCompatibleHubContentTypesResponse instantiates a new ItemListContentTypesGetCompatibleHubContentTypesResponse and sets the default values.
func NewItemListContentTypesGetCompatibleHubContentTypesResponse()(*ItemListContentTypesGetCompatibleHubContentTypesResponse) {
    m := &ItemListContentTypesGetCompatibleHubContentTypesResponse{
        ItemListContentTypesGetCompatibleHubContentTypesgetResponse: *NewItemListContentTypesGetCompatibleHubContentTypesgetResponse(),
    }
    return m
}
// CreateItemListContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListContentTypesGetCompatibleHubContentTypesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListContentTypesGetCompatibleHubContentTypesgetResponseable instead.
type ItemListContentTypesGetCompatibleHubContentTypesResponseable interface {
    ItemListContentTypesGetCompatibleHubContentTypesgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
