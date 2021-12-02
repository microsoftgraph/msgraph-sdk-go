package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubjectRightsRequest 
type SubjectRightsRequest struct {
    Entity
    // Identity that the request is assigned to.
    assignedTo *Identity;
    // The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identity information for the entity that created the request.
    createdBy *IdentitySet;
    // The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Information about the data subject.
    dataSubject *DataSubject;
    // The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
    dataSubjectType *DataSubjectType;
    // Description for the request.
    description *string;
    // The name of the request.
    displayName *string;
    // Collection of history change events.
    history []SubjectRightsRequestHistory;
    // Insight about the request.
    insight *SubjectRightsRequestDetail;
    // The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    internalDueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identity information for the entity that last modified the request.
    lastModifiedBy *IdentitySet;
    // The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of notes associcated with the request.
    notes []AuthoredNote;
    // List of regulations that this request will fulfill.
    regulations []string;
    // Information about the different stages for the request.
    stages []SubjectRightsRequestStageDetail;
    // The status of the request.. Possible values are: active, closed, unknownFutureValue.
    status *SubjectRightsRequestStatus;
    // Information about the Microsoft Teams team that was created for the request.
    team *Team;
    // The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
    type_escaped *SubjectRightsRequestType;
}
// NewSubjectRightsRequest instantiates a new subjectRightsRequest and sets the default values.
func NewSubjectRightsRequest()(*SubjectRightsRequest) {
    m := &SubjectRightsRequest{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignedTo gets the assignedTo property value. Identity that the request is assigned to.
func (m *SubjectRightsRequest) GetAssignedTo()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// GetClosedDateTime gets the closedDateTime property value. The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closedDateTime
    }
}
// GetCreatedBy gets the createdBy property value. Identity information for the entity that created the request.
func (m *SubjectRightsRequest) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDataSubject gets the dataSubject property value. Information about the data subject.
func (m *SubjectRightsRequest) GetDataSubject()(*DataSubject) {
    if m == nil {
        return nil
    } else {
        return m.dataSubject
    }
}
// GetDataSubjectType gets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
func (m *SubjectRightsRequest) GetDataSubjectType()(*DataSubjectType) {
    if m == nil {
        return nil
    } else {
        return m.dataSubjectType
    }
}
// GetDescription gets the description property value. Description for the request.
func (m *SubjectRightsRequest) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name of the request.
func (m *SubjectRightsRequest) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetHistory gets the history property value. Collection of history change events.
func (m *SubjectRightsRequest) GetHistory()([]SubjectRightsRequestHistory) {
    if m == nil {
        return nil
    } else {
        return m.history
    }
}
// GetInsight gets the insight property value. Insight about the request.
func (m *SubjectRightsRequest) GetInsight()(*SubjectRightsRequestDetail) {
    if m == nil {
        return nil
    } else {
        return m.insight
    }
}
// GetInternalDueDateTime gets the internalDueDateTime property value. The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.internalDueDateTime
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity information for the entity that last modified the request.
func (m *SubjectRightsRequest) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNotes gets the notes property value. List of notes associcated with the request.
func (m *SubjectRightsRequest) GetNotes()([]AuthoredNote) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetRegulations gets the regulations property value. List of regulations that this request will fulfill.
func (m *SubjectRightsRequest) GetRegulations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.regulations
    }
}
// GetStages gets the stages property value. Information about the different stages for the request.
func (m *SubjectRightsRequest) GetStages()([]SubjectRightsRequestStageDetail) {
    if m == nil {
        return nil
    } else {
        return m.stages
    }
}
// GetStatus gets the status property value. The status of the request.. Possible values are: active, closed, unknownFutureValue.
func (m *SubjectRightsRequest) GetStatus()(*SubjectRightsRequestStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTeam gets the team property value. Information about the Microsoft Teams team that was created for the request.
func (m *SubjectRightsRequest) GetTeam()(*Team) {
    if m == nil {
        return nil
    } else {
        return m.team
    }
}
// GetType_escaped gets the type_escaped property value. The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
func (m *SubjectRightsRequest) GetType_escaped()(*SubjectRightsRequestType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val.(*Identity))
        }
        return nil
    }
    res["closedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedDateTime(val)
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
    res["dataSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataSubject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSubject(val.(*DataSubject))
        }
        return nil
    }
    res["dataSubjectType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSubjectType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DataSubjectType)
            m.SetDataSubjectType(&cast)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["history"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubjectRightsRequestHistory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestHistory, len(val))
            for i, v := range val {
                res[i] = *(v.(*SubjectRightsRequestHistory))
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["insight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubjectRightsRequestDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsight(val.(*SubjectRightsRequestDetail))
        }
        return nil
    }
    res["internalDueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalDueDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthoredNote() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthoredNote, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthoredNote))
            }
            m.SetNotes(res)
        }
        return nil
    }
    res["regulations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRegulations(res)
        }
        return nil
    }
    res["stages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubjectRightsRequestStageDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestStageDetail, len(val))
            for i, v := range val {
                res[i] = *(v.(*SubjectRightsRequestStageDetail))
            }
            m.SetStages(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SubjectRightsRequestStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["team"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeam() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val.(*Team))
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SubjectRightsRequestType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    return res
}
func (m *SubjectRightsRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SubjectRightsRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closedDateTime", m.GetClosedDateTime())
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
    {
        err = writer.WriteObjectValue("dataSubject", m.GetDataSubject())
        if err != nil {
            return err
        }
    }
    if m.GetDataSubjectType() != nil {
        cast := m.GetDataSubjectType().String()
        err = writer.WriteStringValue("dataSubjectType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("insight", m.GetInsight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("internalDueDateTime", m.GetInternalDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNotes()))
        for i, v := range m.GetNotes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("notes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("regulations", m.GetRegulations())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStages()))
        for i, v := range m.GetStages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("stages", cast)
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
        err = writer.WriteObjectValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. Identity that the request is assigned to.
func (m *SubjectRightsRequest) SetAssignedTo(value *Identity)() {
    if m != nil {
        m.assignedTo = value
    }
}
// SetClosedDateTime sets the closedDateTime property value. The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.closedDateTime = value
    }
}
// SetCreatedBy sets the createdBy property value. Identity information for the entity that created the request.
func (m *SubjectRightsRequest) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDataSubject sets the dataSubject property value. Information about the data subject.
func (m *SubjectRightsRequest) SetDataSubject(value *DataSubject)() {
    if m != nil {
        m.dataSubject = value
    }
}
// SetDataSubjectType sets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
func (m *SubjectRightsRequest) SetDataSubjectType(value *DataSubjectType)() {
    if m != nil {
        m.dataSubjectType = value
    }
}
// SetDescription sets the description property value. Description for the request.
func (m *SubjectRightsRequest) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The name of the request.
func (m *SubjectRightsRequest) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHistory sets the history property value. Collection of history change events.
func (m *SubjectRightsRequest) SetHistory(value []SubjectRightsRequestHistory)() {
    if m != nil {
        m.history = value
    }
}
// SetInsight sets the insight property value. Insight about the request.
func (m *SubjectRightsRequest) SetInsight(value *SubjectRightsRequestDetail)() {
    if m != nil {
        m.insight = value
    }
}
// SetInternalDueDateTime sets the internalDueDateTime property value. The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.internalDueDateTime = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity information for the entity that last modified the request.
func (m *SubjectRightsRequest) SetLastModifiedBy(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNotes sets the notes property value. List of notes associcated with the request.
func (m *SubjectRightsRequest) SetNotes(value []AuthoredNote)() {
    if m != nil {
        m.notes = value
    }
}
// SetRegulations sets the regulations property value. List of regulations that this request will fulfill.
func (m *SubjectRightsRequest) SetRegulations(value []string)() {
    if m != nil {
        m.regulations = value
    }
}
// SetStages sets the stages property value. Information about the different stages for the request.
func (m *SubjectRightsRequest) SetStages(value []SubjectRightsRequestStageDetail)() {
    if m != nil {
        m.stages = value
    }
}
// SetStatus sets the status property value. The status of the request.. Possible values are: active, closed, unknownFutureValue.
func (m *SubjectRightsRequest) SetStatus(value *SubjectRightsRequestStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTeam sets the team property value. Information about the Microsoft Teams team that was created for the request.
func (m *SubjectRightsRequest) SetTeam(value *Team)() {
    if m != nil {
        m.team = value
    }
}
// SetType_escaped sets the type_escaped property value. The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
func (m *SubjectRightsRequest) SetType_escaped(value *SubjectRightsRequestType)() {
    if m != nil {
        m.type_escaped = value
    }
}
