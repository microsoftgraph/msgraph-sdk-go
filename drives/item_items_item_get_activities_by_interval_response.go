package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemGetActivitiesByIntervalResponse 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalGetResponse instead.
type ItemItemsItemGetActivitiesByIntervalResponse struct {
    ItemItemsItemGetActivitiesByIntervalGetResponse
}
// NewItemItemsItemGetActivitiesByIntervalResponse instantiates a new ItemItemsItemGetActivitiesByIntervalResponse and sets the default values.
func NewItemItemsItemGetActivitiesByIntervalResponse()(*ItemItemsItemGetActivitiesByIntervalResponse) {
    m := &ItemItemsItemGetActivitiesByIntervalResponse{
        ItemItemsItemGetActivitiesByIntervalGetResponse: *NewItemItemsItemGetActivitiesByIntervalGetResponse(),
    }
    return m
}
// CreateItemItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemGetActivitiesByIntervalResponse(), nil
}
// ItemItemsItemGetActivitiesByIntervalResponseable 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalGetResponse instead.
type ItemItemsItemGetActivitiesByIntervalResponseable interface {
    ItemItemsItemGetActivitiesByIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
