package deleteuserfromsharedappledevice

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeleteUserFromSharedAppleDeviceRequestBody struct {
    additionalData map[string]interface{};
    userPrincipalName *string;
}
func NewDeleteUserFromSharedAppleDeviceRequestBody()(*DeleteUserFromSharedAppleDeviceRequestBody) {
    m := &DeleteUserFromSharedAppleDeviceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeleteUserFromSharedAppleDeviceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeleteUserFromSharedAppleDeviceRequestBody) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *DeleteUserFromSharedAppleDeviceRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *DeleteUserFromSharedAppleDeviceRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *DeleteUserFromSharedAppleDeviceRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *DeleteUserFromSharedAppleDeviceRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeleteUserFromSharedAppleDeviceRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
