package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RubricLevelable 
type RubricLevelable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDescription()(EducationItemBodyable)
    GetDisplayName()(*string)
    GetGrading()(EducationAssignmentGradeTypeable)
    GetLevelId()(*string)
    SetDescription(value EducationItemBodyable)()
    SetDisplayName(value *string)()
    SetGrading(value EducationAssignmentGradeTypeable)()
    SetLevelId(value *string)()
}
