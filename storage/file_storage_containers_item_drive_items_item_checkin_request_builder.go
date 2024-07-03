package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemCheckinRequestBuilder provides operations to call the checkin method.
type FileStorageContainersItemDriveItemsItemCheckinRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemCheckinRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemCheckinRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemCheckinRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemCheckinRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemCheckinRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemCheckinRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemCheckinRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/checkin", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemCheckinRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemCheckinRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemCheckinRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemCheckinRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemCheckinRequestBuilderInternal(urlParams, requestAdapter)
}
// Post check in a checked out driveItem resource, which makes the version of the document available to others.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-checkin?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveItemsItemCheckinRequestBuilder) Post(ctx context.Context, body FileStorageContainersItemDriveItemsItemCheckinPostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemCheckinRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation check in a checked out driveItem resource, which makes the version of the document available to others.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemCheckinRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageContainersItemDriveItemsItemCheckinPostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemCheckinRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemCheckinRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemCheckinRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemCheckinRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemCheckinRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
