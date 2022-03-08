package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharingInvitationable 
type SharingInvitationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetEmail()(*string)
    GetInvitedBy()(IdentitySetable)
    GetRedeemedBy()(*string)
    GetSignInRequired()(*bool)
    SetEmail(value *string)()
    SetInvitedBy(value IdentitySetable)()
    SetRedeemedBy(value *string)()
    SetSignInRequired(value *bool)()
}
