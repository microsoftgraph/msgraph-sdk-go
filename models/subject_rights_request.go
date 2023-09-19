package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequest 
type SubjectRightsRequest struct {
    Entity
}
// NewSubjectRightsRequest instantiates a new subjectRightsRequest and sets the default values.
func NewSubjectRightsRequest()(*SubjectRightsRequest) {
    m := &SubjectRightsRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSubjectRightsRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequest(), nil
}
// GetApprovers gets the approvers property value. The approvers property
func (m *SubjectRightsRequest) GetApprovers()([]Userable) {
    val, err := m.GetBackingStore().Get("approvers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Userable)
    }
    return nil
}
// GetAssignedTo gets the assignedTo property value. Identity that the request is assigned to.
func (m *SubjectRightsRequest) GetAssignedTo()(Identityable) {
    val, err := m.GetBackingStore().Get("assignedTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetClosedDateTime gets the closedDateTime property value. The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("closedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCollaborators gets the collaborators property value. The collaborators property
func (m *SubjectRightsRequest) GetCollaborators()([]Userable) {
    val, err := m.GetBackingStore().Get("collaborators")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Userable)
    }
    return nil
}
// GetContentQuery gets the contentQuery property value. The contentQuery property
func (m *SubjectRightsRequest) GetContentQuery()(*string) {
    val, err := m.GetBackingStore().Get("contentQuery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. Identity information for the entity that created the request.
func (m *SubjectRightsRequest) GetCreatedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDataSubject gets the dataSubject property value. Information about the data subject.
func (m *SubjectRightsRequest) GetDataSubject()(DataSubjectable) {
    val, err := m.GetBackingStore().Get("dataSubject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DataSubjectable)
    }
    return nil
}
// GetDataSubjectType gets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
func (m *SubjectRightsRequest) GetDataSubjectType()(*DataSubjectType) {
    val, err := m.GetBackingStore().Get("dataSubjectType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DataSubjectType)
    }
    return nil
}
// GetDescription gets the description property value. Description for the request.
func (m *SubjectRightsRequest) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the request.
func (m *SubjectRightsRequest) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalId gets the externalId property value. The externalId property
func (m *SubjectRightsRequest) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetApprovers(res)
        }
        return nil
    }
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val.(Identityable))
        }
        return nil
    }
    res["closedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedDateTime(val)
        }
        return nil
    }
    res["collaborators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetCollaborators(res)
        }
        return nil
    }
    res["contentQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["dataSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDataSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSubject(val.(DataSubjectable))
        }
        return nil
    }
    res["dataSubjectType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSubjectType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSubjectType(val.(*DataSubjectType))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubjectRightsRequestHistoryable)
                }
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["includeAllVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllVersions(val)
        }
        return nil
    }
    res["includeAuthoredContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAuthoredContent(val)
        }
        return nil
    }
    res["insight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubjectRightsRequestDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsight(val.(SubjectRightsRequestDetailable))
        }
        return nil
    }
    res["internalDueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalDueDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["mailboxlocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubjectRightsRequestMailboxLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxlocations(val.(SubjectRightsRequestMailboxLocationable))
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthoredNoteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthoredNoteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthoredNoteable)
                }
            }
            m.SetNotes(res)
        }
        return nil
    }
    res["pauseAfterEstimate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPauseAfterEstimate(val)
        }
        return nil
    }
    res["regulations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRegulations(res)
        }
        return nil
    }
    res["sitelocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubjectRightsRequestSiteLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSitelocations(val.(SubjectRightsRequestSiteLocationable))
        }
        return nil
    }
    res["stages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestStageDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestStageDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubjectRightsRequestStageDetailable)
                }
            }
            m.SetStages(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SubjectRightsRequestStatus))
        }
        return nil
    }
    res["team"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val.(Teamable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*SubjectRightsRequestType))
        }
        return nil
    }
    return res
}
// GetHistory gets the history property value. Collection of history change events.
func (m *SubjectRightsRequest) GetHistory()([]SubjectRightsRequestHistoryable) {
    val, err := m.GetBackingStore().Get("history")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SubjectRightsRequestHistoryable)
    }
    return nil
}
// GetIncludeAllVersions gets the includeAllVersions property value. The includeAllVersions property
func (m *SubjectRightsRequest) GetIncludeAllVersions()(*bool) {
    val, err := m.GetBackingStore().Get("includeAllVersions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIncludeAuthoredContent gets the includeAuthoredContent property value. The includeAuthoredContent property
func (m *SubjectRightsRequest) GetIncludeAuthoredContent()(*bool) {
    val, err := m.GetBackingStore().Get("includeAuthoredContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetInsight gets the insight property value. Insight about the request.
func (m *SubjectRightsRequest) GetInsight()(SubjectRightsRequestDetailable) {
    val, err := m.GetBackingStore().Get("insight")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SubjectRightsRequestDetailable)
    }
    return nil
}
// GetInternalDueDateTime gets the internalDueDateTime property value. The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("internalDueDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity information for the entity that last modified the request.
func (m *SubjectRightsRequest) GetLastModifiedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMailboxlocations gets the mailboxlocations property value. The mailboxlocations property
func (m *SubjectRightsRequest) GetMailboxlocations()(SubjectRightsRequestMailboxLocationable) {
    val, err := m.GetBackingStore().Get("mailboxlocations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SubjectRightsRequestMailboxLocationable)
    }
    return nil
}
// GetNotes gets the notes property value. List of notes associated with the request.
func (m *SubjectRightsRequest) GetNotes()([]AuthoredNoteable) {
    val, err := m.GetBackingStore().Get("notes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthoredNoteable)
    }
    return nil
}
// GetPauseAfterEstimate gets the pauseAfterEstimate property value. The pauseAfterEstimate property
func (m *SubjectRightsRequest) GetPauseAfterEstimate()(*bool) {
    val, err := m.GetBackingStore().Get("pauseAfterEstimate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRegulations gets the regulations property value. List of regulations that this request fulfills.
func (m *SubjectRightsRequest) GetRegulations()([]string) {
    val, err := m.GetBackingStore().Get("regulations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSitelocations gets the sitelocations property value. The sitelocations property
func (m *SubjectRightsRequest) GetSitelocations()(SubjectRightsRequestSiteLocationable) {
    val, err := m.GetBackingStore().Get("sitelocations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SubjectRightsRequestSiteLocationable)
    }
    return nil
}
// GetStages gets the stages property value. Information about the different stages for the request.
func (m *SubjectRightsRequest) GetStages()([]SubjectRightsRequestStageDetailable) {
    val, err := m.GetBackingStore().Get("stages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SubjectRightsRequestStageDetailable)
    }
    return nil
}
// GetStatus gets the status property value. The status of the request. Possible values are: active, closed, unknownFutureValue.
func (m *SubjectRightsRequest) GetStatus()(*SubjectRightsRequestStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubjectRightsRequestStatus)
    }
    return nil
}
// GetTeam gets the team property value. Information about the Microsoft Teams team that was created for the request.
func (m *SubjectRightsRequest) GetTeam()(Teamable) {
    val, err := m.GetBackingStore().Get("team")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Teamable)
    }
    return nil
}
// GetTypeEscaped gets the type property value. The type of the request. Possible values are: export, delete,  access, tagForAction, unknownFutureValue.
func (m *SubjectRightsRequest) GetTypeEscaped()(*SubjectRightsRequestType) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubjectRightsRequestType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SubjectRightsRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApprovers()))
        for i, v := range m.GetApprovers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("approvers", cast)
        if err != nil {
            return err
        }
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
    if m.GetCollaborators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCollaborators()))
        for i, v := range m.GetCollaborators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("collaborators", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
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
        cast := (*m.GetDataSubjectType()).String()
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeAllVersions", m.GetIncludeAllVersions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeAuthoredContent", m.GetIncludeAuthoredContent())
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
        err = writer.WriteObjectValue("mailboxlocations", m.GetMailboxlocations())
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotes()))
        for i, v := range m.GetNotes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("notes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pauseAfterEstimate", m.GetPauseAfterEstimate())
        if err != nil {
            return err
        }
    }
    if m.GetRegulations() != nil {
        err = writer.WriteCollectionOfStringValues("regulations", m.GetRegulations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sitelocations", m.GetSitelocations())
        if err != nil {
            return err
        }
    }
    if m.GetStages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStages()))
        for i, v := range m.GetStages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("stages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
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
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApprovers sets the approvers property value. The approvers property
func (m *SubjectRightsRequest) SetApprovers(value []Userable)() {
    err := m.GetBackingStore().Set("approvers", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignedTo sets the assignedTo property value. Identity that the request is assigned to.
func (m *SubjectRightsRequest) SetAssignedTo(value Identityable)() {
    err := m.GetBackingStore().Set("assignedTo", value)
    if err != nil {
        panic(err)
    }
}
// SetClosedDateTime sets the closedDateTime property value. The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("closedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCollaborators sets the collaborators property value. The collaborators property
func (m *SubjectRightsRequest) SetCollaborators(value []Userable)() {
    err := m.GetBackingStore().Set("collaborators", value)
    if err != nil {
        panic(err)
    }
}
// SetContentQuery sets the contentQuery property value. The contentQuery property
func (m *SubjectRightsRequest) SetContentQuery(value *string)() {
    err := m.GetBackingStore().Set("contentQuery", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. Identity information for the entity that created the request.
func (m *SubjectRightsRequest) SetCreatedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDataSubject sets the dataSubject property value. Information about the data subject.
func (m *SubjectRightsRequest) SetDataSubject(value DataSubjectable)() {
    err := m.GetBackingStore().Set("dataSubject", value)
    if err != nil {
        panic(err)
    }
}
// SetDataSubjectType sets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
func (m *SubjectRightsRequest) SetDataSubjectType(value *DataSubjectType)() {
    err := m.GetBackingStore().Set("dataSubjectType", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description for the request.
func (m *SubjectRightsRequest) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the request.
func (m *SubjectRightsRequest) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The externalId property
func (m *SubjectRightsRequest) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetHistory sets the history property value. Collection of history change events.
func (m *SubjectRightsRequest) SetHistory(value []SubjectRightsRequestHistoryable)() {
    err := m.GetBackingStore().Set("history", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeAllVersions sets the includeAllVersions property value. The includeAllVersions property
func (m *SubjectRightsRequest) SetIncludeAllVersions(value *bool)() {
    err := m.GetBackingStore().Set("includeAllVersions", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeAuthoredContent sets the includeAuthoredContent property value. The includeAuthoredContent property
func (m *SubjectRightsRequest) SetIncludeAuthoredContent(value *bool)() {
    err := m.GetBackingStore().Set("includeAuthoredContent", value)
    if err != nil {
        panic(err)
    }
}
// SetInsight sets the insight property value. Insight about the request.
func (m *SubjectRightsRequest) SetInsight(value SubjectRightsRequestDetailable)() {
    err := m.GetBackingStore().Set("insight", value)
    if err != nil {
        panic(err)
    }
}
// SetInternalDueDateTime sets the internalDueDateTime property value. The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("internalDueDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity information for the entity that last modified the request.
func (m *SubjectRightsRequest) SetLastModifiedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMailboxlocations sets the mailboxlocations property value. The mailboxlocations property
func (m *SubjectRightsRequest) SetMailboxlocations(value SubjectRightsRequestMailboxLocationable)() {
    err := m.GetBackingStore().Set("mailboxlocations", value)
    if err != nil {
        panic(err)
    }
}
// SetNotes sets the notes property value. List of notes associated with the request.
func (m *SubjectRightsRequest) SetNotes(value []AuthoredNoteable)() {
    err := m.GetBackingStore().Set("notes", value)
    if err != nil {
        panic(err)
    }
}
// SetPauseAfterEstimate sets the pauseAfterEstimate property value. The pauseAfterEstimate property
func (m *SubjectRightsRequest) SetPauseAfterEstimate(value *bool)() {
    err := m.GetBackingStore().Set("pauseAfterEstimate", value)
    if err != nil {
        panic(err)
    }
}
// SetRegulations sets the regulations property value. List of regulations that this request fulfills.
func (m *SubjectRightsRequest) SetRegulations(value []string)() {
    err := m.GetBackingStore().Set("regulations", value)
    if err != nil {
        panic(err)
    }
}
// SetSitelocations sets the sitelocations property value. The sitelocations property
func (m *SubjectRightsRequest) SetSitelocations(value SubjectRightsRequestSiteLocationable)() {
    err := m.GetBackingStore().Set("sitelocations", value)
    if err != nil {
        panic(err)
    }
}
// SetStages sets the stages property value. Information about the different stages for the request.
func (m *SubjectRightsRequest) SetStages(value []SubjectRightsRequestStageDetailable)() {
    err := m.GetBackingStore().Set("stages", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status of the request. Possible values are: active, closed, unknownFutureValue.
func (m *SubjectRightsRequest) SetStatus(value *SubjectRightsRequestStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTeam sets the team property value. Information about the Microsoft Teams team that was created for the request.
func (m *SubjectRightsRequest) SetTeam(value Teamable)() {
    err := m.GetBackingStore().Set("team", value)
    if err != nil {
        panic(err)
    }
}
// SetTypeEscaped sets the type property value. The type of the request. Possible values are: export, delete,  access, tagForAction, unknownFutureValue.
func (m *SubjectRightsRequest) SetTypeEscaped(value *SubjectRightsRequestType)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SubjectRightsRequestable 
type SubjectRightsRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovers()([]Userable)
    GetAssignedTo()(Identityable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCollaborators()([]Userable)
    GetContentQuery()(*string)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDataSubject()(DataSubjectable)
    GetDataSubjectType()(*DataSubjectType)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetHistory()([]SubjectRightsRequestHistoryable)
    GetIncludeAllVersions()(*bool)
    GetIncludeAuthoredContent()(*bool)
    GetInsight()(SubjectRightsRequestDetailable)
    GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMailboxlocations()(SubjectRightsRequestMailboxLocationable)
    GetNotes()([]AuthoredNoteable)
    GetPauseAfterEstimate()(*bool)
    GetRegulations()([]string)
    GetSitelocations()(SubjectRightsRequestSiteLocationable)
    GetStages()([]SubjectRightsRequestStageDetailable)
    GetStatus()(*SubjectRightsRequestStatus)
    GetTeam()(Teamable)
    GetTypeEscaped()(*SubjectRightsRequestType)
    SetApprovers(value []Userable)()
    SetAssignedTo(value Identityable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCollaborators(value []Userable)()
    SetContentQuery(value *string)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDataSubject(value DataSubjectable)()
    SetDataSubjectType(value *DataSubjectType)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetHistory(value []SubjectRightsRequestHistoryable)()
    SetIncludeAllVersions(value *bool)()
    SetIncludeAuthoredContent(value *bool)()
    SetInsight(value SubjectRightsRequestDetailable)()
    SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMailboxlocations(value SubjectRightsRequestMailboxLocationable)()
    SetNotes(value []AuthoredNoteable)()
    SetPauseAfterEstimate(value *bool)()
    SetRegulations(value []string)()
    SetSitelocations(value SubjectRightsRequestSiteLocationable)()
    SetStages(value []SubjectRightsRequestStageDetailable)()
    SetStatus(value *SubjectRightsRequestStatus)()
    SetTeam(value Teamable)()
    SetTypeEscaped(value *SubjectRightsRequestType)()
}
