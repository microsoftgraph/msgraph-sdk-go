package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// DeviceEnrollmentConfigurationsItemAssignPostRequestBody 
type DeviceEnrollmentConfigurationsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enrollmentConfigurationAssignments property
    enrollmentConfigurationAssignments []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EnrollmentConfigurationAssignmentable
}
// NewDeviceEnrollmentConfigurationsItemAssignPostRequestBody instantiates a new DeviceEnrollmentConfigurationsItemAssignPostRequestBody and sets the default values.
func NewDeviceEnrollmentConfigurationsItemAssignPostRequestBody()(*DeviceEnrollmentConfigurationsItemAssignPostRequestBody) {
    m := &DeviceEnrollmentConfigurationsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceEnrollmentConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentConfigurationsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnrollmentConfigurationAssignments gets the enrollmentConfigurationAssignments property value. The enrollmentConfigurationAssignments property
func (m *DeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetEnrollmentConfigurationAssignments()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EnrollmentConfigurationAssignmentable) {
    return m.enrollmentConfigurationAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentConfigurationAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEnrollmentConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EnrollmentConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EnrollmentConfigurationAssignmentable)
            }
            m.SetEnrollmentConfigurationAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentConfigurationsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnrollmentConfigurationAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnrollmentConfigurationAssignments()))
        for i, v := range m.GetEnrollmentConfigurationAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
func (m *DeviceEnrollmentConfigurationsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnrollmentConfigurationAssignments sets the enrollmentConfigurationAssignments property value. The enrollmentConfigurationAssignments property
func (m *DeviceEnrollmentConfigurationsItemAssignPostRequestBody) SetEnrollmentConfigurationAssignments(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EnrollmentConfigurationAssignmentable)() {
    m.enrollmentConfigurationAssignments = value
}
