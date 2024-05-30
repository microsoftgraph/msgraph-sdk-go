package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarGetscheduleGetSchedulePostResponseable instead.
type ItemCalendarGetscheduleGetScheduleResponse struct {
    ItemCalendarGetscheduleGetSchedulePostResponse
}
// NewItemCalendarGetscheduleGetScheduleResponse instantiates a new ItemCalendarGetscheduleGetScheduleResponse and sets the default values.
func NewItemCalendarGetscheduleGetScheduleResponse()(*ItemCalendarGetscheduleGetScheduleResponse) {
    m := &ItemCalendarGetscheduleGetScheduleResponse{
        ItemCalendarGetscheduleGetSchedulePostResponse: *NewItemCalendarGetscheduleGetSchedulePostResponse(),
    }
    return m
}
// CreateItemCalendarGetscheduleGetScheduleResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarGetscheduleGetScheduleResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarGetscheduleGetScheduleResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarGetscheduleGetSchedulePostResponseable instead.
type ItemCalendarGetscheduleGetScheduleResponseable interface {
    ItemCalendarGetscheduleGetSchedulePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
