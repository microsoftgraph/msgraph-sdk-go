package validatepermission

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ValidatePermissionRequestBody provides operations to call the validatePermission method.
type ValidatePermissionRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The challengeToken property
    challengeToken *string;
    // The password property
    password *string;
}
// NewValidatePermissionRequestBody instantiates a new validatePermissionRequestBody and sets the default values.
func NewValidatePermissionRequestBody()(*ValidatePermissionRequestBody) {
    m := &ValidatePermissionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateValidatePermissionRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateValidatePermissionRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidatePermissionRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidatePermissionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChallengeToken gets the challengeToken property value. The challengeToken property
func (m *ValidatePermissionRequestBody) GetChallengeToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.challengeToken
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidatePermissionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["challengeToken"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChallengeToken(val)
        }
        return nil
    }
    res["password"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *ValidatePermissionRequestBody) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// Serialize serializes information the current object
func (m *ValidatePermissionRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ValidatePermissionRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChallengeToken sets the challengeToken property value. The challengeToken property
func (m *ValidatePermissionRequestBody) SetChallengeToken(value *string)() {
    if m != nil {
        m.challengeToken = value
    }
}
// SetPassword sets the password property value. The password property
func (m *ValidatePermissionRequestBody) SetPassword(value *string)() {
    if m != nil {
        m.password = value
    }
}
