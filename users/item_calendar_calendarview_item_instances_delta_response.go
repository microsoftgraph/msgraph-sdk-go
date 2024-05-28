package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarCalendarviewItemInstancesDeltaGetResponseable instead.
type ItemCalendarCalendarviewItemInstancesDeltaResponse struct {
    ItemCalendarCalendarviewItemInstancesDeltaGetResponse
}
// NewItemCalendarCalendarviewItemInstancesDeltaResponse instantiates a new ItemCalendarCalendarviewItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarCalendarviewItemInstancesDeltaResponse()(*ItemCalendarCalendarviewItemInstancesDeltaResponse) {
    m := &ItemCalendarCalendarviewItemInstancesDeltaResponse{
        ItemCalendarCalendarviewItemInstancesDeltaGetResponse: *NewItemCalendarCalendarviewItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarCalendarviewItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarCalendarviewItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarCalendarviewItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarCalendarviewItemInstancesDeltaGetResponseable instead.
type ItemCalendarCalendarviewItemInstancesDeltaResponseable interface {
    ItemCalendarCalendarviewItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
