package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSchoolable 
type EducationSchoolable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    EducationOrganizationable
    GetAddress()(PhysicalAddressable)
    GetAdministrativeUnit()(AdministrativeUnitable)
    GetClasses()([]EducationClassable)
    GetCreatedBy()(IdentitySetable)
    GetExternalId()(*string)
    GetExternalPrincipalId()(*string)
    GetFax()(*string)
    GetHighestGrade()(*string)
    GetLowestGrade()(*string)
    GetPhone()(*string)
    GetPrincipalEmail()(*string)
    GetPrincipalName()(*string)
    GetSchoolNumber()(*string)
    GetUsers()([]EducationUserable)
    SetAddress(value PhysicalAddressable)()
    SetAdministrativeUnit(value AdministrativeUnitable)()
    SetClasses(value []EducationClassable)()
    SetCreatedBy(value IdentitySetable)()
    SetExternalId(value *string)()
    SetExternalPrincipalId(value *string)()
    SetFax(value *string)()
    SetHighestGrade(value *string)()
    SetLowestGrade(value *string)()
    SetPhone(value *string)()
    SetPrincipalEmail(value *string)()
    SetPrincipalName(value *string)()
    SetSchoolNumber(value *string)()
    SetUsers(value []EducationUserable)()
}
