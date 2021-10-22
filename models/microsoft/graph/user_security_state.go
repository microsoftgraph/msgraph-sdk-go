package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserSecurityState struct {
    aadUserId *string;
    accountName *string;
    additionalData map[string]interface{};
    domainName *string;
    emailRole *EmailRole;
    isVpn *bool;
    logonDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    logonId *string;
    logonIp *string;
    logonLocation *string;
    logonType *LogonType;
    onPremisesSecurityIdentifier *string;
    riskScore *string;
    userAccountType *UserAccountSecurityType;
    userPrincipalName *string;
}
func NewUserSecurityState()(*UserSecurityState) {
    m := &UserSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserSecurityState) GetAadUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aadUserId
    }
}
func (m *UserSecurityState) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
func (m *UserSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserSecurityState) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
func (m *UserSecurityState) GetEmailRole()(*EmailRole) {
    if m == nil {
        return nil
    } else {
        return m.emailRole
    }
}
func (m *UserSecurityState) GetIsVpn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVpn
    }
}
func (m *UserSecurityState) GetLogonDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.logonDateTime
    }
}
func (m *UserSecurityState) GetLogonId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logonId
    }
}
func (m *UserSecurityState) GetLogonIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logonIp
    }
}
func (m *UserSecurityState) GetLogonLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logonLocation
    }
}
func (m *UserSecurityState) GetLogonType()(*LogonType) {
    if m == nil {
        return nil
    } else {
        return m.logonType
    }
}
func (m *UserSecurityState) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
func (m *UserSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
func (m *UserSecurityState) GetUserAccountType()(*UserAccountSecurityType) {
    if m == nil {
        return nil
    } else {
        return m.userAccountType
    }
}
func (m *UserSecurityState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *UserSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aadUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAadUserId(val)
        return nil
    }
    res["accountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccountName(val)
        return nil
    }
    res["domainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDomainName(val)
        return nil
    }
    res["emailRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailRole)
        if err != nil {
            return err
        }
        cast := val.(EmailRole)
        m.SetEmailRole(&cast)
        return nil
    }
    res["isVpn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsVpn(val)
        return nil
    }
    res["logonDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLogonDateTime(val)
        return nil
    }
    res["logonId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLogonId(val)
        return nil
    }
    res["logonIp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLogonIp(val)
        return nil
    }
    res["logonLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLogonLocation(val)
        return nil
    }
    res["logonType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLogonType)
        if err != nil {
            return err
        }
        cast := val.(LogonType)
        m.SetLogonType(&cast)
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesSecurityIdentifier(val)
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskScore(val)
        return nil
    }
    res["userAccountType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserAccountSecurityType)
        if err != nil {
            return err
        }
        cast := val.(UserAccountSecurityType)
        m.SetUserAccountType(&cast)
        return nil
    }
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
func (m *UserSecurityState) IsNil()(bool) {
    return m == nil
}
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
func (m *UserSecurityState) SetAadUserId(value *string)() {
    m.aadUserId = value
}
func (m *UserSecurityState) SetAccountName(value *string)() {
    m.accountName = value
}
func (m *UserSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserSecurityState) SetDomainName(value *string)() {
    m.domainName = value
}
func (m *UserSecurityState) SetEmailRole(value *EmailRole)() {
    m.emailRole = value
}
func (m *UserSecurityState) SetIsVpn(value *bool)() {
    m.isVpn = value
}
func (m *UserSecurityState) SetLogonDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.logonDateTime = value
}
func (m *UserSecurityState) SetLogonId(value *string)() {
    m.logonId = value
}
func (m *UserSecurityState) SetLogonIp(value *string)() {
    m.logonIp = value
}
func (m *UserSecurityState) SetLogonLocation(value *string)() {
    m.logonLocation = value
}
func (m *UserSecurityState) SetLogonType(value *LogonType)() {
    m.logonType = value
}
func (m *UserSecurityState) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
func (m *UserSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
func (m *UserSecurityState) SetUserAccountType(value *UserAccountSecurityType)() {
    m.userAccountType = value
}
func (m *UserSecurityState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
