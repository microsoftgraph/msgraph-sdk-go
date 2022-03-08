package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AlteredQueryTokenable 
type AlteredQueryTokenable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetLength()(*int32)
    GetOffset()(*int32)
    GetSuggestion()(*string)
    SetLength(value *int32)()
    SetOffset(value *int32)()
    SetSuggestion(value *string)()
}
