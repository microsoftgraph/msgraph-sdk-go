package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i185fb08f41a161b30838e05c0bf932d3d0e8be18c6614c5738399c5c9a98fef8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/series/item/format"
    i2f5cc3fe40966b3534171efab13b74fa6b75918477a3fa311b44472552f6fa28 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/series/item/points"
    i6b96d5bb3da1ad5b185b5a6951fc4e44acf0a170feb6b01bd463de5948e1c130 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/series/item/points/item"
)

// WorkbookChartSeriesRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\series\{workbookChartSeries-id}
type WorkbookChartSeriesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookChartSeriesRequestBuilderDeleteOptions options for Delete
type WorkbookChartSeriesRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookChartSeriesRequestBuilderGetOptions options for Get
type WorkbookChartSeriesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookChartSeriesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookChartSeriesRequestBuilderGetQueryParameters represents either a single series or collection of series in the chart. Read-only.
type WorkbookChartSeriesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookChartSeriesRequestBuilderPatchOptions options for Patch
type WorkbookChartSeriesRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChartSeries;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWorkbookChartSeriesRequestBuilderInternal instantiates a new WorkbookChartSeriesRequestBuilder and sets the default values.
func NewWorkbookChartSeriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartSeriesRequestBuilder) {
    m := &WorkbookChartSeriesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/charts/{workbookChart_id}/series/{workbookChartSeries_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookChartSeriesRequestBuilder instantiates a new WorkbookChartSeriesRequestBuilder and sets the default values.
func NewWorkbookChartSeriesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartSeriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookChartSeriesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) CreateDeleteRequestInformation(options *WorkbookChartSeriesRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) CreateGetRequestInformation(options *WorkbookChartSeriesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) CreatePatchRequestInformation(options *WorkbookChartSeriesRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) Delete(options *WorkbookChartSeriesRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookChartSeriesRequestBuilder) Format()(*i185fb08f41a161b30838e05c0bf932d3d0e8be18c6614c5738399c5c9a98fef8.FormatRequestBuilder) {
    return i185fb08f41a161b30838e05c0bf932d3d0e8be18c6614c5738399c5c9a98fef8.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) Get(options *WorkbookChartSeriesRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChartSeries, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookChartSeries() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChartSeries), nil
}
// Patch represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) Patch(options *WorkbookChartSeriesRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookChartSeriesRequestBuilder) Points()(*i2f5cc3fe40966b3534171efab13b74fa6b75918477a3fa311b44472552f6fa28.PointsRequestBuilder) {
    return i2f5cc3fe40966b3534171efab13b74fa6b75918477a3fa311b44472552f6fa28.NewPointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PointsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item.charts.item.series.item.points.item collection
func (m *WorkbookChartSeriesRequestBuilder) PointsById(id string)(*i6b96d5bb3da1ad5b185b5a6951fc4e44acf0a170feb6b01bd463de5948e1c130.WorkbookChartPointRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartPoint_id"] = id
    }
    return i6b96d5bb3da1ad5b185b5a6951fc4e44acf0a170feb6b01bd463de5948e1c130.NewWorkbookChartPointRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
