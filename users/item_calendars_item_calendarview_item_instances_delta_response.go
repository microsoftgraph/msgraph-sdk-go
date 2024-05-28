package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarsItemCalendarviewItemInstancesDeltaGetResponseable instead.
type ItemCalendarsItemCalendarviewItemInstancesDeltaResponse struct {
    ItemCalendarsItemCalendarviewItemInstancesDeltaGetResponse
}
// NewItemCalendarsItemCalendarviewItemInstancesDeltaResponse instantiates a new ItemCalendarsItemCalendarviewItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarsItemCalendarviewItemInstancesDeltaResponse()(*ItemCalendarsItemCalendarviewItemInstancesDeltaResponse) {
    m := &ItemCalendarsItemCalendarviewItemInstancesDeltaResponse{
        ItemCalendarsItemCalendarviewItemInstancesDeltaGetResponse: *NewItemCalendarsItemCalendarviewItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarsItemCalendarviewItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarsItemCalendarviewItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemCalendarviewItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarsItemCalendarviewItemInstancesDeltaGetResponseable instead.
type ItemCalendarsItemCalendarviewItemInstancesDeltaResponseable interface {
    ItemCalendarsItemCalendarviewItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
