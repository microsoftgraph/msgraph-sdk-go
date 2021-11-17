package charts

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0a27fe160b2cc327adba73524662111269acb368a1e7d859e8a6bc3863486ee1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/count"
    i2432afdbd6c0af6696ff1672e3484ffcf1d37f4e5fbe6e57c7a271e80ba1ff18 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/itematwithindex"
    i476016513e8840bcbc27a43002b9d466702f0cdea2dd3cfe62ab297454d274d2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/add"
    i62b30967630875e1caf62d791333a254430f87c5cfa0fb4191003972602cca4e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/itemwithname"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts
type ChartsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type ChartsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChartsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Returns collection of charts that are part of the worksheet. Read-only.
type ChartsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Options for Post
type ChartsRequestBuilderPostOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ChartsRequestBuilder) Add()(*i476016513e8840bcbc27a43002b9d466702f0cdea2dd3cfe62ab297454d274d2.AddRequestBuilder) {
    return i476016513e8840bcbc27a43002b9d466702f0cdea2dd3cfe62ab297454d274d2.NewAddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ChartsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChartsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChartsRequestBuilder) {
    m := &ChartsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/charts{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ChartsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChartsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChartsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChartsRequestBuilderInternal(urlParams, requestAdapter)
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\microsoft.graph.count()
func (m *ChartsRequestBuilder) Count()(*i0a27fe160b2cc327adba73524662111269acb368a1e7d859e8a6bc3863486ee1.CountRequestBuilder) {
    return i0a27fe160b2cc327adba73524662111269acb368a1e7d859e8a6bc3863486ee1.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ChartsRequestBuilder) CreateGetRequestInformation(options *ChartsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ChartsRequestBuilder) CreatePostRequestInformation(options *ChartsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ChartsRequestBuilder) Get(options *ChartsRequestBuilderGetOptions)(*ChartsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChartsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ChartsResponse), nil
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\microsoft.graph.itemAt(index={index})
// Parameters:
//  - index : Usage: index={index}
func (m *ChartsRequestBuilder) ItemAtWithIndex(index *int32)(*i2432afdbd6c0af6696ff1672e3484ffcf1d37f4e5fbe6e57c7a271e80ba1ff18.ItemAtWithIndexRequestBuilder) {
    return i2432afdbd6c0af6696ff1672e3484ffcf1d37f4e5fbe6e57c7a271e80ba1ff18.NewItemAtWithIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter, index);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\microsoft.graph.item(name='{name}')
// Parameters:
//  - name : Usage: name={name}
func (m *ChartsRequestBuilder) ItemWithName(name *string)(*i62b30967630875e1caf62d791333a254430f87c5cfa0fb4191003972602cca4e.ItemWithNameRequestBuilder) {
    return i62b30967630875e1caf62d791333a254430f87c5cfa0fb4191003972602cca4e.NewItemWithNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, name);
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ChartsRequestBuilder) Post(options *ChartsRequestBuilderPostOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookChart() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart), nil
}
