package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder provides operations to manage the catalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderGetQueryParameters catalog of the access package containing this policy. Read-only.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}/catalog{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderInternal(urlParams, requestAdapter)
}
// Get catalog of the access package containing this policy. Read-only.
// returns a AccessPackageCatalogable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCatalogable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCatalogable), nil
}
// ToGetRequestInformation catalog of the access package containing this policy. Read-only.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesItemCatalogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
