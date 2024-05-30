package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder provides operations to call the cancelMediaProcessing method.
type CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilderInternal instantiates a new CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder and sets the default values.
func NewCallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder) {
    m := &CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/cancelMediaProcessing", pathParameters),
    }
    return m
}
// NewCallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder instantiates a new CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder and sets the default values.
func NewCallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancels processing for any in-progress media operations. Media operations refer to the IVR operations playPrompt and recordResponse, which are by default queued to process in order. The cancelMediaProcessing method cancels any operation that is in-process as well as operations that are queued. For example, this method can be used to clean up the IVR operation queue for a new media operation. However, it will not cancel a subscribeToTone operation because it operates independent of any operation queue.
// returns a CancelMediaProcessingOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/call-cancelmediaprocessing?view=graph-rest-1.0
func (m *CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder) Post(ctx context.Context, body CallsItemCancelmediaprocessingCancelMediaProcessingPostRequestBodyable, requestConfiguration *CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CancelMediaProcessingOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCancelMediaProcessingOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CancelMediaProcessingOperationable), nil
}
// ToPostRequestInformation cancels processing for any in-progress media operations. Media operations refer to the IVR operations playPrompt and recordResponse, which are by default queued to process in order. The cancelMediaProcessing method cancels any operation that is in-process as well as operations that are queued. For example, this method can be used to clean up the IVR operation queue for a new media operation. However, it will not cancel a subscribeToTone operation because it operates independent of any operation queue.
// returns a *RequestInformation when successful
func (m *CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemCancelmediaprocessingCancelMediaProcessingPostRequestBodyable, requestConfiguration *CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder when successful
func (m *CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder) WithUrl(rawUrl string)(*CallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder) {
    return NewCallsItemCancelmediaprocessingCancelMediaProcessingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
