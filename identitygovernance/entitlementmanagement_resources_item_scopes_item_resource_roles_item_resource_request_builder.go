package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder provides operations to manage the resource property of the microsoft.graph.accessPackageResourceRole entity.
type EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderGetQueryParameters get resource from identityGovernance
type EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderGetQueryParameters
}
// EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderInternal instantiates a new EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) {
    m := &EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/resources/{accessPackageResource%2Did}/scopes/{accessPackageResourceScope%2Did}/resource/roles/{accessPackageResourceRole%2Did}/resource{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder instantiates a new EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder and sets the default values.
func NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resource for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Environment provides operations to manage the environment property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder when successful
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) Environment()(*EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilder) {
    return NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get resource from identityGovernance
// returns a AccessPackageResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable), nil
}
// Patch update the navigation property resource in identityGovernance
// returns a AccessPackageResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, requestConfiguration *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable), nil
}
// ToDeleteRequestInformation delete navigation property resource for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get resource from identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resource in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, requestConfiguration *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder when successful
func (m *EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder) {
    return NewEntitlementmanagementResourcesItemScopesItemResourceRolesItemResourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
