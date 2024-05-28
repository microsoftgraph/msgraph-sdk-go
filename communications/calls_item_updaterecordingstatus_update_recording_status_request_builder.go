package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder provides operations to call the updateRecordingStatus method.
type CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilderInternal instantiates a new CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder and sets the default values.
func NewCallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder) {
    m := &CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/updateRecordingStatus", pathParameters),
    }
    return m
}
// NewCallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder instantiates a new CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder and sets the default values.
func NewCallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the application's recording status associated with a call. This requires the use of the Teams policy-based recording solution.
// returns a UpdateRecordingStatusOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/call-updaterecordingstatus?view=graph-rest-1.0
func (m *CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder) Post(ctx context.Context, body CallsItemUpdaterecordingstatusUpdateRecordingStatusPostRequestBodyable, requestConfiguration *CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateRecordingStatusOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUpdateRecordingStatusOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateRecordingStatusOperationable), nil
}
// ToPostRequestInformation update the application's recording status associated with a call. This requires the use of the Teams policy-based recording solution.
// returns a *RequestInformation when successful
func (m *CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemUpdaterecordingstatusUpdateRecordingStatusPostRequestBodyable, requestConfiguration *CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder when successful
func (m *CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder) WithUrl(rawUrl string)(*CallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder) {
    return NewCallsItemUpdaterecordingstatusUpdateRecordingStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
