package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ParticipantInfoable 
type ParticipantInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCountryCode()(*string)
    GetEndpointType()(*EndpointType)
    GetIdentity()(IdentitySetable)
    GetLanguageId()(*string)
    GetParticipantId()(*string)
    GetRegion()(*string)
    SetCountryCode(value *string)()
    SetEndpointType(value *EndpointType)()
    SetIdentity(value IdentitySetable)()
    SetLanguageId(value *string)()
    SetParticipantId(value *string)()
    SetRegion(value *string)()
}
