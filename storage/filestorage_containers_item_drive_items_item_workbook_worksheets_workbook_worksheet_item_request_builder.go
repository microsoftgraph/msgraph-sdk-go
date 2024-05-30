package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters represents a collection of worksheets associated with the workbook. Read-only.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CellWithRowWithColumn provides operations to call the cell method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) CellWithRowWithColumn(column *int32, row *int32)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemCellwithrowwithcolumnCellWithRowWithColumnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, column, row)
}
// Charts provides operations to manage the charts property of the microsoft.graph.workbookWorksheet entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Charts()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property worksheets for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents a collection of worksheets associated with the workbook. Read-only.
// returns a WorkbookWorksheetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, error) {
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
// Names provides operations to manage the names property of the microsoft.graph.workbookWorksheet entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemNamesRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Names()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemNamesRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemNamesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property worksheets in storage
// returns a WorkbookWorksheetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// PivotTables provides operations to manage the pivotTables property of the microsoft.graph.workbookWorksheet entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivottablesPivotTablesRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) PivotTables()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivottablesPivotTablesRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemPivottablesPivotTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Protection provides operations to manage the protection property of the microsoft.graph.workbookWorksheet entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Protection()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemProtectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) RangeEscaped()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeWithAddress provides operations to call the range method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) RangeWithAddress(address *string)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemRangewithaddressRangeWithAddressRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, address)
}
// Tables provides operations to manage the tables property of the microsoft.graph.workbookWorksheet entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) Tables()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property worksheets for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property worksheets in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UsedRange provides operations to call the usedRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemUsedrangeUsedRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) UsedRange()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemUsedrangeUsedRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemUsedrangeUsedRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsedRangeWithValuesOnly provides operations to call the usedRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemUsedrangewithvaluesonlyUsedRangeWithValuesOnlyRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemUsedrangewithvaluesonlyUsedRangeWithValuesOnlyRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemUsedrangewithvaluesonlyUsedRangeWithValuesOnlyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, valuesOnly)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsWorkbookWorksheetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
