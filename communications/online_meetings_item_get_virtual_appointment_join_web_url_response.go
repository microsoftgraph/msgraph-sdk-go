package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlgetResponseable instead.
type OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse struct {
    OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlgetResponse
}
// NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse instantiates a new OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse and sets the default values.
func NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse()(*OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse) {
    m := &OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse{
        OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlgetResponse: *NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlgetResponse(),
    }
    return m
}
// CreateOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse(), nil
}
// Deprecated: This class is obsolete. Use OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlgetResponseable instead.
type OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseable interface {
    OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
