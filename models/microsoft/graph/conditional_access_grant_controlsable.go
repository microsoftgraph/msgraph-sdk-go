package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessGrantControlsable 
type ConditionalAccessGrantControlsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetBuiltInControls()([]ConditionalAccessGrantControl)
    GetCustomAuthenticationFactors()([]string)
    GetOperator()(*string)
    GetTermsOfUse()([]string)
    SetBuiltInControls(value []ConditionalAccessGrantControl)()
    SetCustomAuthenticationFactors(value []string)()
    SetOperator(value *string)()
    SetTermsOfUse(value []string)()
}
