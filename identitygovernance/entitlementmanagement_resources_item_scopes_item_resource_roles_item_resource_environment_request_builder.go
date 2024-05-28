package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder provides operations to manage the environment property of the microsoft.graph.accessPackageResource entity.
type EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderGetQueryParameters contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
type EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderInternal instantiates a new EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder and sets the default values.
func NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder) {
    m := &EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/resources/{accessPackageResource%2Did}/scopes/{accessPackageResourceScope%2Did}/resource/roles/{accessPackageResourceRole%2Did}/resource/environment{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder instantiates a new EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder and sets the default values.
func NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
// returns a AccessPackageResourceEnvironmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable), nil
}
// ToGetRequestInformation contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder when successful
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder) {
    return NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
