package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuthorizationPolicy struct {
    PolicyBase
    allowedToSignUpEmailBasedSubscriptions *bool;
    allowedToUseSSPR *bool;
    allowEmailVerifiedUsersToJoinOrganization *bool;
    allowInvitesFrom *AllowInvitesFrom;
    blockMsolPowerShell *bool;
    defaultUserRolePermissions *DefaultUserRolePermissions;
    guestUserRoleId *string;
}
func NewAuthorizationPolicy()(*AuthorizationPolicy) {
    m := &AuthorizationPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
func (m *AuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToSignUpEmailBasedSubscriptions
    }
}
func (m *AuthorizationPolicy) GetAllowedToUseSSPR()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToUseSSPR
    }
}
func (m *AuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowEmailVerifiedUsersToJoinOrganization
    }
}
func (m *AuthorizationPolicy) GetAllowInvitesFrom()(*AllowInvitesFrom) {
    if m == nil {
        return nil
    } else {
        return m.allowInvitesFrom
    }
}
func (m *AuthorizationPolicy) GetBlockMsolPowerShell()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockMsolPowerShell
    }
}
func (m *AuthorizationPolicy) GetDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    if m == nil {
        return nil
    } else {
        return m.defaultUserRolePermissions
    }
}
func (m *AuthorizationPolicy) GetGuestUserRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.guestUserRoleId
    }
}
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
func (m *AuthorizationPolicy) SetAllowedToSignUpEmailBasedSubscriptions(value *bool)() {
    m.allowedToSignUpEmailBasedSubscriptions = value
}
func (m *AuthorizationPolicy) SetAllowedToUseSSPR(value *bool)() {
    m.allowedToUseSSPR = value
}
func (m *AuthorizationPolicy) SetAllowEmailVerifiedUsersToJoinOrganization(value *bool)() {
    m.allowEmailVerifiedUsersToJoinOrganization = value
}
func (m *AuthorizationPolicy) SetAllowInvitesFrom(value *AllowInvitesFrom)() {
    m.allowInvitesFrom = value
}
func (m *AuthorizationPolicy) SetBlockMsolPowerShell(value *bool)() {
    m.blockMsolPowerShell = value
}
func (m *AuthorizationPolicy) SetDefaultUserRolePermissions(value *DefaultUserRolePermissions)() {
    m.defaultUserRolePermissions = value
}
func (m *AuthorizationPolicy) SetGuestUserRoleId(value *string)() {
    m.guestUserRoleId = value
}
