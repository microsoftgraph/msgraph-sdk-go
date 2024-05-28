package communications

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder provides operations to call the getPstnCalls method.
type CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters get log of PSTN calls as a collection of pstnCallLogRow entries.
type CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters struct {
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
// CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters
}
// NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal instantiates a new CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    m := &CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/microsoft.graph.callRecords.getPstnCalls(fromDateTime={fromDateTime},toDateTime={toDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if fromDateTime != nil {
        m.BaseRequestBuilder.PathParameters["fromDateTime"] = (*fromDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if toDateTime != nil {
        m.BaseRequestBuilder.PathParameters["toDateTime"] = (*toDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder instantiates a new CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get get log of PSTN calls as a collection of pstnCallLogRow entries.
// Deprecated: This method is obsolete. Use GetAsGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse instead.
// returns a CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeResponseable), nil
}
// GetAsGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse get log of PSTN calls as a collection of pstnCallLogRow entries.
// returns a CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder) GetAsGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponseable), nil
}
// ToGetRequestInformation get log of PSTN calls as a collection of pstnCallLogRow entries.
// returns a *RequestInformation when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder) WithUrl(rawUrl string)(*CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnCallsWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
