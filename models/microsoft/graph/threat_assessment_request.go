package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ThreatAssessmentRequest 
type ThreatAssessmentRequest struct {
    Entity
    // The threat category. Possible values are: spam, phishing, malware.
    category *ThreatCategory;
    // The content type of threat assessment. Possible values are: mail, url, file.
    contentType *ThreatAssessmentContentType;
    // The threat assessment request creator.
    createdBy *IdentitySet;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The expected assessment from submitter. Possible values are: block, unblock.
    expectedAssessment *ThreatExpectedAssessment;
    // The source of the threat assessment request. Possible values are: administrator.
    requestSource *ThreatAssessmentRequestSource;
    // A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
    results []ThreatAssessmentResult;
    // The assessment process status. Possible values are: pending, completed.
    status *ThreatAssessmentStatus;
}
// NewThreatAssessmentRequest instantiates a new threatAssessmentRequest and sets the default values.
func NewThreatAssessmentRequest()(*ThreatAssessmentRequest) {
    m := &ThreatAssessmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategory gets the category property value. The threat category. Possible values are: spam, phishing, malware.
func (m *ThreatAssessmentRequest) GetCategory()(*ThreatCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetContentType gets the contentType property value. The content type of threat assessment. Possible values are: mail, url, file.
func (m *ThreatAssessmentRequest) GetContentType()(*ThreatAssessmentContentType) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetCreatedBy gets the createdBy property value. The threat assessment request creator.
func (m *ThreatAssessmentRequest) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ThreatAssessmentRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetExpectedAssessment gets the expectedAssessment property value. The expected assessment from submitter. Possible values are: block, unblock.
func (m *ThreatAssessmentRequest) GetExpectedAssessment()(*ThreatExpectedAssessment) {
    if m == nil {
        return nil
    } else {
        return m.expectedAssessment
    }
}
// GetRequestSource gets the requestSource property value. The source of the threat assessment request. Possible values are: administrator.
func (m *ThreatAssessmentRequest) GetRequestSource()(*ThreatAssessmentRequestSource) {
    if m == nil {
        return nil
    } else {
        return m.requestSource
    }
}
// GetResults gets the results property value. A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
func (m *ThreatAssessmentRequest) GetResults()([]ThreatAssessmentResult) {
    if m == nil {
        return nil
    } else {
        return m.results
    }
}
// GetStatus gets the status property value. The assessment process status. Possible values are: pending, completed.
func (m *ThreatAssessmentRequest) GetStatus()(*ThreatAssessmentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThreatAssessmentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatCategory)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ThreatCategory)
            m.SetCategory(&cast)
        }
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentContentType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ThreatAssessmentContentType)
            m.SetContentType(&cast)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["expectedAssessment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatExpectedAssessment)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ThreatExpectedAssessment)
            m.SetExpectedAssessment(&cast)
        }
        return nil
    }
    res["requestSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentRequestSource)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ThreatAssessmentRequestSource)
            m.SetRequestSource(&cast)
        }
        return nil
    }
    res["results"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThreatAssessmentResult() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThreatAssessmentResult, len(val))
            for i, v := range val {
                res[i] = *(v.(*ThreatAssessmentResult))
            }
            m.SetResults(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ThreatAssessmentStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *ThreatAssessmentRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCategory sets the category property value. The threat category. Possible values are: spam, phishing, malware.
func (m *ThreatAssessmentRequest) SetCategory(value *ThreatCategory)() {
    m.category = value
}
// SetContentType sets the contentType property value. The content type of threat assessment. Possible values are: mail, url, file.
func (m *ThreatAssessmentRequest) SetContentType(value *ThreatAssessmentContentType)() {
    m.contentType = value
}
// SetCreatedBy sets the createdBy property value. The threat assessment request creator.
func (m *ThreatAssessmentRequest) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ThreatAssessmentRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetExpectedAssessment sets the expectedAssessment property value. The expected assessment from submitter. Possible values are: block, unblock.
func (m *ThreatAssessmentRequest) SetExpectedAssessment(value *ThreatExpectedAssessment)() {
    m.expectedAssessment = value
}
// SetRequestSource sets the requestSource property value. The source of the threat assessment request. Possible values are: administrator.
func (m *ThreatAssessmentRequest) SetRequestSource(value *ThreatAssessmentRequestSource)() {
    m.requestSource = value
}
// SetResults sets the results property value. A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
func (m *ThreatAssessmentRequest) SetResults(value []ThreatAssessmentResult)() {
    m.results = value
}
// SetStatus sets the status property value. The assessment process status. Possible values are: pending, completed.
func (m *ThreatAssessmentRequest) SetStatus(value *ThreatAssessmentStatus)() {
    m.status = value
}
