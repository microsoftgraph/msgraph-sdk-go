package changepassword

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChangePasswordRequestBody provides operations to call the changePassword method.
type ChangePasswordRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    currentPassword *string;
    // 
    newPassword *string;
}
// NewChangePasswordRequestBody instantiates a new changePasswordRequestBody and sets the default values.
func NewChangePasswordRequestBody()(*ChangePasswordRequestBody) {
    m := &ChangePasswordRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChangePasswordRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChangePasswordRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewChangePasswordRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangePasswordRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCurrentPassword gets the currentPassword property value. 
func (m *ChangePasswordRequestBody) GetCurrentPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currentPassword
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangePasswordRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["currentPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPassword(val)
        }
        return nil
    }
    res["newPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewPassword(val)
        }
        return nil
    }
    return res
}
// GetNewPassword gets the newPassword property value. 
func (m *ChangePasswordRequestBody) GetNewPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newPassword
    }
}
// Serialize serializes information the current object
func (m *ChangePasswordRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("currentPassword", m.GetCurrentPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newPassword", m.GetNewPassword())
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
func (m *ChangePasswordRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCurrentPassword sets the currentPassword property value. 
func (m *ChangePasswordRequestBody) SetCurrentPassword(value *string)() {
    if m != nil {
        m.currentPassword = value
    }
}
// SetNewPassword sets the newPassword property value. 
func (m *ChangePasswordRequestBody) SetNewPassword(value *string)() {
    if m != nil {
        m.newPassword = value
    }
}
