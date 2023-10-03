package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemContentTypesGetCompatibleHubContentTypesResponse 
// Deprecated: This class is obsolete. Use getCompatibleHubContentTypesGetResponse instead.
type ItemContentTypesGetCompatibleHubContentTypesResponse struct {
    ItemContentTypesGetCompatibleHubContentTypesGetResponse
}
// NewItemContentTypesGetCompatibleHubContentTypesResponse instantiates a new ItemContentTypesGetCompatibleHubContentTypesResponse and sets the default values.
func NewItemContentTypesGetCompatibleHubContentTypesResponse()(*ItemContentTypesGetCompatibleHubContentTypesResponse) {
    m := &ItemContentTypesGetCompatibleHubContentTypesResponse{
        ItemContentTypesGetCompatibleHubContentTypesGetResponse: *NewItemContentTypesGetCompatibleHubContentTypesGetResponse(),
    }
    return m
}
// CreateItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContentTypesGetCompatibleHubContentTypesResponse(), nil
}
// ItemContentTypesGetCompatibleHubContentTypesResponseable 
// Deprecated: This class is obsolete. Use getCompatibleHubContentTypesGetResponse instead.
type ItemContentTypesGetCompatibleHubContentTypesResponseable interface {
    ItemContentTypesGetCompatibleHubContentTypesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
