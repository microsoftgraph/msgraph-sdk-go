package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder provides operations to call the cell method.
type ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, column *int32, row *int32)(*ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/cell(row={row},column={column})", pathParameters),
    }
    if column != nil {
        m.BaseRequestBuilder.PathParameters["column"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*column), 10)
    }
    if row != nil {
        m.BaseRequestBuilder.PathParameters["row"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*row), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get gets the range object containing the single cell based on row and column numbers. The cell can be outside the bounds of its parent range, so long as it's stays within the worksheet grid.
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/worksheet-cell?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable, error) {
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
// ToGetRequestInformation gets the range object containing the single cell based on row and column numbers. The cell can be outside the bounds of its parent range, so long as it's stays within the worksheet grid.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
