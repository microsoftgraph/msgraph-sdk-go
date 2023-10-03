package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemGetActivitiesByIntervalResponse 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalGetResponse instead.
type ItemSitesItemGetActivitiesByIntervalResponse struct {
    ItemSitesItemGetActivitiesByIntervalGetResponse
}
// NewItemSitesItemGetActivitiesByIntervalResponse instantiates a new ItemSitesItemGetActivitiesByIntervalResponse and sets the default values.
func NewItemSitesItemGetActivitiesByIntervalResponse()(*ItemSitesItemGetActivitiesByIntervalResponse) {
    m := &ItemSitesItemGetActivitiesByIntervalResponse{
        ItemSitesItemGetActivitiesByIntervalGetResponse: *NewItemSitesItemGetActivitiesByIntervalGetResponse(),
    }
    return m
}
// CreateItemSitesItemGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemGetActivitiesByIntervalResponse(), nil
}
// ItemSitesItemGetActivitiesByIntervalResponseable 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalGetResponse instead.
type ItemSitesItemGetActivitiesByIntervalResponseable interface {
    ItemSitesItemGetActivitiesByIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
