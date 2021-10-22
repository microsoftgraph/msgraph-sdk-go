package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ServicePrincipal struct {
    DirectoryObject
    accountEnabled *bool;
    addIns []AddIn;
    alternativeNames []string;
    appDescription *string;
    appDisplayName *string;
    appId *string;
    applicationTemplateId *string;
    appOwnerOrganizationId *string;
    appRoleAssignedTo []AppRoleAssignment;
    appRoleAssignmentRequired *bool;
    appRoleAssignments []AppRoleAssignment;
    appRoles []AppRole;
    claimsMappingPolicies []ClaimsMappingPolicy;
    createdObjects []DirectoryObject;
    delegatedPermissionClassifications []DelegatedPermissionClassification;
    description *string;
    disabledByMicrosoftStatus *string;
    displayName *string;
    endpoints []Endpoint;
    homepage *string;
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy;
    info *InformationalUrl;
    keyCredentials []KeyCredential;
    loginUrl *string;
    logoutUrl *string;
    memberOf []DirectoryObject;
    notes *string;
    notificationEmailAddresses []string;
    oauth2PermissionGrants []OAuth2PermissionGrant;
    oauth2PermissionScopes []PermissionScope;
    ownedObjects []DirectoryObject;
    owners []DirectoryObject;
    passwordCredentials []PasswordCredential;
    preferredSingleSignOnMode *string;
    preferredTokenSigningKeyThumbprint *string;
    replyUrls []string;
    samlSingleSignOnSettings *SamlSingleSignOnSettings;
    servicePrincipalNames []string;
    servicePrincipalType *string;
    signInAudience *string;
    tags []string;
    tokenEncryptionKeyId *string;
    tokenIssuancePolicies []TokenIssuancePolicy;
    tokenLifetimePolicies []TokenLifetimePolicy;
    transitiveMemberOf []DirectoryObject;
}
func NewServicePrincipal()(*ServicePrincipal) {
    m := &ServicePrincipal{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
func (m *ServicePrincipal) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
func (m *ServicePrincipal) GetAddIns()([]AddIn) {
    if m == nil {
        return nil
    } else {
        return m.addIns
    }
}
func (m *ServicePrincipal) GetAlternativeNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.alternativeNames
    }
}
func (m *ServicePrincipal) GetAppDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDescription
    }
}
func (m *ServicePrincipal) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
func (m *ServicePrincipal) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
func (m *ServicePrincipal) GetApplicationTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationTemplateId
    }
}
func (m *ServicePrincipal) GetAppOwnerOrganizationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appOwnerOrganizationId
    }
}
func (m *ServicePrincipal) GetAppRoleAssignedTo()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignedTo
    }
}
func (m *ServicePrincipal) GetAppRoleAssignmentRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignmentRequired
    }
}
func (m *ServicePrincipal) GetAppRoleAssignments()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignments
    }
}
func (m *ServicePrincipal) GetAppRoles()([]AppRole) {
    if m == nil {
        return nil
    } else {
        return m.appRoles
    }
}
func (m *ServicePrincipal) GetClaimsMappingPolicies()([]ClaimsMappingPolicy) {
    if m == nil {
        return nil
    } else {
        return m.claimsMappingPolicies
    }
}
func (m *ServicePrincipal) GetCreatedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.createdObjects
    }
}
func (m *ServicePrincipal) GetDelegatedPermissionClassifications()([]DelegatedPermissionClassification) {
    if m == nil {
        return nil
    } else {
        return m.delegatedPermissionClassifications
    }
}
func (m *ServicePrincipal) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *ServicePrincipal) GetDisabledByMicrosoftStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.disabledByMicrosoftStatus
    }
}
func (m *ServicePrincipal) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ServicePrincipal) GetEndpoints()([]Endpoint) {
    if m == nil {
        return nil
    } else {
        return m.endpoints
    }
}
func (m *ServicePrincipal) GetHomepage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homepage
    }
}
func (m *ServicePrincipal) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicy) {
    if m == nil {
        return nil
    } else {
        return m.homeRealmDiscoveryPolicies
    }
}
func (m *ServicePrincipal) GetInfo()(*InformationalUrl) {
    if m == nil {
        return nil
    } else {
        return m.info
    }
}
func (m *ServicePrincipal) GetKeyCredentials()([]KeyCredential) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentials
    }
}
func (m *ServicePrincipal) GetLoginUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loginUrl
    }
}
func (m *ServicePrincipal) GetLogoutUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoutUrl
    }
}
func (m *ServicePrincipal) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
func (m *ServicePrincipal) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
func (m *ServicePrincipal) GetNotificationEmailAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notificationEmailAddresses
    }
}
func (m *ServicePrincipal) GetOauth2PermissionGrants()([]OAuth2PermissionGrant) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionGrants
    }
}
func (m *ServicePrincipal) GetOauth2PermissionScopes()([]PermissionScope) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionScopes
    }
}
func (m *ServicePrincipal) GetOwnedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.ownedObjects
    }
}
func (m *ServicePrincipal) GetOwners()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.owners
    }
}
func (m *ServicePrincipal) GetPasswordCredentials()([]PasswordCredential) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentials
    }
}
func (m *ServicePrincipal) GetPreferredSingleSignOnMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredSingleSignOnMode
    }
}
func (m *ServicePrincipal) GetPreferredTokenSigningKeyThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredTokenSigningKeyThumbprint
    }
}
func (m *ServicePrincipal) GetReplyUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.replyUrls
    }
}
func (m *ServicePrincipal) GetSamlSingleSignOnSettings()(*SamlSingleSignOnSettings) {
    if m == nil {
        return nil
    } else {
        return m.samlSingleSignOnSettings
    }
}
func (m *ServicePrincipal) GetServicePrincipalNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalNames
    }
}
func (m *ServicePrincipal) GetServicePrincipalType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalType
    }
}
func (m *ServicePrincipal) GetSignInAudience()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInAudience
    }
}
func (m *ServicePrincipal) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
func (m *ServicePrincipal) GetTokenEncryptionKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenEncryptionKeyId
    }
}
func (m *ServicePrincipal) GetTokenIssuancePolicies()([]TokenIssuancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenIssuancePolicies
    }
}
func (m *ServicePrincipal) GetTokenLifetimePolicies()([]TokenLifetimePolicy) {
    if m == nil {
        return nil
    } else {
        return m.tokenLifetimePolicies
    }
}
func (m *ServicePrincipal) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
func (m *ServicePrincipal) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
func (m *ServicePrincipal) SetAddIns(value []AddIn)() {
    m.addIns = value
}
func (m *ServicePrincipal) SetAlternativeNames(value []string)() {
    m.alternativeNames = value
}
func (m *ServicePrincipal) SetAppDescription(value *string)() {
    m.appDescription = value
}
func (m *ServicePrincipal) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
func (m *ServicePrincipal) SetAppId(value *string)() {
    m.appId = value
}
func (m *ServicePrincipal) SetApplicationTemplateId(value *string)() {
    m.applicationTemplateId = value
}
func (m *ServicePrincipal) SetAppOwnerOrganizationId(value *string)() {
    m.appOwnerOrganizationId = value
}
func (m *ServicePrincipal) SetAppRoleAssignedTo(value []AppRoleAssignment)() {
    m.appRoleAssignedTo = value
}
func (m *ServicePrincipal) SetAppRoleAssignmentRequired(value *bool)() {
    m.appRoleAssignmentRequired = value
}
func (m *ServicePrincipal) SetAppRoleAssignments(value []AppRoleAssignment)() {
    m.appRoleAssignments = value
}
func (m *ServicePrincipal) SetAppRoles(value []AppRole)() {
    m.appRoles = value
}
func (m *ServicePrincipal) SetClaimsMappingPolicies(value []ClaimsMappingPolicy)() {
    m.claimsMappingPolicies = value
}
func (m *ServicePrincipal) SetCreatedObjects(value []DirectoryObject)() {
    m.createdObjects = value
}
func (m *ServicePrincipal) SetDelegatedPermissionClassifications(value []DelegatedPermissionClassification)() {
    m.delegatedPermissionClassifications = value
}
func (m *ServicePrincipal) SetDescription(value *string)() {
    m.description = value
}
func (m *ServicePrincipal) SetDisabledByMicrosoftStatus(value *string)() {
    m.disabledByMicrosoftStatus = value
}
func (m *ServicePrincipal) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ServicePrincipal) SetEndpoints(value []Endpoint)() {
    m.endpoints = value
}
func (m *ServicePrincipal) SetHomepage(value *string)() {
    m.homepage = value
}
func (m *ServicePrincipal) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicy)() {
    m.homeRealmDiscoveryPolicies = value
}
func (m *ServicePrincipal) SetInfo(value *InformationalUrl)() {
    m.info = value
}
func (m *ServicePrincipal) SetKeyCredentials(value []KeyCredential)() {
    m.keyCredentials = value
}
func (m *ServicePrincipal) SetLoginUrl(value *string)() {
    m.loginUrl = value
}
func (m *ServicePrincipal) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
func (m *ServicePrincipal) SetMemberOf(value []DirectoryObject)() {
    m.memberOf = value
}
func (m *ServicePrincipal) SetNotes(value *string)() {
    m.notes = value
}
func (m *ServicePrincipal) SetNotificationEmailAddresses(value []string)() {
    m.notificationEmailAddresses = value
}
func (m *ServicePrincipal) SetOauth2PermissionGrants(value []OAuth2PermissionGrant)() {
    m.oauth2PermissionGrants = value
}
func (m *ServicePrincipal) SetOauth2PermissionScopes(value []PermissionScope)() {
    m.oauth2PermissionScopes = value
}
func (m *ServicePrincipal) SetOwnedObjects(value []DirectoryObject)() {
    m.ownedObjects = value
}
func (m *ServicePrincipal) SetOwners(value []DirectoryObject)() {
    m.owners = value
}
func (m *ServicePrincipal) SetPasswordCredentials(value []PasswordCredential)() {
    m.passwordCredentials = value
}
func (m *ServicePrincipal) SetPreferredSingleSignOnMode(value *string)() {
    m.preferredSingleSignOnMode = value
}
func (m *ServicePrincipal) SetPreferredTokenSigningKeyThumbprint(value *string)() {
    m.preferredTokenSigningKeyThumbprint = value
}
func (m *ServicePrincipal) SetReplyUrls(value []string)() {
    m.replyUrls = value
}
func (m *ServicePrincipal) SetSamlSingleSignOnSettings(value *SamlSingleSignOnSettings)() {
    m.samlSingleSignOnSettings = value
}
func (m *ServicePrincipal) SetServicePrincipalNames(value []string)() {
    m.servicePrincipalNames = value
}
func (m *ServicePrincipal) SetServicePrincipalType(value *string)() {
    m.servicePrincipalType = value
}
func (m *ServicePrincipal) SetSignInAudience(value *string)() {
    m.signInAudience = value
}
func (m *ServicePrincipal) SetTags(value []string)() {
    m.tags = value
}
func (m *ServicePrincipal) SetTokenEncryptionKeyId(value *string)() {
    m.tokenEncryptionKeyId = value
}
func (m *ServicePrincipal) SetTokenIssuancePolicies(value []TokenIssuancePolicy)() {
    m.tokenIssuancePolicies = value
}
func (m *ServicePrincipal) SetTokenLifetimePolicies(value []TokenLifetimePolicy)() {
    m.tokenLifetimePolicies = value
}
func (m *ServicePrincipal) SetTransitiveMemberOf(value []DirectoryObject)() {
    m.transitiveMemberOf = value
}
