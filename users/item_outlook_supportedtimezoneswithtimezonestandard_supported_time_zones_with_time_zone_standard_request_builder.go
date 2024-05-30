package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder provides operations to call the supportedTimeZones method.
type ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderGetQueryParameters invoke function supportedTimeZones
type ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderGetQueryParameters struct {
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
// ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderGetQueryParameters
}
// NewItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal instantiates a new ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder and sets the default values.
func NewItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, timeZoneStandard *string)(*ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    m := &ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook/supportedTimeZones(TimeZoneStandard='{TimeZoneStandard}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if timeZoneStandard != nil {
        m.BaseRequestBuilder.PathParameters["TimeZoneStandard"] = *timeZoneStandard
    }
    return m
}
// NewItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder instantiates a new ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder and sets the default values.
func NewItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function supportedTimeZones
// Deprecated: This method is obsolete. Use GetAsSupportedTimeZonesWithTimeZoneStandardGetResponse instead.
// returns a ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration)(ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardResponseable), nil
}
// GetAsSupportedTimeZonesWithTimeZoneStandardGetResponse invoke function supportedTimeZones
// returns a ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder) GetAsSupportedTimeZonesWithTimeZoneStandardGetResponse(ctx context.Context, requestConfiguration *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration)(ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardGetResponseable), nil
}
// ToGetRequestInformation invoke function supportedTimeZones
// returns a *RequestInformation when successful
func (m *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder when successful
func (m *ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder) WithUrl(rawUrl string)(*ItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return NewItemOutlookSupportedtimezoneswithtimezonestandardSupportedTimeZonesWithTimeZoneStandardRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
