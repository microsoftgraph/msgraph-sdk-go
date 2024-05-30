package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder provides operations to manage the media for the application entity.
type ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderInternal instantiates a new ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) {
    m := &ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/bulkUpload/$value", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder instantiates a new ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the bulk upload operation for the job.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the bulk upload operation for the job.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put the bulk upload operation for the job.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToDeleteRequestInformation the bulk upload operation for the job.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the bulk upload operation for the job.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the bulk upload operation for the job.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder when successful
func (m *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) {
    return NewItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
