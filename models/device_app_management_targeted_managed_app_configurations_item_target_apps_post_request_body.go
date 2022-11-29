package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
    res["appGroupType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTargetedManagedAppGroupType , m.SetAppGroupType)
    res["apps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue , m.SetApps)
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetApps())
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
