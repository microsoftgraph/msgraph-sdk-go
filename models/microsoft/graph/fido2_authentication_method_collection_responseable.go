package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Fido2AuthenticationMethodCollectionResponseable 
type Fido2AuthenticationMethodCollectionResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetNextLink()(*string)
    GetValue()([]Fido2AuthenticationMethodable)
    SetNextLink(value *string)()
    SetValue(value []Fido2AuthenticationMethodable)()
}
