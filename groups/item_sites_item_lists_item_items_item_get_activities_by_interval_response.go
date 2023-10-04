package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalGetResponse instead.
type ItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse struct {
    ItemSitesItemListsItemItemsItemGetActivitiesByIntervalGetResponse
}
// NewItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse instantiates a new ItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse and sets the default values.
func NewItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse()(*ItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse) {
    m := &ItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse{
        ItemSitesItemListsItemItemsItemGetActivitiesByIntervalGetResponse: *NewItemSitesItemListsItemItemsItemGetActivitiesByIntervalGetResponse(),
    }
    return m
}
// CreateItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse(), nil
}
// ItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponseable 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalGetResponse instead.
type ItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponseable interface {
    ItemSitesItemListsItemItemsItemGetActivitiesByIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
