package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubjectRightsRequestable 
type SubjectRightsRequestable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAssignedTo()(Identityable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDataSubject()(DataSubjectable)
    GetDataSubjectType()(*DataSubjectType)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHistory()([]SubjectRightsRequestHistoryable)
    GetInsight()(SubjectRightsRequestDetailable)
    GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotes()([]AuthoredNoteable)
    GetRegulations()([]string)
    GetStages()([]SubjectRightsRequestStageDetailable)
    GetStatus()(*SubjectRightsRequestStatus)
    GetTeam()(Teamable)
    GetType()(*SubjectRightsRequestType)
    SetAssignedTo(value Identityable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDataSubject(value DataSubjectable)()
    SetDataSubjectType(value *DataSubjectType)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHistory(value []SubjectRightsRequestHistoryable)()
    SetInsight(value SubjectRightsRequestDetailable)()
    SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotes(value []AuthoredNoteable)()
    SetRegulations(value []string)()
    SetStages(value []SubjectRightsRequestStageDetailable)()
    SetStatus(value *SubjectRightsRequestStatus)()
    SetTeam(value Teamable)()
    SetType(value *SubjectRightsRequestType)()
}
