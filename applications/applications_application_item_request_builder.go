package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ApplicationsApplicationItemRequestBuilder provides operations to manage the collection of application entities.
type ApplicationsApplicationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApplicationsApplicationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsApplicationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationsApplicationItemRequestBuilderGetQueryParameters get the properties and relationships of an application object.
type ApplicationsApplicationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApplicationsApplicationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsApplicationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApplicationsApplicationItemRequestBuilderGetQueryParameters
}
// ApplicationsApplicationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsApplicationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddKey provides operations to call the addKey method.
func (m *ApplicationsApplicationItemRequestBuilder) AddKey()(*ApplicationsItemAddKeyRequestBuilder) {
    return NewApplicationsItemAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddPassword provides operations to call the addPassword method.
func (m *ApplicationsApplicationItemRequestBuilder) AddPassword()(*ApplicationsItemAddPasswordRequestBuilder) {
    return NewApplicationsItemAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ApplicationsApplicationItemRequestBuilder) CheckMemberGroups()(*ApplicationsItemCheckMemberGroupsRequestBuilder) {
    return NewApplicationsItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ApplicationsApplicationItemRequestBuilder) CheckMemberObjects()(*ApplicationsItemCheckMemberObjectsRequestBuilder) {
    return NewApplicationsItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewApplicationsApplicationItemRequestBuilderInternal instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationsApplicationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsApplicationItemRequestBuilder) {
    m := &ApplicationsApplicationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationsApplicationItemRequestBuilder instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationsApplicationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsApplicationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationsApplicationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete an application object. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
func (m *ApplicationsApplicationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ApplicationsApplicationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatedOnBehalfOf provides operations to manage the createdOnBehalfOf property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) CreatedOnBehalfOf()(*ApplicationsItemCreatedOnBehalfOfRequestBuilder) {
    return NewApplicationsItemCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the properties and relationships of an application object.
func (m *ApplicationsApplicationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ApplicationsApplicationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of an application object.
func (m *ApplicationsApplicationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, requestConfiguration *ApplicationsApplicationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete an application object. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
func (m *ApplicationsApplicationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApplicationsApplicationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExtensionProperties provides operations to manage the extensionProperties property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) ExtensionProperties()(*ApplicationsItemExtensionPropertiesRequestBuilder) {
    return NewApplicationsItemExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionPropertiesById provides operations to manage the extensionProperties property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) ExtensionPropertiesById(id string)(*ApplicationsItemExtensionPropertiesExtensionPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extensionProperty%2Did"] = id
    }
    return NewApplicationsItemExtensionPropertiesExtensionPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FederatedIdentityCredentials provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) FederatedIdentityCredentials()(*ApplicationsItemFederatedIdentityCredentialsRequestBuilder) {
    return NewApplicationsItemFederatedIdentityCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederatedIdentityCredentialsById provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) FederatedIdentityCredentialsById(id string)(*ApplicationsItemFederatedIdentityCredentialsFederatedIdentityCredentialItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["federatedIdentityCredential%2Did"] = id
    }
    return NewApplicationsItemFederatedIdentityCredentialsFederatedIdentityCredentialItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get the properties and relationships of an application object.
func (m *ApplicationsApplicationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplicationsApplicationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *ApplicationsApplicationItemRequestBuilder) GetMemberGroups()(*ApplicationsItemGetMemberGroupsRequestBuilder) {
    return NewApplicationsItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *ApplicationsApplicationItemRequestBuilder) GetMemberObjects()(*ApplicationsItemGetMemberObjectsRequestBuilder) {
    return NewApplicationsItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) HomeRealmDiscoveryPolicies()(*ApplicationsItemHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewApplicationsItemHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*ApplicationsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return NewApplicationsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Logo provides operations to manage the media for the application entity.
func (m *ApplicationsApplicationItemRequestBuilder) Logo()(*ApplicationsItemLogoRequestBuilder) {
    return NewApplicationsItemLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) Owners()(*ApplicationsItemOwnersRequestBuilder) {
    return NewApplicationsItemOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.applications.item.owners.item collection
func (m *ApplicationsApplicationItemRequestBuilder) OwnersById(id string)(*ApplicationsItemOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewApplicationsItemOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of an application object.
func (m *ApplicationsApplicationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, requestConfiguration *ApplicationsApplicationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Applicationable), nil
}
// RemoveKey provides operations to call the removeKey method.
func (m *ApplicationsApplicationItemRequestBuilder) RemoveKey()(*ApplicationsItemRemoveKeyRequestBuilder) {
    return NewApplicationsItemRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemovePassword provides operations to call the removePassword method.
func (m *ApplicationsApplicationItemRequestBuilder) RemovePassword()(*ApplicationsItemRemovePasswordRequestBuilder) {
    return NewApplicationsItemRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *ApplicationsApplicationItemRequestBuilder) Restore()(*ApplicationsItemRestoreRequestBuilder) {
    return NewApplicationsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetVerifiedPublisher provides operations to call the setVerifiedPublisher method.
func (m *ApplicationsApplicationItemRequestBuilder) SetVerifiedPublisher()(*ApplicationsItemSetVerifiedPublisherRequestBuilder) {
    return NewApplicationsItemSetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) TokenIssuancePolicies()(*ApplicationsItemTokenIssuancePoliciesRequestBuilder) {
    return NewApplicationsItemTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.applications.item.tokenIssuancePolicies.item collection
func (m *ApplicationsApplicationItemRequestBuilder) TokenIssuancePoliciesById(id string)(*ApplicationsItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return NewApplicationsItemTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.application entity.
func (m *ApplicationsApplicationItemRequestBuilder) TokenLifetimePolicies()(*ApplicationsItemTokenLifetimePoliciesRequestBuilder) {
    return NewApplicationsItemTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.applications.item.tokenLifetimePolicies.item collection
func (m *ApplicationsApplicationItemRequestBuilder) TokenLifetimePoliciesById(id string)(*ApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return NewApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnsetVerifiedPublisher provides operations to call the unsetVerifiedPublisher method.
func (m *ApplicationsApplicationItemRequestBuilder) UnsetVerifiedPublisher()(*ApplicationsItemUnsetVerifiedPublisherRequestBuilder) {
    return NewApplicationsItemUnsetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
