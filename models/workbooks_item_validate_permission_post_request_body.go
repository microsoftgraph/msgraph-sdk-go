package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbooksItemValidatePermissionPostRequestBody provides operations to call the validatePermission method.
type WorkbooksItemValidatePermissionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The challengeToken property
    challengeToken *string
    // The password property
    password *string
}
// NewWorkbooksItemValidatePermissionPostRequestBody instantiates a new WorkbooksItemValidatePermissionPostRequestBody and sets the default values.
func NewWorkbooksItemValidatePermissionPostRequestBody()(*WorkbooksItemValidatePermissionPostRequestBody) {
    m := &WorkbooksItemValidatePermissionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbooksItemValidatePermissionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbooksItemValidatePermissionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbooksItemValidatePermissionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbooksItemValidatePermissionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetChallengeToken gets the challengeToken property value. The challengeToken property
func (m *WorkbooksItemValidatePermissionPostRequestBody) GetChallengeToken()(*string) {
    return m.challengeToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbooksItemValidatePermissionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["challengeToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChallengeToken(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
func (m *WorkbooksItemValidatePermissionPostRequestBody) GetPassword()(*string) {
    return m.password
}
// Serialize serializes information the current object
func (m *WorkbooksItemValidatePermissionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WorkbooksItemValidatePermissionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetChallengeToken sets the challengeToken property value. The challengeToken property
func (m *WorkbooksItemValidatePermissionPostRequestBody) SetChallengeToken(value *string)() {
    m.challengeToken = value
}
// SetPassword sets the password property value. The password property
func (m *WorkbooksItemValidatePermissionPostRequestBody) SetPassword(value *string)() {
    m.password = value
}
