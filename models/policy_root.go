package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PolicyRoot provides operations to manage the policyRoot singleton.
type PolicyRoot struct {
    Entity
    // The policy that controls the idle time out for web sessions for applications.
    activityBasedTimeoutPolicies []ActivityBasedTimeoutPolicyable
    // The policy by which consent requests are created and managed for the entire tenant.
    adminConsentRequestPolicy AdminConsentRequestPolicyable
    // The policy configuration of the self-service sign-up experience of external users.
    authenticationFlowsPolicy AuthenticationFlowsPolicyable
    // The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
    authenticationMethodsPolicy AuthenticationMethodsPolicyable
    // The policy that controls Azure AD authorization settings.
    authorizationPolicy AuthorizationPolicyable
    // The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
    claimsMappingPolicies []ClaimsMappingPolicyable
    // The custom rules that define an access scenario.
    conditionalAccessPolicies []ConditionalAccessPolicyable
    // The feature rollout policy associated with a directory object.
    featureRolloutPolicies []FeatureRolloutPolicyable
    // The policy to control Azure AD authentication behavior for federated users.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable
    // The policy that represents the security defaults that protect against common attacks.
    identitySecurityDefaultsEnforcementPolicy IdentitySecurityDefaultsEnforcementPolicyable
    // The policy that specifies the conditions under which consent can be granted.
    permissionGrantPolicies []PermissionGrantPolicyable
    // Represents the role management policies.
    roleManagementPolicies []UnifiedRoleManagementPolicyable
    // Represents the role management policy assignments.
    roleManagementPolicyAssignments []UnifiedRoleManagementPolicyAssignmentable
    // The policy that specifies the characteristics of SAML tokens issued by Azure AD.
    tokenIssuancePolicies []TokenIssuancePolicyable
    // The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
    tokenLifetimePolicies []TokenLifetimePolicyable
}
// NewPolicyRoot instantiates a new policyRoot and sets the default values.
func NewPolicyRoot()(*PolicyRoot) {
    m := &PolicyRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePolicyRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicyRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyRoot(), nil
}
// GetActivityBasedTimeoutPolicies gets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *PolicyRoot) GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.activityBasedTimeoutPolicies
    }
}
// GetAdminConsentRequestPolicy gets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentRequestPolicy
    }
}
// GetAuthenticationFlowsPolicy gets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *PolicyRoot) GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.authenticationFlowsPolicy
    }
}
// GetAuthenticationMethodsPolicy gets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *PolicyRoot) GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsPolicy
    }
}
// GetAuthorizationPolicy gets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *PolicyRoot) GetAuthorizationPolicy()(AuthorizationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.authorizationPolicy
    }
}
// GetClaimsMappingPolicies gets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) GetClaimsMappingPolicies()([]ClaimsMappingPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
// GetConditionalAccessPolicies gets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) GetConditionalAccessPolicies()([]ConditionalAccessPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicies
    }
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PolicyRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityBasedTimeoutPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActivityBasedTimeoutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityBasedTimeoutPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ActivityBasedTimeoutPolicyable)
            }
            m.SetActivityBasedTimeoutPolicies(res)
        }
        return nil
    }
    res["adminConsentRequestPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminConsentRequestPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsentRequestPolicy(val.(AdminConsentRequestPolicyable))
        }
        return nil
    }
    res["authenticationFlowsPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationFlowsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationFlowsPolicy(val.(AuthenticationFlowsPolicyable))
        }
        return nil
    }
    res["authenticationMethodsPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodsPolicy(val.(AuthenticationMethodsPolicyable))
        }
        return nil
    }
    res["authorizationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationPolicy(val.(AuthorizationPolicyable))
        }
        return nil
    }
    res["claimsMappingPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClaimsMappingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClaimsMappingPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ClaimsMappingPolicyable)
            }
            m.SetClaimsMappingPolicies(res)
        }
        return nil
    }
    res["conditionalAccessPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ConditionalAccessPolicyable)
            }
            m.SetConditionalAccessPolicies(res)
        }
        return nil
    }
    res["featureRolloutPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFeatureRolloutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureRolloutPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(FeatureRolloutPolicyable)
            }
            m.SetFeatureRolloutPolicies(res)
        }
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HomeRealmDiscoveryPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(HomeRealmDiscoveryPolicyable)
            }
            m.SetHomeRealmDiscoveryPolicies(res)
        }
        return nil
    }
    res["identitySecurityDefaultsEnforcementPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySecurityDefaultsEnforcementPolicy(val.(IdentitySecurityDefaultsEnforcementPolicyable))
        }
        return nil
    }
    res["permissionGrantPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionGrantPolicyable)
            }
            m.SetPermissionGrantPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementPolicyable)
            }
            m.SetRoleManagementPolicies(res)
        }
        return nil
    }
    res["roleManagementPolicyAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementPolicyAssignmentable)
            }
            m.SetRoleManagementPolicyAssignments(res)
        }
        return nil
    }
    res["tokenIssuancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenIssuancePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TokenIssuancePolicyable)
            }
            m.SetTokenIssuancePolicies(res)
        }
        return nil
    }
    res["tokenLifetimePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenLifetimePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TokenLifetimePolicyable)
            }
            m.SetTokenLifetimePolicies(res)
        }
        return nil
    }
    return res
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *PolicyRoot) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
// GetIdentitySecurityDefaultsEnforcementPolicy gets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.identitySecurityDefaultsEnforcementPolicy
    }
}
// GetPermissionGrantPolicies gets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) GetPermissionGrantPolicies()([]PermissionGrantPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPolicies
    }
}
// GetRoleManagementPolicies gets the roleManagementPolicies property value. Represents the role management policies.
func (m *PolicyRoot) GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.roleManagementPolicies
    }
}
// GetRoleManagementPolicyAssignments gets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *PolicyRoot) GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.roleManagementPolicyAssignments
    }
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *PolicyRoot) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *PolicyRoot) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
// Serialize serializes information the current object
func (m *PolicyRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivityBasedTimeoutPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivityBasedTimeoutPolicies()))
        for i, v := range m.GetActivityBasedTimeoutPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetClaimsMappingPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClaimsMappingPolicies()))
        for i, v := range m.GetClaimsMappingPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("claimsMappingPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConditionalAccessPolicies()))
        for i, v := range m.GetConditionalAccessPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFeatureRolloutPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHomeRealmDiscoveryPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetPermissionGrantPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissionGrantPolicies()))
        for i, v := range m.GetPermissionGrantPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("permissionGrantPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleManagementPolicies()))
        for i, v := range m.GetRoleManagementPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleManagementPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleManagementPolicyAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleManagementPolicyAssignments()))
        for i, v := range m.GetRoleManagementPolicyAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleManagementPolicyAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityBasedTimeoutPolicies sets the activityBasedTimeoutPolicies property value. The policy that controls the idle time out for web sessions for applications.
func (m *PolicyRoot) SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)() {
    if m != nil {
        m.activityBasedTimeoutPolicies = value
    }
}
// SetAdminConsentRequestPolicy sets the adminConsentRequestPolicy property value. The policy by which consent requests are created and managed for the entire tenant.
func (m *PolicyRoot) SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)() {
    if m != nil {
        m.adminConsentRequestPolicy = value
    }
}
// SetAuthenticationFlowsPolicy sets the authenticationFlowsPolicy property value. The policy configuration of the self-service sign-up experience of external users.
func (m *PolicyRoot) SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)() {
    if m != nil {
        m.authenticationFlowsPolicy = value
    }
}
// SetAuthenticationMethodsPolicy sets the authenticationMethodsPolicy property value. The authentication methods and the users that are allowed to use them to sign in and perform multi-factor authentication (MFA) in Azure Active Directory (Azure AD).
func (m *PolicyRoot) SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)() {
    if m != nil {
        m.authenticationMethodsPolicy = value
    }
}
// SetAuthorizationPolicy sets the authorizationPolicy property value. The policy that controls Azure AD authorization settings.
func (m *PolicyRoot) SetAuthorizationPolicy(value AuthorizationPolicyable)() {
    if m != nil {
        m.authorizationPolicy = value
    }
}
// SetClaimsMappingPolicies sets the claimsMappingPolicies property value. The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific application.
func (m *PolicyRoot) SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)() {
    if m != nil {
        m.claimsMappingPolicies = value
    }
}
// SetConditionalAccessPolicies sets the conditionalAccessPolicies property value. The custom rules that define an access scenario.
func (m *PolicyRoot) SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)() {
    if m != nil {
        m.conditionalAccessPolicies = value
    }
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The feature rollout policy associated with a directory object.
func (m *PolicyRoot) SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)() {
    if m != nil {
        m.featureRolloutPolicies = value
    }
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The policy to control Azure AD authentication behavior for federated users.
func (m *PolicyRoot) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    if m != nil {
        m.homeRealmDiscoveryPolicies = value
    }
}
// SetIdentitySecurityDefaultsEnforcementPolicy sets the identitySecurityDefaultsEnforcementPolicy property value. The policy that represents the security defaults that protect against common attacks.
func (m *PolicyRoot) SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)() {
    if m != nil {
        m.identitySecurityDefaultsEnforcementPolicy = value
    }
}
// SetPermissionGrantPolicies sets the permissionGrantPolicies property value. The policy that specifies the conditions under which consent can be granted.
func (m *PolicyRoot) SetPermissionGrantPolicies(value []PermissionGrantPolicyable)() {
    if m != nil {
        m.permissionGrantPolicies = value
    }
}
// SetRoleManagementPolicies sets the roleManagementPolicies property value. Represents the role management policies.
func (m *PolicyRoot) SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)() {
    if m != nil {
        m.roleManagementPolicies = value
    }
}
// SetRoleManagementPolicyAssignments sets the roleManagementPolicyAssignments property value. Represents the role management policy assignments.
func (m *PolicyRoot) SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)() {
    if m != nil {
        m.roleManagementPolicyAssignments = value
    }
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The policy that specifies the characteristics of SAML tokens issued by Azure AD.
func (m *PolicyRoot) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    if m != nil {
        m.tokenIssuancePolicies = value
    }
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Azure AD.
func (m *PolicyRoot) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    if m != nil {
        m.tokenLifetimePolicies = value
    }
}
