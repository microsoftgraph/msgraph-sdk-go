package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarCalendarViewItemInstancesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarCalendarViewItemInstancesDeltaResponse struct {
    ItemCalendarCalendarViewItemInstancesDeltaGetResponse
}
// NewItemCalendarCalendarViewItemInstancesDeltaResponse instantiates a new ItemCalendarCalendarViewItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarCalendarViewItemInstancesDeltaResponse()(*ItemCalendarCalendarViewItemInstancesDeltaResponse) {
    m := &ItemCalendarCalendarViewItemInstancesDeltaResponse{
        ItemCalendarCalendarViewItemInstancesDeltaGetResponse: *NewItemCalendarCalendarViewItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarCalendarViewItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarCalendarViewItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarCalendarViewItemInstancesDeltaResponse(), nil
}
// ItemCalendarCalendarViewItemInstancesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarCalendarViewItemInstancesDeltaResponseable interface {
    ItemCalendarCalendarViewItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
