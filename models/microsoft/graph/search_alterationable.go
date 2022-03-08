package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchAlterationable 
type SearchAlterationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAlteredHighlightedQueryString()(*string)
    GetAlteredQueryString()(*string)
    GetAlteredQueryTokens()([]AlteredQueryTokenable)
    SetAlteredHighlightedQueryString(value *string)()
    SetAlteredQueryString(value *string)()
    SetAlteredQueryTokens(value []AlteredQueryTokenable)()
}
