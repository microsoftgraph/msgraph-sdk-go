package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody provides operations to call the validatePermission method.
type UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The challengeToken property
    challengeToken *string
    // The password property
    password *string
}
// NewUsersItemDrivesItemItemsItemValidatePermissionPostRequestBody instantiates a new UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody and sets the default values.
func NewUsersItemDrivesItemItemsItemValidatePermissionPostRequestBody()(*UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) {
    m := &UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemDrivesItemItemsItemValidatePermissionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemDrivesItemItemsItemValidatePermissionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemDrivesItemItemsItemValidatePermissionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetChallengeToken gets the challengeToken property value. The challengeToken property
func (m *UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) GetChallengeToken()(*string) {
    return m.challengeToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["challengeToken"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetChallengeToken)
    res["password"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPassword)
    return res
}
// GetPassword gets the password property value. The password property
func (m *UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) GetPassword()(*string) {
    return m.password
}
// Serialize serializes information the current object
func (m *UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("challengeToken", m.GetChallengeToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
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
func (m *UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetChallengeToken sets the challengeToken property value. The challengeToken property
func (m *UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) SetChallengeToken(value *string)() {
    m.challengeToken = value
}
// SetPassword sets the password property value. The password property
func (m *UsersItemDrivesItemItemsItemValidatePermissionPostRequestBody) SetPassword(value *string)() {
    m.password = value
}
