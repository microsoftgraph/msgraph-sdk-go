package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder provides operations to manage the worksheet property of the microsoft.graph.workbookPivotTable entity.
type ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderGetQueryParameters the worksheet containing the current PivotTable. Read-only.
type ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderGetQueryParameters
}
// NewItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/pivotTables/{workbookPivotTable%2Did}/worksheet{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the worksheet containing the current PivotTable. Read-only.
// returns a WorkbookWorksheetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookWorksheetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable), nil
}
// ToGetRequestInformation the worksheet containing the current PivotTable. Read-only.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemPivottablesItemWorksheetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
