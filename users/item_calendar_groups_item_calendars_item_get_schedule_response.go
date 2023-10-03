package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarGroupsItemCalendarsItemGetScheduleResponse 
// Deprecated: This class is obsolete. Use getSchedulePostResponse instead.
type ItemCalendarGroupsItemCalendarsItemGetScheduleResponse struct {
    ItemCalendarGroupsItemCalendarsItemGetSchedulePostResponse
}
// NewItemCalendarGroupsItemCalendarsItemGetScheduleResponse instantiates a new ItemCalendarGroupsItemCalendarsItemGetScheduleResponse and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemGetScheduleResponse()(*ItemCalendarGroupsItemCalendarsItemGetScheduleResponse) {
    m := &ItemCalendarGroupsItemCalendarsItemGetScheduleResponse{
        ItemCalendarGroupsItemCalendarsItemGetSchedulePostResponse: *NewItemCalendarGroupsItemCalendarsItemGetSchedulePostResponse(),
    }
    return m
}
// CreateItemCalendarGroupsItemCalendarsItemGetScheduleResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarGroupsItemCalendarsItemGetScheduleResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarGroupsItemCalendarsItemGetScheduleResponse(), nil
}
// ItemCalendarGroupsItemCalendarsItemGetScheduleResponseable 
// Deprecated: This class is obsolete. Use getSchedulePostResponse instead.
type ItemCalendarGroupsItemCalendarsItemGetScheduleResponseable interface {
    ItemCalendarGroupsItemCalendarsItemGetSchedulePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
