package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder provides operations to call the item method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Axes provides operations to manage the axes property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameAxesRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) Axes()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameAxesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameAxesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, name *string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/item(name='{name}')", pathParameters),
    }
    if name != nil {
        m.BaseRequestBuilder.PathParameters["name"] = *name
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// DataLabels provides operations to manage the dataLabels property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameDatalabelsDataLabelsRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) DataLabels()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameDatalabelsDataLabelsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameDatalabelsDataLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Format provides operations to manage the format property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameFormatRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) Format()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameFormatRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function item
// returns a WorkbookChartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartable, error) {
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
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImageRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) Image()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImageRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImageWithWidth provides operations to call the image method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) ImageWithWidth(width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, width)
}
// ImageWithWidthWithHeight provides operations to call the image method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) ImageWithWidthWithHeight(height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, height, width)
}
// ImageWithWidthWithHeightWithFittingMode provides operations to call the image method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) ImageWithWidthWithHeightWithFittingMode(fittingMode *string, height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fittingMode, height, width)
}
// Legend provides operations to manage the legend property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameLegendRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) Legend()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameLegendRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameLegendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Series provides operations to manage the series property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSeriesRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) Series()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSeriesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSeriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetData provides operations to call the setData method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSetdataSetDataRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) SetData()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSetdataSetDataRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSetdataSetDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetPosition provides operations to call the setPosition method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSetpositionSetPositionRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) SetPosition()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSetpositionSetPositionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameSetpositionSetPositionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Title provides operations to manage the title property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameTitleRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) Title()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameTitleRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameTitleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function item
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameWorksheetRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameItemWithNameRequestBuilder) Worksheet()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
