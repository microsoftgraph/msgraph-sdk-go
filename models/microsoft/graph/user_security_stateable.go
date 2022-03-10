package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserSecurityStateable 
type UserSecurityStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAadUserId()(*string)
    GetAccountName()(*string)
    GetDomainName()(*string)
    GetEmailRole()(*EmailRole)
    GetIsVpn()(*bool)
    GetLogonDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogonId()(*string)
    GetLogonIp()(*string)
    GetLogonLocation()(*string)
    GetLogonType()(*LogonType)
    GetOnPremisesSecurityIdentifier()(*string)
    GetRiskScore()(*string)
    GetUserAccountType()(*UserAccountSecurityType)
    GetUserPrincipalName()(*string)
    SetAadUserId(value *string)()
    SetAccountName(value *string)()
    SetDomainName(value *string)()
    SetEmailRole(value *EmailRole)()
    SetIsVpn(value *bool)()
    SetLogonDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogonId(value *string)()
    SetLogonIp(value *string)()
    SetLogonLocation(value *string)()
    SetLogonType(value *LogonType)()
    SetOnPremisesSecurityIdentifier(value *string)()
    SetRiskScore(value *string)()
    SetUserAccountType(value *UserAccountSecurityType)()
    SetUserPrincipalName(value *string)()
}
