package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// userIdentity 
type UserIdentity struct {
    Identity
    // Indicates the client IP address used by user performing the activity (audit log only).
    ipAddress *string;
    // The userPrincipalName attribute of the user.
    userPrincipalName *string;
}
// NewUserIdentity instantiates a new userIdentity and sets the default values.
func NewUserIdentity()(*UserIdentity) {
    m := &UserIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// GetIpAddress gets the ipAddress property value. Indicates the client IP address used by user performing the activity (audit log only).
func (m *UserIdentity) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName attribute of the user.
func (m *UserIdentity) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
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
func (m *UserIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIpAddress sets the ipAddress property value. Indicates the client IP address used by user performing the activity (audit log only).
func (m *UserIdentity) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName attribute of the user.
func (m *UserIdentity) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
