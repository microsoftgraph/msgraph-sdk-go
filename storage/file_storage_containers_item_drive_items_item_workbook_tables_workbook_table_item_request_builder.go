package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder provides operations to manage the tables property of the microsoft.graph.workbook entity.
type FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetQueryParameters represents a collection of tables associated with the workbook. Read-only.
type FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearFilters provides operations to call the clearFilters method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemClearFiltersRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ClearFilters()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemClearFiltersRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemClearFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Columns()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ConvertToRange provides operations to call the convertToRange method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemConvertToRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ConvertToRange()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemConvertToRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemConvertToRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemDataBodyRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) DataBodyRange()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemDataBodyRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property tables for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get represents a collection of tables associated with the workbook. Read-only.
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable), nil
}
// HeaderRowRange provides operations to call the headerRowRange method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemHeaderRowRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) HeaderRowRange()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemHeaderRowRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property tables in storage
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable), nil
}
// RangeEscaped provides operations to call the range method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) RangeEscaped()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReapplyFilters provides operations to call the reapplyFilters method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemReapplyFiltersRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ReapplyFilters()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemReapplyFiltersRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemReapplyFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Rows()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Sort()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemSortRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemSortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property tables for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents a collection of tables associated with the workbook. Read-only.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tables in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemTotalRowRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) TotalRowRange()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemTotalRowRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemWorksheetRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesWorkbookTableItemRequestBuilder) Worksheet()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemWorksheetRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
