package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingParticipantInfoable 
type MeetingParticipantInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetIdentity()(IdentitySetable)
    GetRole()(*OnlineMeetingRole)
    GetUpn()(*string)
    SetIdentity(value IdentitySetable)()
    SetRole(value *OnlineMeetingRole)()
    SetUpn(value *string)()
}
