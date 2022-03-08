package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsAppable 
type TeamsAppable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppDefinitions()([]TeamsAppDefinitionable)
    GetDisplayName()(*string)
    GetDistributionMethod()(*TeamsAppDistributionMethod)
    GetExternalId()(*string)
    SetAppDefinitions(value []TeamsAppDefinitionable)()
    SetDisplayName(value *string)()
    SetDistributionMethod(value *TeamsAppDistributionMethod)()
    SetExternalId(value *string)()
}
