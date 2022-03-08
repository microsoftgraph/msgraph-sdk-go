package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OrgContactable 
type OrgContactable interface {
    DirectoryObjectable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddresses()([]PhysicalOfficeAddressable)
    GetCompanyName()(*string)
    GetDepartment()(*string)
    GetDirectReports()([]DirectoryObjectable)
    GetDisplayName()(*string)
    GetGivenName()(*string)
    GetJobTitle()(*string)
    GetMail()(*string)
    GetMailNickname()(*string)
    GetManager()(DirectoryObjectable)
    GetMemberOf()([]DirectoryObjectable)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable)
    GetOnPremisesSyncEnabled()(*bool)
    GetPhones()([]Phoneable)
    GetProxyAddresses()([]string)
    GetSurname()(*string)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    SetAddresses(value []PhysicalOfficeAddressable)()
    SetCompanyName(value *string)()
    SetDepartment(value *string)()
    SetDirectReports(value []DirectoryObjectable)()
    SetDisplayName(value *string)()
    SetGivenName(value *string)()
    SetJobTitle(value *string)()
    SetMail(value *string)()
    SetMailNickname(value *string)()
    SetManager(value DirectoryObjectable)()
    SetMemberOf(value []DirectoryObjectable)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetPhones(value []Phoneable)()
    SetProxyAddresses(value []string)()
    SetSurname(value *string)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
}
