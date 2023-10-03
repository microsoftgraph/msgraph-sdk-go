package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingBusinessesItemGetStaffAvailabilityResponse 
// Deprecated: This class is obsolete. Use getStaffAvailabilityPostResponse instead.
type BookingBusinessesItemGetStaffAvailabilityResponse struct {
    BookingBusinessesItemGetStaffAvailabilityPostResponse
}
// NewBookingBusinessesItemGetStaffAvailabilityResponse instantiates a new BookingBusinessesItemGetStaffAvailabilityResponse and sets the default values.
func NewBookingBusinessesItemGetStaffAvailabilityResponse()(*BookingBusinessesItemGetStaffAvailabilityResponse) {
    m := &BookingBusinessesItemGetStaffAvailabilityResponse{
        BookingBusinessesItemGetStaffAvailabilityPostResponse: *NewBookingBusinessesItemGetStaffAvailabilityPostResponse(),
    }
    return m
}
// CreateBookingBusinessesItemGetStaffAvailabilityResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingBusinessesItemGetStaffAvailabilityResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingBusinessesItemGetStaffAvailabilityResponse(), nil
}
// BookingBusinessesItemGetStaffAvailabilityResponseable 
// Deprecated: This class is obsolete. Use getStaffAvailabilityPostResponse instead.
type BookingBusinessesItemGetStaffAvailabilityResponseable interface {
    BookingBusinessesItemGetStaffAvailabilityPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
