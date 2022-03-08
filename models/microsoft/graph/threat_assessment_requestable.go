package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ThreatAssessmentRequestable 
type ThreatAssessmentRequestable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetCategory()(*ThreatCategory)
    GetContentType()(*ThreatAssessmentContentType)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpectedAssessment()(*ThreatExpectedAssessment)
    GetRequestSource()(*ThreatAssessmentRequestSource)
    GetResults()([]ThreatAssessmentResultable)
    GetStatus()(*ThreatAssessmentStatus)
    SetCategory(value *ThreatCategory)()
    SetContentType(value *ThreatAssessmentContentType)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpectedAssessment(value *ThreatExpectedAssessment)()
    SetRequestSource(value *ThreatAssessmentRequestSource)()
    SetResults(value []ThreatAssessmentResultable)()
    SetStatus(value *ThreatAssessmentStatus)()
}
