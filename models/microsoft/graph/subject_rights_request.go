package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new subjectRightsRequest and sets the default values.
func NewSubjectRightsRequest()(*SubjectRightsRequest) {
    m := &SubjectRightsRequest{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignedTo property value. Identity that the request is assigned to.
func (m *SubjectRightsRequest) GetAssignedTo()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// Gets the closedDateTime property value. The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closedDateTime
    }
}
// Gets the createdBy property value. Identity information for the entity that created the request.
func (m *SubjectRightsRequest) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the dataSubject property value. Information about the data subject.
func (m *SubjectRightsRequest) GetDataSubject()(*DataSubject) {
    if m == nil {
        return nil
    } else {
        return m.dataSubject
    }
}
// Gets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
func (m *SubjectRightsRequest) GetDataSubjectType()(*DataSubjectType) {
    if m == nil {
        return nil
    } else {
        return m.dataSubjectType
    }
}
// Gets the description property value. Description for the request.
func (m *SubjectRightsRequest) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of the request.
func (m *SubjectRightsRequest) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the history property value. Collection of history change events.
func (m *SubjectRightsRequest) GetHistory()([]SubjectRightsRequestHistory) {
    if m == nil {
        return nil
    } else {
        return m.history
    }
}
// Gets the insight property value. Insight about the request.
func (m *SubjectRightsRequest) GetInsight()(*SubjectRightsRequestDetail) {
    if m == nil {
        return nil
    } else {
        return m.insight
    }
}
// Gets the internalDueDateTime property value. The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.internalDueDateTime
    }
}
// Gets the lastModifiedBy property value. Identity information for the entity that last modified the request.
func (m *SubjectRightsRequest) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedDateTime property value. The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the notes property value. List of notes associcated with the request.
func (m *SubjectRightsRequest) GetNotes()([]AuthoredNote) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the regulations property value. List of regulations that this request will fulfill.
func (m *SubjectRightsRequest) GetRegulations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.regulations
    }
}
// Gets the stages property value. Information about the different stages for the request.
func (m *SubjectRightsRequest) GetStages()([]SubjectRightsRequestStageDetail) {
    if m == nil {
        return nil
    } else {
        return m.stages
    }
}
// Gets the status property value. The status of the request.. Possible values are: active, closed, unknownFutureValue.
func (m *SubjectRightsRequest) GetStatus()(*SubjectRightsRequestStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the team property value. Information about the Microsoft Teams team that was created for the request.
func (m *SubjectRightsRequest) GetTeam()(*Team) {
    if m == nil {
        return nil
    } else {
        return m.team
    }
}
// Gets the type_escaped property value. The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
func (m *SubjectRightsRequest) GetType_escaped()(*SubjectRightsRequestType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *SubjectRightsRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetAssignedTo(val.(*Identity))
        return nil
    }
    res["closedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetClosedDateTime(val)
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
    res["dataSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataSubject() })
        if err != nil {
            return err
        }
        m.SetDataSubject(val.(*DataSubject))
        return nil
    }
    res["dataSubjectType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSubjectType)
        if err != nil {
            return err
        }
        cast := val.(DataSubjectType)
        m.SetDataSubjectType(&cast)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["history"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubjectRightsRequestHistory() })
        if err != nil {
            return err
        }
        res := make([]SubjectRightsRequestHistory, len(val))
        for i, v := range val {
            res[i] = *(v.(*SubjectRightsRequestHistory))
        }
        m.SetHistory(res)
        return nil
    }
    res["insight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubjectRightsRequestDetail() })
        if err != nil {
            return err
        }
        m.SetInsight(val.(*SubjectRightsRequestDetail))
        return nil
    }
    res["internalDueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetInternalDueDateTime(val)
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetLastModifiedBy(val.(*IdentitySet))
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthoredNote() })
        if err != nil {
            return err
        }
        res := make([]AuthoredNote, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuthoredNote))
        }
        m.SetNotes(res)
        return nil
    }
    res["regulations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRegulations(res)
        return nil
    }
    res["stages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubjectRightsRequestStageDetail() })
        if err != nil {
            return err
        }
        res := make([]SubjectRightsRequestStageDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*SubjectRightsRequestStageDetail))
        }
        m.SetStages(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStatus)
        if err != nil {
            return err
        }
        cast := val.(SubjectRightsRequestStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["team"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeam() })
        if err != nil {
            return err
        }
        m.SetTeam(val.(*Team))
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestType)
        if err != nil {
            return err
        }
        cast := val.(SubjectRightsRequestType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *SubjectRightsRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the assignedTo property value. Identity that the request is assigned to.
// Parameters:
//  - value : Value to set for the assignedTo property.
func (m *SubjectRightsRequest) SetAssignedTo(value *Identity)() {
    m.assignedTo = value
}
// Sets the closedDateTime property value. The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the closedDateTime property.
func (m *SubjectRightsRequest) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closedDateTime = value
}
// Sets the createdBy property value. Identity information for the entity that created the request.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *SubjectRightsRequest) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *SubjectRightsRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the dataSubject property value. Information about the data subject.
// Parameters:
//  - value : Value to set for the dataSubject property.
func (m *SubjectRightsRequest) SetDataSubject(value *DataSubject)() {
    m.dataSubject = value
}
// Sets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
// Parameters:
//  - value : Value to set for the dataSubjectType property.
func (m *SubjectRightsRequest) SetDataSubjectType(value *DataSubjectType)() {
    m.dataSubjectType = value
}
// Sets the description property value. Description for the request.
// Parameters:
//  - value : Value to set for the description property.
func (m *SubjectRightsRequest) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The name of the request.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SubjectRightsRequest) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the history property value. Collection of history change events.
// Parameters:
//  - value : Value to set for the history property.
func (m *SubjectRightsRequest) SetHistory(value []SubjectRightsRequestHistory)() {
    m.history = value
}
// Sets the insight property value. Insight about the request.
// Parameters:
//  - value : Value to set for the insight property.
func (m *SubjectRightsRequest) SetInsight(value *SubjectRightsRequestDetail)() {
    m.insight = value
}
// Sets the internalDueDateTime property value. The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the internalDueDateTime property.
func (m *SubjectRightsRequest) SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.internalDueDateTime = value
}
// Sets the lastModifiedBy property value. Identity information for the entity that last modified the request.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *SubjectRightsRequest) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedDateTime property value. The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *SubjectRightsRequest) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the notes property value. List of notes associcated with the request.
// Parameters:
//  - value : Value to set for the notes property.
func (m *SubjectRightsRequest) SetNotes(value []AuthoredNote)() {
    m.notes = value
}
// Sets the regulations property value. List of regulations that this request will fulfill.
// Parameters:
//  - value : Value to set for the regulations property.
func (m *SubjectRightsRequest) SetRegulations(value []string)() {
    m.regulations = value
}
// Sets the stages property value. Information about the different stages for the request.
// Parameters:
//  - value : Value to set for the stages property.
func (m *SubjectRightsRequest) SetStages(value []SubjectRightsRequestStageDetail)() {
    m.stages = value
}
// Sets the status property value. The status of the request.. Possible values are: active, closed, unknownFutureValue.
// Parameters:
//  - value : Value to set for the status property.
func (m *SubjectRightsRequest) SetStatus(value *SubjectRightsRequestStatus)() {
    m.status = value
}
// Sets the team property value. Information about the Microsoft Teams team that was created for the request.
// Parameters:
//  - value : Value to set for the team property.
func (m *SubjectRightsRequest) SetTeam(value *Team)() {
    m.team = value
}
// Sets the type_escaped property value. The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *SubjectRightsRequest) SetType_escaped(value *SubjectRightsRequestType)() {
    m.type_escaped = value
}
