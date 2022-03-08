package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TermsAndConditionsAcceptanceStatusable 
type TermsAndConditionsAcceptanceStatusable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAcceptedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAcceptedVersion()(*int32)
    GetTermsAndConditions()(TermsAndConditionsable)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetAcceptedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAcceptedVersion(value *int32)()
    SetTermsAndConditions(value TermsAndConditionsable)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
