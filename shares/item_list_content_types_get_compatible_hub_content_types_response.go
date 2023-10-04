package shares

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListContentTypesGetCompatibleHubContentTypesResponse 
// Deprecated: This class is obsolete. Use getCompatibleHubContentTypesGetResponse instead.
type ItemListContentTypesGetCompatibleHubContentTypesResponse struct {
    ItemListContentTypesGetCompatibleHubContentTypesGetResponse
}
// NewItemListContentTypesGetCompatibleHubContentTypesResponse instantiates a new ItemListContentTypesGetCompatibleHubContentTypesResponse and sets the default values.
func NewItemListContentTypesGetCompatibleHubContentTypesResponse()(*ItemListContentTypesGetCompatibleHubContentTypesResponse) {
    m := &ItemListContentTypesGetCompatibleHubContentTypesResponse{
        ItemListContentTypesGetCompatibleHubContentTypesGetResponse: *NewItemListContentTypesGetCompatibleHubContentTypesGetResponse(),
    }
    return m
}
// CreateItemListContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListContentTypesGetCompatibleHubContentTypesResponse(), nil
}
// ItemListContentTypesGetCompatibleHubContentTypesResponseable 
// Deprecated: This class is obsolete. Use getCompatibleHubContentTypesGetResponse instead.
type ItemListContentTypesGetCompatibleHubContentTypesResponseable interface {
    ItemListContentTypesGetCompatibleHubContentTypesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
