package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Personable 
type Personable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetBirthday()(*string)
    GetCompanyName()(*string)
    GetDepartment()(*string)
    GetDisplayName()(*string)
    GetGivenName()(*string)
    GetImAddress()(*string)
    GetIsFavorite()(*bool)
    GetJobTitle()(*string)
    GetOfficeLocation()(*string)
    GetPersonNotes()(*string)
    GetPersonType()(PersonTypeable)
    GetPhones()([]Phoneable)
    GetPostalAddresses()([]Locationable)
    GetProfession()(*string)
    GetScoredEmailAddresses()([]ScoredEmailAddressable)
    GetSurname()(*string)
    GetUserPrincipalName()(*string)
    GetWebsites()([]Websiteable)
    GetYomiCompany()(*string)
    SetBirthday(value *string)()
    SetCompanyName(value *string)()
    SetDepartment(value *string)()
    SetDisplayName(value *string)()
    SetGivenName(value *string)()
    SetImAddress(value *string)()
    SetIsFavorite(value *bool)()
    SetJobTitle(value *string)()
    SetOfficeLocation(value *string)()
    SetPersonNotes(value *string)()
    SetPersonType(value PersonTypeable)()
    SetPhones(value []Phoneable)()
    SetPostalAddresses(value []Locationable)()
    SetProfession(value *string)()
    SetScoredEmailAddresses(value []ScoredEmailAddressable)()
    SetSurname(value *string)()
    SetUserPrincipalName(value *string)()
    SetWebsites(value []Websiteable)()
    SetYomiCompany(value *string)()
}
