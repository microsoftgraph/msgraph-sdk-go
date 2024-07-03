package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder provides operations to manage the filter property of the microsoft.graph.workbookTableColumn entity.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetQueryParameters retrieve the filter applied to the column. Read-only.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apply provides operations to call the apply method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) Apply()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyBottomItemsFilter provides operations to call the applyBottomItemsFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomItemsFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyBottomItemsFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomItemsFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomItemsFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyBottomPercentFilter provides operations to call the applyBottomPercentFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyBottomPercentFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyCellColorFilter provides operations to call the applyCellColorFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyCellColorFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCellColorFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyCustomFilter provides operations to call the applyCustomFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCustomFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyCustomFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCustomFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyCustomFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyDynamicFilter provides operations to call the applyDynamicFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyDynamicFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyDynamicFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyDynamicFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyDynamicFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyFontColorFilter provides operations to call the applyFontColorFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyFontColorFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyFontColorFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyIconFilter provides operations to call the applyIconFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyIconFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyIconFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyIconFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyIconFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyTopItemsFilter provides operations to call the applyTopItemsFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyTopItemsFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyTopPercentFilter provides operations to call the applyTopPercentFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyTopPercentFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyTopPercentFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyTopPercentFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyTopPercentFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyValuesFilter provides operations to call the applyValuesFilter method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyValuesFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ApplyValuesFilter()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyValuesFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyValuesFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Clear provides operations to call the clear method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterClearRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) Clear()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterClearRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterClearRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property filter for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the filter applied to the column. Read-only.
// returns a WorkbookFilterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFilterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable), nil
}
// Patch update the navigation property filter in storage
// returns a WorkbookFilterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFilterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable), nil
}
// ToDeleteRequestInformation delete navigation property filter for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the filter applied to the column. Read-only.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property filter in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
