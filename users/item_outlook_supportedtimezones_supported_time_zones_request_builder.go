package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder provides operations to call the supportedTimeZones method.
type ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderGetQueryParameters invoke function supportedTimeZones
type ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderGetQueryParameters
}
// NewItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderInternal instantiates a new ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder and sets the default values.
func NewItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder) {
    m := &ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook/supportedTimeZones(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder instantiates a new ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder and sets the default values.
func NewItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function supportedTimeZones
// Deprecated: This method is obsolete. Use GetAsSupportedTimeZonesGetResponse instead.
// returns a ItemOutlookSupportedtimezonesSupportedTimeZonesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderGetRequestConfiguration)(ItemOutlookSupportedtimezonesSupportedTimeZonesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookSupportedtimezonesSupportedTimeZonesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookSupportedtimezonesSupportedTimeZonesResponseable), nil
}
// GetAsSupportedTimeZonesGetResponse invoke function supportedTimeZones
// returns a ItemOutlookSupportedtimezonesSupportedTimeZonesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder) GetAsSupportedTimeZonesGetResponse(ctx context.Context, requestConfiguration *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderGetRequestConfiguration)(ItemOutlookSupportedtimezonesSupportedTimeZonesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookSupportedtimezonesSupportedTimeZonesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookSupportedtimezonesSupportedTimeZonesGetResponseable), nil
}
// ToGetRequestInformation invoke function supportedTimeZones
// returns a *RequestInformation when successful
func (m *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder when successful
func (m *ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder) WithUrl(rawUrl string)(*ItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder) {
    return NewItemOutlookSupportedtimezonesSupportedTimeZonesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
