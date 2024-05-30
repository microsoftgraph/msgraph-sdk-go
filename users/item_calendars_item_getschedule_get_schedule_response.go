package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarsItemGetscheduleGetSchedulePostResponseable instead.
type ItemCalendarsItemGetscheduleGetScheduleResponse struct {
    ItemCalendarsItemGetscheduleGetSchedulePostResponse
}
// NewItemCalendarsItemGetscheduleGetScheduleResponse instantiates a new ItemCalendarsItemGetscheduleGetScheduleResponse and sets the default values.
func NewItemCalendarsItemGetscheduleGetScheduleResponse()(*ItemCalendarsItemGetscheduleGetScheduleResponse) {
    m := &ItemCalendarsItemGetscheduleGetScheduleResponse{
        ItemCalendarsItemGetscheduleGetSchedulePostResponse: *NewItemCalendarsItemGetscheduleGetSchedulePostResponse(),
    }
    return m
}
// CreateItemCalendarsItemGetscheduleGetScheduleResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarsItemGetscheduleGetScheduleResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemGetscheduleGetScheduleResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarsItemGetscheduleGetSchedulePostResponseable instead.
type ItemCalendarsItemGetscheduleGetScheduleResponseable interface {
    ItemCalendarsItemGetscheduleGetSchedulePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
