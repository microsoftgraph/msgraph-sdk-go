package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder provides operations to count the resources in the collection.
type MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderGetQueryParameters get the number of the resource
type MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderGetQueryParameters
}
// NewMobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderInternal instantiates a new MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder and sets the default values.
func NewMobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder) {
    m := &MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.win32LobApp/contentVersions/{mobileAppContent%2Did}/containedApps/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder instantiates a new MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder and sets the default values.
func NewMobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
func (m *MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder when successful
func (m *MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder) {
    return NewMobileappsItemGraphwin32lobappContentversionsItemContainedappsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
