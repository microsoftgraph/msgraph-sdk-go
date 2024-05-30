package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder provides operations to count the resources in the collection.
type MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderGetQueryParameters get the number of the resource
type MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderGetQueryParameters
}
// NewMobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderInternal instantiates a new MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder) {
    m := &MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedIOSLobApp/assignments/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder instantiates a new MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder when successful
func (m *MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder) {
    return NewMobileappsItemGraphmanagedioslobappAssignmentsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
