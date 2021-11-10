package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PasswordProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // true if the user must change her password on the next login; otherwise false. If not set, default is false. NOTE:  For Azure B2C tenants, set to false and instead use custom policies and user flows to force password reset at first sign in. See Force password reset at first logon.
    forceChangePasswordNextSignIn *bool;
    // If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first perform a multi-factor authentication before password change. After a password change, this property will be automatically reset to false. If not set, default is false.
    forceChangePasswordNextSignInWithMfa *bool;
    // The password for the user. This property is required when a user is created. It can be updated, but the user will be required to change the password on the next login. The password must satisfy minimum requirements as specified by the user’s passwordPolicies property. By default, a strong password is required.
    password *string;
}
// Instantiates a new passwordProfile and sets the default values.
func NewPasswordProfile()(*PasswordProfile) {
    m := &PasswordProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordProfile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the forceChangePasswordNextSignIn property value. true if the user must change her password on the next login; otherwise false. If not set, default is false. NOTE:  For Azure B2C tenants, set to false and instead use custom policies and user flows to force password reset at first sign in. See Force password reset at first logon.
func (m *PasswordProfile) GetForceChangePasswordNextSignIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.forceChangePasswordNextSignIn
    }
}
// Gets the forceChangePasswordNextSignInWithMfa property value. If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first perform a multi-factor authentication before password change. After a password change, this property will be automatically reset to false. If not set, default is false.
func (m *PasswordProfile) GetForceChangePasswordNextSignInWithMfa()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.forceChangePasswordNextSignInWithMfa
    }
}
// Gets the password property value. The password for the user. This property is required when a user is created. It can be updated, but the user will be required to change the password on the next login. The password must satisfy minimum requirements as specified by the user’s passwordPolicies property. By default, a strong password is required.
func (m *PasswordProfile) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// The deserialization information for the current model
func (m *PasswordProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["forceChangePasswordNextSignIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceChangePasswordNextSignIn(val)
        }
        return nil
    }
    res["forceChangePasswordNextSignInWithMfa"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceChangePasswordNextSignInWithMfa(val)
        }
        return nil
    }
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *PasswordProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PasswordProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("forceChangePasswordNextSignIn", m.GetForceChangePasswordNextSignIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("forceChangePasswordNextSignInWithMfa", m.GetForceChangePasswordNextSignInWithMfa())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PasswordProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the forceChangePasswordNextSignIn property value. true if the user must change her password on the next login; otherwise false. If not set, default is false. NOTE:  For Azure B2C tenants, set to false and instead use custom policies and user flows to force password reset at first sign in. See Force password reset at first logon.
// Parameters:
//  - value : Value to set for the forceChangePasswordNextSignIn property.
func (m *PasswordProfile) SetForceChangePasswordNextSignIn(value *bool)() {
    m.forceChangePasswordNextSignIn = value
}
// Sets the forceChangePasswordNextSignInWithMfa property value. If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first perform a multi-factor authentication before password change. After a password change, this property will be automatically reset to false. If not set, default is false.
// Parameters:
//  - value : Value to set for the forceChangePasswordNextSignInWithMfa property.
func (m *PasswordProfile) SetForceChangePasswordNextSignInWithMfa(value *bool)() {
    m.forceChangePasswordNextSignInWithMfa = value
}
// Sets the password property value. The password for the user. This property is required when a user is created. It can be updated, but the user will be required to change the password on the next login. The password must satisfy minimum requirements as specified by the user’s passwordPolicies property. By default, a strong password is required.
// Parameters:
//  - value : Value to set for the password property.
func (m *PasswordProfile) SetPassword(value *string)() {
    m.password = value
}
