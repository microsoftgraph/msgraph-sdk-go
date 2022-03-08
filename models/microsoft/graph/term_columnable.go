package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TermColumnable 
type TermColumnable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAllowMultipleValues()(*bool)
    GetParentTerm()(Termable)
    GetShowFullyQualifiedName()(*bool)
    GetTermSet()(Setable)
    SetAllowMultipleValues(value *bool)()
    SetParentTerm(value Termable)()
    SetShowFullyQualifiedName(value *bool)()
    SetTermSet(value Setable)()
}
