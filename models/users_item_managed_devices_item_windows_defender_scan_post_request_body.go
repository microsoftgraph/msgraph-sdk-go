package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody provides operations to call the windowsDefenderScan method.
type UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The quickScan property
    quickScan *bool
}
// NewUsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody instantiates a new UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody and sets the default values.
func NewUsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody()(*UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody) {
    m := &UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemManagedDevicesItemWindowsDefenderScanPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemManagedDevicesItemWindowsDefenderScanPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["quickScan"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetQuickScan)
    return res
}
// GetQuickScan gets the quickScan property value. The quickScan property
func (m *UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody) GetQuickScan()(*bool) {
    return m.quickScan
}
// Serialize serializes information the current object
func (m *UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("quickScan", m.GetQuickScan())
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
func (m *UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetQuickScan sets the quickScan property value. The quickScan property
func (m *UsersItemManagedDevicesItemWindowsDefenderScanPostRequestBody) SetQuickScan(value *bool)() {
    m.quickScan = value
}
