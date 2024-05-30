package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
type EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderGetQueryParameters retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
type EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderGetQueryParameters
}
// ByAccessPackageId1 gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.incompatibleAccessPackages.item collection
// returns a *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesAccessPackageItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) ByAccessPackageId1(accessPackageId1 string)(*EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageId1 != "" {
        urlTplParams["accessPackage%2Did1"] = accessPackageId1
    }
    return NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleAccessPackages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) Count()(*EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
// returns a AccessPackageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-list-incompatibleaccesspackages?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCollectionResponseable), nil
}
// Ref provides operations to manage the collection of identityGovernance entities.
// returns a *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesRefRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) Ref()(*EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesRefRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
