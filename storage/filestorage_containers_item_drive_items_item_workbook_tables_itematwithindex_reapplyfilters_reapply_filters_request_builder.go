package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder provides operations to call the reapplyFilters method.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/itemAt(index={index})/reapplyFilters", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reapplies all the filters currently on the table.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/table-reapplyfilters?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reapplies all the filters currently on the table.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
