package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesItemCleanWindowsDevicePostRequestBody provides operations to call the cleanWindowsDevice method.
type MeManagedDevicesItemCleanWindowsDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The keepUserData property
    keepUserData *bool
}
// NewMeManagedDevicesItemCleanWindowsDevicePostRequestBody instantiates a new MeManagedDevicesItemCleanWindowsDevicePostRequestBody and sets the default values.
func NewMeManagedDevicesItemCleanWindowsDevicePostRequestBody()(*MeManagedDevicesItemCleanWindowsDevicePostRequestBody) {
    m := &MeManagedDevicesItemCleanWindowsDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesItemCleanWindowsDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesItemCleanWindowsDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesItemCleanWindowsDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesItemCleanWindowsDevicePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesItemCleanWindowsDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keepUserData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    return res
}
// GetKeepUserData gets the keepUserData property value. The keepUserData property
func (m *MeManagedDevicesItemCleanWindowsDevicePostRequestBody) GetKeepUserData()(*bool) {
    return m.keepUserData
}
// Serialize serializes information the current object
func (m *MeManagedDevicesItemCleanWindowsDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
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
func (m *MeManagedDevicesItemCleanWindowsDevicePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *MeManagedDevicesItemCleanWindowsDevicePostRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
