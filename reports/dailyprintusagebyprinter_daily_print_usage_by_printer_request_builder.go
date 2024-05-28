package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
type DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderGetQueryParameters retrieve a list of daily print usage summaries, grouped by printer.
type DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderGetQueryParameters struct {
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
// DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderGetQueryParameters
}
// DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrintUsageByPrinterId provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
// returns a *DailyprintusagebyprinterPrintUsageByPrinterItemRequestBuilder when successful
func (m *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) ByPrintUsageByPrinterId(printUsageByPrinterId string)(*DailyprintusagebyprinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if printUsageByPrinterId != "" {
        urlTplParams["printUsageByPrinter%2Did"] = printUsageByPrinterId
    }
    return NewDailyprintusagebyprinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderInternal instantiates a new DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder and sets the default values.
func NewDailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) {
    m := &DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/dailyPrintUsageByPrinter{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder instantiates a new DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder and sets the default values.
func NewDailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DailyprintusagebyprinterCountRequestBuilder when successful
func (m *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) Count()(*DailyprintusagebyprinterCountRequestBuilder) {
    return NewDailyprintusagebyprinterCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of daily print usage summaries, grouped by printer.
// returns a PrintUsageByPrinterCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-list-dailyprintusagebyprinter?view=graph-rest-1.0
func (m *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) Get(ctx context.Context, requestConfiguration *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByPrinterCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintUsageByPrinterCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByPrinterCollectionResponseable), nil
}
// Post create new navigation property to dailyPrintUsageByPrinter for reports
// returns a PrintUsageByPrinterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByPrinterable, requestConfiguration *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByPrinterable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintUsageByPrinterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByPrinterable), nil
}
// ToGetRequestInformation retrieve a list of daily print usage summaries, grouped by printer.
// returns a *RequestInformation when successful
func (m *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dailyPrintUsageByPrinter for reports
// returns a *RequestInformation when successful
func (m *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintUsageByPrinterable, requestConfiguration *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder when successful
func (m *DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) WithUrl(rawUrl string)(*DailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder) {
    return NewDailyprintusagebyprinterDailyPrintUsageByPrinterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
