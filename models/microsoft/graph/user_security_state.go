package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserSecurityState struct {
    // AAD User object identifier (GUID) - represents the physical/multi-account user entity.
    aadUserId *string;
    // Account name of user account (without Active Directory domain or DNS domain) - (also called mailNickName).
    accountName *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // NetBIOS/Active Directory domain of user account (that is, domain/account format).
    domainName *string;
    // For email-related alerts - user account's email 'role'. Possible values are: unknown, sender, recipient.
    emailRole *EmailRole;
    // Indicates whether the user logged on through a VPN.
    isVpn *bool;
    // Time at which the sign-in occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    logonDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // User sign-in ID.
    logonId *string;
    // IP Address the sign-in request originated from.
    logonIp *string;
    // Location (by IP address mapping) associated with a user sign-in event by this user.
    logonLocation *string;
    // Method of user sign in. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
    logonType *LogonType;
    // Active Directory (on-premises) Security Identifier (SID) of the user.
    onPremisesSecurityIdentifier *string;
    // Provider-generated/calculated risk score of the user account. Recommended value range of 0-1, which equates to a percentage.
    riskScore *string;
    // User account type (group membership), per Windows definition. Possible values are: unknown, standard, power, administrator.
    userAccountType *UserAccountSecurityType;
    // User sign-in name - internet format: (user account name)@(user account DNS domain name).
    userPrincipalName *string;
}
// Instantiates a new userSecurityState and sets the default values.
func NewUserSecurityState()(*UserSecurityState) {
    m := &UserSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the aadUserId property value. AAD User object identifier (GUID) - represents the physical/multi-account user entity.
func (m *UserSecurityState) GetAadUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aadUserId
    }
}
// Gets the accountName property value. Account name of user account (without Active Directory domain or DNS domain) - (also called mailNickName).
func (m *UserSecurityState) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the domainName property value. NetBIOS/Active Directory domain of user account (that is, domain/account format).
func (m *UserSecurityState) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
// Gets the emailRole property value. For email-related alerts - user account's email 'role'. Possible values are: unknown, sender, recipient.
func (m *UserSecurityState) GetEmailRole()(*EmailRole) {
    if m == nil {
        return nil
    } else {
        return m.emailRole
    }
}
// Gets the isVpn property value. Indicates whether the user logged on through a VPN.
func (m *UserSecurityState) GetIsVpn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVpn
    }
}
// Gets the logonDateTime property value. Time at which the sign-in occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *UserSecurityState) GetLogonDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.logonDateTime
    }
}
// Gets the logonId property value. User sign-in ID.
func (m *UserSecurityState) GetLogonId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logonId
    }
}
// Gets the logonIp property value. IP Address the sign-in request originated from.
func (m *UserSecurityState) GetLogonIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logonIp
    }
}
// Gets the logonLocation property value. Location (by IP address mapping) associated with a user sign-in event by this user.
func (m *UserSecurityState) GetLogonLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logonLocation
    }
}
// Gets the logonType property value. Method of user sign in. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
func (m *UserSecurityState) GetLogonType()(*LogonType) {
    if m == nil {
        return nil
    } else {
        return m.logonType
    }
}
// Gets the onPremisesSecurityIdentifier property value. Active Directory (on-premises) Security Identifier (SID) of the user.
func (m *UserSecurityState) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
// Gets the riskScore property value. Provider-generated/calculated risk score of the user account. Recommended value range of 0-1, which equates to a percentage.
func (m *UserSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// Gets the userAccountType property value. User account type (group membership), per Windows definition. Possible values are: unknown, standard, power, administrator.
func (m *UserSecurityState) GetUserAccountType()(*UserAccountSecurityType) {
    if m == nil {
        return nil
    } else {
        return m.userAccountType
    }
}
// Gets the userPrincipalName property value. User sign-in name - internet format: (user account name)@(user account DNS domain name).
func (m *UserSecurityState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *UserSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aadUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadUserId(val)
        }
        return nil
    }
    res["accountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["domainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["emailRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailRole)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EmailRole)
            m.SetEmailRole(&cast)
        }
        return nil
    }
    res["isVpn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVpn(val)
        }
        return nil
    }
    res["logonDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonDateTime(val)
        }
        return nil
    }
    res["logonId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonId(val)
        }
        return nil
    }
    res["logonIp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonIp(val)
        }
        return nil
    }
    res["logonLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonLocation(val)
        }
        return nil
    }
    res["logonType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLogonType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(LogonType)
            m.SetLogonType(&cast)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
        }
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["userAccountType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserAccountSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(UserAccountSecurityType)
            m.SetUserAccountType(&cast)
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
func (m *UserSecurityState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("aadUserId", m.GetAadUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    if m.GetEmailRole() != nil {
        cast := m.GetEmailRole().String()
        err := writer.WriteStringValue("emailRole", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVpn", m.GetIsVpn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("logonDateTime", m.GetLogonDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logonId", m.GetLogonId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logonIp", m.GetLogonIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logonLocation", m.GetLogonLocation())
        if err != nil {
            return err
        }
    }
    if m.GetLogonType() != nil {
        cast := m.GetLogonType().String()
        err := writer.WriteStringValue("logonType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountType() != nil {
        cast := m.GetUserAccountType().String()
        err := writer.WriteStringValue("userAccountType", &cast)
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
// Sets the aadUserId property value. AAD User object identifier (GUID) - represents the physical/multi-account user entity.
// Parameters:
//  - value : Value to set for the aadUserId property.
func (m *UserSecurityState) SetAadUserId(value *string)() {
    m.aadUserId = value
}
// Sets the accountName property value. Account name of user account (without Active Directory domain or DNS domain) - (also called mailNickName).
// Parameters:
//  - value : Value to set for the accountName property.
func (m *UserSecurityState) SetAccountName(value *string)() {
    m.accountName = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the domainName property value. NetBIOS/Active Directory domain of user account (that is, domain/account format).
// Parameters:
//  - value : Value to set for the domainName property.
func (m *UserSecurityState) SetDomainName(value *string)() {
    m.domainName = value
}
// Sets the emailRole property value. For email-related alerts - user account's email 'role'. Possible values are: unknown, sender, recipient.
// Parameters:
//  - value : Value to set for the emailRole property.
func (m *UserSecurityState) SetEmailRole(value *EmailRole)() {
    m.emailRole = value
}
// Sets the isVpn property value. Indicates whether the user logged on through a VPN.
// Parameters:
//  - value : Value to set for the isVpn property.
func (m *UserSecurityState) SetIsVpn(value *bool)() {
    m.isVpn = value
}
// Sets the logonDateTime property value. Time at which the sign-in occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the logonDateTime property.
func (m *UserSecurityState) SetLogonDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.logonDateTime = value
}
// Sets the logonId property value. User sign-in ID.
// Parameters:
//  - value : Value to set for the logonId property.
func (m *UserSecurityState) SetLogonId(value *string)() {
    m.logonId = value
}
// Sets the logonIp property value. IP Address the sign-in request originated from.
// Parameters:
//  - value : Value to set for the logonIp property.
func (m *UserSecurityState) SetLogonIp(value *string)() {
    m.logonIp = value
}
// Sets the logonLocation property value. Location (by IP address mapping) associated with a user sign-in event by this user.
// Parameters:
//  - value : Value to set for the logonLocation property.
func (m *UserSecurityState) SetLogonLocation(value *string)() {
    m.logonLocation = value
}
// Sets the logonType property value. Method of user sign in. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
// Parameters:
//  - value : Value to set for the logonType property.
func (m *UserSecurityState) SetLogonType(value *LogonType)() {
    m.logonType = value
}
// Sets the onPremisesSecurityIdentifier property value. Active Directory (on-premises) Security Identifier (SID) of the user.
// Parameters:
//  - value : Value to set for the onPremisesSecurityIdentifier property.
func (m *UserSecurityState) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// Sets the riskScore property value. Provider-generated/calculated risk score of the user account. Recommended value range of 0-1, which equates to a percentage.
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *UserSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
// Sets the userAccountType property value. User account type (group membership), per Windows definition. Possible values are: unknown, standard, power, administrator.
// Parameters:
//  - value : Value to set for the userAccountType property.
func (m *UserSecurityState) SetUserAccountType(value *UserAccountSecurityType)() {
    m.userAccountType = value
}
// Sets the userPrincipalName property value. User sign-in name - internet format: (user account name)@(user account DNS domain name).
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *UserSecurityState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
