package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearFilters provides operations to call the clearFilters method.
// returns a *ItemItemsItemWorkbookTablesItematwithindexClearfiltersClearFiltersRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) ClearFilters()(*ItemItemsItemWorkbookTablesItematwithindexClearfiltersClearFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexClearfiltersClearFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookTablesItematwithindexColumnsRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Columns()(*ItemItemsItemWorkbookTablesItematwithindexColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder instantiates a new ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ConvertToRange provides operations to call the convertToRange method.
// returns a *ItemItemsItemWorkbookTablesItematwithindexConverttorangeConvertToRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) ConvertToRange()(*ItemItemsItemWorkbookTablesItematwithindexConverttorangeConvertToRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexConverttorangeConvertToRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *ItemItemsItemWorkbookTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) DataBodyRange()(*ItemItemsItemWorkbookTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function itemAt
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, error) {
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
// returns a *ItemItemsItemWorkbookTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) HeaderRowRange()(*ItemItemsItemWorkbookTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *ItemItemsItemWorkbookTablesItematwithindexRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) RangeEscaped()(*ItemItemsItemWorkbookTablesItematwithindexRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReapplyFilters provides operations to call the reapplyFilters method.
// returns a *ItemItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) ReapplyFilters()(*ItemItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookTablesItematwithindexRowsRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Rows()(*ItemItemsItemWorkbookTablesItematwithindexRowsRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookTablesItematwithindexSortRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Sort()(*ItemItemsItemWorkbookTablesItematwithindexSortRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexSortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function itemAt
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *ItemItemsItemWorkbookTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) TotalRowRange()(*ItemItemsItemWorkbookTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookTablesItematwithindexWorksheetRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItematwithindexItemAtWithIndexRequestBuilder) Worksheet()(*ItemItemsItemWorkbookTablesItematwithindexWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItematwithindexWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
