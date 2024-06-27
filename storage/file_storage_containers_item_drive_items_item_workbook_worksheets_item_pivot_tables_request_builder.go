package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder provides operations to manage the pivotTables property of the microsoft.graph.workbookWorksheet entity.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderGetQueryParameters collection of PivotTables that are part of the worksheet.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWorkbookPivotTableId provides operations to manage the pivotTables property of the microsoft.graph.workbookWorksheet entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesWorkbookPivotTableItemRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) ByWorkbookPivotTableId(workbookPivotTableId string)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesWorkbookPivotTableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workbookPivotTableId != "" {
        urlTplParams["workbookPivotTable%2Did"] = workbookPivotTableId
    }
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesWorkbookPivotTableItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/pivotTables{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesCountRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) Count()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesCountRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of PivotTables that are part of the worksheet.
// returns a WorkbookPivotTableCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookPivotTableCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookPivotTableCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookPivotTableCollectionResponseable), nil
}
// Post create new navigation property to pivotTables for storage
// returns a WorkbookPivotTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookPivotTableable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookPivotTableable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookPivotTableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookPivotTableable), nil
}
// RefreshAll provides operations to call the refreshAll method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRefreshAllRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) RefreshAll()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRefreshAllRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRefreshAllRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation collection of PivotTables that are part of the worksheet.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to pivotTables for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookPivotTableable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
