package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder provides operations to call the applyFontColorFilter method.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter/applyFontColorFilter", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyFontColorFilter
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) Post(ctx context.Context, body FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterPostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action applyFontColorFilter
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterPostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
