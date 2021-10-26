package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i055d8a318e859fd25af71635255a303d5dfdee609eafb446b05814c8061989a5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/format"
    i3256fa41856ae702abd2e3394e4e3e397692ca3a8395bcbf514818029e267a3b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/datalabels"
    i55e86de6a1e2a3a70a7c241d9c10bf1fa49ab588f242d5a0cc5a298534e8c949 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/imagewithwidthwithheightwithfittingmode"
    i58a04b78818648cf7c8a7ab65f8df4551f0847340aa3932086d767021047bc62 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/title"
    i6867a9f2d99ed64479a33b179d5fed6dd7e15fa6a426efb7beb418ac4166a0cd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/setdata"
    i85570cc44db9719f8e7e52d250d145a46c8fe80c8d71ff340e6cb8447a028027 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/imagewithwidthwithheight"
    i8897f95808fc2c88f0ddf5e22d363cb43c7cf7d34319861a1f0a06bbce734098 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/series"
    i8dc75bd90205c69802ce36355504fc05aecd15925b62b67e19ca84f1981d52f3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/legend"
    i90cee6397a0770400ff5db04576558b1d5dc7dfd65bbd2c74c768ae6bd44ac1f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/axes"
    iadd7f62a2393c2ba269750786e0a00a09cee287374186025e02a53bda168e998 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/worksheet"
    ib4453c1162ab48e638072867f7e3e1b34576893ed0c10b2a4f4abe8d33b768fc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/image"
    ib4be22823f431ae2468ae39341f0bd48c2f9011d7692b26df6c6e45397fe0abc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/setposition"
    ic492b1beee78723a3aeb64d5c7f38725feb3e2ada9c801d8227c16386e20cba9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/imagewithwidth"
    i1b5e1516477c8a38cc2f74fceebc339bf483f4c58ebd3b03e5b53a8dee57e32e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/series/item"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}
type WorkbookChartRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Returns collection of charts that are part of the worksheet. Read-only.
type WorkbookChartRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *WorkbookChartRequestBuilder) Axes()(*i90cee6397a0770400ff5db04576558b1d5dc7dfd65bbd2c74c768ae6bd44ac1f.AxesRequestBuilder) {
    return i90cee6397a0770400ff5db04576558b1d5dc7dfd65bbd2c74c768ae6bd44ac1f.NewAxesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new WorkbookChartRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookChartRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartRequestBuilder) {
    m := &WorkbookChartRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/charts/{workbookChart_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorkbookChartRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookChartRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookChartRequestBuilderInternal(urlParams, requestAdapter)
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *WorkbookChartRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *WorkbookChartRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookChartRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookChartRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *WorkbookChartRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookChartRequestBuilder) DataLabels()(*i3256fa41856ae702abd2e3394e4e3e397692ca3a8395bcbf514818029e267a3b.DataLabelsRequestBuilder) {
    return i3256fa41856ae702abd2e3394e4e3e397692ca3a8395bcbf514818029e267a3b.NewDataLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookChartRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookChartRequestBuilder) Format()(*i055d8a318e859fd25af71635255a303d5dfdee609eafb446b05814c8061989a5.FormatRequestBuilder) {
    return i055d8a318e859fd25af71635255a303d5dfdee609eafb446b05814c8061989a5.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookChartRequestBuilder) Get(q func (value *WorkbookChartRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookChart() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart), nil
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\microsoft.graph.image()
func (m *WorkbookChartRequestBuilder) Image()(*ib4453c1162ab48e638072867f7e3e1b34576893ed0c10b2a4f4abe8d33b768fc.ImageRequestBuilder) {
    return ib4453c1162ab48e638072867f7e3e1b34576893ed0c10b2a4f4abe8d33b768fc.NewImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\microsoft.graph.image(width={width})
// Parameters:
//  - width : Usage: width={width}
func (m *WorkbookChartRequestBuilder) ImageWithWidth(width *int32)(*ic492b1beee78723a3aeb64d5c7f38725feb3e2ada9c801d8227c16386e20cba9.ImageWithWidthRequestBuilder) {
    return ic492b1beee78723a3aeb64d5c7f38725feb3e2ada9c801d8227c16386e20cba9.NewImageWithWidthRequestBuilderInternal(m.pathParameters, m.requestAdapter, width);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\microsoft.graph.image(width={width},height={height})
// Parameters:
//  - height : Usage: height={height}
//  - width : Usage: width={width}
func (m *WorkbookChartRequestBuilder) ImageWithWidthWithHeight(width *int32, height *int32)(*i85570cc44db9719f8e7e52d250d145a46c8fe80c8d71ff340e6cb8447a028027.ImageWithWidthWithHeightRequestBuilder) {
    return i85570cc44db9719f8e7e52d250d145a46c8fe80c8d71ff340e6cb8447a028027.NewImageWithWidthWithHeightRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\microsoft.graph.image(width={width},height={height},fittingMode='{fittingMode}')
// Parameters:
//  - fittingMode : Usage: fittingMode={fittingMode}
//  - height : Usage: height={height}
//  - width : Usage: width={width}
func (m *WorkbookChartRequestBuilder) ImageWithWidthWithHeightWithFittingMode(width *int32, height *int32, fittingMode *string)(*i55e86de6a1e2a3a70a7c241d9c10bf1fa49ab588f242d5a0cc5a298534e8c949.ImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return i55e86de6a1e2a3a70a7c241d9c10bf1fa49ab588f242d5a0cc5a298534e8c949.NewImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height, fittingMode);
}
func (m *WorkbookChartRequestBuilder) Legend()(*i8dc75bd90205c69802ce36355504fc05aecd15925b62b67e19ca84f1981d52f3.LegendRequestBuilder) {
    return i8dc75bd90205c69802ce36355504fc05aecd15925b62b67e19ca84f1981d52f3.NewLegendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookChartRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookChartRequestBuilder) Series()(*i8897f95808fc2c88f0ddf5e22d363cb43c7cf7d34319861a1f0a06bbce734098.SeriesRequestBuilder) {
    return i8897f95808fc2c88f0ddf5e22d363cb43c7cf7d34319861a1f0a06bbce734098.NewSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item.charts.item.series.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookChartRequestBuilder) SeriesById(id string)(*i1b5e1516477c8a38cc2f74fceebc339bf483f4c58ebd3b03e5b53a8dee57e32e.WorkbookChartSeriesRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartSeries_id"] = id
    }
    return i1b5e1516477c8a38cc2f74fceebc339bf483f4c58ebd3b03e5b53a8dee57e32e.NewWorkbookChartSeriesRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SetData()(*i6867a9f2d99ed64479a33b179d5fed6dd7e15fa6a426efb7beb418ac4166a0cd.SetDataRequestBuilder) {
    return i6867a9f2d99ed64479a33b179d5fed6dd7e15fa6a426efb7beb418ac4166a0cd.NewSetDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SetPosition()(*ib4be22823f431ae2468ae39341f0bd48c2f9011d7692b26df6c6e45397fe0abc.SetPositionRequestBuilder) {
    return ib4be22823f431ae2468ae39341f0bd48c2f9011d7692b26df6c6e45397fe0abc.NewSetPositionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Title()(*i58a04b78818648cf7c8a7ab65f8df4551f0847340aa3932086d767021047bc62.TitleRequestBuilder) {
    return i58a04b78818648cf7c8a7ab65f8df4551f0847340aa3932086d767021047bc62.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Worksheet()(*iadd7f62a2393c2ba269750786e0a00a09cee287374186025e02a53bda168e998.WorksheetRequestBuilder) {
    return iadd7f62a2393c2ba269750786e0a00a09cee287374186025e02a53bda168e998.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
