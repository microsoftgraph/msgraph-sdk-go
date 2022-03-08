package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSubmissionable 
type EducationSubmissionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetOutcomes()([]EducationOutcomeable)
    GetReassignedBy()(IdentitySetable)
    GetReassignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipient()(EducationSubmissionRecipientable)
    GetResources()([]EducationSubmissionResourceable)
    GetResourcesFolderUrl()(*string)
    GetReturnedBy()(IdentitySetable)
    GetReturnedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*EducationSubmissionStatus)
    GetSubmittedBy()(IdentitySetable)
    GetSubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubmittedResources()([]EducationSubmissionResourceable)
    GetUnsubmittedBy()(IdentitySetable)
    GetUnsubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetOutcomes(value []EducationOutcomeable)()
    SetReassignedBy(value IdentitySetable)()
    SetReassignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipient(value EducationSubmissionRecipientable)()
    SetResources(value []EducationSubmissionResourceable)()
    SetResourcesFolderUrl(value *string)()
    SetReturnedBy(value IdentitySetable)()
    SetReturnedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *EducationSubmissionStatus)()
    SetSubmittedBy(value IdentitySetable)()
    SetSubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubmittedResources(value []EducationSubmissionResourceable)()
    SetUnsubmittedBy(value IdentitySetable)()
    SetUnsubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
