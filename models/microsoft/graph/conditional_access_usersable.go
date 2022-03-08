package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessUsersable 
type ConditionalAccessUsersable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExcludeGroups()([]string)
    GetExcludeRoles()([]string)
    GetExcludeUsers()([]string)
    GetIncludeGroups()([]string)
    GetIncludeRoles()([]string)
    GetIncludeUsers()([]string)
    SetExcludeGroups(value []string)()
    SetExcludeRoles(value []string)()
    SetExcludeUsers(value []string)()
    SetIncludeGroups(value []string)()
    SetIncludeRoles(value []string)()
    SetIncludeUsers(value []string)()
}
