package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// B2xIdentityUserFlowable 
type B2xIdentityUserFlowable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    IdentityUserFlowable
    GetApiConnectorConfiguration()(UserFlowApiConnectorConfigurationable)
    GetIdentityProviders()([]IdentityProviderable)
    GetLanguages()([]UserFlowLanguageConfigurationable)
    GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignmentable)
    GetUserFlowIdentityProviders()([]IdentityProviderBaseable)
    SetApiConnectorConfiguration(value UserFlowApiConnectorConfigurationable)()
    SetIdentityProviders(value []IdentityProviderable)()
    SetLanguages(value []UserFlowLanguageConfigurationable)()
    SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignmentable)()
    SetUserFlowIdentityProviders(value []IdentityProviderBaseable)()
}
