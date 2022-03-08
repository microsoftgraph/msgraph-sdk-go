package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LocationConstraintable 
type LocationConstraintable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetIsRequired()(*bool)
    GetLocations()([]LocationConstraintItemable)
    GetSuggestLocation()(*bool)
    SetIsRequired(value *bool)()
    SetLocations(value []LocationConstraintItemable)()
    SetSuggestLocation(value *bool)()
}
