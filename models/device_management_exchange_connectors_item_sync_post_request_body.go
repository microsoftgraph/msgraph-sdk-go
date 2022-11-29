package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExchangeConnectorsItemSyncPostRequestBody provides operations to call the sync method.
type DeviceManagementExchangeConnectorsItemSyncPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type of Exchange Connector sync requested.
    syncType *DeviceManagementExchangeConnectorSyncType
}
// NewDeviceManagementExchangeConnectorsItemSyncPostRequestBody instantiates a new DeviceManagementExchangeConnectorsItemSyncPostRequestBody and sets the default values.
func NewDeviceManagementExchangeConnectorsItemSyncPostRequestBody()(*DeviceManagementExchangeConnectorsItemSyncPostRequestBody) {
    m := &DeviceManagementExchangeConnectorsItemSyncPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementExchangeConnectorsItemSyncPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementExchangeConnectorsItemSyncPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementExchangeConnectorsItemSyncPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementExchangeConnectorsItemSyncPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeConnectorsItemSyncPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["syncType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementExchangeConnectorSyncType , m.SetSyncType)
    return res
}
// GetSyncType gets the syncType property value. The type of Exchange Connector sync requested.
func (m *DeviceManagementExchangeConnectorsItemSyncPostRequestBody) GetSyncType()(*DeviceManagementExchangeConnectorSyncType) {
    return m.syncType
}
// Serialize serializes information the current object
func (m *DeviceManagementExchangeConnectorsItemSyncPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSyncType() != nil {
        cast := (*m.GetSyncType()).String()
        err := writer.WriteStringValue("syncType", &cast)
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
func (m *DeviceManagementExchangeConnectorsItemSyncPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSyncType sets the syncType property value. The type of Exchange Connector sync requested.
func (m *DeviceManagementExchangeConnectorsItemSyncPostRequestBody) SetSyncType(value *DeviceManagementExchangeConnectorSyncType)() {
    m.syncType = value
}
