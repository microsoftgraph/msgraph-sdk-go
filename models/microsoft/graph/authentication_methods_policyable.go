package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthenticationMethodsPolicyable 
type AuthenticationMethodsPolicyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAuthenticationMethodConfigurations()([]AuthenticationMethodConfigurationable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPolicyVersion()(*string)
    GetReconfirmationInDays()(*int32)
    GetRegistrationEnforcement()(RegistrationEnforcementable)
    SetAuthenticationMethodConfigurations(value []AuthenticationMethodConfigurationable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPolicyVersion(value *string)()
    SetReconfirmationInDays(value *int32)()
    SetRegistrationEnforcement(value RegistrationEnforcementable)()
}
