package batchrecorddecisions

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BatchRecordDecisionsRequestBodyable 
type BatchRecordDecisionsRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDecision()(*string)
    GetJustification()(*string)
    GetPrincipalId()(*string)
    GetResourceId()(*string)
    SetDecision(value *string)()
    SetJustification(value *string)()
    SetPrincipalId(value *string)()
    SetResourceId(value *string)()
}
