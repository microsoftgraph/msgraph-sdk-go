package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable instead.
type ItemListsItemContenttypesItemIspublishedIsPublishedResponse struct {
    ItemListsItemContenttypesItemIspublishedIsPublishedGetResponse
}
// NewItemListsItemContenttypesItemIspublishedIsPublishedResponse instantiates a new ItemListsItemContenttypesItemIspublishedIsPublishedResponse and sets the default values.
func NewItemListsItemContenttypesItemIspublishedIsPublishedResponse()(*ItemListsItemContenttypesItemIspublishedIsPublishedResponse) {
    m := &ItemListsItemContenttypesItemIspublishedIsPublishedResponse{
        ItemListsItemContenttypesItemIspublishedIsPublishedGetResponse: *NewItemListsItemContenttypesItemIspublishedIsPublishedGetResponse(),
    }
    return m
}
// CreateItemListsItemContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListsItemContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemContenttypesItemIspublishedIsPublishedResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable instead.
type ItemListsItemContenttypesItemIspublishedIsPublishedResponseable interface {
    ItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
