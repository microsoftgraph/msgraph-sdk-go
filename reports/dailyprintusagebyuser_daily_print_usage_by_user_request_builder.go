package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
type DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderGetQueryParameters retrieve a list of daily print usage summaries, grouped by user.
type DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderGetQueryParameters struct {
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
// DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderGetQueryParameters
}
// DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrintUsageByUserId provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
// returns a *DailyprintusagebyuserPrintUsageByUserItemRequestBuilder when successful
func (m *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) ByPrintUsageByUserId(printUsageByUserId string)(*DailyprintusagebyuserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if printUsageByUserId != "" {
        urlTplParams["printUsageByUser%2Did"] = printUsageByUserId
    }
    return NewDailyprintusagebyuserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDailyprintusagebyuserDailyPrintUsageByUserRequestBuilderInternal instantiates a new DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder and sets the default values.
func NewDailyprintusagebyuserDailyPrintUsageByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) {
    m := &DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/dailyPrintUsageByUser{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDailyprintusagebyuserDailyPrintUsageByUserRequestBuilder instantiates a new DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder and sets the default values.
func NewDailyprintusagebyuserDailyPrintUsageByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDailyprintusagebyuserDailyPrintUsageByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DailyprintusagebyuserCountRequestBuilder when successful
func (m *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) Count()(*DailyprintusagebyuserCountRequestBuilder) {
    return NewDailyprintusagebyuserCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of daily print usage summaries, grouped by user.
// returns a PrintUsageByUserCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-list-dailyprintusagebyuser?view=graph-rest-1.0
func (m *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByUserCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintUsageByUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByUserCollectionResponseable), nil
}
// Post create new navigation property to dailyPrintUsageByUser for reports
// returns a PrintUsageByUserable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByUserable, requestConfiguration *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByUserable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintUsageByUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByUserable), nil
}
// ToGetRequestInformation retrieve a list of daily print usage summaries, grouped by user.
// returns a *RequestInformation when successful
func (m *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dailyPrintUsageByUser for reports
// returns a *RequestInformation when successful
func (m *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByUserable, requestConfiguration *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder when successful
func (m *DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) WithUrl(rawUrl string)(*DailyprintusagebyuserDailyPrintUsageByUserRequestBuilder) {
    return NewDailyprintusagebyuserDailyPrintUsageByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
