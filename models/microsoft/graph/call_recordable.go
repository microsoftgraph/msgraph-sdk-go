package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/callrecords"
)

// CallRecordable 
type CallRecordable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetJoinWebUrl()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetModalities()([]i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality)
    GetOrganizer()(IdentitySetable)
    GetParticipants()([]IdentitySetable)
    GetSessions()([]Sessionable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetType()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.CallType)
    GetVersion()(*int64)
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetJoinWebUrl(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetModalities(value []i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality)()
    SetOrganizer(value IdentitySetable)()
    SetParticipants(value []IdentitySetable)()
    SetSessions(value []Sessionable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetType(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.CallType)()
    SetVersion(value *int64)()
}
