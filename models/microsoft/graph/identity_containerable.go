package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityContainerable 
type IdentityContainerable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApiConnectors()([]IdentityApiConnectorable)
    GetB2xUserFlows()([]B2xIdentityUserFlowable)
    GetConditionalAccess()(ConditionalAccessRootable)
    GetIdentityProviders()([]IdentityProviderBaseable)
    GetUserFlowAttributes()([]IdentityUserFlowAttributeable)
    SetApiConnectors(value []IdentityApiConnectorable)()
    SetB2xUserFlows(value []B2xIdentityUserFlowable)()
    SetConditionalAccess(value ConditionalAccessRootable)()
    SetIdentityProviders(value []IdentityProviderBaseable)()
    SetUserFlowAttributes(value []IdentityUserFlowAttributeable)()
}
