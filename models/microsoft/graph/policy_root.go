package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PolicyRoot struct {
    Entity
    // The policy that controls the idle time out for web sessions for applications.
    activityBasedTimeoutPolicies []ActivityBasedTimeoutPolicy;
    // The policy by which consent requests are created and managed for the entire tenant.
    adminConsentRequestPolicy *AdminConsentRequestPolicy;
    // The policy configuration of the self-service sign-up experience of external users.
    authenticationFlowsPolicy *AuthenticationFlowsPolicy;
    // The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
    authenticationMethodsPolicy *AuthenticationMethodsPolicy;
    // The policy that controls Azure AD authorization settings.
    authorizationPolicy *AuthorizationPolicy;
    // The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
    claimsMappingPolicies []ClaimsMappingPolicy;
    // The custom rules that define an access scenario.
    conditionalAccessPolicies []ConditionalAccessPolicy;
    // The feature rollout policy associated with a directory object.
    featureRolloutPolicies []FeatureRolloutPolicy;
    // The policy to control Azure AD authentication behavior for federated users.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy;
    // The policy that represents the security defaults that protect against common attacks.
    identitySecurityDefaultsEnforcementPolicy *IdentitySecurityDefaultsEnforcementPolicy;
    // The policy that specifies the conditions under which consent can be granted.
    permissionGrantPolicies []PermissionGrantPolicy;
    // The policy that specifies the characteristics of SAML tokens issued by Azure AD.
    tokenIssuancePolicies []TokenIssuancePolicy;
    // The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
    tokenLifetimePolicies []TokenLifetimePolicy;
}
// Instantiates a new policyRoot and sets the default values.
func NewPolicyRoot()(*PolicyRoot) {
    m := &PolicyRoot{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *PolicyRoot) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.activityBasedTimeoutPolicies
    }
}
// Gets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) GetAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentRequestPolicy
    }
}
// Gets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *PolicyRoot) GetAuthenticationFlowsPolicy()(*AuthenticationFlowsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationFlowsPolicy
    }
}
// Gets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *PolicyRoot) GetAuthenticationMethodsPolicy()(*AuthenticationMethodsPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsPolicy
    }
}
// Gets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *PolicyRoot) GetAuthorizationPolicy()(*AuthorizationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.authorizationPolicy
    }
}
// Gets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) GetClaimsMappingPolicies()([]ClaimsMappingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
// Gets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) GetConditionalAccessPolicies()([]ConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicies
    }
}
// Gets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) GetFeatureRolloutPolicies()([]FeatureRolloutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
// Gets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *PolicyRoot) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicy) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
// Gets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicy) {
    if m == nil {
        return nil
    } else {
        return m.identitySecurityDefaultsEnforcementPolicy
    }
}
// Gets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) GetPermissionGrantPolicies()([]PermissionGrantPolicy) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPolicies
    }
}
// Gets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *PolicyRoot) GetTokenIssuancePolicies()([]TokenIssuancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
// Gets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *PolicyRoot) GetTokenLifetimePolicies()([]TokenLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
// The deserialization information for the current model
func (m *PolicyRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityBasedTimeoutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActivityBasedTimeoutPolicy() })
        if err != nil {
            return err
        }
        res := make([]ActivityBasedTimeoutPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ActivityBasedTimeoutPolicy))
        }
        m.SetActivityBasedTimeoutPolicies(res)
        return nil
    }
    res["adminConsentRequestPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdminConsentRequestPolicy() })
        if err != nil {
            return err
        }
        m.SetAdminConsentRequestPolicy(val.(*AdminConsentRequestPolicy))
        return nil
    }
    res["authenticationFlowsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationFlowsPolicy() })
        if err != nil {
            return err
        }
        m.SetAuthenticationFlowsPolicy(val.(*AuthenticationFlowsPolicy))
        return nil
    }
    res["authenticationMethodsPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsPolicy() })
        if err != nil {
            return err
        }
        m.SetAuthenticationMethodsPolicy(val.(*AuthenticationMethodsPolicy))
        return nil
    }
    res["authorizationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthorizationPolicy() })
        if err != nil {
            return err
        }
        m.SetAuthorizationPolicy(val.(*AuthorizationPolicy))
        return nil
    }
    res["claimsMappingPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClaimsMappingPolicy() })
        if err != nil {
            return err
        }
        res := make([]ClaimsMappingPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ClaimsMappingPolicy))
        }
        m.SetClaimsMappingPolicies(res)
        return nil
    }
    res["conditionalAccessPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPolicy() })
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessPolicy))
        }
        m.SetConditionalAccessPolicies(res)
        return nil
    }
    res["featureRolloutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFeatureRolloutPolicy() })
        if err != nil {
            return err
        }
        res := make([]FeatureRolloutPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*FeatureRolloutPolicy))
        }
        m.SetFeatureRolloutPolicies(res)
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHomeRealmDiscoveryPolicy() })
        if err != nil {
            return err
        }
        res := make([]HomeRealmDiscoveryPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*HomeRealmDiscoveryPolicy))
        }
        m.SetHomeRealmDiscoveryPolicies(res)
        return nil
    }
    res["identitySecurityDefaultsEnforcementPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySecurityDefaultsEnforcementPolicy() })
        if err != nil {
            return err
        }
        m.SetIdentitySecurityDefaultsEnforcementPolicy(val.(*IdentitySecurityDefaultsEnforcementPolicy))
        return nil
    }
    res["permissionGrantPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionGrantPolicy() })
        if err != nil {
            return err
        }
        res := make([]PermissionGrantPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*PermissionGrantPolicy))
        }
        m.SetPermissionGrantPolicies(res)
        return nil
    }
    res["tokenIssuancePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenIssuancePolicy() })
        if err != nil {
            return err
        }
        res := make([]TokenIssuancePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*TokenIssuancePolicy))
        }
        m.SetTokenIssuancePolicies(res)
        return nil
    }
    res["tokenLifetimePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTokenLifetimePolicy() })
        if err != nil {
            return err
        }
        res := make([]TokenLifetimePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*TokenLifetimePolicy))
        }
        m.SetTokenLifetimePolicies(res)
        return nil
    }
    return res
}
func (m *PolicyRoot) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PolicyRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivityBasedTimeoutPolicies()))
        for i, v := range m.GetActivityBasedTimeoutPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("activityBasedTimeoutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("adminConsentRequestPolicy", m.GetAdminConsentRequestPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationFlowsPolicy", m.GetAuthenticationFlowsPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationMethodsPolicy", m.GetAuthenticationMethodsPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizationPolicy", m.GetAuthorizationPolicy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClaimsMappingPolicies()))
        for i, v := range m.GetClaimsMappingPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConditionalAccessPolicies()))
        for i, v := range m.GetConditionalAccessPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("homeRealmDiscoveryPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identitySecurityDefaultsEnforcementPolicy", m.GetIdentitySecurityDefaultsEnforcementPolicy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissionGrantPolicies()))
        for i, v := range m.GetPermissionGrantPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("permissionGrantPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
// Parameters:
//  - value : Value to set for the activityBasedTimeoutPolicies property.
func (m *PolicyRoot) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicy)() {
    m.activityBasedTimeoutPolicies = value
}
// Sets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
// Parameters:
//  - value : Value to set for the adminConsentRequestPolicy property.
func (m *PolicyRoot) SetAdminConsentRequestPolicy(value *AdminConsentRequestPolicy)() {
    m.adminConsentRequestPolicy = value
}
// Sets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
// Parameters:
//  - value : Value to set for the authenticationFlowsPolicy property.
func (m *PolicyRoot) SetAuthenticationFlowsPolicy(value *AuthenticationFlowsPolicy)() {
    m.authenticationFlowsPolicy = value
}
// Sets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
// Parameters:
//  - value : Value to set for the authenticationMethodsPolicy property.
func (m *PolicyRoot) SetAuthenticationMethodsPolicy(value *AuthenticationMethodsPolicy)() {
    m.authenticationMethodsPolicy = value
}
// Sets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
// Parameters:
//  - value : Value to set for the authorizationPolicy property.
func (m *PolicyRoot) SetAuthorizationPolicy(value *AuthorizationPolicy)() {
    m.authorizationPolicy = value
}
// Sets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
// Parameters:
//  - value : Value to set for the claimsMappingPolicies property.
func (m *PolicyRoot) SetClaimsMappingPolicies(value []ClaimsMappingPolicy)() {
    m.claimsMappingPolicies = value
}
// Sets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
// Parameters:
//  - value : Value to set for the conditionalAccessPolicies property.
func (m *PolicyRoot) SetConditionalAccessPolicies(value []ConditionalAccessPolicy)() {
    m.conditionalAccessPolicies = value
}
// Sets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
// Parameters:
//  - value : Value to set for the featureRolloutPolicies property.
func (m *PolicyRoot) SetFeatureRolloutPolicies(value []FeatureRolloutPolicy)() {
    m.featureRolloutPolicies = value
}
// Sets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
// Parameters:
//  - value : Value to set for the homeRealmDiscoveryPolicies property.
func (m *PolicyRoot) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicy)() {
    m.homeRealmDiscoveryPolicies = value
}
// Sets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
// Parameters:
//  - value : Value to set for the identitySecurityDefaultsEnforcementPolicy property.
func (m *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(value *IdentitySecurityDefaultsEnforcementPolicy)() {
    m.identitySecurityDefaultsEnforcementPolicy = value
}
// Sets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
// Parameters:
//  - value : Value to set for the permissionGrantPolicies property.
func (m *PolicyRoot) SetPermissionGrantPolicies(value []PermissionGrantPolicy)() {
    m.permissionGrantPolicies = value
}
// Sets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
// Parameters:
//  - value : Value to set for the tokenIssuancePolicies property.
func (m *PolicyRoot) SetTokenIssuancePolicies(value []TokenIssuancePolicy)() {
    m.tokenIssuancePolicies = value
}
// Sets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
// Parameters:
//  - value : Value to set for the tokenLifetimePolicies property.
func (m *PolicyRoot) SetTokenLifetimePolicies(value []TokenLifetimePolicy)() {
    m.tokenLifetimePolicies = value
}
