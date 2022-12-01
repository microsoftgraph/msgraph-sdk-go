package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServicePrincipalsServicePrincipalItemRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ServicePrincipalsServicePrincipalItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalsServicePrincipalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsServicePrincipalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalsServicePrincipalItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a servicePrincipal object.
type ServicePrincipalsServicePrincipalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalsServicePrincipalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsServicePrincipalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalsServicePrincipalItemRequestBuilderGetQueryParameters
}
// ServicePrincipalsServicePrincipalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsServicePrincipalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddKey provides operations to call the addKey method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) AddKey()(*ServicePrincipalsItemAddKeyRequestBuilder) {
    return NewServicePrincipalsItemAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddPassword provides operations to call the addPassword method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) AddPassword()(*ServicePrincipalsItemAddPasswordRequestBuilder) {
    return NewServicePrincipalsItemAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddTokenSigningCertificate provides operations to call the addTokenSigningCertificate method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) AddTokenSigningCertificate()(*ServicePrincipalsItemAddTokenSigningCertificateRequestBuilder) {
    return NewServicePrincipalsItemAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedTo provides operations to manage the appRoleAssignedTo property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) AppRoleAssignedTo()(*ServicePrincipalsItemAppRoleAssignedToRequestBuilder) {
    return NewServicePrincipalsItemAppRoleAssignedToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedToById provides operations to manage the appRoleAssignedTo property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) AppRoleAssignedToById(id string)(*ServicePrincipalsItemAppRoleAssignedToAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewServicePrincipalsItemAppRoleAssignedToAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) AppRoleAssignments()(*ServicePrincipalsItemAppRoleAssignmentsRequestBuilder) {
    return NewServicePrincipalsItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) AppRoleAssignmentsById(id string)(*ServicePrincipalsItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewServicePrincipalsItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) CheckMemberGroups()(*ServicePrincipalsItemCheckMemberGroupsRequestBuilder) {
    return NewServicePrincipalsItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) CheckMemberObjects()(*ServicePrincipalsItemCheckMemberObjectsRequestBuilder) {
    return NewServicePrincipalsItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPolicies provides operations to manage the claimsMappingPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) ClaimsMappingPolicies()(*ServicePrincipalsItemClaimsMappingPoliciesRequestBuilder) {
    return NewServicePrincipalsItemClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go.servicePrincipals.item.claimsMappingPolicies.item collection
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) ClaimsMappingPoliciesById(id string)(*ServicePrincipalsItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy%2Did"] = id
    }
    return NewServicePrincipalsItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewServicePrincipalsServicePrincipalItemRequestBuilderInternal instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalsServicePrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsServicePrincipalItemRequestBuilder) {
    m := &ServicePrincipalsServicePrincipalItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsServicePrincipalItemRequestBuilder instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalsServicePrincipalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsServicePrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsServicePrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete a servicePrincipal object.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) CreatedObjects()(*ServicePrincipalsItemCreatedObjectsRequestBuilder) {
    return NewServicePrincipalsItemCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatedObjectsById provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) CreatedObjectsById(id string)(*ServicePrincipalsItemCreatedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewServicePrincipalsItemCreatedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation retrieve the properties and relationships of a servicePrincipal object.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsServicePrincipalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in servicePrincipals
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, requestConfiguration *ServicePrincipalsServicePrincipalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DelegatedPermissionClassifications provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) DelegatedPermissionClassifications()(*ServicePrincipalsItemDelegatedPermissionClassificationsRequestBuilder) {
    return NewServicePrincipalsItemDelegatedPermissionClassificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DelegatedPermissionClassificationsById provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) DelegatedPermissionClassificationsById(id string)(*ServicePrincipalsItemDelegatedPermissionClassificationsDelegatedPermissionClassificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedPermissionClassification%2Did"] = id
    }
    return NewServicePrincipalsItemDelegatedPermissionClassificationsDelegatedPermissionClassificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete a servicePrincipal object.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalsServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Endpoints provides operations to manage the endpoints property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) Endpoints()(*ServicePrincipalsItemEndpointsRequestBuilder) {
    return NewServicePrincipalsItemEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById provides operations to manage the endpoints property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) EndpointsById(id string)(*ServicePrincipalsItemEndpointsEndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint%2Did"] = id
    }
    return NewServicePrincipalsItemEndpointsEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FederatedIdentityCredentials provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) FederatedIdentityCredentials()(*ServicePrincipalsItemFederatedIdentityCredentialsRequestBuilder) {
    return NewServicePrincipalsItemFederatedIdentityCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederatedIdentityCredentialsById provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) FederatedIdentityCredentialsById(id string)(*ServicePrincipalsItemFederatedIdentityCredentialsFederatedIdentityCredentialItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["federatedIdentityCredential%2Did"] = id
    }
    return NewServicePrincipalsItemFederatedIdentityCredentialsFederatedIdentityCredentialItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a servicePrincipal object.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalsServicePrincipalItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) GetMemberGroups()(*ServicePrincipalsItemGetMemberGroupsRequestBuilder) {
    return NewServicePrincipalsItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) GetMemberObjects()(*ServicePrincipalsItemGetMemberObjectsRequestBuilder) {
    return NewServicePrincipalsItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPolicies()(*ServicePrincipalsItemHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewServicePrincipalsItemHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go.servicePrincipals.item.homeRealmDiscoveryPolicies.item collection
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*ServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return NewServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) MemberOf()(*ServicePrincipalsItemMemberOfRequestBuilder) {
    return NewServicePrincipalsItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) MemberOfById(id string)(*ServicePrincipalsItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewServicePrincipalsItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) Oauth2PermissionGrants()(*ServicePrincipalsItemOauth2PermissionGrantsRequestBuilder) {
    return NewServicePrincipalsItemOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) Oauth2PermissionGrantsById(id string)(*ServicePrincipalsItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return NewServicePrincipalsItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) OwnedObjects()(*ServicePrincipalsItemOwnedObjectsRequestBuilder) {
    return NewServicePrincipalsItemOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedObjectsById provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) OwnedObjectsById(id string)(*ServicePrincipalsItemOwnedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewServicePrincipalsItemOwnedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) Owners()(*ServicePrincipalsItemOwnersRequestBuilder) {
    return NewServicePrincipalsItemOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go.servicePrincipals.item.owners.item collection
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) OwnersById(id string)(*ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in servicePrincipals
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, requestConfiguration *ServicePrincipalsServicePrincipalItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// RemoveKey provides operations to call the removeKey method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) RemoveKey()(*ServicePrincipalsItemRemoveKeyRequestBuilder) {
    return NewServicePrincipalsItemRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemovePassword provides operations to call the removePassword method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) RemovePassword()(*ServicePrincipalsItemRemovePasswordRequestBuilder) {
    return NewServicePrincipalsItemRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) Restore()(*ServicePrincipalsItemRestoreRequestBuilder) {
    return NewServicePrincipalsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) TokenIssuancePolicies()(*ServicePrincipalsItemTokenIssuancePoliciesRequestBuilder) {
    return NewServicePrincipalsItemTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) TokenIssuancePoliciesById(id string)(*ServicePrincipalsItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return NewServicePrincipalsItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) TokenLifetimePolicies()(*ServicePrincipalsItemTokenLifetimePoliciesRequestBuilder) {
    return NewServicePrincipalsItemTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) TokenLifetimePoliciesById(id string)(*ServicePrincipalsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return NewServicePrincipalsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) TransitiveMemberOf()(*ServicePrincipalsItemTransitiveMemberOfRequestBuilder) {
    return NewServicePrincipalsItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalsServicePrincipalItemRequestBuilder) TransitiveMemberOfById(id string)(*ServicePrincipalsItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewServicePrincipalsItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
