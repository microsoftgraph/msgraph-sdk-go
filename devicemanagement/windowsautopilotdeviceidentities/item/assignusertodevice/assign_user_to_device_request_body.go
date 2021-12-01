package assignusertodevice

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignUserToDeviceRequestBody 
type AssignUserToDeviceRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    addressableUserName *string;
    // 
    userPrincipalName *string;
}
// NewAssignUserToDeviceRequestBody instantiates a new assignUserToDeviceRequestBody and sets the default values.
func NewAssignUserToDeviceRequestBody()(*AssignUserToDeviceRequestBody) {
    m := &AssignUserToDeviceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignUserToDeviceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddressableUserName gets the addressableUserName property value. 
func (m *AssignUserToDeviceRequestBody) GetAddressableUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.addressableUserName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *AssignUserToDeviceRequestBody) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignUserToDeviceRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addressableUserName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressableUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *AssignUserToDeviceRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssignUserToDeviceRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignUserToDeviceRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddressableUserName sets the addressableUserName property value. 
func (m *AssignUserToDeviceRequestBody) SetAddressableUserName(value *string)() {
    if m != nil {
        m.addressableUserName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *AssignUserToDeviceRequestBody) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
