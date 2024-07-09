package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder provides operations to call the restore method.
type FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal instantiates a new FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder) {
    m := &FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/items/{listItem%2Did}/documentSetVersions/{documentSetVersion%2Did}/restore", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder instantiates a new FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a document set version.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/documentsetversion-restore?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation restore a document set version.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder when successful
func (m *FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder) {
    return NewFileStorageContainersItemDriveListItemsItemDocumentSetVersionsItemRestoreRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
