package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DevicecompliancepoliciesItemAssignPostResponseable instead.
type DevicecompliancepoliciesItemAssignResponse struct {
    DevicecompliancepoliciesItemAssignPostResponse
}
// NewDevicecompliancepoliciesItemAssignResponse instantiates a new DevicecompliancepoliciesItemAssignResponse and sets the default values.
func NewDevicecompliancepoliciesItemAssignResponse()(*DevicecompliancepoliciesItemAssignResponse) {
    m := &DevicecompliancepoliciesItemAssignResponse{
        DevicecompliancepoliciesItemAssignPostResponse: *NewDevicecompliancepoliciesItemAssignPostResponse(),
    }
    return m
}
// CreateDevicecompliancepoliciesItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDevicecompliancepoliciesItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevicecompliancepoliciesItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use DevicecompliancepoliciesItemAssignPostResponseable instead.
type DevicecompliancepoliciesItemAssignResponseable interface {
    DevicecompliancepoliciesItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
