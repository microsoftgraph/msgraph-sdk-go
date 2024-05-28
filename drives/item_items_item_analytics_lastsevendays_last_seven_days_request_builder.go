package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder provides operations to manage the lastSevenDays property of the microsoft.graph.itemAnalytics entity.
type ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderGetQueryParameters get lastSevenDays from drives
type ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderGetQueryParameters
}
// NewItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderInternal instantiates a new ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder and sets the default values.
func NewItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder) {
    m := &ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/analytics/lastSevenDays{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder instantiates a new ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder and sets the default values.
func NewItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get lastSevenDays from drives
// returns a ItemActivityStatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// ToGetRequestInformation get lastSevenDays from drives
// returns a *RequestInformation when successful
func (m *ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder when successful
func (m *ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder) {
    return NewItemItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
