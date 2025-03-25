package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListsItemContentTypesGetCompatibleHubContentTypesgetResponseable instead.
type ItemListsItemContentTypesGetCompatibleHubContentTypesResponse struct {
    ItemListsItemContentTypesGetCompatibleHubContentTypesgetResponse
}
// NewItemListsItemContentTypesGetCompatibleHubContentTypesResponse instantiates a new ItemListsItemContentTypesGetCompatibleHubContentTypesResponse and sets the default values.
func NewItemListsItemContentTypesGetCompatibleHubContentTypesResponse()(*ItemListsItemContentTypesGetCompatibleHubContentTypesResponse) {
    m := &ItemListsItemContentTypesGetCompatibleHubContentTypesResponse{
        ItemListsItemContentTypesGetCompatibleHubContentTypesgetResponse: *NewItemListsItemContentTypesGetCompatibleHubContentTypesgetResponse(),
    }
    return m
}
// CreateItemListsItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListsItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemContentTypesGetCompatibleHubContentTypesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListsItemContentTypesGetCompatibleHubContentTypesgetResponseable instead.
type ItemListsItemContentTypesGetCompatibleHubContentTypesResponseable interface {
    ItemListsItemContentTypesGetCompatibleHubContentTypesgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
