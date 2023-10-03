package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetActivitiesByIntervalResponse 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalGetResponse instead.
type ItemGetActivitiesByIntervalResponse struct {
    ItemGetActivitiesByIntervalGetResponse
}
// NewItemGetActivitiesByIntervalResponse instantiates a new ItemGetActivitiesByIntervalResponse and sets the default values.
func NewItemGetActivitiesByIntervalResponse()(*ItemGetActivitiesByIntervalResponse) {
    m := &ItemGetActivitiesByIntervalResponse{
        ItemGetActivitiesByIntervalGetResponse: *NewItemGetActivitiesByIntervalGetResponse(),
    }
    return m
}
// CreateItemGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetActivitiesByIntervalResponse(), nil
}
// ItemGetActivitiesByIntervalResponseable 
// Deprecated: This class is obsolete. Use getActivitiesByIntervalGetResponse instead.
type ItemGetActivitiesByIntervalResponseable interface {
    ItemGetActivitiesByIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
