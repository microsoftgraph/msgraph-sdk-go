package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ScoredEmailAddressable 
type ScoredEmailAddressable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddress()(*string)
    GetItemId()(*string)
    GetRelevanceScore()(*float64)
    GetSelectionLikelihood()(*SelectionLikelihoodInfo)
    SetAddress(value *string)()
    SetItemId(value *string)()
    SetRelevanceScore(value *float64)()
    SetSelectionLikelihood(value *SelectionLikelihoodInfo)()
}
