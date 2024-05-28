package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexDatabodyrangeDataBodyRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) DataBodyRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexDatabodyrangeDataBodyRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexDatabodyrangeDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Filter provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexFilterRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) Filter()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexFilterRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets a column based on its position in the collection.
// returns a WorkbookTableColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablecolumncollection-itemat?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable), nil
}
// HeaderRowRange provides operations to call the headerRowRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) HeaderRowRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) RangeEscaped()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets a column based on its position in the collection.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexTotalrowrangeTotalRowRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) TotalRowRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexTotalrowrangeTotalRowRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexTotalrowrangeTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItematwithindexItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
