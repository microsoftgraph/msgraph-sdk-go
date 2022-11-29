package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder provides operations to manage the internalSponsors property of the microsoft.graph.connectedOrganization entity.
type IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetQueryParameters retrieve a list of a connectedOrganization's internal sponsors.  The internal sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
type IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderInternal instantiates a new InternalSponsorsRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/internalSponsors{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder instantiates a new InternalSponsorsRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) Count()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsCountRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation retrieve a list of a connectedOrganization's internal sponsors.  The internal sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to internalSponsors for identityGovernance
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get retrieve a list of a connectedOrganization's internal sponsors.  The internal sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// GetAvailableExtensionProperties provides operations to call the getAvailableExtensionProperties method.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) GetAvailableExtensionProperties()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetAvailableExtensionPropertiesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetByIds provides operations to call the getByIds method.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) GetByIds()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to internalSponsors for identityGovernance
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) Ref()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRefRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) ValidateProperties()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsValidatePropertiesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemInternalSponsorsValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
