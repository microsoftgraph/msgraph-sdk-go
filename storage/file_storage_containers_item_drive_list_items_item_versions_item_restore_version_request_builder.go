package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder provides operations to call the restoreVersion method.
type FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilderInternal instantiates a new FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder) {
    m := &FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/items/{listItem%2Did}/versions/{listItemVersion%2Did}/restoreVersion", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder instantiates a new FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a previous version of a ListItem to be the current version. This will create a new version with the contents of the previous version, but preserves all existing versions of the item.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/listitemversion-restore?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation restore a previous version of a ListItem to be the current version. This will create a new version with the contents of the previous version, but preserves all existing versions of the item.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder when successful
func (m *FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder) {
    return NewFileStorageContainersItemDriveListItemsItemVersionsItemRestoreVersionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
