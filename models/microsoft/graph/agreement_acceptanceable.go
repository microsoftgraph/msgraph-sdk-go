package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgreementAcceptanceable 
type AgreementAcceptanceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAgreementFileId()(*string)
    GetAgreementId()(*string)
    GetDeviceDisplayName()(*string)
    GetDeviceId()(*string)
    GetDeviceOSType()(*string)
    GetDeviceOSVersion()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetState()(*AgreementAcceptanceState)
    GetUserDisplayName()(*string)
    GetUserEmail()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetAgreementFileId(value *string)()
    SetAgreementId(value *string)()
    SetDeviceDisplayName(value *string)()
    SetDeviceId(value *string)()
    SetDeviceOSType(value *string)()
    SetDeviceOSVersion(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetState(value *AgreementAcceptanceState)()
    SetUserDisplayName(value *string)()
    SetUserEmail(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
