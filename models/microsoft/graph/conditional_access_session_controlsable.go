package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessSessionControlsable 
type ConditionalAccessSessionControlsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationEnforcedRestrictions()(ApplicationEnforcedRestrictionsSessionControlable)
    GetCloudAppSecurity()(CloudAppSecuritySessionControlable)
    GetDisableResilienceDefaults()(*bool)
    GetPersistentBrowser()(PersistentBrowserSessionControlable)
    GetSignInFrequency()(SignInFrequencySessionControlable)
    SetApplicationEnforcedRestrictions(value ApplicationEnforcedRestrictionsSessionControlable)()
    SetCloudAppSecurity(value CloudAppSecuritySessionControlable)()
    SetDisableResilienceDefaults(value *bool)()
    SetPersistentBrowser(value PersistentBrowserSessionControlable)()
    SetSignInFrequency(value SignInFrequencySessionControlable)()
}
