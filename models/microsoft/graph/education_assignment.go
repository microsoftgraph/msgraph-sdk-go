package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationAssignment struct {
    Entity
    // Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none value. Currently supports only two values: none or assignIfOpen.
    addedStudentAction *EducationAddedStudentAction;
    // Identifies whether students can submit after the due date. If this property isn't specified during create, it defaults to true.
    allowLateSubmissions *bool;
    // Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
    allowStudentsToAddResourcesToSubmission *bool;
    // The date when the assignment should become active.  If in the future, the assignment isn't shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    assignDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    assignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Which users, or whole class should receive a submission object once the assignment is published.
    assignTo *EducationAssignmentRecipient;
    // When set, enables users to easily find assignments of a given type.  Read-only. Nullable.
    categories []EducationCategory;
    // Class which this assignment belongs.
    classId *string;
    // Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    closeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Who created the assignment.
    createdBy *IdentitySet;
    // Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the assignment.
    displayName *string;
    // Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    dueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // How the assignment will be graded.
    grading *EducationAssignmentGradeType;
    // Instructions for the assignment.  This along with the display name tell the student what to do.
    instructions *EducationItemBody;
    // Who last modified the assignment.
    lastModifiedBy *IdentitySet;
    // Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl isn't allowed after the assignment has been published.
    notificationChannelUrl *string;
    // Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
    resources []EducationAssignmentResource;
    // Folder URL where all the file resources for this assignment are stored.
    resourcesFolderUrl *string;
    // When set, the grading rubric attached to this assignment.
    rubric *EducationRubric;
    // Status of the Assignment.  You can't PATCH this value.  Possible values are: draft, scheduled, published, assigned.
    status *EducationAssignmentStatus;
    // Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
    submissions []EducationSubmission;
    // The deep link URL for the given assignment.
    webUrl *string;
}
// Instantiates a new educationAssignment and sets the default values.
func NewEducationAssignment()(*EducationAssignment) {
    m := &EducationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the addedStudentAction property value. Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none value. Currently supports only two values: none or assignIfOpen.
func (m *EducationAssignment) GetAddedStudentAction()(*EducationAddedStudentAction) {
    if m == nil {
        return nil
    } else {
        return m.addedStudentAction
    }
}
// Gets the allowLateSubmissions property value. Identifies whether students can submit after the due date. If this property isn't specified during create, it defaults to true.
func (m *EducationAssignment) GetAllowLateSubmissions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowLateSubmissions
    }
}
// Gets the allowStudentsToAddResourcesToSubmission property value. Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
func (m *EducationAssignment) GetAllowStudentsToAddResourcesToSubmission()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowStudentsToAddResourcesToSubmission
    }
}
// Gets the assignDateTime property value. The date when the assignment should become active.  If in the future, the assignment isn't shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetAssignDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.assignDateTime
    }
}
// Gets the assignedDateTime property value. The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.assignedDateTime
    }
}
// Gets the assignTo property value. Which users, or whole class should receive a submission object once the assignment is published.
func (m *EducationAssignment) GetAssignTo()(*EducationAssignmentRecipient) {
    if m == nil {
        return nil
    } else {
        return m.assignTo
    }
}
// Gets the categories property value. When set, enables users to easily find assignments of a given type.  Read-only. Nullable.
func (m *EducationAssignment) GetCategories()([]EducationCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the classId property value. Class which this assignment belongs.
func (m *EducationAssignment) GetClassId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classId
    }
}
// Gets the closeDateTime property value. Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetCloseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closeDateTime
    }
}
// Gets the createdBy property value. Who created the assignment.
func (m *EducationAssignment) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the displayName property value. Name of the assignment.
func (m *EducationAssignment) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the dueDateTime property value. Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// Gets the grading property value. How the assignment will be graded.
func (m *EducationAssignment) GetGrading()(*EducationAssignmentGradeType) {
    if m == nil {
        return nil
    } else {
        return m.grading
    }
}
// Gets the instructions property value. Instructions for the assignment.  This along with the display name tell the student what to do.
func (m *EducationAssignment) GetInstructions()(*EducationItemBody) {
    if m == nil {
        return nil
    } else {
        return m.instructions
    }
}
// Gets the lastModifiedBy property value. Who last modified the assignment.
func (m *EducationAssignment) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedDateTime property value. Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the notificationChannelUrl property value. Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl isn't allowed after the assignment has been published.
func (m *EducationAssignment) GetNotificationChannelUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationChannelUrl
    }
}
// Gets the resources property value. Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
func (m *EducationAssignment) GetResources()([]EducationAssignmentResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// Gets the resourcesFolderUrl property value. Folder URL where all the file resources for this assignment are stored.
func (m *EducationAssignment) GetResourcesFolderUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourcesFolderUrl
    }
}
// Gets the rubric property value. When set, the grading rubric attached to this assignment.
func (m *EducationAssignment) GetRubric()(*EducationRubric) {
    if m == nil {
        return nil
    } else {
        return m.rubric
    }
}
// Gets the status property value. Status of the Assignment.  You can't PATCH this value.  Possible values are: draft, scheduled, published, assigned.
func (m *EducationAssignment) GetStatus()(*EducationAssignmentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the submissions property value. Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationAssignment) GetSubmissions()([]EducationSubmission) {
    if m == nil {
        return nil
    } else {
        return m.submissions
    }
}
// Gets the webUrl property value. The deep link URL for the given assignment.
func (m *EducationAssignment) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *EducationAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedStudentAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddedStudentAction)
        if err != nil {
            return err
        }
        cast := val.(EducationAddedStudentAction)
        m.SetAddedStudentAction(&cast)
        return nil
    }
    res["allowLateSubmissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowLateSubmissions(val)
        return nil
    }
    res["allowStudentsToAddResourcesToSubmission"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowStudentsToAddResourcesToSubmission(val)
        return nil
    }
    res["assignDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAssignDateTime(val)
        return nil
    }
    res["assignedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAssignedDateTime(val)
        return nil
    }
    res["assignTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentRecipient() })
        if err != nil {
            return err
        }
        m.SetAssignTo(val.(*EducationAssignmentRecipient))
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationCategory() })
        if err != nil {
            return err
        }
        res := make([]EducationCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationCategory))
        }
        m.SetCategories(res)
        return nil
    }
    res["classId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClassId(val)
        return nil
    }
    res["closeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCloseDateTime(val)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
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
    res["grading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentGradeType() })
        if err != nil {
            return err
        }
        m.SetGrading(val.(*EducationAssignmentGradeType))
        return nil
    }
    res["instructions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationItemBody() })
        if err != nil {
            return err
        }
        m.SetInstructions(val.(*EducationItemBody))
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
    res["notificationChannelUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationChannelUrl(val)
        return nil
    }
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentResource() })
        if err != nil {
            return err
        }
        res := make([]EducationAssignmentResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationAssignmentResource))
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
    res["rubric"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationRubric() })
        if err != nil {
            return err
        }
        m.SetRubric(val.(*EducationRubric))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAssignmentStatus)
        if err != nil {
            return err
        }
        cast := val.(EducationAssignmentStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["submissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSubmission() })
        if err != nil {
            return err
        }
        res := make([]EducationSubmission, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationSubmission))
        }
        m.SetSubmissions(res)
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *EducationAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EducationAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddedStudentAction() != nil {
        cast := m.GetAddedStudentAction().String()
        err = writer.WriteStringValue("addedStudentAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowLateSubmissions", m.GetAllowLateSubmissions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowStudentsToAddResourcesToSubmission", m.GetAllowStudentsToAddResourcesToSubmission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("assignDateTime", m.GetAssignDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("assignedDateTime", m.GetAssignedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignTo", m.GetAssignTo())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("classId", m.GetClassId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closeDateTime", m.GetCloseDateTime())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteObjectValue("grading", m.GetGrading())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("instructions", m.GetInstructions())
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
        err = writer.WriteStringValue("notificationChannelUrl", m.GetNotificationChannelUrl())
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
        err = writer.WriteObjectValue("rubric", m.GetRubric())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSubmissions()))
        for i, v := range m.GetSubmissions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("submissions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the addedStudentAction property value. Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none value. Currently supports only two values: none or assignIfOpen.
// Parameters:
//  - value : Value to set for the addedStudentAction property.
func (m *EducationAssignment) SetAddedStudentAction(value *EducationAddedStudentAction)() {
    m.addedStudentAction = value
}
// Sets the allowLateSubmissions property value. Identifies whether students can submit after the due date. If this property isn't specified during create, it defaults to true.
// Parameters:
//  - value : Value to set for the allowLateSubmissions property.
func (m *EducationAssignment) SetAllowLateSubmissions(value *bool)() {
    m.allowLateSubmissions = value
}
// Sets the allowStudentsToAddResourcesToSubmission property value. Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
// Parameters:
//  - value : Value to set for the allowStudentsToAddResourcesToSubmission property.
func (m *EducationAssignment) SetAllowStudentsToAddResourcesToSubmission(value *bool)() {
    m.allowStudentsToAddResourcesToSubmission = value
}
// Sets the assignDateTime property value. The date when the assignment should become active.  If in the future, the assignment isn't shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the assignDateTime property.
func (m *EducationAssignment) SetAssignDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignDateTime = value
}
// Sets the assignedDateTime property value. The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the assignedDateTime property.
func (m *EducationAssignment) SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignedDateTime = value
}
// Sets the assignTo property value. Which users, or whole class should receive a submission object once the assignment is published.
// Parameters:
//  - value : Value to set for the assignTo property.
func (m *EducationAssignment) SetAssignTo(value *EducationAssignmentRecipient)() {
    m.assignTo = value
}
// Sets the categories property value. When set, enables users to easily find assignments of a given type.  Read-only. Nullable.
// Parameters:
//  - value : Value to set for the categories property.
func (m *EducationAssignment) SetCategories(value []EducationCategory)() {
    m.categories = value
}
// Sets the classId property value. Class which this assignment belongs.
// Parameters:
//  - value : Value to set for the classId property.
func (m *EducationAssignment) SetClassId(value *string)() {
    m.classId = value
}
// Sets the closeDateTime property value. Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the closeDateTime property.
func (m *EducationAssignment) SetCloseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closeDateTime = value
}
// Sets the createdBy property value. Who created the assignment.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *EducationAssignment) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *EducationAssignment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the displayName property value. Name of the assignment.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EducationAssignment) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the dueDateTime property value. Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the dueDateTime property.
func (m *EducationAssignment) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDateTime = value
}
// Sets the grading property value. How the assignment will be graded.
// Parameters:
//  - value : Value to set for the grading property.
func (m *EducationAssignment) SetGrading(value *EducationAssignmentGradeType)() {
    m.grading = value
}
// Sets the instructions property value. Instructions for the assignment.  This along with the display name tell the student what to do.
// Parameters:
//  - value : Value to set for the instructions property.
func (m *EducationAssignment) SetInstructions(value *EducationItemBody)() {
    m.instructions = value
}
// Sets the lastModifiedBy property value. Who last modified the assignment.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *EducationAssignment) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedDateTime property value. Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *EducationAssignment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the notificationChannelUrl property value. Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl isn't allowed after the assignment has been published.
// Parameters:
//  - value : Value to set for the notificationChannelUrl property.
func (m *EducationAssignment) SetNotificationChannelUrl(value *string)() {
    m.notificationChannelUrl = value
}
// Sets the resources property value. Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
// Parameters:
//  - value : Value to set for the resources property.
func (m *EducationAssignment) SetResources(value []EducationAssignmentResource)() {
    m.resources = value
}
// Sets the resourcesFolderUrl property value. Folder URL where all the file resources for this assignment are stored.
// Parameters:
//  - value : Value to set for the resourcesFolderUrl property.
func (m *EducationAssignment) SetResourcesFolderUrl(value *string)() {
    m.resourcesFolderUrl = value
}
// Sets the rubric property value. When set, the grading rubric attached to this assignment.
// Parameters:
//  - value : Value to set for the rubric property.
func (m *EducationAssignment) SetRubric(value *EducationRubric)() {
    m.rubric = value
}
// Sets the status property value. Status of the Assignment.  You can't PATCH this value.  Possible values are: draft, scheduled, published, assigned.
// Parameters:
//  - value : Value to set for the status property.
func (m *EducationAssignment) SetStatus(value *EducationAssignmentStatus)() {
    m.status = value
}
// Sets the submissions property value. Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
// Parameters:
//  - value : Value to set for the submissions property.
func (m *EducationAssignment) SetSubmissions(value []EducationSubmission)() {
    m.submissions = value
}
// Sets the webUrl property value. The deep link URL for the given assignment.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *EducationAssignment) SetWebUrl(value *string)() {
    m.webUrl = value
}
