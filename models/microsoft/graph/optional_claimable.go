package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OptionalClaimable 
type OptionalClaimable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalProperties()([]string)
    GetEssential()(*bool)
    GetName()(*string)
    GetSource()(*string)
    SetAdditionalProperties(value []string)()
    SetEssential(value *bool)()
    SetName(value *string)()
    SetSource(value *string)()
}
