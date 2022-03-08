package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationCourseable 
type EducationCourseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCourseNumber()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetSubject()(*string)
    SetCourseNumber(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetSubject(value *string)()
}
