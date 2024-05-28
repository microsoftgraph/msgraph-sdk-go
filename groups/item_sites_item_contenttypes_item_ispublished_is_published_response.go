package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemContenttypesItemIspublishedIsPublishedGetResponseable instead.
type ItemSitesItemContenttypesItemIspublishedIsPublishedResponse struct {
    ItemSitesItemContenttypesItemIspublishedIsPublishedGetResponse
}
// NewItemSitesItemContenttypesItemIspublishedIsPublishedResponse instantiates a new ItemSitesItemContenttypesItemIspublishedIsPublishedResponse and sets the default values.
func NewItemSitesItemContenttypesItemIspublishedIsPublishedResponse()(*ItemSitesItemContenttypesItemIspublishedIsPublishedResponse) {
    m := &ItemSitesItemContenttypesItemIspublishedIsPublishedResponse{
        ItemSitesItemContenttypesItemIspublishedIsPublishedGetResponse: *NewItemSitesItemContenttypesItemIspublishedIsPublishedGetResponse(),
    }
    return m
}
// CreateItemSitesItemContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemContenttypesItemIspublishedIsPublishedResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemContenttypesItemIspublishedIsPublishedGetResponseable instead.
type ItemSitesItemContenttypesItemIspublishedIsPublishedResponseable interface {
    ItemSitesItemContenttypesItemIspublishedIsPublishedGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
