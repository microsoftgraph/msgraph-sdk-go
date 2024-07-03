package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Axes provides operations to manage the axes property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexAxesRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) Axes()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexAxesRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexAxesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// DataLabels provides operations to manage the dataLabels property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexDataLabelsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) DataLabels()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexDataLabelsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexDataLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Format provides operations to manage the format property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexFormatRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) Format()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexFormatRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets a chart based on its position in the collection.
// returns a WorkbookChartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartcollection-itemat?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) Image()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImageWithWidth provides operations to call the image method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) ImageWithWidth(width *int32)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, width)
}
// ImageWithWidthWithHeight provides operations to call the image method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) ImageWithWidthWithHeight(height *int32, width *int32)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, height, width)
}
// ImageWithWidthWithHeightWithFittingMode provides operations to call the image method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) ImageWithWidthWithHeightWithFittingMode(fittingMode *string, height *int32, width *int32)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fittingMode, height, width)
}
// Legend provides operations to manage the legend property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexLegendRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) Legend()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexLegendRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexLegendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Series provides operations to manage the series property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSeriesRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) Series()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSeriesRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSeriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetData provides operations to call the setData method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetDataRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) SetData()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetDataRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetPosition provides operations to call the setPosition method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) SetPosition()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Title provides operations to manage the title property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexTitleRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) Title()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexTitleRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexTitleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets a chart based on its position in the collection.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookChart entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexWorksheetRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexRequestBuilder) Worksheet()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexWorksheetRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
