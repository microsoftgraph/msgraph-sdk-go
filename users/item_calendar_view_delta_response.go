package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarViewDeltagetResponseable instead.
type ItemCalendarViewDeltaResponse struct {
    ItemCalendarViewDeltagetResponse
}
// NewItemCalendarViewDeltaResponse instantiates a new ItemCalendarViewDeltaResponse and sets the default values.
func NewItemCalendarViewDeltaResponse()(*ItemCalendarViewDeltaResponse) {
    m := &ItemCalendarViewDeltaResponse{
        ItemCalendarViewDeltagetResponse: *NewItemCalendarViewDeltagetResponse(),
    }
    return m
}
// CreateItemCalendarViewDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarViewDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarViewDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarViewDeltagetResponseable instead.
type ItemCalendarViewDeltaResponseable interface {
    ItemCalendarViewDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
