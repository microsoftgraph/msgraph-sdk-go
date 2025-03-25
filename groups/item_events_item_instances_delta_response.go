package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemEventsItemInstancesDeltagetResponseable instead.
type ItemEventsItemInstancesDeltaResponse struct {
    ItemEventsItemInstancesDeltagetResponse
}
// NewItemEventsItemInstancesDeltaResponse instantiates a new ItemEventsItemInstancesDeltaResponse and sets the default values.
func NewItemEventsItemInstancesDeltaResponse()(*ItemEventsItemInstancesDeltaResponse) {
    m := &ItemEventsItemInstancesDeltaResponse{
        ItemEventsItemInstancesDeltagetResponse: *NewItemEventsItemInstancesDeltagetResponse(),
    }
    return m
}
// CreateItemEventsItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemEventsItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemEventsItemInstancesDeltagetResponseable instead.
type ItemEventsItemInstancesDeltaResponseable interface {
    ItemEventsItemInstancesDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
