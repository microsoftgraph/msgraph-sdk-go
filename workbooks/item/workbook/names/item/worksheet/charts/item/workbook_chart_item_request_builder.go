package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i00e8b29c40b942b88edeefb77f583774c2e2eb092dcf6a28dce32e2d1fe4fc0c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/imagewithwidthwithheightwithfittingmode"
    i06c195de7e7d7504e8d2fd06d3d06b7832c23410d54911f6d79c0fc5f5344181 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/image"
    i0f12a52f5b069a5e287b9786908af85e0c37a7d21e2efee92a13a48ce673f620 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/worksheet"
    i190fc68794fd31b35481b18576d02094cecc66fa9e10e2054b9c0ce732679777 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/format"
    i79d4b6ff4e1095059dfbc737d00e85c022a1ca498b3ce6a2f51805b28ca6f2e9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/setdata"
    i803b11d6b67f5bfb7df64e669c9d9de29285ba27ad00ae2df93138499c96aaff "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/imagewithwidthwithheight"
    i8ede4b1ea544372cbc8cf6bebff454d8f704a58cb230d686887e0448d25f02c4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/setposition"
    i9ab643d4316779d9d3e7f2a4601d1474677e3dde7e71141f93827ca672ab09ae "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/datalabels"
    ibdb04c205108790f70d735ec7e42c4bb0b32eda884aeb88a304e1a8d428b1972 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/imagewithwidth"
    icf0de94eab5eedf9f84adb5d3d336bd077ac4f85385efb6dfc9e6c1e5a6c41a9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/legend"
    id155a9459c408272b4278fe60218f0fba19a3d42c64913a174b3eeadc5250ee9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes"
    ieb36f4a14687ad9ba54ea3f7fbe801dbe8d51f9907fedceac961358074b0e234 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/title"
    if6c63be159d2471eea0c0924705e2ad59adb302e3ca5fe0eebe61fdca03392e4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/series"
    i28bf140701aeebdde7c593f39208093698c7f5d1b66bc90e8a19aaad06b16b21 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/series/item"
)

// WorkbookChartItemRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}
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
func (m *WorkbookChartItemRequestBuilder) Axes()(*id155a9459c408272b4278fe60218f0fba19a3d42c64913a174b3eeadc5250ee9.AxesRequestBuilder) {
    return id155a9459c408272b4278fe60218f0fba19a3d42c64913a174b3eeadc5250ee9.NewAxesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookChartItemRequestBuilderInternal instantiates a new WorkbookChartItemRequestBuilder and sets the default values.
func NewWorkbookChartItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartItemRequestBuilder) {
    m := &WorkbookChartItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/charts/{workbookChart_id}{?select,expand}";
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
func (m *WorkbookChartItemRequestBuilder) DataLabels()(*i9ab643d4316779d9d3e7f2a4601d1474677e3dde7e71141f93827ca672ab09ae.DataLabelsRequestBuilder) {
    return i9ab643d4316779d9d3e7f2a4601d1474677e3dde7e71141f93827ca672ab09ae.NewDataLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *WorkbookChartItemRequestBuilder) Format()(*i190fc68794fd31b35481b18576d02094cecc66fa9e10e2054b9c0ce732679777.FormatRequestBuilder) {
    return i190fc68794fd31b35481b18576d02094cecc66fa9e10e2054b9c0ce732679777.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Image builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image()
func (m *WorkbookChartItemRequestBuilder) Image()(*i06c195de7e7d7504e8d2fd06d3d06b7832c23410d54911f6d79c0fc5f5344181.ImageRequestBuilder) {
    return i06c195de7e7d7504e8d2fd06d3d06b7832c23410d54911f6d79c0fc5f5344181.NewImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImageWithWidth builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width})
func (m *WorkbookChartItemRequestBuilder) ImageWithWidth(width *int32)(*ibdb04c205108790f70d735ec7e42c4bb0b32eda884aeb88a304e1a8d428b1972.ImageWithWidthRequestBuilder) {
    return ibdb04c205108790f70d735ec7e42c4bb0b32eda884aeb88a304e1a8d428b1972.NewImageWithWidthRequestBuilderInternal(m.pathParameters, m.requestAdapter, width);
}
// ImageWithWidthWithHeight builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width},height={height})
func (m *WorkbookChartItemRequestBuilder) ImageWithWidthWithHeight(width *int32, height *int32)(*i803b11d6b67f5bfb7df64e669c9d9de29285ba27ad00ae2df93138499c96aaff.ImageWithWidthWithHeightRequestBuilder) {
    return i803b11d6b67f5bfb7df64e669c9d9de29285ba27ad00ae2df93138499c96aaff.NewImageWithWidthWithHeightRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height);
}
// ImageWithWidthWithHeightWithFittingMode builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width},height={height},fittingMode='{fittingMode}')
func (m *WorkbookChartItemRequestBuilder) ImageWithWidthWithHeightWithFittingMode(width *int32, height *int32, fittingMode *string)(*i00e8b29c40b942b88edeefb77f583774c2e2eb092dcf6a28dce32e2d1fe4fc0c.ImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return i00e8b29c40b942b88edeefb77f583774c2e2eb092dcf6a28dce32e2d1fe4fc0c.NewImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height, fittingMode);
}
func (m *WorkbookChartItemRequestBuilder) Legend()(*icf0de94eab5eedf9f84adb5d3d336bd077ac4f85385efb6dfc9e6c1e5a6c41a9.LegendRequestBuilder) {
    return icf0de94eab5eedf9f84adb5d3d336bd077ac4f85385efb6dfc9e6c1e5a6c41a9.NewLegendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *WorkbookChartItemRequestBuilder) Series()(*if6c63be159d2471eea0c0924705e2ad59adb302e3ca5fe0eebe61fdca03392e4.SeriesRequestBuilder) {
    return if6c63be159d2471eea0c0924705e2ad59adb302e3ca5fe0eebe61fdca03392e4.NewSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SeriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.names.item.worksheet.charts.item.series.item collection
func (m *WorkbookChartItemRequestBuilder) SeriesById(id string)(*i28bf140701aeebdde7c593f39208093698c7f5d1b66bc90e8a19aaad06b16b21.WorkbookChartSeriesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartSeries_id"] = id
    }
    return i28bf140701aeebdde7c593f39208093698c7f5d1b66bc90e8a19aaad06b16b21.NewWorkbookChartSeriesItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookChartItemRequestBuilder) SetData()(*i79d4b6ff4e1095059dfbc737d00e85c022a1ca498b3ce6a2f51805b28ca6f2e9.SetDataRequestBuilder) {
    return i79d4b6ff4e1095059dfbc737d00e85c022a1ca498b3ce6a2f51805b28ca6f2e9.NewSetDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartItemRequestBuilder) SetPosition()(*i8ede4b1ea544372cbc8cf6bebff454d8f704a58cb230d686887e0448d25f02c4.SetPositionRequestBuilder) {
    return i8ede4b1ea544372cbc8cf6bebff454d8f704a58cb230d686887e0448d25f02c4.NewSetPositionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartItemRequestBuilder) Title()(*ieb36f4a14687ad9ba54ea3f7fbe801dbe8d51f9907fedceac961358074b0e234.TitleRequestBuilder) {
    return ieb36f4a14687ad9ba54ea3f7fbe801dbe8d51f9907fedceac961358074b0e234.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartItemRequestBuilder) Worksheet()(*i0f12a52f5b069a5e287b9786908af85e0c37a7d21e2efee92a13a48ce673f620.WorksheetRequestBuilder) {
    return i0f12a52f5b069a5e287b9786908af85e0c37a7d21e2efee92a13a48ce673f620.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
