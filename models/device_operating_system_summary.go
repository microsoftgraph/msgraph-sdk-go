package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceOperatingSystemSummary device operating system summary.
type DeviceOperatingSystemSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Number of android device count.
    androidCount *int32
    // Number of iOS device count.
    iosCount *int32
    // Number of Mac OS X device count.
    macOSCount *int32
    // The OdataType property
    odataType *string
    // Number of unknown device count.
    unknownCount *int32
    // Number of Windows device count.
    windowsCount *int32
    // Number of Windows mobile device count.
    windowsMobileCount *int32
}
// NewDeviceOperatingSystemSummary instantiates a new deviceOperatingSystemSummary and sets the default values.
func NewDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    m := &DeviceOperatingSystemSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceOperatingSystemSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceOperatingSystemSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceOperatingSystemSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceOperatingSystemSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceOperatingSystemSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAndroidCount gets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) GetAndroidCount()(*int32) {
    return m.androidCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceOperatingSystemSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["androidCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetAndroidCount)
    res["iosCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetIosCount)
    res["macOSCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMacOSCount)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["unknownCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUnknownCount)
    res["windowsCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWindowsCount)
    res["windowsMobileCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWindowsMobileCount)
    return res
}
// GetIosCount gets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) GetIosCount()(*int32) {
    return m.iosCount
}
// GetMacOSCount gets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) GetMacOSCount()(*int32) {
    return m.macOSCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceOperatingSystemSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetUnknownCount gets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) GetUnknownCount()(*int32) {
    return m.unknownCount
}
// GetWindowsCount gets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) GetWindowsCount()(*int32) {
    return m.windowsCount
}
// GetWindowsMobileCount gets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) GetWindowsMobileCount()(*int32) {
    return m.windowsMobileCount
}
// Serialize serializes information the current object
func (m *DeviceOperatingSystemSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("androidCount", m.GetAndroidCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("iosCount", m.GetIosCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("macOSCount", m.GetMacOSCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownCount", m.GetUnknownCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsCount", m.GetWindowsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsMobileCount", m.GetWindowsMobileCount())
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
func (m *DeviceOperatingSystemSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAndroidCount sets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) SetAndroidCount(value *int32)() {
    m.androidCount = value
}
// SetIosCount sets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) SetIosCount(value *int32)() {
    m.iosCount = value
}
// SetMacOSCount sets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) SetMacOSCount(value *int32)() {
    m.macOSCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceOperatingSystemSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUnknownCount sets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) SetUnknownCount(value *int32)() {
    m.unknownCount = value
}
// SetWindowsCount sets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) SetWindowsCount(value *int32)() {
    m.windowsCount = value
}
// SetWindowsMobileCount sets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) SetWindowsMobileCount(value *int32)() {
    m.windowsMobileCount = value
}
