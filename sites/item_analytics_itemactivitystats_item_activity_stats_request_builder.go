package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder provides operations to manage the itemActivityStats property of the microsoft.graph.itemAnalytics entity.
type ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetQueryParameters get itemActivityStats from sites
type ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetQueryParameters struct {
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
// ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetQueryParameters
}
// ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByItemActivityStatId provides operations to manage the itemActivityStats property of the microsoft.graph.itemAnalytics entity.
// returns a *ItemAnalyticsItemactivitystatsItemActivityStatItemRequestBuilder when successful
func (m *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) ByItemActivityStatId(itemActivityStatId string)(*ItemAnalyticsItemactivitystatsItemActivityStatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if itemActivityStatId != "" {
        urlTplParams["itemActivityStat%2Did"] = itemActivityStatId
    }
    return NewItemAnalyticsItemactivitystatsItemActivityStatItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderInternal instantiates a new ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder and sets the default values.
func NewItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) {
    m := &ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/analytics/itemActivityStats{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder instantiates a new ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder and sets the default values.
func NewItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAnalyticsItemactivitystatsCountRequestBuilder when successful
func (m *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) Count()(*ItemAnalyticsItemactivitystatsCountRequestBuilder) {
    return NewItemAnalyticsItemactivitystatsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get itemActivityStats from sites
// returns a ItemActivityStatCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemActivityStatCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatCollectionResponseable), nil
}
// Post create new navigation property to itemActivityStats for sites
// returns a ItemActivityStatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemActivityStatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable), nil
}
// ToGetRequestInformation get itemActivityStats from sites
// returns a *RequestInformation when successful
func (m *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to itemActivityStats for sites
// returns a *RequestInformation when successful
func (m *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder when successful
func (m *ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) WithUrl(rawUrl string)(*ItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) {
    return NewItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
