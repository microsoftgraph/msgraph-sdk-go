package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PublicErrorable 
type PublicErrorable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCode()(*string)
    GetDetails()([]PublicErrorDetailable)
    GetInnerError()(PublicInnerErrorable)
    GetMessage()(*string)
    GetTarget()(*string)
    SetCode(value *string)()
    SetDetails(value []PublicErrorDetailable)()
    SetInnerError(value PublicInnerErrorable)()
    SetMessage(value *string)()
    SetTarget(value *string)()
}
