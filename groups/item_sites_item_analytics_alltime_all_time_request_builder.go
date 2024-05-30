package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder provides operations to manage the allTime property of the microsoft.graph.itemAnalytics entity.
type ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderGetQueryParameters get allTime from groups
type ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderGetQueryParameters
}
// NewItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderInternal instantiates a new ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder and sets the default values.
func NewItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder) {
    m := &ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/analytics/allTime{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder instantiates a new ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder and sets the default values.
func NewItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get allTime from groups
// returns a ItemActivityStatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable, error) {
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
// ToGetRequestInformation get allTime from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder when successful
func (m *ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder) {
    return NewItemSitesItemAnalyticsAlltimeAllTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
