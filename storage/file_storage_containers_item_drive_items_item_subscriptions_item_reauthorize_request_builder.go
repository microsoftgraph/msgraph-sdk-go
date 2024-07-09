package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder provides operations to call the reauthorize method.
type FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/subscriptions/{subscription%2Did}/reauthorize", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reauthorize a subscription when you receive a reauthorizationRequired challenge.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/subscription-reauthorize?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reauthorize a subscription when you receive a reauthorizationRequired challenge.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemSubscriptionsItemReauthorizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
