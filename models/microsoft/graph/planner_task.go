package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerTask struct {
    Entity
    activeChecklistItemCount *int32;
    appliedCategories *PlannerAppliedCategories;
    assignedToTaskBoardFormat *PlannerAssignedToTaskBoardTaskFormat;
    assigneePriority *string;
    assignments *PlannerAssignments;
    bucketId *string;
    bucketTaskBoardFormat *PlannerBucketTaskBoardTaskFormat;
    checklistItemCount *int32;
    completedBy *IdentitySet;
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    conversationThreadId *string;
    createdBy *IdentitySet;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    details *PlannerTaskDetails;
    dueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    hasDescription *bool;
    orderHint *string;
    percentComplete *int32;
    planId *string;
    previewType *PlannerPreviewType;
    progressTaskBoardFormat *PlannerProgressTaskBoardTaskFormat;
    referenceCount *int32;
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    title *string;
}
func NewPlannerTask()(*PlannerTask) {
    m := &PlannerTask{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PlannerTask) GetActiveChecklistItemCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeChecklistItemCount
    }
}
func (m *PlannerTask) GetAppliedCategories()(*PlannerAppliedCategories) {
    if m == nil {
        return nil
    } else {
        return m.appliedCategories
    }
}
func (m *PlannerTask) GetAssignedToTaskBoardFormat()(*PlannerAssignedToTaskBoardTaskFormat) {
    if m == nil {
        return nil
    } else {
        return m.assignedToTaskBoardFormat
    }
}
func (m *PlannerTask) GetAssigneePriority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assigneePriority
    }
}
func (m *PlannerTask) GetAssignments()(*PlannerAssignments) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *PlannerTask) GetBucketId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bucketId
    }
}
func (m *PlannerTask) GetBucketTaskBoardFormat()(*PlannerBucketTaskBoardTaskFormat) {
    if m == nil {
        return nil
    } else {
        return m.bucketTaskBoardFormat
    }
}
func (m *PlannerTask) GetChecklistItemCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.checklistItemCount
    }
}
func (m *PlannerTask) GetCompletedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.completedBy
    }
}
func (m *PlannerTask) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
func (m *PlannerTask) GetConversationThreadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conversationThreadId
    }
}
func (m *PlannerTask) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
func (m *PlannerTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *PlannerTask) GetDetails()(*PlannerTaskDetails) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
func (m *PlannerTask) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
func (m *PlannerTask) GetHasDescription()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasDescription
    }
}
func (m *PlannerTask) GetOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderHint
    }
}
func (m *PlannerTask) GetPercentComplete()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.percentComplete
    }
}
func (m *PlannerTask) GetPlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.planId
    }
}
func (m *PlannerTask) GetPreviewType()(*PlannerPreviewType) {
    if m == nil {
        return nil
    } else {
        return m.previewType
    }
}
func (m *PlannerTask) GetProgressTaskBoardFormat()(*PlannerProgressTaskBoardTaskFormat) {
    if m == nil {
        return nil
    } else {
        return m.progressTaskBoardFormat
    }
}
func (m *PlannerTask) GetReferenceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.referenceCount
    }
}
func (m *PlannerTask) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
func (m *PlannerTask) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
func (m *PlannerTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeChecklistItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActiveChecklistItemCount(val)
        return nil
    }
    res["appliedCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerAppliedCategories() })
        if err != nil {
            return err
        }
        m.SetAppliedCategories(val.(*PlannerAppliedCategories))
        return nil
    }
    res["assignedToTaskBoardFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerAssignedToTaskBoardTaskFormat() })
        if err != nil {
            return err
        }
        m.SetAssignedToTaskBoardFormat(val.(*PlannerAssignedToTaskBoardTaskFormat))
        return nil
    }
    res["assigneePriority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssigneePriority(val)
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerAssignments() })
        if err != nil {
            return err
        }
        m.SetAssignments(val.(*PlannerAssignments))
        return nil
    }
    res["bucketId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBucketId(val)
        return nil
    }
    res["bucketTaskBoardFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerBucketTaskBoardTaskFormat() })
        if err != nil {
            return err
        }
        m.SetBucketTaskBoardFormat(val.(*PlannerBucketTaskBoardTaskFormat))
        return nil
    }
    res["checklistItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetChecklistItemCount(val)
        return nil
    }
    res["completedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetCompletedBy(val.(*IdentitySet))
        return nil
    }
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCompletedDateTime(val)
        return nil
    }
    res["conversationThreadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConversationThreadId(val)
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
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerTaskDetails() })
        if err != nil {
            return err
        }
        m.SetDetails(val.(*PlannerTaskDetails))
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDueDateTime(val)
        return nil
    }
    res["hasDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasDescription(val)
        return nil
    }
    res["orderHint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrderHint(val)
        return nil
    }
    res["percentComplete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPercentComplete(val)
        return nil
    }
    res["planId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPlanId(val)
        return nil
    }
    res["previewType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerPreviewType)
        if err != nil {
            return err
        }
        cast := val.(PlannerPreviewType)
        m.SetPreviewType(&cast)
        return nil
    }
    res["progressTaskBoardFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerProgressTaskBoardTaskFormat() })
        if err != nil {
            return err
        }
        m.SetProgressTaskBoardFormat(val.(*PlannerProgressTaskBoardTaskFormat))
        return nil
    }
    res["referenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetReferenceCount(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    return res
}
func (m *PlannerTask) IsNil()(bool) {
    return m == nil
}
func (m *PlannerTask) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeChecklistItemCount", m.GetActiveChecklistItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appliedCategories", m.GetAppliedCategories())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignedToTaskBoardFormat", m.GetAssignedToTaskBoardFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assigneePriority", m.GetAssigneePriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignments", m.GetAssignments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bucketId", m.GetBucketId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bucketTaskBoardFormat", m.GetBucketTaskBoardFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("checklistItemCount", m.GetChecklistItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("completedBy", m.GetCompletedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("conversationThreadId", m.GetConversationThreadId())
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
        err = writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasDescription", m.GetHasDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("orderHint", m.GetOrderHint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentComplete", m.GetPercentComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("planId", m.GetPlanId())
        if err != nil {
            return err
        }
    }
    if m.GetPreviewType() != nil {
        cast := m.GetPreviewType().String()
        err = writer.WriteStringValue("previewType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("progressTaskBoardFormat", m.GetProgressTaskBoardFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("referenceCount", m.GetReferenceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PlannerTask) SetActiveChecklistItemCount(value *int32)() {
    m.activeChecklistItemCount = value
}
func (m *PlannerTask) SetAppliedCategories(value *PlannerAppliedCategories)() {
    m.appliedCategories = value
}
func (m *PlannerTask) SetAssignedToTaskBoardFormat(value *PlannerAssignedToTaskBoardTaskFormat)() {
    m.assignedToTaskBoardFormat = value
}
func (m *PlannerTask) SetAssigneePriority(value *string)() {
    m.assigneePriority = value
}
func (m *PlannerTask) SetAssignments(value *PlannerAssignments)() {
    m.assignments = value
}
func (m *PlannerTask) SetBucketId(value *string)() {
    m.bucketId = value
}
func (m *PlannerTask) SetBucketTaskBoardFormat(value *PlannerBucketTaskBoardTaskFormat)() {
    m.bucketTaskBoardFormat = value
}
func (m *PlannerTask) SetChecklistItemCount(value *int32)() {
    m.checklistItemCount = value
}
func (m *PlannerTask) SetCompletedBy(value *IdentitySet)() {
    m.completedBy = value
}
func (m *PlannerTask) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
func (m *PlannerTask) SetConversationThreadId(value *string)() {
    m.conversationThreadId = value
}
func (m *PlannerTask) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
func (m *PlannerTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *PlannerTask) SetDetails(value *PlannerTaskDetails)() {
    m.details = value
}
func (m *PlannerTask) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDateTime = value
}
func (m *PlannerTask) SetHasDescription(value *bool)() {
    m.hasDescription = value
}
func (m *PlannerTask) SetOrderHint(value *string)() {
    m.orderHint = value
}
func (m *PlannerTask) SetPercentComplete(value *int32)() {
    m.percentComplete = value
}
func (m *PlannerTask) SetPlanId(value *string)() {
    m.planId = value
}
func (m *PlannerTask) SetPreviewType(value *PlannerPreviewType)() {
    m.previewType = value
}
func (m *PlannerTask) SetProgressTaskBoardFormat(value *PlannerProgressTaskBoardTaskFormat)() {
    m.progressTaskBoardFormat = value
}
func (m *PlannerTask) SetReferenceCount(value *int32)() {
    m.referenceCount = value
}
func (m *PlannerTask) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
func (m *PlannerTask) SetTitle(value *string)() {
    m.title = value
}
