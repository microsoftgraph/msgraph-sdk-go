package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewInstanceDecisionItem struct {
    Entity
    accessReviewId *string;
    appliedBy *UserIdentity;
    appliedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    applyResult *string;
    decision *string;
    justification *string;
    principal *Identity;
    principalLink *string;
    recommendation *string;
    resource *AccessReviewInstanceDecisionItemResource;
    resourceLink *string;
    reviewedBy *UserIdentity;
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewAccessReviewInstanceDecisionItem()(*AccessReviewInstanceDecisionItem) {
    m := &AccessReviewInstanceDecisionItem{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AccessReviewInstanceDecisionItem) GetAccessReviewId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewId
    }
}
func (m *AccessReviewInstanceDecisionItem) GetAppliedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.appliedBy
    }
}
func (m *AccessReviewInstanceDecisionItem) GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.appliedDateTime
    }
}
func (m *AccessReviewInstanceDecisionItem) GetApplyResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applyResult
    }
}
func (m *AccessReviewInstanceDecisionItem) GetDecision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decision
    }
}
func (m *AccessReviewInstanceDecisionItem) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
func (m *AccessReviewInstanceDecisionItem) GetPrincipal()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.principal
    }
}
func (m *AccessReviewInstanceDecisionItem) GetPrincipalLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalLink
    }
}
func (m *AccessReviewInstanceDecisionItem) GetRecommendation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendation
    }
}
func (m *AccessReviewInstanceDecisionItem) GetResource()(*AccessReviewInstanceDecisionItemResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *AccessReviewInstanceDecisionItem) GetResourceLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceLink
    }
}
func (m *AccessReviewInstanceDecisionItem) GetReviewedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.reviewedBy
    }
}
func (m *AccessReviewInstanceDecisionItem) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewedDateTime
    }
}
func (m *AccessReviewInstanceDecisionItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessReviewId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccessReviewId(val)
        return nil
    }
    res["appliedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetAppliedBy(val.(*UserIdentity))
        return nil
    }
    res["appliedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAppliedDateTime(val)
        return nil
    }
    res["applyResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplyResult(val)
        return nil
    }
    res["decision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDecision(val)
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustification(val)
        return nil
    }
    res["principal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetPrincipal(val.(*Identity))
        return nil
    }
    res["principalLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrincipalLink(val)
        return nil
    }
    res["recommendation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecommendation(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewInstanceDecisionItemResource() })
        if err != nil {
            return err
        }
        m.SetResource(val.(*AccessReviewInstanceDecisionItemResource))
        return nil
    }
    res["resourceLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceLink(val)
        return nil
    }
    res["reviewedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetReviewedBy(val.(*UserIdentity))
        return nil
    }
    res["reviewedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReviewedDateTime(val)
        return nil
    }
    return res
}
func (m *AccessReviewInstanceDecisionItem) IsNil()(bool) {
    return m == nil
}
func (m *AccessReviewInstanceDecisionItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accessReviewId", m.GetAccessReviewId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appliedBy", m.GetAppliedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("appliedDateTime", m.GetAppliedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applyResult", m.GetApplyResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("decision", m.GetDecision())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalLink", m.GetPrincipalLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendation", m.GetRecommendation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceLink", m.GetResourceLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewedBy", m.GetReviewedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewedDateTime", m.GetReviewedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AccessReviewInstanceDecisionItem) SetAccessReviewId(value *string)() {
    m.accessReviewId = value
}
func (m *AccessReviewInstanceDecisionItem) SetAppliedBy(value *UserIdentity)() {
    m.appliedBy = value
}
func (m *AccessReviewInstanceDecisionItem) SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.appliedDateTime = value
}
func (m *AccessReviewInstanceDecisionItem) SetApplyResult(value *string)() {
    m.applyResult = value
}
func (m *AccessReviewInstanceDecisionItem) SetDecision(value *string)() {
    m.decision = value
}
func (m *AccessReviewInstanceDecisionItem) SetJustification(value *string)() {
    m.justification = value
}
func (m *AccessReviewInstanceDecisionItem) SetPrincipal(value *Identity)() {
    m.principal = value
}
func (m *AccessReviewInstanceDecisionItem) SetPrincipalLink(value *string)() {
    m.principalLink = value
}
func (m *AccessReviewInstanceDecisionItem) SetRecommendation(value *string)() {
    m.recommendation = value
}
func (m *AccessReviewInstanceDecisionItem) SetResource(value *AccessReviewInstanceDecisionItemResource)() {
    m.resource = value
}
func (m *AccessReviewInstanceDecisionItem) SetResourceLink(value *string)() {
    m.resourceLink = value
}
func (m *AccessReviewInstanceDecisionItem) SetReviewedBy(value *UserIdentity)() {
    m.reviewedBy = value
}
func (m *AccessReviewInstanceDecisionItem) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
