package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearFilters provides operations to call the clearFilters method.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexClearfiltersClearFiltersRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) ClearFilters()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexClearfiltersClearFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexClearfiltersClearFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexColumnsRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) Columns()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ConvertToRange provides operations to call the convertToRange method.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexConverttorangeConvertToRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) ConvertToRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexConverttorangeConvertToRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexConverttorangeConvertToRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) DataBodyRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexDatabodyrangeDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function itemAt
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, error) {
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
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) HeaderRowRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexHeaderrowrangeHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) RangeEscaped()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReapplyFilters provides operations to call the reapplyFilters method.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) ReapplyFilters()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexReapplyfiltersReapplyFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexRowsRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) Rows()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexRowsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexSortRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) Sort()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexSortRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexSortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function itemAt
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) TotalRowRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexTotalrowrangeTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexWorksheetRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexItemAtWithIndexRequestBuilder) Worksheet()(*ItemItemsItemWorkbookWorksheetsItemTablesItematwithindexWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItematwithindexWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
