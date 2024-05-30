package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder provides operations to call the unfollow method.
type FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemUnfollowRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemUnfollowRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemUnfollowRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemUnfollowRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/unfollow", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemUnfollowRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemUnfollowRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemUnfollowRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unfollow a driveItem.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-unfollow?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemUnfollowRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unfollow a driveItem.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemUnfollowRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemUnfollowRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemUnfollowRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
