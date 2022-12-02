package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody provides operations to call the targetApps method.
type DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appGroupType property
    appGroupType *TargetedManagedAppGroupType
    // The apps property
    apps []ManagedMobileAppable
}
// NewDeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody instantiates a new DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody and sets the default values.
func NewDeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody()(*DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) {
    m := &DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppGroupType gets the appGroupType property value. The appGroupType property
func (m *DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) GetAppGroupType()(*TargetedManagedAppGroupType) {
    return m.appGroupType
}
// GetApps gets the apps property value. The apps property
func (m *DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) GetApps()([]ManagedMobileAppable) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appGroupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTargetedManagedAppGroupType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppGroupType(val.(*TargetedManagedAppGroupType))
        }
        return nil
    }
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
func (m *DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppGroupType() != nil {
        cast := (*m.GetAppGroupType()).String()
        err := writer.WriteStringValue("appGroupType", &cast)
        if err != nil {
            return err
        }
    }
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
func (m *DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppGroupType sets the appGroupType property value. The appGroupType property
func (m *DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) SetAppGroupType(value *TargetedManagedAppGroupType)() {
    m.appGroupType = value
}
// SetApps sets the apps property value. The apps property
func (m *DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsPostRequestBody) SetApps(value []ManagedMobileAppable)() {
    m.apps = value
}
