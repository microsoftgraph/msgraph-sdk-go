package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VerifiedDomainable 
type VerifiedDomainable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCapabilities()(*string)
    GetIsDefault()(*bool)
    GetIsInitial()(*bool)
    GetName()(*string)
    GetType()(*string)
    SetCapabilities(value *string)()
    SetIsDefault(value *bool)()
    SetIsInitial(value *bool)()
    SetName(value *string)()
    SetType(value *string)()
}
