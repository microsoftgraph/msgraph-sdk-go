package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsTabable 
type TeamsTabable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfiguration()(TeamsTabConfigurationable)
    GetDisplayName()(*string)
    GetTeamsApp()(TeamsAppable)
    GetWebUrl()(*string)
    SetConfiguration(value TeamsTabConfigurationable)()
    SetDisplayName(value *string)()
    SetTeamsApp(value TeamsAppable)()
    SetWebUrl(value *string)()
}
