package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementMobileAppsItemAssignPostRequestBody provides operations to call the assign method.
type DeviceAppManagementMobileAppsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The mobileAppAssignments property
    mobileAppAssignments []MobileAppAssignmentable
}
// NewDeviceAppManagementMobileAppsItemAssignPostRequestBody instantiates a new DeviceAppManagementMobileAppsItemAssignPostRequestBody and sets the default values.
func NewDeviceAppManagementMobileAppsItemAssignPostRequestBody()(*DeviceAppManagementMobileAppsItemAssignPostRequestBody) {
    m := &DeviceAppManagementMobileAppsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementMobileAppsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementMobileAppsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementMobileAppsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementMobileAppsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementMobileAppsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mobileAppAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppAssignmentable)
            }
            m.SetMobileAppAssignments(res)
        }
        return nil
    }
    return res
}
// GetMobileAppAssignments gets the mobileAppAssignments property value. The mobileAppAssignments property
func (m *DeviceAppManagementMobileAppsItemAssignPostRequestBody) GetMobileAppAssignments()([]MobileAppAssignmentable) {
    return m.mobileAppAssignments
}
// Serialize serializes information the current object
func (m *DeviceAppManagementMobileAppsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMobileAppAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppAssignments()))
        for i, v := range m.GetMobileAppAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("mobileAppAssignments", cast)
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
func (m *DeviceAppManagementMobileAppsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMobileAppAssignments sets the mobileAppAssignments property value. The mobileAppAssignments property
func (m *DeviceAppManagementMobileAppsItemAssignPostRequestBody) SetMobileAppAssignments(value []MobileAppAssignmentable)() {
    m.mobileAppAssignments = value
}
