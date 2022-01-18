package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationAssignment 
type EducationAssignment struct {
    Entity
    // Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none value. Currently supports only two values: none or assignIfOpen.
    addedStudentAction *EducationAddedStudentAction;
    // Optional field to control the assignment behavior  for adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
    addToCalendarAction *EducationAddToCalendarOptions;
    // Identifies whether students can submit after the due date. If this property is not specified during create, it defaults to true.
    allowLateSubmissions *bool;
    // Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
    allowStudentsToAddResourcesToSubmission *bool;
    // The date when the assignment should become active.  If in the future, the assignment is not shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
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
    // Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl is not allowed after the assignment has been published.
    notificationChannelUrl *string;
    // Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
    resources []EducationAssignmentResource;
    // Folder URL where all the file resources for this assignment are stored.
    resourcesFolderUrl *string;
    // When set, the grading rubric attached to this assignment.
    rubric *EducationRubric;
    // Status of the Assignment.  You can not PATCH this value.  Possible values are: draft, scheduled, published, assigned.
    status *EducationAssignmentStatus;
    // Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
    submissions []EducationSubmission;
    // The deep link URL for the given assignment.
    webUrl *string;
}
// NewEducationAssignment instantiates a new educationAssignment and sets the default values.
func NewEducationAssignment()(*EducationAssignment) {
    m := &EducationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetAddedStudentAction gets the addedStudentAction property value. Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none value. Currently supports only two values: none or assignIfOpen.
func (m *EducationAssignment) GetAddedStudentAction()(*EducationAddedStudentAction) {
    if m == nil {
        return nil
    } else {
        return m.addedStudentAction
    }
}
// GetAddToCalendarAction gets the addToCalendarAction property value. Optional field to control the assignment behavior  for adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
func (m *EducationAssignment) GetAddToCalendarAction()(*EducationAddToCalendarOptions) {
    if m == nil {
        return nil
    } else {
        return m.addToCalendarAction
    }
}
// GetAllowLateSubmissions gets the allowLateSubmissions property value. Identifies whether students can submit after the due date. If this property is not specified during create, it defaults to true.
func (m *EducationAssignment) GetAllowLateSubmissions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowLateSubmissions
    }
}
// GetAllowStudentsToAddResourcesToSubmission gets the allowStudentsToAddResourcesToSubmission property value. Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
func (m *EducationAssignment) GetAllowStudentsToAddResourcesToSubmission()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowStudentsToAddResourcesToSubmission
    }
}
// GetAssignDateTime gets the assignDateTime property value. The date when the assignment should become active.  If in the future, the assignment is not shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetAssignDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.assignDateTime
    }
}
// GetAssignedDateTime gets the assignedDateTime property value. The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.assignedDateTime
    }
}
// GetAssignTo gets the assignTo property value. Which users, or whole class should receive a submission object once the assignment is published.
func (m *EducationAssignment) GetAssignTo()(*EducationAssignmentRecipient) {
    if m == nil {
        return nil
    } else {
        return m.assignTo
    }
}
// GetCategories gets the categories property value. When set, enables users to easily find assignments of a given type.  Read-only. Nullable.
func (m *EducationAssignment) GetCategories()([]EducationCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetClassId gets the classId property value. Class which this assignment belongs.
func (m *EducationAssignment) GetClassId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classId
    }
}
// GetCloseDateTime gets the closeDateTime property value. Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetCloseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closeDateTime
    }
}
// GetCreatedBy gets the createdBy property value. Who created the assignment.
func (m *EducationAssignment) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. Name of the assignment.
func (m *EducationAssignment) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDueDateTime gets the dueDateTime property value. Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// GetGrading gets the grading property value. How the assignment will be graded.
func (m *EducationAssignment) GetGrading()(*EducationAssignmentGradeType) {
    if m == nil {
        return nil
    } else {
        return m.grading
    }
}
// GetInstructions gets the instructions property value. Instructions for the assignment.  This along with the display name tell the student what to do.
func (m *EducationAssignment) GetInstructions()(*EducationItemBody) {
    if m == nil {
        return nil
    } else {
        return m.instructions
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. Who last modified the assignment.
func (m *EducationAssignment) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNotificationChannelUrl gets the notificationChannelUrl property value. Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl is not allowed after the assignment has been published.
func (m *EducationAssignment) GetNotificationChannelUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationChannelUrl
    }
}
// GetResources gets the resources property value. Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
func (m *EducationAssignment) GetResources()([]EducationAssignmentResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// GetResourcesFolderUrl gets the resourcesFolderUrl property value. Folder URL where all the file resources for this assignment are stored.
func (m *EducationAssignment) GetResourcesFolderUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourcesFolderUrl
    }
}
// GetRubric gets the rubric property value. When set, the grading rubric attached to this assignment.
func (m *EducationAssignment) GetRubric()(*EducationRubric) {
    if m == nil {
        return nil
    } else {
        return m.rubric
    }
}
// GetStatus gets the status property value. Status of the Assignment.  You can not PATCH this value.  Possible values are: draft, scheduled, published, assigned.
func (m *EducationAssignment) GetStatus()(*EducationAssignmentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSubmissions gets the submissions property value. Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationAssignment) GetSubmissions()([]EducationSubmission) {
    if m == nil {
        return nil
    } else {
        return m.submissions
    }
}
// GetWebUrl gets the webUrl property value. The deep link URL for the given assignment.
func (m *EducationAssignment) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedStudentAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddedStudentAction)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationAddedStudentAction)
            m.SetAddedStudentAction(&cast)
        }
        return nil
    }
    res["addToCalendarAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddToCalendarOptions)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationAddToCalendarOptions)
            m.SetAddToCalendarAction(&cast)
        }
        return nil
    }
    res["allowLateSubmissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLateSubmissions(val)
        }
        return nil
    }
    res["allowStudentsToAddResourcesToSubmission"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowStudentsToAddResourcesToSubmission(val)
        }
        return nil
    }
    res["assignDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignDateTime(val)
        }
        return nil
    }
    res["assignedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedDateTime(val)
        }
        return nil
    }
    res["assignTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignTo(val.(*EducationAssignmentRecipient))
        }
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationCategory))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["classId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassId(val)
        }
        return nil
    }
    res["closeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloseDateTime(val)
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
    res["grading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentGradeType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrading(val.(*EducationAssignmentGradeType))
        }
        return nil
    }
    res["instructions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstructions(val.(*EducationItemBody))
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
    res["notificationChannelUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationChannelUrl(val)
        }
        return nil
    }
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationAssignmentResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationAssignmentResource))
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
    res["rubric"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationRubric() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRubric(val.(*EducationRubric))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAssignmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationAssignmentStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["submissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSubmission() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSubmission, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationSubmission))
            }
            m.SetSubmissions(res)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
func (m *EducationAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetAddToCalendarAction() != nil {
        cast := m.GetAddToCalendarAction().String()
        err = writer.WriteStringValue("addToCalendarAction", &cast)
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
// SetAddedStudentAction sets the addedStudentAction property value. Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none value. Currently supports only two values: none or assignIfOpen.
func (m *EducationAssignment) SetAddedStudentAction(value *EducationAddedStudentAction)() {
    if m != nil {
        m.addedStudentAction = value
    }
}
// SetAddToCalendarAction sets the addToCalendarAction property value. Optional field to control the assignment behavior  for adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
func (m *EducationAssignment) SetAddToCalendarAction(value *EducationAddToCalendarOptions)() {
    if m != nil {
        m.addToCalendarAction = value
    }
}
// SetAllowLateSubmissions sets the allowLateSubmissions property value. Identifies whether students can submit after the due date. If this property is not specified during create, it defaults to true.
func (m *EducationAssignment) SetAllowLateSubmissions(value *bool)() {
    if m != nil {
        m.allowLateSubmissions = value
    }
}
// SetAllowStudentsToAddResourcesToSubmission sets the allowStudentsToAddResourcesToSubmission property value. Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
func (m *EducationAssignment) SetAllowStudentsToAddResourcesToSubmission(value *bool)() {
    if m != nil {
        m.allowStudentsToAddResourcesToSubmission = value
    }
}
// SetAssignDateTime sets the assignDateTime property value. The date when the assignment should become active.  If in the future, the assignment is not shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetAssignDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.assignDateTime = value
    }
}
// SetAssignedDateTime sets the assignedDateTime property value. The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.assignedDateTime = value
    }
}
// SetAssignTo sets the assignTo property value. Which users, or whole class should receive a submission object once the assignment is published.
func (m *EducationAssignment) SetAssignTo(value *EducationAssignmentRecipient)() {
    if m != nil {
        m.assignTo = value
    }
}
// SetCategories sets the categories property value. When set, enables users to easily find assignments of a given type.  Read-only. Nullable.
func (m *EducationAssignment) SetCategories(value []EducationCategory)() {
    if m != nil {
        m.categories = value
    }
}
// SetClassId sets the classId property value. Class which this assignment belongs.
func (m *EducationAssignment) SetClassId(value *string)() {
    if m != nil {
        m.classId = value
    }
}
// SetCloseDateTime sets the closeDateTime property value. Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetCloseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.closeDateTime = value
    }
}
// SetCreatedBy sets the createdBy property value. Who created the assignment.
func (m *EducationAssignment) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. Name of the assignment.
func (m *EducationAssignment) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDueDateTime sets the dueDateTime property value. Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.dueDateTime = value
    }
}
// SetGrading sets the grading property value. How the assignment will be graded.
func (m *EducationAssignment) SetGrading(value *EducationAssignmentGradeType)() {
    if m != nil {
        m.grading = value
    }
}
// SetInstructions sets the instructions property value. Instructions for the assignment.  This along with the display name tell the student what to do.
func (m *EducationAssignment) SetInstructions(value *EducationItemBody)() {
    if m != nil {
        m.instructions = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Who last modified the assignment.
func (m *EducationAssignment) SetLastModifiedBy(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNotificationChannelUrl sets the notificationChannelUrl property value. Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl is not allowed after the assignment has been published.
func (m *EducationAssignment) SetNotificationChannelUrl(value *string)() {
    if m != nil {
        m.notificationChannelUrl = value
    }
}
// SetResources sets the resources property value. Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
func (m *EducationAssignment) SetResources(value []EducationAssignmentResource)() {
    if m != nil {
        m.resources = value
    }
}
// SetResourcesFolderUrl sets the resourcesFolderUrl property value. Folder URL where all the file resources for this assignment are stored.
func (m *EducationAssignment) SetResourcesFolderUrl(value *string)() {
    if m != nil {
        m.resourcesFolderUrl = value
    }
}
// SetRubric sets the rubric property value. When set, the grading rubric attached to this assignment.
func (m *EducationAssignment) SetRubric(value *EducationRubric)() {
    if m != nil {
        m.rubric = value
    }
}
// SetStatus sets the status property value. Status of the Assignment.  You can not PATCH this value.  Possible values are: draft, scheduled, published, assigned.
func (m *EducationAssignment) SetStatus(value *EducationAssignmentStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetSubmissions sets the submissions property value. Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationAssignment) SetSubmissions(value []EducationSubmission)() {
    if m != nil {
        m.submissions = value
    }
}
// SetWebUrl sets the webUrl property value. The deep link URL for the given assignment.
func (m *EducationAssignment) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
