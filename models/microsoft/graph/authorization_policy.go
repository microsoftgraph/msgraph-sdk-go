package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthorizationPolicy 
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
// NewAuthorizationPolicy instantiates a new authorizationPolicy and sets the default values.
func NewAuthorizationPolicy()(*AuthorizationPolicy) {
    m := &AuthorizationPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// GetAllowedToSignUpEmailBasedSubscriptions gets the allowedToSignUpEmailBasedSubscriptions property value. Indicates whether users can sign up for email based subscriptions.
func (m *AuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToSignUpEmailBasedSubscriptions
    }
}
// GetAllowedToUseSSPR gets the allowedToUseSSPR property value. Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
func (m *AuthorizationPolicy) GetAllowedToUseSSPR()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToUseSSPR
    }
}
// GetAllowEmailVerifiedUsersToJoinOrganization gets the allowEmailVerifiedUsersToJoinOrganization property value. Indicates whether a user can join the tenant by email validation.
func (m *AuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowEmailVerifiedUsersToJoinOrganization
    }
}
// GetAllowInvitesFrom gets the allowInvitesFrom property value. Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.  everyone is the default setting for all cloud environments except US Government. See more in the table below.
func (m *AuthorizationPolicy) GetAllowInvitesFrom()(*AllowInvitesFrom) {
    if m == nil {
        return nil
    } else {
        return m.allowInvitesFrom
    }
}
// GetBlockMsolPowerShell gets the blockMsolPowerShell property value. To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph.
func (m *AuthorizationPolicy) GetBlockMsolPowerShell()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockMsolPowerShell
    }
}
// GetDefaultUserRolePermissions gets the defaultUserRolePermissions property value. 
func (m *AuthorizationPolicy) GetDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    if m == nil {
        return nil
    } else {
        return m.defaultUserRolePermissions
    }
}
// GetGuestUserRoleId gets the guestUserRoleId property value. Represents role templateId for the role that should be granted to guest user. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
func (m *AuthorizationPolicy) GetGuestUserRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.guestUserRoleId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthorizationPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["allowedToSignUpEmailBasedSubscriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToSignUpEmailBasedSubscriptions(val)
        }
        return nil
    }
    res["allowedToUseSSPR"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToUseSSPR(val)
        }
        return nil
    }
    res["allowEmailVerifiedUsersToJoinOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEmailVerifiedUsersToJoinOrganization(val)
        }
        return nil
    }
    res["allowInvitesFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowInvitesFrom)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowInvitesFrom(val.(*AllowInvitesFrom))
        }
        return nil
    }
    res["blockMsolPowerShell"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockMsolPowerShell(val)
        }
        return nil
    }
    res["defaultUserRolePermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDefaultUserRolePermissions() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultUserRolePermissions(val.(*DefaultUserRolePermissions))
        }
        return nil
    }
    res["guestUserRoleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestUserRoleId(val)
        }
        return nil
    }
    return res
}
func (m *AuthorizationPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetAllowInvitesFrom()).String()
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
// SetAllowedToSignUpEmailBasedSubscriptions sets the allowedToSignUpEmailBasedSubscriptions property value. Indicates whether users can sign up for email based subscriptions.
func (m *AuthorizationPolicy) SetAllowedToSignUpEmailBasedSubscriptions(value *bool)() {
    if m != nil {
        m.allowedToSignUpEmailBasedSubscriptions = value
    }
}
// SetAllowedToUseSSPR sets the allowedToUseSSPR property value. Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
func (m *AuthorizationPolicy) SetAllowedToUseSSPR(value *bool)() {
    if m != nil {
        m.allowedToUseSSPR = value
    }
}
// SetAllowEmailVerifiedUsersToJoinOrganization sets the allowEmailVerifiedUsersToJoinOrganization property value. Indicates whether a user can join the tenant by email validation.
func (m *AuthorizationPolicy) SetAllowEmailVerifiedUsersToJoinOrganization(value *bool)() {
    if m != nil {
        m.allowEmailVerifiedUsersToJoinOrganization = value
    }
}
// SetAllowInvitesFrom sets the allowInvitesFrom property value. Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.  everyone is the default setting for all cloud environments except US Government. See more in the table below.
func (m *AuthorizationPolicy) SetAllowInvitesFrom(value *AllowInvitesFrom)() {
    if m != nil {
        m.allowInvitesFrom = value
    }
}
// SetBlockMsolPowerShell sets the blockMsolPowerShell property value. To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph.
func (m *AuthorizationPolicy) SetBlockMsolPowerShell(value *bool)() {
    if m != nil {
        m.blockMsolPowerShell = value
    }
}
// SetDefaultUserRolePermissions sets the defaultUserRolePermissions property value. 
func (m *AuthorizationPolicy) SetDefaultUserRolePermissions(value *DefaultUserRolePermissions)() {
    if m != nil {
        m.defaultUserRolePermissions = value
    }
}
// SetGuestUserRoleId sets the guestUserRoleId property value. Represents role templateId for the role that should be granted to guest user. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
func (m *AuthorizationPolicy) SetGuestUserRoleId(value *string)() {
    if m != nil {
        m.guestUserRoleId = value
    }
}
