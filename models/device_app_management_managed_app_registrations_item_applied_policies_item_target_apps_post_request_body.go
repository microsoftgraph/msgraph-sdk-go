package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody provides operations to call the targetApps method.
type DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The apps property
    apps []ManagedMobileAppable
}
// NewDeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody instantiates a new DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody and sets the default values.
func NewDeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody()(*DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody) {
    m := &DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApps gets the apps property value. The apps property
func (m *DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody) GetApps()([]ManagedMobileAppable) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("apps", cast)
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
func (m *DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApps sets the apps property value. The apps property
func (m *DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBody) SetApps(value []ManagedMobileAppable)() {
    m.apps = value
}
