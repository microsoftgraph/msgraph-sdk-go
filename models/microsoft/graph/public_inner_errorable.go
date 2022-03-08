package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PublicInnerErrorable 
type PublicInnerErrorable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCode()(*string)
    GetDetails()([]PublicErrorDetailable)
    GetMessage()(*string)
    GetTarget()(*string)
    SetCode(value *string)()
    SetDetails(value []PublicErrorDetailable)()
    SetMessage(value *string)()
    SetTarget(value *string)()
}
