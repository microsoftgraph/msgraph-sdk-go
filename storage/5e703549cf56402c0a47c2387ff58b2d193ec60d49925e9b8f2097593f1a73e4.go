package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder provides operations to call the item method.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Axes provides operations to manage the axes property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameAxesRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) Axes()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameAxesRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameAxesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, name *string)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/item(name='{name}')", pathParameters),
    }
    if name != nil {
        m.BaseRequestBuilder.PathParameters["name"] = *name
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// DataLabels provides operations to manage the dataLabels property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameDataLabelsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) DataLabels()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameDataLabelsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameDataLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Format provides operations to manage the format property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameFormatRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) Format()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameFormatRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function item
// returns a WorkbookChartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable), nil
}
// Image provides operations to call the image method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) Image()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImageWithWidth provides operations to call the image method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) ImageWithWidth(width *int32)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, width)
}
// ImageWithWidthWithHeight provides operations to call the image method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) ImageWithWidthWithHeight(height *int32, width *int32)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, height, width)
}
// ImageWithWidthWithHeightWithFittingMode provides operations to call the image method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) ImageWithWidthWithHeightWithFittingMode(fittingMode *string, height *int32, width *int32)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fittingMode, height, width)
}
// Legend provides operations to manage the legend property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameLegendRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) Legend()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameLegendRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameLegendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Series provides operations to manage the series property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSeriesRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) Series()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSeriesRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSeriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetData provides operations to call the setData method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSetDataRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) SetData()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSetDataRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSetDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetPosition provides operations to call the setPosition method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) SetPosition()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameSetPositionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Title provides operations to manage the title property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameTitleRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) Title()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameTitleRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameTitleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function item
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameWorksheetRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameRequestBuilder) Worksheet()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameWorksheetRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
