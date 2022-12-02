package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody provides operations to call the updateWindowsDeviceAccount method.
type UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The updateWindowsDeviceAccountActionParameter property
    updateWindowsDeviceAccountActionParameter UpdateWindowsDeviceAccountActionParameterable
}
// NewUsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody instantiates a new UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody and sets the default values.
func NewUsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody()(*UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) {
    m := &UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["updateWindowsDeviceAccountActionParameter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUpdateWindowsDeviceAccountActionParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindowsDeviceAccountActionParameter(val.(UpdateWindowsDeviceAccountActionParameterable))
        }
        return nil
    }
    return res
}
// GetUpdateWindowsDeviceAccountActionParameter gets the updateWindowsDeviceAccountActionParameter property value. The updateWindowsDeviceAccountActionParameter property
func (m *UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) GetUpdateWindowsDeviceAccountActionParameter()(UpdateWindowsDeviceAccountActionParameterable) {
    return m.updateWindowsDeviceAccountActionParameter
}
// Serialize serializes information the current object
func (m *UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("updateWindowsDeviceAccountActionParameter", m.GetUpdateWindowsDeviceAccountActionParameter())
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
func (m *UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetUpdateWindowsDeviceAccountActionParameter sets the updateWindowsDeviceAccountActionParameter property value. The updateWindowsDeviceAccountActionParameter property
func (m *UsersItemManagedDevicesItemUpdateWindowsDeviceAccountPostRequestBody) SetUpdateWindowsDeviceAccountActionParameter(value UpdateWindowsDeviceAccountActionParameterable)() {
    m.updateWindowsDeviceAccountActionParameter = value
}
