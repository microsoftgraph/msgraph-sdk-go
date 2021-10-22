package validatepermission

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ValidatePermissionRequestBody struct {
    additionalData map[string]interface{};
    challengeToken *string;
    password *string;
}
func NewValidatePermissionRequestBody()(*ValidatePermissionRequestBody) {
    m := &ValidatePermissionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ValidatePermissionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ValidatePermissionRequestBody) GetChallengeToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.challengeToken
    }
}
func (m *ValidatePermissionRequestBody) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
func (m *ValidatePermissionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["challengeToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChallengeToken(val)
        return nil
    }
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPassword(val)
        return nil
    }
    return res
}
func (m *ValidatePermissionRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ValidatePermissionRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *ValidatePermissionRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ValidatePermissionRequestBody) SetChallengeToken(value *string)() {
    m.challengeToken = value
}
func (m *ValidatePermissionRequestBody) SetPassword(value *string)() {
    m.password = value
}
