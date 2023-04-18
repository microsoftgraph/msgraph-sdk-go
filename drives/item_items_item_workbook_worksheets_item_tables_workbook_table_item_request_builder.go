package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder provides operations to manage the tables property of the microsoft.graph.workbookWorksheet entity.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetQueryParameters collection of tables that are part of the worksheet. Read-only.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearFilters provides operations to call the clearFilters method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ClearFilters()(*ItemItemsItemWorkbookWorksheetsItemTablesItemClearFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemClearFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Columns()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderInternal instantiates a new WorkbookTableItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder instantiates a new WorkbookTableItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ConvertToRange provides operations to call the convertToRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ConvertToRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemConvertToRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemConvertToRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) DataBodyRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemDataBodyRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property tables for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of tables that are part of the worksheet. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) HeaderRowRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemHeaderRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property tables in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) RangeEscaped()(*ItemItemsItemWorkbookWorksheetsItemTablesItemRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReapplyFilters provides operations to call the reapplyFilters method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ReapplyFilters()(*ItemItemsItemWorkbookWorksheetsItemTablesItemReapplyFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemReapplyFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Rows()(*ItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Sort()(*ItemItemsItemWorkbookWorksheetsItemTablesItemSortRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemSortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property tables for drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation collection of tables that are part of the worksheet. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tables in drives
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TotalRowRange provides operations to call the totalRowRange method.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) TotalRowRange()(*ItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesWorkbookTableItemRequestBuilder) Worksheet()(*ItemItemsItemWorkbookWorksheetsItemTablesItemWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
