package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContentTypesGetCompatibleHubContentTypesgetResponseable instead.
type ItemContentTypesGetCompatibleHubContentTypesResponse struct {
    ItemContentTypesGetCompatibleHubContentTypesgetResponse
}
// NewItemContentTypesGetCompatibleHubContentTypesResponse instantiates a new ItemContentTypesGetCompatibleHubContentTypesResponse and sets the default values.
func NewItemContentTypesGetCompatibleHubContentTypesResponse()(*ItemContentTypesGetCompatibleHubContentTypesResponse) {
    m := &ItemContentTypesGetCompatibleHubContentTypesResponse{
        ItemContentTypesGetCompatibleHubContentTypesgetResponse: *NewItemContentTypesGetCompatibleHubContentTypesgetResponse(),
    }
    return m
}
// CreateItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContentTypesGetCompatibleHubContentTypesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContentTypesGetCompatibleHubContentTypesgetResponseable instead.
type ItemContentTypesGetCompatibleHubContentTypesResponseable interface {
    ItemContentTypesGetCompatibleHubContentTypesgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
