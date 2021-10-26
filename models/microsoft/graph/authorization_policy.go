package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthorizationPolicy struct {
    PolicyBase
    // Indicates whether users can sign up for email based subscriptions.
    allowedToSignUpEmailBasedSubscriptions *bool;
    // Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
    allowedToUseSSPR *bool;
    // Indicates whether a user can join the tenant by email validation.
    allowEmailVerifiedUsersToJoinOrganization *bool;
    // Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.  everyone is the default setting for all cloud environments except US Government. See more in the table below.
    allowInvitesFrom *AllowInvitesFrom;
    // To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph.
    blockMsolPowerShell *bool;
    // 
    defaultUserRolePermissions *DefaultUserRolePermissions;
    // Represents role templateId for the role that should be granted to guest user. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
    guestUserRoleId *string;
}
// Instantiates a new authorizationPolicy and sets the default values.
func NewAuthorizationPolicy()(*AuthorizationPolicy) {
    m := &AuthorizationPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// Gets the allowedToSignUpEmailBasedSubscriptions property value. Indicates whether users can sign up for email based subscriptions.
func (m *AuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToSignUpEmailBasedSubscriptions
    }
}
// Gets the allowedToUseSSPR property value. Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
func (m *AuthorizationPolicy) GetAllowedToUseSSPR()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToUseSSPR
    }
}
// Gets the allowEmailVerifiedUsersToJoinOrganization property value. Indicates whether a user can join the tenant by email validation.
func (m *AuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowEmailVerifiedUsersToJoinOrganization
    }
}
// Gets the allowInvitesFrom property value. Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.  everyone is the default setting for all cloud environments except US Government. See more in the table below.
func (m *AuthorizationPolicy) GetAllowInvitesFrom()(*AllowInvitesFrom) {
    if m == nil {
        return nil
    } else {
        return m.allowInvitesFrom
    }
}
// Gets the blockMsolPowerShell property value. To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph.
func (m *AuthorizationPolicy) GetBlockMsolPowerShell()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockMsolPowerShell
    }
}
// Gets the defaultUserRolePermissions property value. 
func (m *AuthorizationPolicy) GetDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    if m == nil {
        return nil
    } else {
        return m.defaultUserRolePermissions
    }
}
// Gets the guestUserRoleId property value. Represents role templateId for the role that should be granted to guest user. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
func (m *AuthorizationPolicy) GetGuestUserRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.guestUserRoleId
    }
}
// The deserialization information for the current model
func (m *AuthorizationPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["allowedToSignUpEmailBasedSubscriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowedToSignUpEmailBasedSubscriptions(val)
        return nil
    }
    res["allowedToUseSSPR"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowedToUseSSPR(val)
        return nil
    }
    res["allowEmailVerifiedUsersToJoinOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowEmailVerifiedUsersToJoinOrganization(val)
        return nil
    }
    res["allowInvitesFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowInvitesFrom)
        if err != nil {
            return err
        }
        cast := val.(AllowInvitesFrom)
        m.SetAllowInvitesFrom(&cast)
        return nil
    }
    res["blockMsolPowerShell"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetBlockMsolPowerShell(val)
        return nil
    }
    res["defaultUserRolePermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDefaultUserRolePermissions() })
        if err != nil {
            return err
        }
        m.SetDefaultUserRolePermissions(val.(*DefaultUserRolePermissions))
        return nil
    }
    res["guestUserRoleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGuestUserRoleId(val)
        return nil
    }
    return res
}
func (m *AuthorizationPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AuthorizationPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowedToSignUpEmailBasedSubscriptions", m.GetAllowedToSignUpEmailBasedSubscriptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowedToUseSSPR", m.GetAllowedToUseSSPR())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowEmailVerifiedUsersToJoinOrganization", m.GetAllowEmailVerifiedUsersToJoinOrganization())
        if err != nil {
            return err
        }
    }
    if m.GetAllowInvitesFrom() != nil {
        cast := m.GetAllowInvitesFrom().String()
        err = writer.WriteStringValue("allowInvitesFrom", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockMsolPowerShell", m.GetBlockMsolPowerShell())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultUserRolePermissions", m.GetDefaultUserRolePermissions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("guestUserRoleId", m.GetGuestUserRoleId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowedToSignUpEmailBasedSubscriptions property value. Indicates whether users can sign up for email based subscriptions.
// Parameters:
//  - value : Value to set for the allowedToSignUpEmailBasedSubscriptions property.
func (m *AuthorizationPolicy) SetAllowedToSignUpEmailBasedSubscriptions(value *bool)() {
    m.allowedToSignUpEmailBasedSubscriptions = value
}
// Sets the allowedToUseSSPR property value. Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
// Parameters:
//  - value : Value to set for the allowedToUseSSPR property.
func (m *AuthorizationPolicy) SetAllowedToUseSSPR(value *bool)() {
    m.allowedToUseSSPR = value
}
// Sets the allowEmailVerifiedUsersToJoinOrganization property value. Indicates whether a user can join the tenant by email validation.
// Parameters:
//  - value : Value to set for the allowEmailVerifiedUsersToJoinOrganization property.
func (m *AuthorizationPolicy) SetAllowEmailVerifiedUsersToJoinOrganization(value *bool)() {
    m.allowEmailVerifiedUsersToJoinOrganization = value
}
// Sets the allowInvitesFrom property value. Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.  everyone is the default setting for all cloud environments except US Government. See more in the table below.
// Parameters:
//  - value : Value to set for the allowInvitesFrom property.
func (m *AuthorizationPolicy) SetAllowInvitesFrom(value *AllowInvitesFrom)() {
    m.allowInvitesFrom = value
}
// Sets the blockMsolPowerShell property value. To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph.
// Parameters:
//  - value : Value to set for the blockMsolPowerShell property.
func (m *AuthorizationPolicy) SetBlockMsolPowerShell(value *bool)() {
    m.blockMsolPowerShell = value
}
// Sets the defaultUserRolePermissions property value. 
// Parameters:
//  - value : Value to set for the defaultUserRolePermissions property.
func (m *AuthorizationPolicy) SetDefaultUserRolePermissions(value *DefaultUserRolePermissions)() {
    m.defaultUserRolePermissions = value
}
// Sets the guestUserRoleId property value. Represents role templateId for the role that should be granted to guest user. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
// Parameters:
//  - value : Value to set for the guestUserRoleId property.
func (m *AuthorizationPolicy) SetGuestUserRoleId(value *string)() {
    m.guestUserRoleId = value
}
