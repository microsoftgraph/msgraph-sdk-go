package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SignInable 
type SignInable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicyable)
    GetClientAppUsed()(*string)
    GetConditionalAccessStatus()(*ConditionalAccessStatus)
    GetCorrelationId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceDetail()(DeviceDetailable)
    GetIpAddress()(*string)
    GetIsInteractive()(*bool)
    GetLocation()(SignInLocationable)
    GetResourceDisplayName()(*string)
    GetResourceId()(*string)
    GetRiskDetail()(*RiskDetail)
    GetRiskEventTypes()([]RiskEventType)
    GetRiskEventTypes_v2()([]string)
    GetRiskLevelAggregated()(*RiskLevel)
    GetRiskLevelDuringSignIn()(*RiskLevel)
    GetRiskState()(*RiskState)
    GetStatus()(SignInStatusable)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicyable)()
    SetClientAppUsed(value *string)()
    SetConditionalAccessStatus(value *ConditionalAccessStatus)()
    SetCorrelationId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceDetail(value DeviceDetailable)()
    SetIpAddress(value *string)()
    SetIsInteractive(value *bool)()
    SetLocation(value SignInLocationable)()
    SetResourceDisplayName(value *string)()
    SetResourceId(value *string)()
    SetRiskDetail(value *RiskDetail)()
    SetRiskEventTypes(value []RiskEventType)()
    SetRiskEventTypes_v2(value []string)()
    SetRiskLevelAggregated(value *RiskLevel)()
    SetRiskLevelDuringSignIn(value *RiskLevel)()
    SetRiskState(value *RiskState)()
    SetStatus(value SignInStatusable)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
