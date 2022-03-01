package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i07c4bbb38bd9316f4c3b1d3c547b665236f9b96d10d86c393ac1830072b24391 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/setposition"
    i25cf65d8e047c204308390327f2e23e310366ff2da03d485b9ef4cb40794620a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/axes"
    i2cb7996ef32859e96e272e57a9595c7b9bb3bb056d4868f1e8d0f75acf212dd6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/series"
    i4d51bab51a3d07772718202eb8db775bb66330cc9810d787996960ce3a79a025 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/imagewithwidthwithheight"
    i59a93d5d7e5c00eda39b7cf48dc032649a699ffdb1bd29131a71795be2293ef1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/format"
    i65c9db5ca9c9daa141a06172f63c930adec909b53ad532b6cded188edeb98dfe "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/image"
    i6962e66e0b17b305d7d2b1a0c2e8b86f535632b470ad80fa9b1a2c36add7ee73 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/title"
    i91a03885e524b14f1259be2b90fb2a4c1ee551d417a9b7e689e4b0ab43b49a03 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/imagewithwidth"
    i9f3c684cc34946f71c857f745d6e75ac29bebda1dfb8375fb5b124589c2ecde3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/datalabels"
    ic12e0965af410a589a2970f97a99dc3fe3cffd1a4c2db93c84d0e433e9ac4546 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/setdata"
    ida044d87f72d570fac7f79191715c65b304ef5f4b4c3da9463f7061c6cd2b9ff "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/legend"
    idd913e004bad35a74201fe41c4b1a3e70cbfdf150334fabc56c49293e156ea1a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/imagewithwidthwithheightwithfittingmode"
    ie523eef030298b9d77e396023ddb5bc10942eba30b606ac5c57370d59bffefb1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/worksheet"
    iaf085d8f220124a45599446e8798e49e294c460506a99d7a90b633fbe38af6f7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/series/item"
)

// WorkbookChartItemRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\{workbookChart-id}
type WorkbookChartItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookChartItemRequestBuilderDeleteOptions options for Delete
type WorkbookChartItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookChartItemRequestBuilderGetOptions options for Get
type WorkbookChartItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookChartItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookChartItemRequestBuilderGetQueryParameters returns collection of charts that are part of the worksheet. Read-only.
type WorkbookChartItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookChartItemRequestBuilderPatchOptions options for Patch
type WorkbookChartItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WorkbookChartItemRequestBuilder) Axes()(*i25cf65d8e047c204308390327f2e23e310366ff2da03d485b9ef4cb40794620a.AxesRequestBuilder) {
    return i25cf65d8e047c204308390327f2e23e310366ff2da03d485b9ef4cb40794620a.NewAxesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookChartItemRequestBuilderInternal instantiates a new WorkbookChartItemRequestBuilder and sets the default values.
func NewWorkbookChartItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartItemRequestBuilder) {
    m := &WorkbookChartItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet/charts/{workbookChart_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookChartItemRequestBuilder instantiates a new WorkbookChartItemRequestBuilder and sets the default values.
func NewWorkbookChartItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookChartItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartItemRequestBuilder) CreateDeleteRequestInformation(options *WorkbookChartItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartItemRequestBuilder) CreateGetRequestInformation(options *WorkbookChartItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartItemRequestBuilder) CreatePatchRequestInformation(options *WorkbookChartItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookChartItemRequestBuilder) DataLabels()(*i9f3c684cc34946f71c857f745d6e75ac29bebda1dfb8375fb5b124589c2ecde3.DataLabelsRequestBuilder) {
    return i9f3c684cc34946f71c857f745d6e75ac29bebda1dfb8375fb5b124589c2ecde3.NewDataLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartItemRequestBuilder) Delete(options *WorkbookChartItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookChartItemRequestBuilder) Format()(*i59a93d5d7e5c00eda39b7cf48dc032649a699ffdb1bd29131a71795be2293ef1.FormatRequestBuilder) {
    return i59a93d5d7e5c00eda39b7cf48dc032649a699ffdb1bd29131a71795be2293ef1.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartItemRequestBuilder) Get(options *WorkbookChartItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookChart() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart), nil
}
// Image builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image()
func (m *WorkbookChartItemRequestBuilder) Image()(*i65c9db5ca9c9daa141a06172f63c930adec909b53ad532b6cded188edeb98dfe.ImageRequestBuilder) {
    return i65c9db5ca9c9daa141a06172f63c930adec909b53ad532b6cded188edeb98dfe.NewImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImageWithWidth builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width})
func (m *WorkbookChartItemRequestBuilder) ImageWithWidth(width *int32)(*i91a03885e524b14f1259be2b90fb2a4c1ee551d417a9b7e689e4b0ab43b49a03.ImageWithWidthRequestBuilder) {
    return i91a03885e524b14f1259be2b90fb2a4c1ee551d417a9b7e689e4b0ab43b49a03.NewImageWithWidthRequestBuilderInternal(m.pathParameters, m.requestAdapter, width);
}
// ImageWithWidthWithHeight builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width},height={height})
func (m *WorkbookChartItemRequestBuilder) ImageWithWidthWithHeight(width *int32, height *int32)(*i4d51bab51a3d07772718202eb8db775bb66330cc9810d787996960ce3a79a025.ImageWithWidthWithHeightRequestBuilder) {
    return i4d51bab51a3d07772718202eb8db775bb66330cc9810d787996960ce3a79a025.NewImageWithWidthWithHeightRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height);
}
// ImageWithWidthWithHeightWithFittingMode builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width},height={height},fittingMode='{fittingMode}')
func (m *WorkbookChartItemRequestBuilder) ImageWithWidthWithHeightWithFittingMode(width *int32, height *int32, fittingMode *string)(*idd913e004bad35a74201fe41c4b1a3e70cbfdf150334fabc56c49293e156ea1a.ImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return idd913e004bad35a74201fe41c4b1a3e70cbfdf150334fabc56c49293e156ea1a.NewImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height, fittingMode);
}
func (m *WorkbookChartItemRequestBuilder) Legend()(*ida044d87f72d570fac7f79191715c65b304ef5f4b4c3da9463f7061c6cd2b9ff.LegendRequestBuilder) {
    return ida044d87f72d570fac7f79191715c65b304ef5f4b4c3da9463f7061c6cd2b9ff.NewLegendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartItemRequestBuilder) Patch(options *WorkbookChartItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookChartItemRequestBuilder) Series()(*i2cb7996ef32859e96e272e57a9595c7b9bb3bb056d4868f1e8d0f75acf212dd6.SeriesRequestBuilder) {
    return i2cb7996ef32859e96e272e57a9595c7b9bb3bb056d4868f1e8d0f75acf212dd6.NewSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SeriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.tables.item.worksheet.charts.item.series.item collection
func (m *WorkbookChartItemRequestBuilder) SeriesById(id string)(*iaf085d8f220124a45599446e8798e49e294c460506a99d7a90b633fbe38af6f7.WorkbookChartSeriesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartSeries_id"] = id
    }
    return iaf085d8f220124a45599446e8798e49e294c460506a99d7a90b633fbe38af6f7.NewWorkbookChartSeriesItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookChartItemRequestBuilder) SetData()(*ic12e0965af410a589a2970f97a99dc3fe3cffd1a4c2db93c84d0e433e9ac4546.SetDataRequestBuilder) {
    return ic12e0965af410a589a2970f97a99dc3fe3cffd1a4c2db93c84d0e433e9ac4546.NewSetDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartItemRequestBuilder) SetPosition()(*i07c4bbb38bd9316f4c3b1d3c547b665236f9b96d10d86c393ac1830072b24391.SetPositionRequestBuilder) {
    return i07c4bbb38bd9316f4c3b1d3c547b665236f9b96d10d86c393ac1830072b24391.NewSetPositionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartItemRequestBuilder) Title()(*i6962e66e0b17b305d7d2b1a0c2e8b86f535632b470ad80fa9b1a2c36add7ee73.TitleRequestBuilder) {
    return i6962e66e0b17b305d7d2b1a0c2e8b86f535632b470ad80fa9b1a2c36add7ee73.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartItemRequestBuilder) Worksheet()(*ie523eef030298b9d77e396023ddb5bc10942eba30b606ac5c57370d59bffefb1.WorksheetRequestBuilder) {
    return ie523eef030298b9d77e396023ddb5bc10942eba30b606ac5c57370d59bffefb1.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
