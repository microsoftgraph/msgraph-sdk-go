package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder provides operations to call the restoreVersion method.
type FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/versions/{driveItemVersion%2Did}/restoreVersion", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a previous version of a DriveItem to be the current version. This will create a new version with the contents of the previous version, but preserves all existing versions of the file.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitemversion-restore?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation restore a previous version of a DriveItem to be the current version. This will create a new version with the contents of the previous version, but preserves all existing versions of the file.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemVersionsItemRestoreversionRestoreVersionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
