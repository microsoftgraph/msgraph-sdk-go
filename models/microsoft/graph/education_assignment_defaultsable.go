package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationAssignmentDefaultsable 
type EducationAssignmentDefaultsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddedStudentAction()(*EducationAddedStudentAction)
    GetAddToCalendarAction()(*EducationAddToCalendarOptions)
    GetDueTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)
    GetNotificationChannelUrl()(*string)
    SetAddedStudentAction(value *EducationAddedStudentAction)()
    SetAddToCalendarAction(value *EducationAddToCalendarOptions)()
    SetDueTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)()
    SetNotificationChannelUrl(value *string)()
}
