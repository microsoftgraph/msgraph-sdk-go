package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IncomingContextable 
type IncomingContextable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetObservedParticipantId()(*string)
    GetOnBehalfOf()(IdentitySetable)
    GetSourceParticipantId()(*string)
    GetTransferor()(IdentitySetable)
    SetObservedParticipantId(value *string)()
    SetOnBehalfOf(value IdentitySetable)()
    SetSourceParticipantId(value *string)()
    SetTransferor(value IdentitySetable)()
}
