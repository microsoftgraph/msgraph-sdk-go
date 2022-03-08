package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationClassable 
type EducationClassable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAssignmentCategories()([]EducationCategoryable)
    GetAssignmentDefaults()(EducationAssignmentDefaultsable)
    GetAssignments()([]EducationAssignmentable)
    GetAssignmentSettings()(EducationAssignmentSettingsable)
    GetClassCode()(*string)
    GetCourse()(EducationCourseable)
    GetCreatedBy()(IdentitySetable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetExternalName()(*string)
    GetExternalSource()(*EducationExternalSource)
    GetExternalSourceDetail()(*string)
    GetGrade()(*string)
    GetGroup()(Groupable)
    GetMailNickname()(*string)
    GetMembers()([]EducationUserable)
    GetSchools()([]EducationSchoolable)
    GetTeachers()([]EducationUserable)
    GetTerm()(EducationTermable)
    SetAssignmentCategories(value []EducationCategoryable)()
    SetAssignmentDefaults(value EducationAssignmentDefaultsable)()
    SetAssignments(value []EducationAssignmentable)()
    SetAssignmentSettings(value EducationAssignmentSettingsable)()
    SetClassCode(value *string)()
    SetCourse(value EducationCourseable)()
    SetCreatedBy(value IdentitySetable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetExternalName(value *string)()
    SetExternalSource(value *EducationExternalSource)()
    SetExternalSourceDetail(value *string)()
    SetGrade(value *string)()
    SetGroup(value Groupable)()
    SetMailNickname(value *string)()
    SetMembers(value []EducationUserable)()
    SetSchools(value []EducationSchoolable)()
    SetTeachers(value []EducationUserable)()
    SetTerm(value EducationTermable)()
}
