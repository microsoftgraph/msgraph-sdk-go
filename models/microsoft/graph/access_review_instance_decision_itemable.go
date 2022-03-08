package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewInstanceDecisionItemable 
type AccessReviewInstanceDecisionItemable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessReviewId()(*string)
    GetAppliedBy()(UserIdentityable)
    GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetApplyResult()(*string)
    GetDecision()(*string)
    GetJustification()(*string)
    GetPrincipal()(Identityable)
    GetPrincipalLink()(*string)
    GetRecommendation()(*string)
    GetResource()(AccessReviewInstanceDecisionItemResourceable)
    GetResourceLink()(*string)
    GetReviewedBy()(UserIdentityable)
    GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessReviewId(value *string)()
    SetAppliedBy(value UserIdentityable)()
    SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetApplyResult(value *string)()
    SetDecision(value *string)()
    SetJustification(value *string)()
    SetPrincipal(value Identityable)()
    SetPrincipalLink(value *string)()
    SetRecommendation(value *string)()
    SetResource(value AccessReviewInstanceDecisionItemResourceable)()
    SetResourceLink(value *string)()
    SetReviewedBy(value UserIdentityable)()
    SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
