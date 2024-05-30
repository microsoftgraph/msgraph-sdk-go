package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder provides operations to manage the principal property of the microsoft.graph.unifiedRoleAssignment entity.
type EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderGetQueryParameters referencing the assigned principal. Read-only. Supports $expand.
type EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderInternal instantiates a new EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder) {
    m := &EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleAssignments/{unifiedRoleAssignment%2Did}/principal{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder instantiates a new EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get referencing the assigned principal. Read-only. Supports $expand.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// ToGetRequestInformation referencing the assigned principal. Read-only. Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
