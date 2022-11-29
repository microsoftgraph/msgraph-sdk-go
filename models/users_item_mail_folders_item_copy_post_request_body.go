package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemMailFoldersItemCopyPostRequestBody provides operations to call the copy method.
type UsersItemMailFoldersItemCopyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The DestinationId property
    destinationId *string
}
// NewUsersItemMailFoldersItemCopyPostRequestBody instantiates a new UsersItemMailFoldersItemCopyPostRequestBody and sets the default values.
func NewUsersItemMailFoldersItemCopyPostRequestBody()(*UsersItemMailFoldersItemCopyPostRequestBody) {
    m := &UsersItemMailFoldersItemCopyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemMailFoldersItemCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemMailFoldersItemCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemMailFoldersItemCopyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemMailFoldersItemCopyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDestinationId gets the destinationId property value. The DestinationId property
func (m *UsersItemMailFoldersItemCopyPostRequestBody) GetDestinationId()(*string) {
    return m.destinationId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemMailFoldersItemCopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDestinationId)
    return res
}
// Serialize serializes information the current object
func (m *UsersItemMailFoldersItemCopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationId", m.GetDestinationId())
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
func (m *UsersItemMailFoldersItemCopyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDestinationId sets the destinationId property value. The DestinationId property
func (m *UsersItemMailFoldersItemCopyPostRequestBody) SetDestinationId(value *string)() {
    m.destinationId = value
}
