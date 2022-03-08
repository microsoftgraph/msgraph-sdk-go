package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ParentalControlSettingsable 
type ParentalControlSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCountriesBlockedForMinors()([]string)
    GetLegalAgeGroupRule()(*string)
    SetCountriesBlockedForMinors(value []string)()
    SetLegalAgeGroupRule(value *string)()
}
