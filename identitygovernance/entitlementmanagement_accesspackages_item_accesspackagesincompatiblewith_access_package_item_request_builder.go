package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
type EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderGetQueryParameters the access packages that are incompatible with this package. Read-only.
type EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackagesIncompatibleWith/{accessPackage%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the access packages that are incompatible with this package. Read-only.
// returns a AccessPackageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable), nil
}
// ToGetRequestInformation the access packages that are incompatible with this package. Read-only.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
