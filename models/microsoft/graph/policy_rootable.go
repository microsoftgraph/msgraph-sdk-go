package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PolicyRootable 
type PolicyRootable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable)
    GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable)
    GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable)
    GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable)
    GetAuthorizationPolicy()(AuthorizationPolicyable)
    GetClaimsMappingPolicies()([]ClaimsMappingPolicyable)
    GetConditionalAccessPolicies()([]ConditionalAccessPolicyable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable)
    GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable)
    GetPermissionGrantPolicies()([]PermissionGrantPolicyable)
    GetTokenIssuancePolicies()([]TokenIssuancePolicyable)
    GetTokenLifetimePolicies()([]TokenLifetimePolicyable)
    SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)()
    SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)()
    SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)()
    SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)()
    SetAuthorizationPolicy(value AuthorizationPolicyable)()
    SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)()
    SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)()
    SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)()
    SetPermissionGrantPolicies(value []PermissionGrantPolicyable)()
    SetTokenIssuancePolicies(value []TokenIssuancePolicyable)()
    SetTokenLifetimePolicies(value []TokenLifetimePolicyable)()
}
