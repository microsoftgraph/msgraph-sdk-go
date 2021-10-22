package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationSubmission struct {
    Entity
    outcomes []EducationOutcome;
    recipient *EducationSubmissionRecipient;
    resources []EducationSubmissionResource;
    resourcesFolderUrl *string;
    returnedBy *IdentitySet;
    returnedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    status *EducationSubmissionStatus;
    submittedBy *IdentitySet;
    submittedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    submittedResources []EducationSubmissionResource;
    unsubmittedBy *IdentitySet;
    unsubmittedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewEducationSubmission()(*EducationSubmission) {
    m := &EducationSubmission{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EducationSubmission) GetOutcomes()([]EducationOutcome) {
    if m == nil {
        return nil
    } else {
        return m.outcomes
    }
}
func (m *EducationSubmission) GetRecipient()(*EducationSubmissionRecipient) {
    if m == nil {
        return nil
    } else {
        return m.recipient
    }
}
func (m *EducationSubmission) GetResources()([]EducationSubmissionResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
func (m *EducationSubmission) GetResourcesFolderUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourcesFolderUrl
    }
}
func (m *EducationSubmission) GetReturnedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.returnedBy
    }
}
func (m *EducationSubmission) GetReturnedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.returnedDateTime
    }
}
func (m *EducationSubmission) GetStatus()(*EducationSubmissionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *EducationSubmission) GetSubmittedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.submittedBy
    }
}
func (m *EducationSubmission) GetSubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.submittedDateTime
    }
}
func (m *EducationSubmission) GetSubmittedResources()([]EducationSubmissionResource) {
    if m == nil {
        return nil
    } else {
        return m.submittedResources
    }
}
func (m *EducationSubmission) GetUnsubmittedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.unsubmittedBy
    }
}
func (m *EducationSubmission) GetUnsubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.unsubmittedDateTime
    }
}
func (m *EducationSubmission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["outcomes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationOutcome() })
        if err != nil {
            return err
        }
        res := make([]EducationOutcome, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationOutcome))
        }
        m.SetOutcomes(res)
        return nil
    }
    res["recipient"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSubmissionRecipient() })
        if err != nil {
            return err
        }
        m.SetRecipient(val.(*EducationSubmissionRecipient))
        return nil
    }
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSubmissionResource() })
        if err != nil {
            return err
        }
        res := make([]EducationSubmissionResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationSubmissionResource))
        }
        m.SetResources(res)
        return nil
    }
    res["resourcesFolderUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourcesFolderUrl(val)
        return nil
    }
    res["returnedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetReturnedBy(val.(*IdentitySet))
        return nil
    }
    res["returnedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReturnedDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationSubmissionStatus)
        if err != nil {
            return err
        }
        cast := val.(EducationSubmissionStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["submittedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetSubmittedBy(val.(*IdentitySet))
        return nil
    }
    res["submittedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSubmittedDateTime(val)
        return nil
    }
    res["submittedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSubmissionResource() })
        if err != nil {
            return err
        }
        res := make([]EducationSubmissionResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationSubmissionResource))
        }
        m.SetSubmittedResources(res)
        return nil
    }
    res["unsubmittedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetUnsubmittedBy(val.(*IdentitySet))
        return nil
    }
    res["unsubmittedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetUnsubmittedDateTime(val)
        return nil
    }
    return res
}
func (m *EducationSubmission) IsNil()(bool) {
    return m == nil
}
func (m *EducationSubmission) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOutcomes()))
        for i, v := range m.GetOutcomes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("outcomes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recipient", m.GetRecipient())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourcesFolderUrl", m.GetResourcesFolderUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("returnedBy", m.GetReturnedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("returnedDateTime", m.GetReturnedDateTime())
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
    {
        err = writer.WriteObjectValue("submittedBy", m.GetSubmittedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("submittedDateTime", m.GetSubmittedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSubmittedResources()))
        for i, v := range m.GetSubmittedResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("submittedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("unsubmittedBy", m.GetUnsubmittedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("unsubmittedDateTime", m.GetUnsubmittedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *EducationSubmission) SetOutcomes(value []EducationOutcome)() {
    m.outcomes = value
}
func (m *EducationSubmission) SetRecipient(value *EducationSubmissionRecipient)() {
    m.recipient = value
}
func (m *EducationSubmission) SetResources(value []EducationSubmissionResource)() {
    m.resources = value
}
func (m *EducationSubmission) SetResourcesFolderUrl(value *string)() {
    m.resourcesFolderUrl = value
}
func (m *EducationSubmission) SetReturnedBy(value *IdentitySet)() {
    m.returnedBy = value
}
func (m *EducationSubmission) SetReturnedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.returnedDateTime = value
}
func (m *EducationSubmission) SetStatus(value *EducationSubmissionStatus)() {
    m.status = value
}
func (m *EducationSubmission) SetSubmittedBy(value *IdentitySet)() {
    m.submittedBy = value
}
func (m *EducationSubmission) SetSubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.submittedDateTime = value
}
func (m *EducationSubmission) SetSubmittedResources(value []EducationSubmissionResource)() {
    m.submittedResources = value
}
func (m *EducationSubmission) SetUnsubmittedBy(value *IdentitySet)() {
    m.unsubmittedBy = value
}
func (m *EducationSubmission) SetUnsubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.unsubmittedDateTime = value
}
