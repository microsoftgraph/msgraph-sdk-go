package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse provides operations to call the getUserIdsWithFlaggedAppRegistration method.
type DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []string
}
// NewDeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse instantiates a new DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse and sets the default values.
func NewDeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse()(*DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse) {
    m := &DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponse) SetValue(value []string)() {
    m.value = value
}
