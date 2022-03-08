package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApprovalStageable 
type ApprovalStageable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignedToMe()(*bool)
    GetDisplayName()(*string)
    GetJustification()(*string)
    GetReviewedBy()(Identityable)
    GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewResult()(*string)
    GetStatus()(*string)
    SetAssignedToMe(value *bool)()
    SetDisplayName(value *string)()
    SetJustification(value *string)()
    SetReviewedBy(value Identityable)()
    SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewResult(value *string)()
    SetStatus(value *string)()
}
