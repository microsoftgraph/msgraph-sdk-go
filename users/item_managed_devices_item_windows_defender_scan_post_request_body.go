package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemWindowsDefenderScanPostRequestBody provides operations to call the windowsDefenderScan method.
type ItemManagedDevicesItemWindowsDefenderScanPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The quickScan property
    quickScan *bool
}
// NewItemManagedDevicesItemWindowsDefenderScanPostRequestBody instantiates a new ItemManagedDevicesItemWindowsDefenderScanPostRequestBody and sets the default values.
func NewItemManagedDevicesItemWindowsDefenderScanPostRequestBody()(*ItemManagedDevicesItemWindowsDefenderScanPostRequestBody) {
    m := &ItemManagedDevicesItemWindowsDefenderScanPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemManagedDevicesItemWindowsDefenderScanPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemWindowsDefenderScanPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemWindowsDefenderScanPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemWindowsDefenderScanPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemWindowsDefenderScanPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["quickScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickScan(val)
        }
        return nil
    }
    return res
}
// GetQuickScan gets the quickScan property value. The quickScan property
func (m *ItemManagedDevicesItemWindowsDefenderScanPostRequestBody) GetQuickScan()(*bool) {
    return m.quickScan
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemWindowsDefenderScanPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemManagedDevicesItemWindowsDefenderScanPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetQuickScan sets the quickScan property value. The quickScan property
func (m *ItemManagedDevicesItemWindowsDefenderScanPostRequestBody) SetQuickScan(value *bool)() {
    m.quickScan = value
}
