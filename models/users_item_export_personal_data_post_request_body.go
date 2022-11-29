package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemExportPersonalDataPostRequestBody provides operations to call the exportPersonalData method.
type UsersItemExportPersonalDataPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The storageLocation property
    storageLocation *string
}
// NewUsersItemExportPersonalDataPostRequestBody instantiates a new UsersItemExportPersonalDataPostRequestBody and sets the default values.
func NewUsersItemExportPersonalDataPostRequestBody()(*UsersItemExportPersonalDataPostRequestBody) {
    m := &UsersItemExportPersonalDataPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemExportPersonalDataPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemExportPersonalDataPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemExportPersonalDataPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemExportPersonalDataPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemExportPersonalDataPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["storageLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStorageLocation)
    return res
}
// GetStorageLocation gets the storageLocation property value. The storageLocation property
func (m *UsersItemExportPersonalDataPostRequestBody) GetStorageLocation()(*string) {
    return m.storageLocation
}
// Serialize serializes information the current object
func (m *UsersItemExportPersonalDataPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("storageLocation", m.GetStorageLocation())
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
func (m *UsersItemExportPersonalDataPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetStorageLocation sets the storageLocation property value. The storageLocation property
func (m *UsersItemExportPersonalDataPostRequestBody) SetStorageLocation(value *string)() {
    m.storageLocation = value
}
