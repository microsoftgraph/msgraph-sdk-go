package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSubmission 
type EducationSubmission struct {
    Entity
    // Read-Write. Nullable.
    outcomes []EducationOutcome;
    // Who this submission is assigned to.
    recipient *EducationSubmissionRecipient;
    // Nullable.
    resources []EducationSubmissionResource;
    // Folder where all file resources for this submission need to be stored.
    resourcesFolderUrl *string;
    // User who moved the status of this submission to returned.
    returnedBy *IdentitySet;
    // Moment in time when the submission was returned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    returnedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-Only. Possible values are: working, submitted, released, returned.
    status *EducationSubmissionStatus;
    // User who moved the resource into the submitted state.
    submittedBy *IdentitySet;
    // Moment in time when the submission was moved into the submitted state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    submittedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. Nullable.
    submittedResources []EducationSubmissionResource;
    // User who moved the resource from submitted into the working state.
    unsubmittedBy *IdentitySet;
    // Moment in time when the submission was moved from submitted into the working state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    unsubmittedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewEducationSubmission instantiates a new educationSubmission and sets the default values.
func NewEducationSubmission()(*EducationSubmission) {
    m := &EducationSubmission{
        Entity: *NewEntity(),
    }
    return m
}
// GetOutcomes gets the outcomes property value. Read-Write. Nullable.
func (m *EducationSubmission) GetOutcomes()([]EducationOutcome) {
    if m == nil {
        return nil
    } else {
        return m.outcomes
    }
}
// GetRecipient gets the recipient property value. Who this submission is assigned to.
func (m *EducationSubmission) GetRecipient()(*EducationSubmissionRecipient) {
    if m == nil {
        return nil
    } else {
        return m.recipient
    }
}
// GetResources gets the resources property value. Nullable.
func (m *EducationSubmission) GetResources()([]EducationSubmissionResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// GetResourcesFolderUrl gets the resourcesFolderUrl property value. Folder where all file resources for this submission need to be stored.
func (m *EducationSubmission) GetResourcesFolderUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourcesFolderUrl
    }
}
// GetReturnedBy gets the returnedBy property value. User who moved the status of this submission to returned.
func (m *EducationSubmission) GetReturnedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.returnedBy
    }
}
// GetReturnedDateTime gets the returnedDateTime property value. Moment in time when the submission was returned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationSubmission) GetReturnedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.returnedDateTime
    }
}
// GetStatus gets the status property value. Read-Only. Possible values are: working, submitted, released, returned.
func (m *EducationSubmission) GetStatus()(*EducationSubmissionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSubmittedBy gets the submittedBy property value. User who moved the resource into the submitted state.
func (m *EducationSubmission) GetSubmittedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.submittedBy
    }
}
// GetSubmittedDateTime gets the submittedDateTime property value. Moment in time when the submission was moved into the submitted state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationSubmission) GetSubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.submittedDateTime
    }
}
// GetSubmittedResources gets the submittedResources property value. Read-only. Nullable.
func (m *EducationSubmission) GetSubmittedResources()([]EducationSubmissionResource) {
    if m == nil {
        return nil
    } else {
        return m.submittedResources
    }
}
// GetUnsubmittedBy gets the unsubmittedBy property value. User who moved the resource from submitted into the working state.
func (m *EducationSubmission) GetUnsubmittedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.unsubmittedBy
    }
}
// GetUnsubmittedDateTime gets the unsubmittedDateTime property value. Moment in time when the submission was moved from submitted into the working state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationSubmission) GetUnsubmittedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.unsubmittedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSubmission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["outcomes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationOutcome() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationOutcome, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationOutcome))
            }
            m.SetOutcomes(res)
        }
        return nil
    }
    res["recipient"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSubmissionRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipient(val.(*EducationSubmissionRecipient))
        }
        return nil
    }
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSubmissionResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSubmissionResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationSubmissionResource))
            }
            m.SetResources(res)
        }
        return nil
    }
    res["resourcesFolderUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourcesFolderUrl(val)
        }
        return nil
    }
    res["returnedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["returnedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnedDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationSubmissionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationSubmissionStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["submittedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmittedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["submittedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmittedDateTime(val)
        }
        return nil
    }
    res["submittedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSubmissionResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSubmissionResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationSubmissionResource))
            }
            m.SetSubmittedResources(res)
        }
        return nil
    }
    res["unsubmittedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnsubmittedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["unsubmittedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnsubmittedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *EducationSubmission) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetOutcomes sets the outcomes property value. Read-Write. Nullable.
func (m *EducationSubmission) SetOutcomes(value []EducationOutcome)() {
    m.outcomes = value
}
// SetRecipient sets the recipient property value. Who this submission is assigned to.
func (m *EducationSubmission) SetRecipient(value *EducationSubmissionRecipient)() {
    m.recipient = value
}
// SetResources sets the resources property value. Nullable.
func (m *EducationSubmission) SetResources(value []EducationSubmissionResource)() {
    m.resources = value
}
// SetResourcesFolderUrl sets the resourcesFolderUrl property value. Folder where all file resources for this submission need to be stored.
func (m *EducationSubmission) SetResourcesFolderUrl(value *string)() {
    m.resourcesFolderUrl = value
}
// SetReturnedBy sets the returnedBy property value. User who moved the status of this submission to returned.
func (m *EducationSubmission) SetReturnedBy(value *IdentitySet)() {
    m.returnedBy = value
}
// SetReturnedDateTime sets the returnedDateTime property value. Moment in time when the submission was returned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationSubmission) SetReturnedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.returnedDateTime = value
}
// SetStatus sets the status property value. Read-Only. Possible values are: working, submitted, released, returned.
func (m *EducationSubmission) SetStatus(value *EducationSubmissionStatus)() {
    m.status = value
}
// SetSubmittedBy sets the submittedBy property value. User who moved the resource into the submitted state.
func (m *EducationSubmission) SetSubmittedBy(value *IdentitySet)() {
    m.submittedBy = value
}
// SetSubmittedDateTime sets the submittedDateTime property value. Moment in time when the submission was moved into the submitted state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationSubmission) SetSubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.submittedDateTime = value
}
// SetSubmittedResources sets the submittedResources property value. Read-only. Nullable.
func (m *EducationSubmission) SetSubmittedResources(value []EducationSubmissionResource)() {
    m.submittedResources = value
}
// SetUnsubmittedBy sets the unsubmittedBy property value. User who moved the resource from submitted into the working state.
func (m *EducationSubmission) SetUnsubmittedBy(value *IdentitySet)() {
    m.unsubmittedBy = value
}
// SetUnsubmittedDateTime sets the unsubmittedDateTime property value. Moment in time when the submission was moved from submitted into the working state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationSubmission) SetUnsubmittedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.unsubmittedDateTime = value
}
