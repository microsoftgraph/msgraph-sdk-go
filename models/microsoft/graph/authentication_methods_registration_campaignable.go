package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthenticationMethodsRegistrationCampaignable 
type AuthenticationMethodsRegistrationCampaignable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetExcludeTargets()([]ExcludeTargetable)
    GetIncludeTargets()([]AuthenticationMethodsRegistrationCampaignIncludeTargetable)
    GetSnoozeDurationInDays()(*int32)
    GetState()(*AdvancedConfigState)
    SetExcludeTargets(value []ExcludeTargetable)()
    SetIncludeTargets(value []AuthenticationMethodsRegistrationCampaignIncludeTargetable)()
    SetSnoozeDurationInDays(value *int32)()
    SetState(value *AdvancedConfigState)()
}
