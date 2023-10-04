package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse 
// Deprecated: This class is obsolete. Use getVirtualAppointmentJoinWebUrlGetResponse instead.
type OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse struct {
    OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlGetResponse
}
// NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse instantiates a new OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse and sets the default values.
func NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse()(*OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse) {
    m := &OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse{
        OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlGetResponse: *NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlGetResponse(),
    }
    return m
}
// CreateOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponse(), nil
}
// OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseable 
// Deprecated: This class is obsolete. Use getVirtualAppointmentJoinWebUrlGetResponse instead.
type OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlResponseable interface {
    OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
