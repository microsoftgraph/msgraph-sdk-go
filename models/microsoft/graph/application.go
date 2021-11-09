package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Application struct {
    DirectoryObject
    // Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Office 365 call the application in the context of a document the user is working on.
    addIns []AddIn;
    // Specifies settings for an application that implements a web API.
    api *ApiApplication;
    // The unique identifier for the application that is assigned to an application by Azure AD. Not nullable. Read-only.
    appId *string;
    // Unique identifier of the applicationTemplate.
    applicationTemplateId *string;
    // The collection of roles assigned to the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
    appRoles []AppRole;
    // The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, NOT, ge, le, in) and $orderBy.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only.
    createdOnBehalfOf *DirectoryObject;
    // An optional description of the application. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
    description *string;
    // Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, NOT).
    disabledByMicrosoftStatus *string;
    // The display name for the application. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
    displayName *string;
    // Read-only. Nullable.
    extensionProperties []ExtensionProperty;
    // Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following valid string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all of the security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
    groupMembershipClaims *string;
    // 
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy;
    // Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
    identifierUris []string;
    // Basic profile information of the application such as  app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, NOT, ge, le).
    info *InformationalUrl;
    // Specifies whether this application supports device authentication without a user. The default is false.
    isDeviceOnlyAuthSupported *bool;
    // Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where it is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
    isFallbackPublicClient *bool;
    // The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, NOT, ge, le).
    keyCredentials []KeyCredential;
    // The main logo for the application. Not nullable.
    logo []byte;
    // Notes relevant for the management of the application.
    notes *string;
    // 
    oauth2RequirePostResponse *bool;
    // Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
    optionalClaims *OptionalClaims;
    // Directory objects that are owners of the application. Read-only. Nullable. Supports $expand.
    owners []DirectoryObject;
    // Specifies parental control settings for an application.
    parentalControlSettings *ParentalControlSettings;
    // The collection of password credentials associated with the application. Not nullable.
    passwordCredentials []PasswordCredential;
    // Specifies settings for installed clients such as desktop or mobile devices.
    publicClient *PublicClientApplication;
    // The verified publisher domain for the application. Read-only. For more information, see How to: Configure an application's publisher domain. Supports $filter (eq, ne, ge, le, startsWith).
    publisherDomain *string;
    // Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, NOT, ge, le).
    requiredResourceAccess []RequiredResourceAccess;
    // Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, NOT).
    signInAudience *string;
    // Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
    spa *SpaApplication;
    // Custom strings that can be used to categorize and identify the application. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
    tags []string;
    // Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
    tokenEncryptionKeyId *string;
    // 
    tokenIssuancePolicies []TokenIssuancePolicy;
    // The tokenLifetimePolicies assigned to this application. Supports $expand.
    tokenLifetimePolicies []TokenLifetimePolicy;
    // Specifies the verified publisher of the application.
    verifiedPublisher *VerifiedPublisher;
    // Specifies settings for a web application.
    web *WebApplication;
}
// Instantiates a new application and sets the default values.
func NewApplication()(*Application) {
    m := &Application{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Office 365 call the application in the context of a document the user is working on.
func (m *Application) GetAddIns()([]AddIn) {
    if m == nil {
        return nil
    } else {
        return m.addIns
    }
}
// Gets the api property value. Specifies settings for an application that implements a web API.
func (m *Application) GetApi()(*ApiApplication) {
    if m == nil {
        return nil
    } else {
        return m.api
    }
}
// Gets the appId property value. The unique identifier for the application that is assigned to an application by Azure AD. Not nullable. Read-only.
func (m *Application) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the applicationTemplateId property value. Unique identifier of the applicationTemplate.
func (m *Application) GetApplicationTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationTemplateId
    }
}
// Gets the appRoles property value. The collection of roles assigned to the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
func (m *Application) GetAppRoles()([]AppRole) {
    if m == nil {
        return nil
    } else {
        return m.appRoles
    }
}
// Gets the createdDateTime property value. The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, NOT, ge, le, in) and $orderBy.
func (m *Application) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the createdOnBehalfOf property value. Read-only.
func (m *Application) GetCreatedOnBehalfOf()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.createdOnBehalfOf
    }
}
// Gets the description property value. An optional description of the application. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
func (m *Application) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, NOT).
func (m *Application) GetDisabledByMicrosoftStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledByMicrosoftStatus
    }
}
// Gets the displayName property value. The display name for the application. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
func (m *Application) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the extensionProperties property value. Read-only. Nullable.
func (m *Application) GetExtensionProperties()([]ExtensionProperty) {
    if m == nil {
        return nil
    } else {
        return m.extensionProperties
    }
}
// Gets the groupMembershipClaims property value. Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following valid string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all of the security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
func (m *Application) GetGroupMembershipClaims()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupMembershipClaims
    }
}
// Gets the homeRealmDiscoveryPolicies property value. 
func (m *Application) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicy) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
// Gets the identifierUris property value. Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) GetIdentifierUris()([]string) {
    if m == nil {
        return nil
    } else {
        return m.identifierUris
    }
}
// Gets the info property value. Basic profile information of the application such as  app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, NOT, ge, le).
func (m *Application) GetInfo()(*InformationalUrl) {
    if m == nil {
        return nil
    } else {
        return m.info
    }
}
// Gets the isDeviceOnlyAuthSupported property value. Specifies whether this application supports device authentication without a user. The default is false.
func (m *Application) GetIsDeviceOnlyAuthSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeviceOnlyAuthSupported
    }
}
// Gets the isFallbackPublicClient property value. Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where it is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
func (m *Application) GetIsFallbackPublicClient()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFallbackPublicClient
    }
}
// Gets the keyCredentials property value. The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, NOT, ge, le).
func (m *Application) GetKeyCredentials()([]KeyCredential) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentials
    }
}
// Gets the logo property value. The main logo for the application. Not nullable.
func (m *Application) GetLogo()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.logo
    }
}
// Gets the notes property value. Notes relevant for the management of the application.
func (m *Application) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the oauth2RequirePostResponse property value. 
func (m *Application) GetOauth2RequirePostResponse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.oauth2RequirePostResponse
    }
}
// Gets the optionalClaims property value. Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
func (m *Application) GetOptionalClaims()(*OptionalClaims) {
    if m == nil {
        return nil
    } else {
        return m.optionalClaims
    }
}
// Gets the owners property value. Directory objects that are owners of the application. Read-only. Nullable. Supports $expand.
func (m *Application) GetOwners()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.owners
    }
}
// Gets the parentalControlSettings property value. Specifies parental control settings for an application.
func (m *Application) GetParentalControlSettings()(*ParentalControlSettings) {
    if m == nil {
        return nil
    } else {
        return m.parentalControlSettings
    }
}
// Gets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *Application) GetPasswordCredentials()([]PasswordCredential) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentials
    }
}
// Gets the publicClient property value. Specifies settings for installed clients such as desktop or mobile devices.
func (m *Application) GetPublicClient()(*PublicClientApplication) {
    if m == nil {
        return nil
    } else {
        return m.publicClient
    }
}
// Gets the publisherDomain property value. The verified publisher domain for the application. Read-only. For more information, see How to: Configure an application's publisher domain. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) GetPublisherDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherDomain
    }
}
// Gets the requiredResourceAccess property value. Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, NOT, ge, le).
func (m *Application) GetRequiredResourceAccess()([]RequiredResourceAccess) {
    if m == nil {
        return nil
    } else {
        return m.requiredResourceAccess
    }
}
// Gets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, NOT).
func (m *Application) GetSignInAudience()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInAudience
    }
}
// Gets the spa property value. Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
func (m *Application) GetSpa()(*SpaApplication) {
    if m == nil {
        return nil
    } else {
        return m.spa
    }
}
// Gets the tags property value. Custom strings that can be used to categorize and identify the application. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *Application) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Gets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *Application) GetTokenEncryptionKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenEncryptionKeyId
    }
}
// Gets the tokenIssuancePolicies property value. 
func (m *Application) GetTokenIssuancePolicies()([]TokenIssuancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
// Gets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this application. Supports $expand.
func (m *Application) GetTokenLifetimePolicies()([]TokenLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
// Gets the verifiedPublisher property value. Specifies the verified publisher of the application.
func (m *Application) GetVerifiedPublisher()(*VerifiedPublisher) {
    if m == nil {
        return nil
    } else {
        return m.verifiedPublisher
    }
}
// Gets the web property value. Specifies settings for a web application.
func (m *Application) GetWeb()(*WebApplication) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
// The deserialization information for the current model
func (m *Application) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
    res["api"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApiApplication() })
        if err != nil {
            return err
        }
        m.SetApi(val.(*ApiApplication))
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
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["createdOnBehalfOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        m.SetCreatedOnBehalfOf(val.(*DirectoryObject))
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
    res["extensionProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtensionProperty() })
        if err != nil {
            return err
        }
        res := make([]ExtensionProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExtensionProperty))
        }
        m.SetExtensionProperties(res)
        return nil
    }
    res["groupMembershipClaims"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupMembershipClaims(val)
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
    res["identifierUris"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetIdentifierUris(res)
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
    res["isDeviceOnlyAuthSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeviceOnlyAuthSupported(val)
        return nil
    }
    res["isFallbackPublicClient"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsFallbackPublicClient(val)
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
    res["logo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetLogo(val)
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
    res["oauth2RequirePostResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOauth2RequirePostResponse(val)
        return nil
    }
    res["optionalClaims"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOptionalClaims() })
        if err != nil {
            return err
        }
        m.SetOptionalClaims(val.(*OptionalClaims))
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
    res["parentalControlSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParentalControlSettings() })
        if err != nil {
            return err
        }
        m.SetParentalControlSettings(val.(*ParentalControlSettings))
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
    res["publicClient"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicClientApplication() })
        if err != nil {
            return err
        }
        m.SetPublicClient(val.(*PublicClientApplication))
        return nil
    }
    res["publisherDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublisherDomain(val)
        return nil
    }
    res["requiredResourceAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequiredResourceAccess() })
        if err != nil {
            return err
        }
        res := make([]RequiredResourceAccess, len(val))
        for i, v := range val {
            res[i] = *(v.(*RequiredResourceAccess))
        }
        m.SetRequiredResourceAccess(res)
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
    res["spa"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSpaApplication() })
        if err != nil {
            return err
        }
        m.SetSpa(val.(*SpaApplication))
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
    res["verifiedPublisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVerifiedPublisher() })
        if err != nil {
            return err
        }
        m.SetVerifiedPublisher(val.(*VerifiedPublisher))
        return nil
    }
    res["web"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWebApplication() })
        if err != nil {
            return err
        }
        m.SetWeb(val.(*WebApplication))
        return nil
    }
    return res
}
func (m *Application) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Application) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteObjectValue("api", m.GetApi())
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdOnBehalfOf", m.GetCreatedOnBehalfOf())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensionProperties()))
        for i, v := range m.GetExtensionProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensionProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupMembershipClaims", m.GetGroupMembershipClaims())
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
        err = writer.WriteCollectionOfStringValues("identifierUris", m.GetIdentifierUris())
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
        err = writer.WriteBoolValue("isDeviceOnlyAuthSupported", m.GetIsDeviceOnlyAuthSupported())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFallbackPublicClient", m.GetIsFallbackPublicClient())
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
        err = writer.WriteByteArrayValue("logo", m.GetLogo())
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
        err = writer.WriteBoolValue("oauth2RequirePostResponse", m.GetOauth2RequirePostResponse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("optionalClaims", m.GetOptionalClaims())
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
        err = writer.WriteObjectValue("parentalControlSettings", m.GetParentalControlSettings())
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
        err = writer.WriteObjectValue("publicClient", m.GetPublicClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherDomain", m.GetPublisherDomain())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRequiredResourceAccess()))
        for i, v := range m.GetRequiredResourceAccess() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("requiredResourceAccess", cast)
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
        err = writer.WriteObjectValue("spa", m.GetSpa())
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
        err = writer.WriteObjectValue("verifiedPublisher", m.GetVerifiedPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Office 365 call the application in the context of a document the user is working on.
// Parameters:
//  - value : Value to set for the addIns property.
func (m *Application) SetAddIns(value []AddIn)() {
    m.addIns = value
}
// Sets the api property value. Specifies settings for an application that implements a web API.
// Parameters:
//  - value : Value to set for the api property.
func (m *Application) SetApi(value *ApiApplication)() {
    m.api = value
}
// Sets the appId property value. The unique identifier for the application that is assigned to an application by Azure AD. Not nullable. Read-only.
// Parameters:
//  - value : Value to set for the appId property.
func (m *Application) SetAppId(value *string)() {
    m.appId = value
}
// Sets the applicationTemplateId property value. Unique identifier of the applicationTemplate.
// Parameters:
//  - value : Value to set for the applicationTemplateId property.
func (m *Application) SetApplicationTemplateId(value *string)() {
    m.applicationTemplateId = value
}
// Sets the appRoles property value. The collection of roles assigned to the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
// Parameters:
//  - value : Value to set for the appRoles property.
func (m *Application) SetAppRoles(value []AppRole)() {
    m.appRoles = value
}
// Sets the createdDateTime property value. The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, NOT, ge, le, in) and $orderBy.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Application) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the createdOnBehalfOf property value. Read-only.
// Parameters:
//  - value : Value to set for the createdOnBehalfOf property.
func (m *Application) SetCreatedOnBehalfOf(value *DirectoryObject)() {
    m.createdOnBehalfOf = value
}
// Sets the description property value. An optional description of the application. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
// Parameters:
//  - value : Value to set for the description property.
func (m *Application) SetDescription(value *string)() {
    m.description = value
}
// Sets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, NOT).
// Parameters:
//  - value : Value to set for the disabledByMicrosoftStatus property.
func (m *Application) SetDisabledByMicrosoftStatus(value *string)() {
    m.disabledByMicrosoftStatus = value
}
// Sets the displayName property value. The display name for the application. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Application) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the extensionProperties property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the extensionProperties property.
func (m *Application) SetExtensionProperties(value []ExtensionProperty)() {
    m.extensionProperties = value
}
// Sets the groupMembershipClaims property value. Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following valid string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all of the security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
// Parameters:
//  - value : Value to set for the groupMembershipClaims property.
func (m *Application) SetGroupMembershipClaims(value *string)() {
    m.groupMembershipClaims = value
}
// Sets the homeRealmDiscoveryPolicies property value. 
// Parameters:
//  - value : Value to set for the homeRealmDiscoveryPolicies property.
func (m *Application) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicy)() {
    m.homeRealmDiscoveryPolicies = value
}
// Sets the identifierUris property value. Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the identifierUris property.
func (m *Application) SetIdentifierUris(value []string)() {
    m.identifierUris = value
}
// Sets the info property value. Basic profile information of the application such as  app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, NOT, ge, le).
// Parameters:
//  - value : Value to set for the info property.
func (m *Application) SetInfo(value *InformationalUrl)() {
    m.info = value
}
// Sets the isDeviceOnlyAuthSupported property value. Specifies whether this application supports device authentication without a user. The default is false.
// Parameters:
//  - value : Value to set for the isDeviceOnlyAuthSupported property.
func (m *Application) SetIsDeviceOnlyAuthSupported(value *bool)() {
    m.isDeviceOnlyAuthSupported = value
}
// Sets the isFallbackPublicClient property value. Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where it is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
// Parameters:
//  - value : Value to set for the isFallbackPublicClient property.
func (m *Application) SetIsFallbackPublicClient(value *bool)() {
    m.isFallbackPublicClient = value
}
// Sets the keyCredentials property value. The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, NOT, ge, le).
// Parameters:
//  - value : Value to set for the keyCredentials property.
func (m *Application) SetKeyCredentials(value []KeyCredential)() {
    m.keyCredentials = value
}
// Sets the logo property value. The main logo for the application. Not nullable.
// Parameters:
//  - value : Value to set for the logo property.
func (m *Application) SetLogo(value []byte)() {
    m.logo = value
}
// Sets the notes property value. Notes relevant for the management of the application.
// Parameters:
//  - value : Value to set for the notes property.
func (m *Application) SetNotes(value *string)() {
    m.notes = value
}
// Sets the oauth2RequirePostResponse property value. 
// Parameters:
//  - value : Value to set for the oauth2RequirePostResponse property.
func (m *Application) SetOauth2RequirePostResponse(value *bool)() {
    m.oauth2RequirePostResponse = value
}
// Sets the optionalClaims property value. Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
// Parameters:
//  - value : Value to set for the optionalClaims property.
func (m *Application) SetOptionalClaims(value *OptionalClaims)() {
    m.optionalClaims = value
}
// Sets the owners property value. Directory objects that are owners of the application. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the owners property.
func (m *Application) SetOwners(value []DirectoryObject)() {
    m.owners = value
}
// Sets the parentalControlSettings property value. Specifies parental control settings for an application.
// Parameters:
//  - value : Value to set for the parentalControlSettings property.
func (m *Application) SetParentalControlSettings(value *ParentalControlSettings)() {
    m.parentalControlSettings = value
}
// Sets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
// Parameters:
//  - value : Value to set for the passwordCredentials property.
func (m *Application) SetPasswordCredentials(value []PasswordCredential)() {
    m.passwordCredentials = value
}
// Sets the publicClient property value. Specifies settings for installed clients such as desktop or mobile devices.
// Parameters:
//  - value : Value to set for the publicClient property.
func (m *Application) SetPublicClient(value *PublicClientApplication)() {
    m.publicClient = value
}
// Sets the publisherDomain property value. The verified publisher domain for the application. Read-only. For more information, see How to: Configure an application's publisher domain. Supports $filter (eq, ne, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the publisherDomain property.
func (m *Application) SetPublisherDomain(value *string)() {
    m.publisherDomain = value
}
// Sets the requiredResourceAccess property value. Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, NOT, ge, le).
// Parameters:
//  - value : Value to set for the requiredResourceAccess property.
func (m *Application) SetRequiredResourceAccess(value []RequiredResourceAccess)() {
    m.requiredResourceAccess = value
}
// Sets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, NOT).
// Parameters:
//  - value : Value to set for the signInAudience property.
func (m *Application) SetSignInAudience(value *string)() {
    m.signInAudience = value
}
// Sets the spa property value. Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
// Parameters:
//  - value : Value to set for the spa property.
func (m *Application) SetSpa(value *SpaApplication)() {
    m.spa = value
}
// Sets the tags property value. Custom strings that can be used to categorize and identify the application. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the tags property.
func (m *Application) SetTags(value []string)() {
    m.tags = value
}
// Sets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
// Parameters:
//  - value : Value to set for the tokenEncryptionKeyId property.
func (m *Application) SetTokenEncryptionKeyId(value *string)() {
    m.tokenEncryptionKeyId = value
}
// Sets the tokenIssuancePolicies property value. 
// Parameters:
//  - value : Value to set for the tokenIssuancePolicies property.
func (m *Application) SetTokenIssuancePolicies(value []TokenIssuancePolicy)() {
    m.tokenIssuancePolicies = value
}
// Sets the tokenLifetimePolicies property value. The tokenLifetimePolicies assigned to this application. Supports $expand.
// Parameters:
//  - value : Value to set for the tokenLifetimePolicies property.
func (m *Application) SetTokenLifetimePolicies(value []TokenLifetimePolicy)() {
    m.tokenLifetimePolicies = value
}
// Sets the verifiedPublisher property value. Specifies the verified publisher of the application.
// Parameters:
//  - value : Value to set for the verifiedPublisher property.
func (m *Application) SetVerifiedPublisher(value *VerifiedPublisher)() {
    m.verifiedPublisher = value
}
// Sets the web property value. Specifies settings for a web application.
// Parameters:
//  - value : Value to set for the web property.
func (m *Application) SetWeb(value *WebApplication)() {
    m.web = value
}
