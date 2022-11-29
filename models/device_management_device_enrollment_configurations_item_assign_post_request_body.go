package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody provides operations to call the assign method.
type DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enrollmentConfigurationAssignments property
    enrollmentConfigurationAssignments []EnrollmentConfigurationAssignmentable
}
// NewDeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody instantiates a new DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody and sets the default values.
func NewDeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody()(*DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody) {
    m := &DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnrollmentConfigurationAssignments gets the enrollmentConfigurationAssignments property value. The enrollmentConfigurationAssignments property
func (m *DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetEnrollmentConfigurationAssignments()([]EnrollmentConfigurationAssignmentable) {
    return m.enrollmentConfigurationAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentConfigurationAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEnrollmentConfigurationAssignmentFromDiscriminatorValue , m.SetEnrollmentConfigurationAssignments)
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnrollmentConfigurationAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEnrollmentConfigurationAssignments())
        err := writer.WriteCollectionOfObjectValues("enrollmentConfigurationAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnrollmentConfigurationAssignments sets the enrollmentConfigurationAssignments property value. The enrollmentConfigurationAssignments property
func (m *DeviceManagementDeviceEnrollmentConfigurationsItemAssignPostRequestBody) SetEnrollmentConfigurationAssignments(value []EnrollmentConfigurationAssignmentable)() {
    m.enrollmentConfigurationAssignments = value
}
