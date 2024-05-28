package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarsItemCalendarviewDeltaGetResponseable instead.
type ItemCalendarsItemCalendarviewDeltaResponse struct {
    ItemCalendarsItemCalendarviewDeltaGetResponse
}
// NewItemCalendarsItemCalendarviewDeltaResponse instantiates a new ItemCalendarsItemCalendarviewDeltaResponse and sets the default values.
func NewItemCalendarsItemCalendarviewDeltaResponse()(*ItemCalendarsItemCalendarviewDeltaResponse) {
    m := &ItemCalendarsItemCalendarviewDeltaResponse{
        ItemCalendarsItemCalendarviewDeltaGetResponse: *NewItemCalendarsItemCalendarviewDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarsItemCalendarviewDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarsItemCalendarviewDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemCalendarviewDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarsItemCalendarviewDeltaGetResponseable instead.
type ItemCalendarsItemCalendarviewDeltaResponseable interface {
    ItemCalendarsItemCalendarviewDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
