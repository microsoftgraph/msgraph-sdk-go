package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemStartRequestBuilder provides operations to call the start method.
type ItemSynchronizationJobsItemStartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemStartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemStartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemStartRequestBuilderInternal instantiates a new StartRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemStartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemStartRequestBuilder) {
    m := &ItemSynchronizationJobsItemStartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/start", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemStartRequestBuilder instantiates a new StartRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemStartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemStartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemStartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start an existing synchronization job. If the job is in a paused state, it continues processing changes from the point where it was paused. If the job is in quarantine, the quarantine status is cleared. Don't create scripts to call the start job continuously while it's running because that can cause the service to stop running. Use the start job only when the job is currently paused or in quarantine.  This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationjob-start?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemStartRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemStartRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation start an existing synchronization job. If the job is in a paused state, it continues processing changes from the point where it was paused. If the job is in quarantine, the quarantine status is cleared. Don't create scripts to call the start job continuously while it's running because that can cause the service to stop running. Use the start job only when the job is currently paused or in quarantine.  This API is available in the following national cloud deployments.
func (m *ItemSynchronizationJobsItemStartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemStartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSynchronizationJobsItemStartRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemStartRequestBuilder) {
    return NewItemSynchronizationJobsItemStartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
