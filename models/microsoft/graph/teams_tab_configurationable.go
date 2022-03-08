package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsTabConfigurationable 
type TeamsTabConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContentUrl()(*string)
    GetEntityId()(*string)
    GetRemoveUrl()(*string)
    GetWebsiteUrl()(*string)
    SetContentUrl(value *string)()
    SetEntityId(value *string)()
    SetRemoveUrl(value *string)()
    SetWebsiteUrl(value *string)()
}
