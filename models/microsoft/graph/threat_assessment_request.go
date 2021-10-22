package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ThreatAssessmentRequest struct {
    Entity
    category *ThreatCategory;
    contentType *ThreatAssessmentContentType;
    createdBy *IdentitySet;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    expectedAssessment *ThreatExpectedAssessment;
    requestSource *ThreatAssessmentRequestSource;
    results []ThreatAssessmentResult;
    status *ThreatAssessmentStatus;
}
func NewThreatAssessmentRequest()(*ThreatAssessmentRequest) {
    m := &ThreatAssessmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ThreatAssessmentRequest) GetCategory()(*ThreatCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *ThreatAssessmentRequest) GetContentType()(*ThreatAssessmentContentType) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
func (m *ThreatAssessmentRequest) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
func (m *ThreatAssessmentRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *ThreatAssessmentRequest) GetExpectedAssessment()(*ThreatExpectedAssessment) {
    if m == nil {
        return nil
    } else {
        return m.expectedAssessment
    }
}
func (m *ThreatAssessmentRequest) GetRequestSource()(*ThreatAssessmentRequestSource) {
    if m == nil {
        return nil
    } else {
        return m.requestSource
    }
}
func (m *ThreatAssessmentRequest) GetResults()([]ThreatAssessmentResult) {
    if m == nil {
        return nil
    } else {
        return m.results
    }
}
func (m *ThreatAssessmentRequest) GetStatus()(*ThreatAssessmentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ThreatAssessmentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatCategory)
        if err != nil {
            return err
        }
        cast := val.(ThreatCategory)
        m.SetCategory(&cast)
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentContentType)
        if err != nil {
            return err
        }
        cast := val.(ThreatAssessmentContentType)
        m.SetContentType(&cast)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*IdentitySet))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["expectedAssessment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatExpectedAssessment)
        if err != nil {
            return err
        }
        cast := val.(ThreatExpectedAssessment)
        m.SetExpectedAssessment(&cast)
        return nil
    }
    res["requestSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentRequestSource)
        if err != nil {
            return err
        }
        cast := val.(ThreatAssessmentRequestSource)
        m.SetRequestSource(&cast)
        return nil
    }
    res["results"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThreatAssessmentResult() })
        if err != nil {
            return err
        }
        res := make([]ThreatAssessmentResult, len(val))
        for i, v := range val {
            res[i] = *(v.(*ThreatAssessmentResult))
        }
        m.SetResults(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentStatus)
        if err != nil {
            return err
        }
        cast := val.(ThreatAssessmentStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *ThreatAssessmentRequest) IsNil()(bool) {
    return m == nil
}
func (m *ThreatAssessmentRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := m.GetCategory().String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentType() != nil {
        cast := m.GetContentType().String()
        err = writer.WriteStringValue("contentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetExpectedAssessment() != nil {
        cast := m.GetExpectedAssessment().String()
        err = writer.WriteStringValue("expectedAssessment", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequestSource() != nil {
        cast := m.GetRequestSource().String()
        err = writer.WriteStringValue("requestSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResults()))
        for i, v := range m.GetResults() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("results", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ThreatAssessmentRequest) SetCategory(value *ThreatCategory)() {
    m.category = value
}
func (m *ThreatAssessmentRequest) SetContentType(value *ThreatAssessmentContentType)() {
    m.contentType = value
}
func (m *ThreatAssessmentRequest) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
func (m *ThreatAssessmentRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *ThreatAssessmentRequest) SetExpectedAssessment(value *ThreatExpectedAssessment)() {
    m.expectedAssessment = value
}
func (m *ThreatAssessmentRequest) SetRequestSource(value *ThreatAssessmentRequestSource)() {
    m.requestSource = value
}
func (m *ThreatAssessmentRequest) SetResults(value []ThreatAssessmentResult)() {
    m.results = value
}
func (m *ThreatAssessmentRequest) SetStatus(value *ThreatAssessmentStatus)() {
    m.status = value
}
