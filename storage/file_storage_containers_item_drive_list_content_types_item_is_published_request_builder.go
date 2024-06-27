package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder provides operations to call the isPublished method.
type FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilderInternal instantiates a new FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) {
    m := &FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}/isPublished()", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder instantiates a new FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get check the publishing status of a contentType in a content type hub site.
// Deprecated: This method is obsolete. Use GetAsIsPublishedGetResponse instead.
// returns a FileStorageContainersItemDriveListContentTypesItemIsPublishedResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveListContentTypesItemIsPublishedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveListContentTypesItemIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveListContentTypesItemIsPublishedResponseable), nil
}
// GetAsIsPublishedGetResponse check the publishing status of a contentType in a content type hub site.
// returns a FileStorageContainersItemDriveListContentTypesItemIsPublishedGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) GetAsIsPublishedGetResponse(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveListContentTypesItemIsPublishedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveListContentTypesItemIsPublishedGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveListContentTypesItemIsPublishedGetResponseable), nil
}
// ToGetRequestInformation check the publishing status of a contentType in a content type hub site.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder when successful
func (m *FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) {
    return NewFileStorageContainersItemDriveListContentTypesItemIsPublishedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
