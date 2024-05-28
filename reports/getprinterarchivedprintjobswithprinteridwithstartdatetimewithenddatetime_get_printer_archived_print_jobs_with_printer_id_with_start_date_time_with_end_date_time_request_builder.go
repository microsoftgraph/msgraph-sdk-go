package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the getPrinterArchivedPrintJobs method.
type GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters get a list of archived print jobs that were queued for particular printer.
type GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
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
// GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewGetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewGetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getPrinterArchivedPrintJobs(printerId='{printerId}',startDateTime={startDateTime},endDateTime={endDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if printerId != nil {
        m.BaseRequestBuilder.PathParameters["printerId"] = *printerId
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewGetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewGetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get get a list of archived print jobs that were queued for particular printer.
// Deprecated: This method is obsolete. Use GetAsGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetResponse instead.
// returns a GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reports-getprinterarchivedprintjobs?view=graph-rest-1.0
func (m *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetResponse get a list of archived print jobs that were queued for particular printer.
// returns a GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reports-getprinterarchivedprintjobs?view=graph-rest-1.0
func (m *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation get a list of archived print jobs that were queued for particular printer.
// returns a *RequestInformation when successful
func (m *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*GetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewGetprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
