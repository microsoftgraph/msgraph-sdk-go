package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessApplicationsable 
type ConditionalAccessApplicationsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExcludeApplications()([]string)
    GetIncludeApplications()([]string)
    GetIncludeAuthenticationContextClassReferences()([]string)
    GetIncludeUserActions()([]string)
    SetExcludeApplications(value []string)()
    SetIncludeApplications(value []string)()
    SetIncludeAuthenticationContextClassReferences(value []string)()
    SetIncludeUserActions(value []string)()
}
