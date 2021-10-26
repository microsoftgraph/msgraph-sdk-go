package changepassword

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChangePasswordRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    currentPassword *string;
    // 
    newPassword *string;
}
// Instantiates a new changePasswordRequestBody and sets the default values.
func NewChangePasswordRequestBody()(*ChangePasswordRequestBody) {
    m := &ChangePasswordRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangePasswordRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the currentPassword property value. 
func (m *ChangePasswordRequestBody) GetCurrentPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currentPassword
    }
}
// Gets the newPassword property value. 
func (m *ChangePasswordRequestBody) GetNewPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newPassword
    }
}
// The deserialization information for the current model
func (m *ChangePasswordRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["currentPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCurrentPassword(val)
        return nil
    }
    res["newPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNewPassword(val)
        return nil
    }
    return res
}
func (m *ChangePasswordRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ChangePasswordRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the currentPassword property value. 
// Parameters:
//  - value : Value to set for the currentPassword property.
func (m *ChangePasswordRequestBody) SetCurrentPassword(value *string)() {
    m.currentPassword = value
}
// Sets the newPassword property value. 
// Parameters:
//  - value : Value to set for the newPassword property.
func (m *ChangePasswordRequestBody) SetNewPassword(value *string)() {
    m.newPassword = value
}
