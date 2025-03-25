package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarViewItemInstancesDeltagetResponseable instead.
type ItemCalendarViewItemInstancesDeltaResponse struct {
    ItemCalendarViewItemInstancesDeltagetResponse
}
// NewItemCalendarViewItemInstancesDeltaResponse instantiates a new ItemCalendarViewItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarViewItemInstancesDeltaResponse()(*ItemCalendarViewItemInstancesDeltaResponse) {
    m := &ItemCalendarViewItemInstancesDeltaResponse{
        ItemCalendarViewItemInstancesDeltagetResponse: *NewItemCalendarViewItemInstancesDeltagetResponse(),
    }
    return m
}
// CreateItemCalendarViewItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarViewItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarViewItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarViewItemInstancesDeltagetResponseable instead.
type ItemCalendarViewItemInstancesDeltaResponseable interface {
    ItemCalendarViewItemInstancesDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
