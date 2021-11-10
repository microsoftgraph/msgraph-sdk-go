package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerTask struct {
    Entity
    // Number of checklist items with value set to false, representing incomplete items.
    activeChecklistItemCount *int32;
    // The categories to which the task has been applied. See applied Categories for possible values.
    appliedCategories *PlannerAppliedCategories;
    // Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
    assignedToTaskBoardFormat *PlannerAssignedToTaskBoardTaskFormat;
    // Hint used to order items of this type in a list view. The format is defined as outlined here.
    assigneePriority *string;
    // The set of assignees the task is assigned to.
    assignments *PlannerAssignments;
    // Bucket ID to which the task belongs. The bucket needs to be in the plan that the task is in. It is 28 characters long and case-sensitive. Format validation is done on the service.
    bucketId *string;
    // Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
    bucketTaskBoardFormat *PlannerBucketTaskBoardTaskFormat;
    // Number of checklist items that are present on the task.
    checklistItemCount *int32;
    // Identity of the user that completed the task.
    completedBy *IdentitySet;
    // Read-only. Date and time at which the 'percentComplete' of the task is set to '100'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group.
    conversationThreadId *string;
    // Identity of the user that created the task.
    createdBy *IdentitySet;
    // Read-only. Date and time at which the task is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. Nullable. Additional details about the task.
    details *PlannerTaskDetails;
    // Date and time at which the task is due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    dueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. Value is true if the details object of the task has a non-empty description and false otherwise.
    hasDescription *bool;
    // Hint used to order items of this type in a list view. The format is defined as outlined here.
    orderHint *string;
    // Percentage of task completion. When set to 100, the task is considered completed.
    percentComplete *int32;
    // Plan ID to which the task belongs.
    planId *string;
    // This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference.
    previewType *PlannerPreviewType;
    // Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
    progressTaskBoardFormat *PlannerProgressTaskBoardTaskFormat;
    // Number of external references that exist on the task.
    referenceCount *int32;
    // Date and time at which the task starts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Title of the task.
    title *string;
}
// Instantiates a new plannerTask and sets the default values.
func NewPlannerTask()(*PlannerTask) {
    m := &PlannerTask{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeChecklistItemCount property value. Number of checklist items with value set to false, representing incomplete items.
func (m *PlannerTask) GetActiveChecklistItemCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeChecklistItemCount
    }
}
// Gets the appliedCategories property value. The categories to which the task has been applied. See applied Categories for possible values.
func (m *PlannerTask) GetAppliedCategories()(*PlannerAppliedCategories) {
    if m == nil {
        return nil
    } else {
        return m.appliedCategories
    }
}
// Gets the assignedToTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
func (m *PlannerTask) GetAssignedToTaskBoardFormat()(*PlannerAssignedToTaskBoardTaskFormat) {
    if m == nil {
        return nil
    } else {
        return m.assignedToTaskBoardFormat
    }
}
// Gets the assigneePriority property value. Hint used to order items of this type in a list view. The format is defined as outlined here.
func (m *PlannerTask) GetAssigneePriority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assigneePriority
    }
}
// Gets the assignments property value. The set of assignees the task is assigned to.
func (m *PlannerTask) GetAssignments()(*PlannerAssignments) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the bucketId property value. Bucket ID to which the task belongs. The bucket needs to be in the plan that the task is in. It is 28 characters long and case-sensitive. Format validation is done on the service.
func (m *PlannerTask) GetBucketId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bucketId
    }
}
// Gets the bucketTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
func (m *PlannerTask) GetBucketTaskBoardFormat()(*PlannerBucketTaskBoardTaskFormat) {
    if m == nil {
        return nil
    } else {
        return m.bucketTaskBoardFormat
    }
}
// Gets the checklistItemCount property value. Number of checklist items that are present on the task.
func (m *PlannerTask) GetChecklistItemCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.checklistItemCount
    }
}
// Gets the completedBy property value. Identity of the user that completed the task.
func (m *PlannerTask) GetCompletedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.completedBy
    }
}
// Gets the completedDateTime property value. Read-only. Date and time at which the 'percentComplete' of the task is set to '100'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerTask) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// Gets the conversationThreadId property value. Thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group.
func (m *PlannerTask) GetConversationThreadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conversationThreadId
    }
}
// Gets the createdBy property value. Identity of the user that created the task.
func (m *PlannerTask) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. Read-only. Date and time at which the task is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the details property value. Read-only. Nullable. Additional details about the task.
func (m *PlannerTask) GetDetails()(*PlannerTaskDetails) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// Gets the dueDateTime property value. Date and time at which the task is due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerTask) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// Gets the hasDescription property value. Read-only. Value is true if the details object of the task has a non-empty description and false otherwise.
func (m *PlannerTask) GetHasDescription()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasDescription
    }
}
// Gets the orderHint property value. Hint used to order items of this type in a list view. The format is defined as outlined here.
func (m *PlannerTask) GetOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderHint
    }
}
// Gets the percentComplete property value. Percentage of task completion. When set to 100, the task is considered completed.
func (m *PlannerTask) GetPercentComplete()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.percentComplete
    }
}
// Gets the planId property value. Plan ID to which the task belongs.
func (m *PlannerTask) GetPlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.planId
    }
}
// Gets the previewType property value. This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference.
func (m *PlannerTask) GetPreviewType()(*PlannerPreviewType) {
    if m == nil {
        return nil
    } else {
        return m.previewType
    }
}
// Gets the progressTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
func (m *PlannerTask) GetProgressTaskBoardFormat()(*PlannerProgressTaskBoardTaskFormat) {
    if m == nil {
        return nil
    } else {
        return m.progressTaskBoardFormat
    }
}
// Gets the referenceCount property value. Number of external references that exist on the task.
func (m *PlannerTask) GetReferenceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.referenceCount
    }
}
// Gets the startDateTime property value. Date and time at which the task starts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerTask) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the title property value. Title of the task.
func (m *PlannerTask) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// The deserialization information for the current model
func (m *PlannerTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeChecklistItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveChecklistItemCount(val)
        }
        return nil
    }
    res["appliedCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerAppliedCategories() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedCategories(val.(*PlannerAppliedCategories))
        }
        return nil
    }
    res["assignedToTaskBoardFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerAssignedToTaskBoardTaskFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedToTaskBoardFormat(val.(*PlannerAssignedToTaskBoardTaskFormat))
        }
        return nil
    }
    res["assigneePriority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigneePriority(val)
        }
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerAssignments() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignments(val.(*PlannerAssignments))
        }
        return nil
    }
    res["bucketId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucketId(val)
        }
        return nil
    }
    res["bucketTaskBoardFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerBucketTaskBoardTaskFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucketTaskBoardFormat(val.(*PlannerBucketTaskBoardTaskFormat))
        }
        return nil
    }
    res["checklistItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecklistItemCount(val)
        }
        return nil
    }
    res["completedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["conversationThreadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationThreadId(val)
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
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerTaskDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(*PlannerTaskDetails))
        }
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val)
        }
        return nil
    }
    res["hasDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasDescription(val)
        }
        return nil
    }
    res["orderHint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderHint(val)
        }
        return nil
    }
    res["percentComplete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentComplete(val)
        }
        return nil
    }
    res["planId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanId(val)
        }
        return nil
    }
    res["previewType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerPreviewType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PlannerPreviewType)
            m.SetPreviewType(&cast)
        }
        return nil
    }
    res["progressTaskBoardFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerProgressTaskBoardTaskFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgressTaskBoardFormat(val.(*PlannerProgressTaskBoardTaskFormat))
        }
        return nil
    }
    res["referenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceCount(val)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
func (m *PlannerTask) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activeChecklistItemCount property value. Number of checklist items with value set to false, representing incomplete items.
// Parameters:
//  - value : Value to set for the activeChecklistItemCount property.
func (m *PlannerTask) SetActiveChecklistItemCount(value *int32)() {
    m.activeChecklistItemCount = value
}
// Sets the appliedCategories property value. The categories to which the task has been applied. See applied Categories for possible values.
// Parameters:
//  - value : Value to set for the appliedCategories property.
func (m *PlannerTask) SetAppliedCategories(value *PlannerAppliedCategories)() {
    m.appliedCategories = value
}
// Sets the assignedToTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
// Parameters:
//  - value : Value to set for the assignedToTaskBoardFormat property.
func (m *PlannerTask) SetAssignedToTaskBoardFormat(value *PlannerAssignedToTaskBoardTaskFormat)() {
    m.assignedToTaskBoardFormat = value
}
// Sets the assigneePriority property value. Hint used to order items of this type in a list view. The format is defined as outlined here.
// Parameters:
//  - value : Value to set for the assigneePriority property.
func (m *PlannerTask) SetAssigneePriority(value *string)() {
    m.assigneePriority = value
}
// Sets the assignments property value. The set of assignees the task is assigned to.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *PlannerTask) SetAssignments(value *PlannerAssignments)() {
    m.assignments = value
}
// Sets the bucketId property value. Bucket ID to which the task belongs. The bucket needs to be in the plan that the task is in. It is 28 characters long and case-sensitive. Format validation is done on the service.
// Parameters:
//  - value : Value to set for the bucketId property.
func (m *PlannerTask) SetBucketId(value *string)() {
    m.bucketId = value
}
// Sets the bucketTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
// Parameters:
//  - value : Value to set for the bucketTaskBoardFormat property.
func (m *PlannerTask) SetBucketTaskBoardFormat(value *PlannerBucketTaskBoardTaskFormat)() {
    m.bucketTaskBoardFormat = value
}
// Sets the checklistItemCount property value. Number of checklist items that are present on the task.
// Parameters:
//  - value : Value to set for the checklistItemCount property.
func (m *PlannerTask) SetChecklistItemCount(value *int32)() {
    m.checklistItemCount = value
}
// Sets the completedBy property value. Identity of the user that completed the task.
// Parameters:
//  - value : Value to set for the completedBy property.
func (m *PlannerTask) SetCompletedBy(value *IdentitySet)() {
    m.completedBy = value
}
// Sets the completedDateTime property value. Read-only. Date and time at which the 'percentComplete' of the task is set to '100'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the completedDateTime property.
func (m *PlannerTask) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// Sets the conversationThreadId property value. Thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group.
// Parameters:
//  - value : Value to set for the conversationThreadId property.
func (m *PlannerTask) SetConversationThreadId(value *string)() {
    m.conversationThreadId = value
}
// Sets the createdBy property value. Identity of the user that created the task.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *PlannerTask) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. Read-only. Date and time at which the task is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *PlannerTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the details property value. Read-only. Nullable. Additional details about the task.
// Parameters:
//  - value : Value to set for the details property.
func (m *PlannerTask) SetDetails(value *PlannerTaskDetails)() {
    m.details = value
}
// Sets the dueDateTime property value. Date and time at which the task is due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the dueDateTime property.
func (m *PlannerTask) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDateTime = value
}
// Sets the hasDescription property value. Read-only. Value is true if the details object of the task has a non-empty description and false otherwise.
// Parameters:
//  - value : Value to set for the hasDescription property.
func (m *PlannerTask) SetHasDescription(value *bool)() {
    m.hasDescription = value
}
// Sets the orderHint property value. Hint used to order items of this type in a list view. The format is defined as outlined here.
// Parameters:
//  - value : Value to set for the orderHint property.
func (m *PlannerTask) SetOrderHint(value *string)() {
    m.orderHint = value
}
// Sets the percentComplete property value. Percentage of task completion. When set to 100, the task is considered completed.
// Parameters:
//  - value : Value to set for the percentComplete property.
func (m *PlannerTask) SetPercentComplete(value *int32)() {
    m.percentComplete = value
}
// Sets the planId property value. Plan ID to which the task belongs.
// Parameters:
//  - value : Value to set for the planId property.
func (m *PlannerTask) SetPlanId(value *string)() {
    m.planId = value
}
// Sets the previewType property value. This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference.
// Parameters:
//  - value : Value to set for the previewType property.
func (m *PlannerTask) SetPreviewType(value *PlannerPreviewType)() {
    m.previewType = value
}
// Sets the progressTaskBoardFormat property value. Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
// Parameters:
//  - value : Value to set for the progressTaskBoardFormat property.
func (m *PlannerTask) SetProgressTaskBoardFormat(value *PlannerProgressTaskBoardTaskFormat)() {
    m.progressTaskBoardFormat = value
}
// Sets the referenceCount property value. Number of external references that exist on the task.
// Parameters:
//  - value : Value to set for the referenceCount property.
func (m *PlannerTask) SetReferenceCount(value *int32)() {
    m.referenceCount = value
}
// Sets the startDateTime property value. Date and time at which the task starts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *PlannerTask) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the title property value. Title of the task.
// Parameters:
//  - value : Value to set for the title property.
func (m *PlannerTask) SetTitle(value *string)() {
    m.title = value
}
