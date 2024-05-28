package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarCalendarviewDeltaGetResponseable instead.
type ItemCalendarCalendarviewDeltaResponse struct {
    ItemCalendarCalendarviewDeltaGetResponse
}
// NewItemCalendarCalendarviewDeltaResponse instantiates a new ItemCalendarCalendarviewDeltaResponse and sets the default values.
func NewItemCalendarCalendarviewDeltaResponse()(*ItemCalendarCalendarviewDeltaResponse) {
    m := &ItemCalendarCalendarviewDeltaResponse{
        ItemCalendarCalendarviewDeltaGetResponse: *NewItemCalendarCalendarviewDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarCalendarviewDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarCalendarviewDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarCalendarviewDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarCalendarviewDeltaGetResponseable instead.
type ItemCalendarCalendarviewDeltaResponseable interface {
    ItemCalendarCalendarviewDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
