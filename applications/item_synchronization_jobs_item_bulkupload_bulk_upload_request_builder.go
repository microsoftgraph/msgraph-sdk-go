package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder provides operations to manage the bulkUpload property of the microsoft.graph.synchronizationJob entity.
type ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderGetQueryParameters the bulk upload operation for the job.
type ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderGetQueryParameters
}
// ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderInternal instantiates a new ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) {
    m := &ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/bulkUpload{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder instantiates a new ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the application entity.
// returns a *ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder when successful
func (m *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) Content()(*ItemSynchronizationJobsItemBulkuploadValueContentRequestBuilder) {
    return NewItemSynchronizationJobsItemBulkuploadValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property bulkUpload for applications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a BulkUploadable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BulkUploadable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBulkUploadFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BulkUploadable), nil
}
// Patch update the navigation property bulkUpload in applications
// returns a BulkUploadable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BulkUploadable, requestConfiguration *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BulkUploadable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBulkUploadFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BulkUploadable), nil
}
// ToDeleteRequestInformation delete navigation property bulkUpload for applications
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property bulkUpload in applications
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BulkUploadable, requestConfiguration *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder when successful
func (m *ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder) {
    return NewItemSynchronizationJobsItemBulkuploadBulkUploadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
