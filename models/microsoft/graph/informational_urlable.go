package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationalUrlable 
type InformationalUrlable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetLogoUrl()(*string)
    GetMarketingUrl()(*string)
    GetPrivacyStatementUrl()(*string)
    GetSupportUrl()(*string)
    GetTermsOfServiceUrl()(*string)
    SetLogoUrl(value *string)()
    SetMarketingUrl(value *string)()
    SetPrivacyStatementUrl(value *string)()
    SetSupportUrl(value *string)()
    SetTermsOfServiceUrl(value *string)()
}
