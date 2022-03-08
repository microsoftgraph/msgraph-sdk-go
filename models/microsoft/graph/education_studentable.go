package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationStudentable 
type EducationStudentable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBirthDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetExternalId()(*string)
    GetGender()(*EducationGender)
    GetGrade()(*string)
    GetGraduationYear()(*string)
    GetStudentNumber()(*string)
    SetBirthDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetExternalId(value *string)()
    SetGender(value *EducationGender)()
    SetGrade(value *string)()
    SetGraduationYear(value *string)()
    SetStudentNumber(value *string)()
}
