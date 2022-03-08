package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharePointIdentitySetable 
type SharePointIdentitySetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    IdentitySetable
    GetGroup()(Identityable)
    GetSiteGroup()(SharePointIdentityable)
    GetSiteUser()(SharePointIdentityable)
    SetGroup(value Identityable)()
    SetSiteGroup(value SharePointIdentityable)()
    SetSiteUser(value SharePointIdentityable)()
}
