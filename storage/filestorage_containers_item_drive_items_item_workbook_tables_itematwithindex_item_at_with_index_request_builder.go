package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearFilters provides operations to call the clearFilters method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexClearfiltersClearFiltersRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) ClearFilters()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexClearfiltersClearFiltersRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexClearfiltersClearFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexColumnsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Columns()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexColumnsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ConvertToRange provides operations to call the convertToRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexConverttorangeConvertToRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) ConvertToRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexConverttorangeConvertToRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexConverttorangeConvertToRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) DataBodyRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function itemAt
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) HeaderRowRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) RangeEscaped()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReapplyFilters provides operations to call the reapplyFilters method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) ReapplyFilters()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexRowsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Rows()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexRowsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexSortRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Sort()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexSortRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexSortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function itemAt
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) TotalRowRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexWorksheetRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Worksheet()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexWorksheetRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItematwithindexWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
