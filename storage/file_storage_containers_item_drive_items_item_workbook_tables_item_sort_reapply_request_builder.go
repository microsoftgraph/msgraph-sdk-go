package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder provides operations to call the reapply method.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/sort/reapply", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reapplies the current sorting parameters to the table.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablesort-reapply?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reapplies the current sorting parameters to the table.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemSortReapplyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
