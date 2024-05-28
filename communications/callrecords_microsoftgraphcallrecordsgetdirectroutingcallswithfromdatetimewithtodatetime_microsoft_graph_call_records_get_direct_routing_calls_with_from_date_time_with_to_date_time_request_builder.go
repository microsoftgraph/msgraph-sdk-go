package communications

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder provides operations to call the getDirectRoutingCalls method.
type CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters get a log of direct routing calls as a collection of directRoutingLogRow entries.
type CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters struct {
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
// CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters
}
// NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal instantiates a new CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    m := &CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/microsoft.graph.callRecords.getDirectRoutingCalls(fromDateTime={fromDateTime},toDateTime={toDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if fromDateTime != nil {
        m.BaseRequestBuilder.PathParameters["fromDateTime"] = (*fromDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if toDateTime != nil {
        m.BaseRequestBuilder.PathParameters["toDateTime"] = (*toDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder instantiates a new CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get get a log of direct routing calls as a collection of directRoutingLogRow entries.
// Deprecated: This method is obsolete. Use GetAsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse instead.
// returns a CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseable), nil
}
// GetAsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse get a log of direct routing calls as a collection of directRoutingLogRow entries.
// returns a CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) GetAsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponse(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetResponseable), nil
}
// ToGetRequestInformation get a log of direct routing calls as a collection of directRoutingLogRow entries.
// returns a *RequestInformation when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) WithUrl(rawUrl string)(*CallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallrecordsMicrosoftgraphcallrecordsgetdirectroutingcallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
