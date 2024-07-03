package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder provides operations to call the count method.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/rows/count()", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function count
// Deprecated: This method is obsolete. Use GetAsCountGetResponse instead.
// returns a FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponseable), nil
}
// GetAsCountGetResponse invoke function count
// returns a FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder) GetAsCountGetResponse(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponseable), nil
}
// ToGetRequestInformation invoke function count
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
