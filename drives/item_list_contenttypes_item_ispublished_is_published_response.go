package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListContenttypesItemIspublishedIsPublishedGetResponseable instead.
type ItemListContenttypesItemIspublishedIsPublishedResponse struct {
    ItemListContenttypesItemIspublishedIsPublishedGetResponse
}
// NewItemListContenttypesItemIspublishedIsPublishedResponse instantiates a new ItemListContenttypesItemIspublishedIsPublishedResponse and sets the default values.
func NewItemListContenttypesItemIspublishedIsPublishedResponse()(*ItemListContenttypesItemIspublishedIsPublishedResponse) {
    m := &ItemListContenttypesItemIspublishedIsPublishedResponse{
        ItemListContenttypesItemIspublishedIsPublishedGetResponse: *NewItemListContenttypesItemIspublishedIsPublishedGetResponse(),
    }
    return m
}
// CreateItemListContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListContenttypesItemIspublishedIsPublishedResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListContenttypesItemIspublishedIsPublishedGetResponseable instead.
type ItemListContenttypesItemIspublishedIsPublishedResponseable interface {
    ItemListContenttypesItemIspublishedIsPublishedGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
