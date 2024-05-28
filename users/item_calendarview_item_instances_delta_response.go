package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarviewItemInstancesDeltaGetResponseable instead.
type ItemCalendarviewItemInstancesDeltaResponse struct {
    ItemCalendarviewItemInstancesDeltaGetResponse
}
// NewItemCalendarviewItemInstancesDeltaResponse instantiates a new ItemCalendarviewItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarviewItemInstancesDeltaResponse()(*ItemCalendarviewItemInstancesDeltaResponse) {
    m := &ItemCalendarviewItemInstancesDeltaResponse{
        ItemCalendarviewItemInstancesDeltaGetResponse: *NewItemCalendarviewItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarviewItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarviewItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarviewItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarviewItemInstancesDeltaGetResponseable instead.
type ItemCalendarviewItemInstancesDeltaResponseable interface {
    ItemCalendarviewItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
