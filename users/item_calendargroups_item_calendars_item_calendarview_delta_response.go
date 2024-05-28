package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendargroupsItemCalendarsItemCalendarviewDeltaGetResponseable instead.
type ItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponse struct {
    ItemCalendargroupsItemCalendarsItemCalendarviewDeltaGetResponse
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponse instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponse and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponse()(*ItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponse) {
    m := &ItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponse{
        ItemCalendargroupsItemCalendarsItemCalendarviewDeltaGetResponse: *NewItemCalendargroupsItemCalendarsItemCalendarviewDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendargroupsItemCalendarsItemCalendarviewDeltaGetResponseable instead.
type ItemCalendargroupsItemCalendarsItemCalendarviewDeltaResponseable interface {
    ItemCalendargroupsItemCalendarsItemCalendarviewDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
