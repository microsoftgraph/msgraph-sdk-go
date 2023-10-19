package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder provides operations to manage the resources property of the microsoft.graph.accessPackageResourceEnvironment entity.
type EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderGetQueryParameters read-only. Required.
type EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderGetQueryParameters
}
// EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderInternal instantiates a new AccessPackageResourceItemRequestBuilder and sets the default values.
func NewEntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) {
    m := &EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/resourceEnvironments/{accessPackageResourceEnvironment%2Did}/resources/{accessPackageResource%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder instantiates a new AccessPackageResourceItemRequestBuilder and sets the default values.
func NewEntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resources for identityGovernance
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Environment provides operations to manage the environment property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) Environment()(*EntitlementManagementResourceEnvironmentsItemResourcesItemEnvironmentRequestBuilder) {
    return NewEntitlementManagementResourceEnvironmentsItemResourcesItemEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Required.
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property resources in identityGovernance
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, requestConfiguration *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Roles provides operations to manage the roles property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) Roles()(*EntitlementManagementResourceEnvironmentsItemResourcesItemRolesRequestBuilder) {
    return NewEntitlementManagementResourceEnvironmentsItemResourcesItemRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Scopes provides operations to manage the scopes property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) Scopes()(*EntitlementManagementResourceEnvironmentsItemResourcesItemScopesRequestBuilder) {
    return NewEntitlementManagementResourceEnvironmentsItemResourcesItemScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property resources for identityGovernance
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Required.
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property resources in identityGovernance
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceable, requestConfiguration *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder) {
    return NewEntitlementManagementResourceEnvironmentsItemResourcesAccessPackageResourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
