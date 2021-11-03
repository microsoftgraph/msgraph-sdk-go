package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ServicePrincipal struct {
    DirectoryObject
    // true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, NOT, in).
    accountEnabled *bool;
    // Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
    addIns []AddIn;
    // Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, NOT, ge, le, startsWith).
    alternativeNames []string;
    // The description exposed by the associated application.
    appDescription *string;
    // The display name exposed by the associated application.
    appDisplayName *string;
    // The unique identifier for the associated application (its appId property).
    appId *string;
    // Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
    applicationTemplateId *string;
    // Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
    appOwnerOrganizationId *string;
    // App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
    appRoleAssignedTo []AppRoleAssignment;
    // Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
    appRoleAssignmentRequired *bool;
    // App role assignment for another app or service, granted to this service principal. Supports $expand.
    appRoleAssignments []AppRoleAssignment;
    // The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
    appRoles []AppRole;
    // The claimsMappingPolicies assigned to this service principal. Supports $expand.
    claimsMappingPolicies []ClaimsMappingPolicy;
    // Directory objects created by this service principal. Read-only. Nullable.
    createdObjects []DirectoryObject;
    // The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
    delegatedPermissionClassifications []DelegatedPermissionClassification;
    // Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
    description *string;
    // Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, NOT).
    disabledByMicrosoftStatus *string;
    // The display name for the service principal. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
    displayName *string;
    // Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
    endpoints []Endpoint;
    // Home page or landing page of the application.
    homepage *string;
    // The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy;
    // Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, NOT, ge, le).
    info *InformationalUrl;
    // The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, NOT, ge, le).
    keyCredentials []KeyCredential;
    // Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
    loginUrl *string;
    // Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
    logoutUrl *string;
    // Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObject;
    // Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
    notes *string;
    // Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
    notificationEmailAddresses []string;
    // Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
    oauth2PermissionGrants []OAuth2PermissionGrant;
    // The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable.
    oauth2PermissionScopes []PermissionScope;
    // Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
    ownedObjects []DirectoryObject;
    // Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
    owners []DirectoryObject;
    // The collection of password credentials associated with the application. Not nullable.
    passwordCredentials []PasswordCredential;
    // Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
    preferredSingleSignOnMode *string;
    // Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
    preferredTokenSigningKeyThumbprint *string;
    // The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
    replyUrls []string;
    // The collection for settings related to saml single sign-on.
    samlSingleSignOnSettings *SamlSingleSignOnSettings;
    // Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, NOT, ge, le, startsWith).
    servicePrincipalNames []string;
    // Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally. The servicePrincipalType property can be set to three different values: __Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens are not issued for the service principal.__ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but cannot be updated or modified directly.__Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. Legacy service principal can have credentials, service principal names, reply URLs, and other properties which are editable by an authorized user, but does not have an associated app registration. The appId value does not associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.__SocialIdp - For internal use.
    servicePrincipalType *string;
    // Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
    signInAudience *string;
    // Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
    tags []string;
    // Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
    tokenEncryptionKeyId *string;
    // The tokenIssuancePolicies assigned to this service principal.
    tokenIssuancePolicies []TokenIssuancePolicy;
    // The tokenLifetimePolicies assigned to this service principal.
    tokenLifetimePolicies []TokenLifetimePolicy;
    // 
    transitiveMemberOf []DirectoryObject;
}
// Instantiates a new servicePrincipal and sets the default values.
func NewServicePrincipal()(*ServicePrincipal) {
    m := &ServicePrincipal{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the accountEnabled property value. true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, NOT, in).
func (m *ServicePrincipal) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
// Gets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
func (m *ServicePrincipal) GetAddIns()([]AddIn) {
    if m == nil {
        return nil
    } else {
        return m.addIns
    }
}
// Gets the alternativeNames property value. Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *ServicePrincipal) GetAlternativeNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.alternativeNames
    }
}
// Gets the appDescription property value. The description exposed by the associated application.
func (m *ServicePrincipal) GetAppDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDescription
    }
}
// Gets the appDisplayName property value. The display name exposed by the associated application.
func (m *ServicePrincipal) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appId property value. The unique identifier for the associated application (its appId property).
func (m *ServicePrincipal) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the applicationTemplateId property value. Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
func (m *ServicePrincipal) GetApplicationTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationTemplateId
    }
}
// Gets the appOwnerOrganizationId property value. Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
func (m *ServicePrincipal) GetAppOwnerOrganizationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appOwnerOrganizationId
    }
}
// Gets the appRoleAssignedTo property value. App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
func (m *ServicePrincipal) GetAppRoleAssignedTo()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignedTo
    }
}
// Gets the appRoleAssignmentRequired property value. Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
func (m *ServicePrincipal) GetAppRoleAssignmentRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignmentRequired
    }
}
// Gets the appRoleAssignments property value. App role assignment for another app or service, granted to this service principal. Supports $expand.
func (m *ServicePrincipal) GetAppRoleAssignments()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignments
    }
}
// Gets the appRoles property value. The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
func (m *ServicePrincipal) GetAppRoles()([]AppRole) {
    if m == nil {
        return nil
    } else {
        return m.appRoles
    }
}
// Gets the claimsMappingPolicies property value. The claimsMappingPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) GetClaimsMappingPolicies()([]ClaimsMappingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
// Gets the createdObjects property value. Directory objects created by this service principal. Read-only. Nullable.
func (m *ServicePrincipal) GetCreatedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.createdObjects
    }
}
// Gets the delegatedPermissionClassifications property value. The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
func (m *ServicePrincipal) GetDelegatedPermissionClassifications()([]DelegatedPermissionClassification) {
    if m == nil {
        return nil
    } else {
        return m.delegatedPermissionClassifications
    }
}
// Gets the description property value. Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
func (m *ServicePrincipal) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, NOT).
func (m *ServicePrincipal) GetDisabledByMicrosoftStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledByMicrosoftStatus
    }
}
// Gets the displayName property value. The display name for the service principal. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
func (m *ServicePrincipal) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endpoints property value. Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
func (m *ServicePrincipal) GetEndpoints()([]Endpoint) {
    if m == nil {
        return nil
    } else {
        return m.endpoints
    }
}
// Gets the homepage property value. Home page or landing page of the application.
func (m *ServicePrincipal) GetHomepage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homepage
    }
}
// Gets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
func (m *ServicePrincipal) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicy) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
// Gets the info property value. Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, NOT, ge, le).
func (m *ServicePrincipal) GetInfo()(*InformationalUrl) {
    if m == nil {
        return nil
    } else {
        return m.info
    }
}
// Gets the keyCredentials property value. The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, NOT, ge, le).
func (m *ServicePrincipal) GetKeyCredentials()([]KeyCredential) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentials
    }
}
// Gets the loginUrl property value. Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
func (m *ServicePrincipal) GetLoginUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loginUrl
    }
}
// Gets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
func (m *ServicePrincipal) GetLogoutUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoutUrl
    }
}
// Gets the memberOf property value. Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// Gets the notes property value. Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
func (m *ServicePrincipal) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the notificationEmailAddresses property value. Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
func (m *ServicePrincipal) GetNotificationEmailAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notificationEmailAddresses
    }
}
// Gets the oauth2PermissionGrants property value. Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
func (m *ServicePrincipal) GetOauth2PermissionGrants()([]OAuth2PermissionGrant) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionGrants
    }
}
// Gets the oauth2PermissionScopes property value. The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable.
func (m *ServicePrincipal) GetOauth2PermissionScopes()([]PermissionScope) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionScopes
    }
}
// Gets the ownedObjects property value. Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetOwnedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.ownedObjects
    }
}
// Gets the owners property value. Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
func (m *ServicePrincipal) GetOwners()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.owners
    }
}
// Gets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *ServicePrincipal) GetPasswordCredentials()([]PasswordCredential) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentials
    }
}
// Gets the preferredSingleSignOnMode property value. Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
func (m *ServicePrincipal) GetPreferredSingleSignOnMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredSingleSignOnMode
    }
}
// Gets the preferredTokenSigningKeyThumbprint property value. Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
func (m *ServicePrincipal) GetPreferredTokenSigningKeyThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredTokenSigningKeyThumbprint
    }
}
// Gets the replyUrls property value. The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
func (m *ServicePrincipal) GetReplyUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.replyUrls
    }
}
// Gets the samlSingleSignOnSettings property value. The collection for settings related to saml single sign-on.
func (m *ServicePrincipal) GetSamlSingleSignOnSettings()(*SamlSingleSignOnSettings) {
    if m == nil {
        return nil
    } else {
        return m.samlSingleSignOnSettings
    }
}
// Gets the servicePrincipalNames property value. Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, NOT, ge, le, startsWith).
func (m *ServicePrincipal) GetServicePrincipalNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalNames
    }
}
// Gets the servicePrincipalType property value. Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally. The servicePrincipalType property can be set to three different values: __Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens are not issued for the service principal.__ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but cannot be updated or modified directly.__Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. Legacy service principal can have credentials, service principal names, reply URLs, and other properties which are editable by an authorized user, but does not have an associated app registration. The appId value does not associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.__SocialIdp - For internal use.
func (m *ServicePrincipal) GetServicePrincipalType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalType
    }
}
// Gets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
func (m *ServicePrincipal) GetSignInAudience()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInAudience
    }
}
// Gets the tags property value. Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *ServicePrincipal) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Gets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *ServicePrincipal) GetTokenEncryptionKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenEncryptionKeyId
    }
}
// Gets the tokenIssuancePolicies property value. The tokenIssuancePolicies assigned to this service principal.
func (m *ServicePrincipal) GetTokenIssuancePolicies()([]TokenIssuancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
// Gets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this service principal.
func (m *ServicePrincipal) GetTokenLifetimePolicies()([]TokenLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
// Gets the transitiveMemberOf property value. 
func (m *ServicePrincipal) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// The deserialization information for the current model
func (m *ServicePrincipal) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAccountEnabled(val)
        return nil
    }
    res["addIns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAddIn() })
        if err != nil {
            return err
        }
        res := make([]AddIn, len(val))
        for i, v := range val {
            res[i] = *(v.(*AddIn))
        }
        m.SetAddIns(res)
        return nil
    }
    res["alternativeNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetAlternativeNames(res)
        return nil
    }
    res["appDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDescription(val)
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["applicationTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationTemplateId(val)
        return nil
    }
    res["appOwnerOrganizationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppOwnerOrganizationId(val)
        return nil
    }
    res["appRoleAssignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRoleAssignment() })
        if err != nil {
            return err
        }
        res := make([]AppRoleAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppRoleAssignment))
        }
        m.SetAppRoleAssignedTo(res)
        return nil
    }
    res["appRoleAssignmentRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAppRoleAssignmentRequired(val)
        return nil
    }
    res["appRoleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRoleAssignment() })
        if err != nil {
            return err
        }
        res := make([]AppRoleAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppRoleAssignment))
        }
        m.SetAppRoleAssignments(res)
        return nil
    }
    res["appRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRole() })
        if err != nil {
            return err
        }
        res := make([]AppRole, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppRole))
        }
        m.SetAppRoles(res)
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
    res["createdObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetCreatedObjects(res)
        return nil
    }
    res["delegatedPermissionClassifications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedPermissionClassification() })
        if err != nil {
            return err
        }
        res := make([]DelegatedPermissionClassification, len(val))
        for i, v := range val {
            res[i] = *(v.(*DelegatedPermissionClassification))
        }
        m.SetDelegatedPermissionClassifications(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["disabledByMicrosoftStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisabledByMicrosoftStatus(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["endpoints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEndpoint() })
        if err != nil {
            return err
        }
        res := make([]Endpoint, len(val))
        for i, v := range val {
            res[i] = *(v.(*Endpoint))
        }
        m.SetEndpoints(res)
        return nil
    }
    res["homepage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHomepage(val)
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
    res["info"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInformationalUrl() })
        if err != nil {
            return err
        }
        m.SetInfo(val.(*InformationalUrl))
        return nil
    }
    res["keyCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyCredential() })
        if err != nil {
            return err
        }
        res := make([]KeyCredential, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyCredential))
        }
        m.SetKeyCredentials(res)
        return nil
    }
    res["loginUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLoginUrl(val)
        return nil
    }
    res["logoutUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLogoutUrl(val)
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetMemberOf(res)
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotes(val)
        return nil
    }
    res["notificationEmailAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetNotificationEmailAddresses(res)
        return nil
    }
    res["oauth2PermissionGrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOAuth2PermissionGrant() })
        if err != nil {
            return err
        }
        res := make([]OAuth2PermissionGrant, len(val))
        for i, v := range val {
            res[i] = *(v.(*OAuth2PermissionGrant))
        }
        m.SetOauth2PermissionGrants(res)
        return nil
    }
    res["oauth2PermissionScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionScope() })
        if err != nil {
            return err
        }
        res := make([]PermissionScope, len(val))
        for i, v := range val {
            res[i] = *(v.(*PermissionScope))
        }
        m.SetOauth2PermissionScopes(res)
        return nil
    }
    res["ownedObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetOwnedObjects(res)
        return nil
    }
    res["owners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetOwners(res)
        return nil
    }
    res["passwordCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordCredential() })
        if err != nil {
            return err
        }
        res := make([]PasswordCredential, len(val))
        for i, v := range val {
            res[i] = *(v.(*PasswordCredential))
        }
        m.SetPasswordCredentials(res)
        return nil
    }
    res["preferredSingleSignOnMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreferredSingleSignOnMode(val)
        return nil
    }
    res["preferredTokenSigningKeyThumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreferredTokenSigningKeyThumbprint(val)
        return nil
    }
    res["replyUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetReplyUrls(res)
        return nil
    }
    res["samlSingleSignOnSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSamlSingleSignOnSettings() })
        if err != nil {
            return err
        }
        m.SetSamlSingleSignOnSettings(val.(*SamlSingleSignOnSettings))
        return nil
    }
    res["servicePrincipalNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetServicePrincipalNames(res)
        return nil
    }
    res["servicePrincipalType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePrincipalType(val)
        return nil
    }
    res["signInAudience"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSignInAudience(val)
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetTags(res)
        return nil
    }
    res["tokenEncryptionKeyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTokenEncryptionKeyId(val)
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
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetTransitiveMemberOf(res)
        return nil
    }
    return res
}
func (m *ServicePrincipal) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ServicePrincipal) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAddIns()))
        for i, v := range m.GetAddIns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("addIns", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("alternativeNames", m.GetAlternativeNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDescription", m.GetAppDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applicationTemplateId", m.GetApplicationTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appOwnerOrganizationId", m.GetAppOwnerOrganizationId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoleAssignedTo()))
        for i, v := range m.GetAppRoleAssignedTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appRoleAssignedTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appRoleAssignmentRequired", m.GetAppRoleAssignmentRequired())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoleAssignments()))
        for i, v := range m.GetAppRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoles()))
        for i, v := range m.GetAppRoles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appRoles", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCreatedObjects()))
        for i, v := range m.GetCreatedObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("createdObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDelegatedPermissionClassifications()))
        for i, v := range m.GetDelegatedPermissionClassifications() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("delegatedPermissionClassifications", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("disabledByMicrosoftStatus", m.GetDisabledByMicrosoftStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEndpoints()))
        for i, v := range m.GetEndpoints() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("endpoints", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homepage", m.GetHomepage())
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
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loginUrl", m.GetLoginUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("logoutUrl", m.GetLogoutUrl())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("notificationEmailAddresses", m.GetNotificationEmailAddresses())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOauth2PermissionGrants()))
        for i, v := range m.GetOauth2PermissionGrants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("oauth2PermissionGrants", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOauth2PermissionScopes()))
        for i, v := range m.GetOauth2PermissionScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("oauth2PermissionScopes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwnedObjects()))
        for i, v := range m.GetOwnedObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ownedObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredSingleSignOnMode", m.GetPreferredSingleSignOnMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredTokenSigningKeyThumbprint", m.GetPreferredTokenSigningKeyThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("replyUrls", m.GetReplyUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("samlSingleSignOnSettings", m.GetSamlSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("servicePrincipalNames", m.GetServicePrincipalNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalType", m.GetServicePrincipalType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInAudience", m.GetSignInAudience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenEncryptionKeyId", m.GetTokenEncryptionKeyId())
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accountEnabled property value. true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the accountEnabled property.
func (m *ServicePrincipal) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// Sets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on.
// Parameters:
//  - value : Value to set for the addIns property.
func (m *ServicePrincipal) SetAddIns(value []AddIn)() {
    m.addIns = value
}
// Sets the alternativeNames property value. Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the alternativeNames property.
func (m *ServicePrincipal) SetAlternativeNames(value []string)() {
    m.alternativeNames = value
}
// Sets the appDescription property value. The description exposed by the associated application.
// Parameters:
//  - value : Value to set for the appDescription property.
func (m *ServicePrincipal) SetAppDescription(value *string)() {
    m.appDescription = value
}
// Sets the appDisplayName property value. The display name exposed by the associated application.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *ServicePrincipal) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appId property value. The unique identifier for the associated application (its appId property).
// Parameters:
//  - value : Value to set for the appId property.
func (m *ServicePrincipal) SetAppId(value *string)() {
    m.appId = value
}
// Sets the applicationTemplateId property value. Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith).
// Parameters:
//  - value : Value to set for the applicationTemplateId property.
func (m *ServicePrincipal) SetApplicationTemplateId(value *string)() {
    m.applicationTemplateId = value
}
// Sets the appOwnerOrganizationId property value. Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le).
// Parameters:
//  - value : Value to set for the appOwnerOrganizationId property.
func (m *ServicePrincipal) SetAppOwnerOrganizationId(value *string)() {
    m.appOwnerOrganizationId = value
}
// Sets the appRoleAssignedTo property value. App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand.
// Parameters:
//  - value : Value to set for the appRoleAssignedTo property.
func (m *ServicePrincipal) SetAppRoleAssignedTo(value []AppRoleAssignment)() {
    m.appRoleAssignedTo = value
}
// Sets the appRoleAssignmentRequired property value. Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT).
// Parameters:
//  - value : Value to set for the appRoleAssignmentRequired property.
func (m *ServicePrincipal) SetAppRoleAssignmentRequired(value *bool)() {
    m.appRoleAssignmentRequired = value
}
// Sets the appRoleAssignments property value. App role assignment for another app or service, granted to this service principal. Supports $expand.
// Parameters:
//  - value : Value to set for the appRoleAssignments property.
func (m *ServicePrincipal) SetAppRoleAssignments(value []AppRoleAssignment)() {
    m.appRoleAssignments = value
}
// Sets the appRoles property value. The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable.
// Parameters:
//  - value : Value to set for the appRoles property.
func (m *ServicePrincipal) SetAppRoles(value []AppRole)() {
    m.appRoles = value
}
// Sets the claimsMappingPolicies property value. The claimsMappingPolicies assigned to this service principal. Supports $expand.
// Parameters:
//  - value : Value to set for the claimsMappingPolicies property.
func (m *ServicePrincipal) SetClaimsMappingPolicies(value []ClaimsMappingPolicy)() {
    m.claimsMappingPolicies = value
}
// Sets the createdObjects property value. Directory objects created by this service principal. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the createdObjects property.
func (m *ServicePrincipal) SetCreatedObjects(value []DirectoryObject)() {
    m.createdObjects = value
}
// Sets the delegatedPermissionClassifications property value. The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
// Parameters:
//  - value : Value to set for the delegatedPermissionClassifications property.
func (m *ServicePrincipal) SetDelegatedPermissionClassifications(value []DelegatedPermissionClassification)() {
    m.delegatedPermissionClassifications = value
}
// Sets the description property value. Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
// Parameters:
//  - value : Value to set for the description property.
func (m *ServicePrincipal) SetDescription(value *string)() {
    m.description = value
}
// Sets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, NOT).
// Parameters:
//  - value : Value to set for the disabledByMicrosoftStatus property.
func (m *ServicePrincipal) SetDisabledByMicrosoftStatus(value *string)() {
    m.disabledByMicrosoftStatus = value
}
// Sets the displayName property value. The display name for the service principal. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ServicePrincipal) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endpoints property value. Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences.
// Parameters:
//  - value : Value to set for the endpoints property.
func (m *ServicePrincipal) SetEndpoints(value []Endpoint)() {
    m.endpoints = value
}
// Sets the homepage property value. Home page or landing page of the application.
// Parameters:
//  - value : Value to set for the homepage property.
func (m *ServicePrincipal) SetHomepage(value *string)() {
    m.homepage = value
}
// Sets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
// Parameters:
//  - value : Value to set for the homeRealmDiscoveryPolicies property.
func (m *ServicePrincipal) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicy)() {
    m.homeRealmDiscoveryPolicies = value
}
// Sets the info property value. Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, NOT, ge, le).
// Parameters:
//  - value : Value to set for the info property.
func (m *ServicePrincipal) SetInfo(value *InformationalUrl)() {
    m.info = value
}
// Sets the keyCredentials property value. The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, NOT, ge, le).
// Parameters:
//  - value : Value to set for the keyCredentials property.
func (m *ServicePrincipal) SetKeyCredentials(value []KeyCredential)() {
    m.keyCredentials = value
}
// Sets the loginUrl property value. Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL.
// Parameters:
//  - value : Value to set for the loginUrl property.
func (m *ServicePrincipal) SetLoginUrl(value *string)() {
    m.loginUrl = value
}
// Sets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.
// Parameters:
//  - value : Value to set for the logoutUrl property.
func (m *ServicePrincipal) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
// Sets the memberOf property value. Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the memberOf property.
func (m *ServicePrincipal) SetMemberOf(value []DirectoryObject)() {
    m.memberOf = value
}
// Sets the notes property value. Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters.
// Parameters:
//  - value : Value to set for the notes property.
func (m *ServicePrincipal) SetNotes(value *string)() {
    m.notes = value
}
// Sets the notificationEmailAddresses property value. Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
// Parameters:
//  - value : Value to set for the notificationEmailAddresses property.
func (m *ServicePrincipal) SetNotificationEmailAddresses(value []string)() {
    m.notificationEmailAddresses = value
}
// Sets the oauth2PermissionGrants property value. Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the oauth2PermissionGrants property.
func (m *ServicePrincipal) SetOauth2PermissionGrants(value []OAuth2PermissionGrant)() {
    m.oauth2PermissionGrants = value
}
// Sets the oauth2PermissionScopes property value. The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity's api property. Not nullable.
// Parameters:
//  - value : Value to set for the oauth2PermissionScopes property.
func (m *ServicePrincipal) SetOauth2PermissionScopes(value []PermissionScope)() {
    m.oauth2PermissionScopes = value
}
// Sets the ownedObjects property value. Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the ownedObjects property.
func (m *ServicePrincipal) SetOwnedObjects(value []DirectoryObject)() {
    m.ownedObjects = value
}
// Sets the owners property value. Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the owners property.
func (m *ServicePrincipal) SetOwners(value []DirectoryObject)() {
    m.owners = value
}
// Sets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
// Parameters:
//  - value : Value to set for the passwordCredentials property.
func (m *ServicePrincipal) SetPasswordCredentials(value []PasswordCredential)() {
    m.passwordCredentials = value
}
// Sets the preferredSingleSignOnMode property value. Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc.
// Parameters:
//  - value : Value to set for the preferredSingleSignOnMode property.
func (m *ServicePrincipal) SetPreferredSingleSignOnMode(value *string)() {
    m.preferredSingleSignOnMode = value
}
// Sets the preferredTokenSigningKeyThumbprint property value. Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions.
// Parameters:
//  - value : Value to set for the preferredTokenSigningKeyThumbprint property.
func (m *ServicePrincipal) SetPreferredTokenSigningKeyThumbprint(value *string)() {
    m.preferredTokenSigningKeyThumbprint = value
}
// Sets the replyUrls property value. The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
// Parameters:
//  - value : Value to set for the replyUrls property.
func (m *ServicePrincipal) SetReplyUrls(value []string)() {
    m.replyUrls = value
}
// Sets the samlSingleSignOnSettings property value. The collection for settings related to saml single sign-on.
// Parameters:
//  - value : Value to set for the samlSingleSignOnSettings property.
func (m *ServicePrincipal) SetSamlSingleSignOnSettings(value *SamlSingleSignOnSettings)() {
    m.samlSingleSignOnSettings = value
}
// Sets the servicePrincipalNames property value. Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the servicePrincipalNames property.
func (m *ServicePrincipal) SetServicePrincipalNames(value []string)() {
    m.servicePrincipalNames = value
}
// Sets the servicePrincipalType property value. Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally. The servicePrincipalType property can be set to three different values: __Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens are not issued for the service principal.__ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but cannot be updated or modified directly.__Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. Legacy service principal can have credentials, service principal names, reply URLs, and other properties which are editable by an authorized user, but does not have an associated app registration. The appId value does not associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.__SocialIdp - For internal use.
// Parameters:
//  - value : Value to set for the servicePrincipalType property.
func (m *ServicePrincipal) SetServicePrincipalType(value *string)() {
    m.servicePrincipalType = value
}
// Sets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization’s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization’s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization’s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only.
// Parameters:
//  - value : Value to set for the signInAudience property.
func (m *ServicePrincipal) SetSignInAudience(value *string)() {
    m.signInAudience = value
}
// Sets the tags property value. Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the tags property.
func (m *ServicePrincipal) SetTags(value []string)() {
    m.tags = value
}
// Sets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
// Parameters:
//  - value : Value to set for the tokenEncryptionKeyId property.
func (m *ServicePrincipal) SetTokenEncryptionKeyId(value *string)() {
    m.tokenEncryptionKeyId = value
}
// Sets the tokenIssuancePolicies property value. The tokenIssuancePolicies assigned to this service principal.
// Parameters:
//  - value : Value to set for the tokenIssuancePolicies property.
func (m *ServicePrincipal) SetTokenIssuancePolicies(value []TokenIssuancePolicy)() {
    m.tokenIssuancePolicies = value
}
// Sets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this service principal.
// Parameters:
//  - value : Value to set for the tokenLifetimePolicies property.
func (m *ServicePrincipal) SetTokenLifetimePolicies(value []TokenLifetimePolicy)() {
    m.tokenLifetimePolicies = value
}
// Sets the transitiveMemberOf property value. 
// Parameters:
//  - value : Value to set for the transitiveMemberOf property.
func (m *ServicePrincipal) SetTransitiveMemberOf(value []DirectoryObject)() {
    m.transitiveMemberOf = value
}
