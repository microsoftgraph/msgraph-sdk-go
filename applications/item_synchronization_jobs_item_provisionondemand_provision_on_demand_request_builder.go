package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder provides operations to call the provisionOnDemand method.
type ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilderInternal instantiates a new ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder) {
    m := &ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/provisionOnDemand", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder instantiates a new ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilderInternal(urlParams, requestAdapter)
}
// Post select a user and provision the account on-demand. The rate limit for this API is 5 requests per 10 seconds.
// returns a StringKeyStringValuePairable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationjob-provisionondemand?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder) Post(ctx context.Context, body ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringKeyStringValuePairable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateStringKeyStringValuePairFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringKeyStringValuePairable), nil
}
// ToPostRequestInformation select a user and provision the account on-demand. The rate limit for this API is 5 requests per 10 seconds.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder when successful
func (m *ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder) {
    return NewItemSynchronizationJobsItemProvisionondemandProvisionOnDemandRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
