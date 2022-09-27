package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentable 
type EducationAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedStudentAction()(*EducationAddedStudentAction)
    GetAddToCalendarAction()(*EducationAddToCalendarOptions)
    GetAllowLateSubmissions()(*bool)
    GetAllowStudentsToAddResourcesToSubmission()(*bool)
    GetAssignDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAssignTo()(EducationAssignmentRecipientable)
    GetCategories()([]EducationCategoryable)
    GetClassId()(*string)
    GetCloseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGrading()(EducationAssignmentGradeTypeable)
    GetInstructions()(EducationItemBodyable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotificationChannelUrl()(*string)
    GetResources()([]EducationAssignmentResourceable)
    GetResourcesFolderUrl()(*string)
    GetRubric()(EducationRubricable)
    GetStatus()(*EducationAssignmentStatus)
    GetSubmissions()([]EducationSubmissionable)
    GetWebUrl()(*string)
    SetAddedStudentAction(value *EducationAddedStudentAction)()
    SetAddToCalendarAction(value *EducationAddToCalendarOptions)()
    SetAllowLateSubmissions(value *bool)()
    SetAllowStudentsToAddResourcesToSubmission(value *bool)()
    SetAssignTo(value EducationAssignmentRecipientable)()
    SetCategories(value []EducationCategoryable)()
    SetClassId(value *string)()
    SetCloseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGrading(value EducationAssignmentGradeTypeable)()
    SetInstructions(value EducationItemBodyable)()
    SetNotificationChannelUrl(value *string)()
    SetResources(value []EducationAssignmentResourceable)()
    SetRubric(value EducationRubricable)()
    SetSubmissions(value []EducationSubmissionable)()
}
