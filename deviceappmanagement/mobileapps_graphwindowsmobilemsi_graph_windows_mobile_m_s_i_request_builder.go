package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder casts the previous resource to windowsMobileMSI.
type MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetQueryParameters get the items of type microsoft.graph.windowsMobileMSI in the microsoft.graph.mobileApp collection
type MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetQueryParameters struct {
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
// MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetQueryParameters
}
// NewMobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderInternal instantiates a new MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder and sets the default values.
func NewMobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) {
    m := &MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/graph.windowsMobileMSI{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder instantiates a new MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder and sets the default values.
func NewMobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappsGraphwindowsmobilemsiCountRequestBuilder when successful
func (m *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) Count()(*MobileappsGraphwindowsmobilemsiCountRequestBuilder) {
    return NewMobileappsGraphwindowsmobilemsiCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.windowsMobileMSI in the microsoft.graph.mobileApp collection
// returns a WindowsMobileMSICollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsMobileMSICollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsMobileMSICollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsMobileMSICollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.windowsMobileMSI in the microsoft.graph.mobileApp collection
// returns a *RequestInformation when successful
func (m *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder when successful
func (m *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) WithUrl(rawUrl string)(*MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) {
    return NewMobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
