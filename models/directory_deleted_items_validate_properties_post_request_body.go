package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryDeletedItemsValidatePropertiesPostRequestBody provides operations to call the validateProperties method.
type DirectoryDeletedItemsValidatePropertiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The displayName property
    displayName *string
    // The entityType property
    entityType *string
    // The mailNickname property
    mailNickname *string
    // The onBehalfOfUserId property
    onBehalfOfUserId *string
}
// NewDirectoryDeletedItemsValidatePropertiesPostRequestBody instantiates a new DirectoryDeletedItemsValidatePropertiesPostRequestBody and sets the default values.
func NewDirectoryDeletedItemsValidatePropertiesPostRequestBody()(*DirectoryDeletedItemsValidatePropertiesPostRequestBody) {
    m := &DirectoryDeletedItemsValidatePropertiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDirectoryDeletedItemsValidatePropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryDeletedItemsValidatePropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryDeletedItemsValidatePropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetEntityType gets the entityType property value. The entityType property
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) GetEntityType()(*string) {
    return m.entityType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["entityType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEntityType)
    res["mailNickname"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMailNickname)
    res["onBehalfOfUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnBehalfOfUserId)
    return res
}
// GetMailNickname gets the mailNickname property value. The mailNickname property
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetOnBehalfOfUserId gets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) GetOnBehalfOfUserId()(*string) {
    return m.onBehalfOfUserId
}
// Serialize serializes information the current object
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityType", m.GetEntityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onBehalfOfUserId", m.GetOnBehalfOfUserId())
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
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEntityType sets the entityType property value. The entityType property
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) SetEntityType(value *string)() {
    m.entityType = value
}
// SetMailNickname sets the mailNickname property value. The mailNickname property
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetOnBehalfOfUserId sets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *DirectoryDeletedItemsValidatePropertiesPostRequestBody) SetOnBehalfOfUserId(value *string)() {
    m.onBehalfOfUserId = value
}
