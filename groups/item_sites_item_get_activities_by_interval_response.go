package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemGetActivitiesByIntervalgetResponseable instead.
type ItemSitesItemGetActivitiesByIntervalResponse struct {
    ItemSitesItemGetActivitiesByIntervalgetResponse
}
// NewItemSitesItemGetActivitiesByIntervalResponse instantiates a new ItemSitesItemGetActivitiesByIntervalResponse and sets the default values.
func NewItemSitesItemGetActivitiesByIntervalResponse()(*ItemSitesItemGetActivitiesByIntervalResponse) {
    m := &ItemSitesItemGetActivitiesByIntervalResponse{
        ItemSitesItemGetActivitiesByIntervalgetResponse: *NewItemSitesItemGetActivitiesByIntervalgetResponse(),
    }
    return m
}
// CreateItemSitesItemGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemGetActivitiesByIntervalResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemGetActivitiesByIntervalgetResponseable instead.
type ItemSitesItemGetActivitiesByIntervalResponseable interface {
    ItemSitesItemGetActivitiesByIntervalgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
