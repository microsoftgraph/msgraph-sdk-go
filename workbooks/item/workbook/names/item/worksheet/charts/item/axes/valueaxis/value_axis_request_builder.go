package valueaxis

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0f3b263924e6f63c3c7c7d197e332dce64b134b16077bfccddf41d4d2edd8ac5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/valueaxis/title"
    i63c442ed4d988e897a2c67fd376140d5a305d55dad41d9fab61c53819eb8b462 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/valueaxis/majorgridlines"
    i7aa424fd0b92b4c1be60dfb61c78530efeea11fc69faa34c818111488e87a7ef "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/valueaxis/format"
    id19812d3c15d67ca408c6d22bae95718befae2619fc35b133ece3b47e7024ca9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/valueaxis/minorgridlines"
)

// ValueAxisRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\axes\valueAxis
type ValueAxisRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ValueAxisRequestBuilderDeleteOptions options for Delete
type ValueAxisRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ValueAxisRequestBuilderGetOptions options for Get
type ValueAxisRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ValueAxisRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ValueAxisRequestBuilderGetQueryParameters represents the value axis in an axis. Read-only.
type ValueAxisRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ValueAxisRequestBuilderPatchOptions options for Patch
type ValueAxisRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChartAxis;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewValueAxisRequestBuilderInternal instantiates a new ValueAxisRequestBuilder and sets the default values.
func NewValueAxisRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValueAxisRequestBuilder) {
    m := &ValueAxisRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/charts/{workbookChart_id}/axes/valueAxis{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewValueAxisRequestBuilder instantiates a new ValueAxisRequestBuilder and sets the default values.
func NewValueAxisRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValueAxisRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewValueAxisRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) CreateDeleteRequestInformation(options *ValueAxisRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) CreateGetRequestInformation(options *ValueAxisRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) CreatePatchRequestInformation(options *ValueAxisRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) Delete(options *ValueAxisRequestBuilderDeleteOptions)(error) {
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
func (m *ValueAxisRequestBuilder) Format()(*i7aa424fd0b92b4c1be60dfb61c78530efeea11fc69faa34c818111488e87a7ef.FormatRequestBuilder) {
    return i7aa424fd0b92b4c1be60dfb61c78530efeea11fc69faa34c818111488e87a7ef.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) Get(options *ValueAxisRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChartAxis, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookChartAxis() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChartAxis), nil
}
func (m *ValueAxisRequestBuilder) MajorGridlines()(*i63c442ed4d988e897a2c67fd376140d5a305d55dad41d9fab61c53819eb8b462.MajorGridlinesRequestBuilder) {
    return i63c442ed4d988e897a2c67fd376140d5a305d55dad41d9fab61c53819eb8b462.NewMajorGridlinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ValueAxisRequestBuilder) MinorGridlines()(*id19812d3c15d67ca408c6d22bae95718befae2619fc35b133ece3b47e7024ca9.MinorGridlinesRequestBuilder) {
    return id19812d3c15d67ca408c6d22bae95718befae2619fc35b133ece3b47e7024ca9.NewMinorGridlinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) Patch(options *ValueAxisRequestBuilderPatchOptions)(error) {
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
func (m *ValueAxisRequestBuilder) Title()(*i0f3b263924e6f63c3c7c7d197e332dce64b134b16077bfccddf41d4d2edd8ac5.TitleRequestBuilder) {
    return i0f3b263924e6f63c3c7c7d197e332dce64b134b16077bfccddf41d4d2edd8ac5.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
