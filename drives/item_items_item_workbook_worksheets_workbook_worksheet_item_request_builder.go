package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters retrieve the properties and relationships of worksheet object.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CellWithRowWithColumn provides operations to call the cell method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) CellWithRowWithColumn(column *int32, row *int32)(*ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, column, row)
}
// Charts provides operations to manage the charts property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Charts()(*ItemItemsItemWorkbookWorksheetsItemChartsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal instantiates a new WorkbookWorksheetItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder instantiates a new WorkbookWorksheetItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes the worksheet from the workbook.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/worksheet-delete?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties and relationships of worksheet object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/worksheet-get?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Names provides operations to manage the names property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Names()(*ItemItemsItemWorkbookWorksheetsItemNamesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemNamesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of worksheet object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/worksheet-update?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// PivotTables provides operations to manage the pivotTables property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) PivotTables()(*ItemItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemPivotTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Protection provides operations to manage the protection property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Protection()(*ItemItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemProtectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) RangeEscaped()(*ItemItemsItemWorkbookWorksheetsItemRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeWithAddress provides operations to call the range method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) RangeWithAddress(address *string)(*ItemItemsItemWorkbookWorksheetsItemRangeWithAddressRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemRangeWithAddressRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, address)
}
// Tables provides operations to manage the tables property of the microsoft.graph.workbookWorksheet entity.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Tables()(*ItemItemsItemWorkbookWorksheetsItemTablesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes the worksheet from the workbook.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of worksheet object.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of worksheet object.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, requestConfiguration *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UsedRange provides operations to call the usedRange method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) UsedRange()(*ItemItemsItemWorkbookWorksheetsItemUsedRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemUsedRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsedRangeWithValuesOnly provides operations to call the usedRange method.
func (m *ItemItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ItemItemsItemWorkbookWorksheetsItemUsedRangeWithValuesOnlyRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemUsedRangeWithValuesOnlyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, valuesOnly)
}
