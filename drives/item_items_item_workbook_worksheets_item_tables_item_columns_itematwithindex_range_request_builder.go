package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder provides operations to call the range method.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/columns/itemAt(index={index})/range()", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the range object associated with the entire column.
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablecolumn-range?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookRangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable), nil
}
// ToGetRequestInformation gets the range object associated with the entire column.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsItematwithindexRangeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
