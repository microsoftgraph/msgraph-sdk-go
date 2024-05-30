package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters represents a collection of all the columns in the table. Read-only.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemDatabodyrangeDataBodyRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) DataBodyRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemDatabodyrangeDataBodyRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemDatabodyrangeDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property columns for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Filter provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Filter()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents a collection of all the columns in the table. Read-only.
// returns a WorkbookTableColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemHeaderrowrangeHeaderRowRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) HeaderRowRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemHeaderrowrangeHeaderRowRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemHeaderrowrangeHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property columns in storage
// returns a WorkbookTableColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// RangeEscaped provides operations to call the range method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) RangeEscaped()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property columns for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents a collection of all the columns in the table. Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property columns in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemTotalrowrangeTotalRowRangeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) TotalRowRange()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemTotalrowrangeTotalRowRangeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemTotalrowrangeTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsWorkbookTableColumnItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
