package charts

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i350ab6d962bbc4771226932eb9316a81de3fa793ff6bba0bbc11fedea9ab01ca "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/itemwithname"
    i4eda0d7006fe63e4157b5f6581a412cac8f0e666eac58b7243ab20e9a64662af "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/add"
    ic7726546589b142e02e65ac275d1115f2d5bfc64e6e571013e1eb18144d10bda "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/count"
    ied7b61575feaa2e13a72ce52d62b8a82b8bb62ec4ff1bfa69cb45de337dc7c7b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/itematwithindex"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts
type ChartsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Returns collection of charts that are part of the worksheet. Read-only.
type ChartsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
func (m *ChartsRequestBuilder) Add()(*i4eda0d7006fe63e4157b5f6581a412cac8f0e666eac58b7243ab20e9a64662af.AddRequestBuilder) {
    return i4eda0d7006fe63e4157b5f6581a412cac8f0e666eac58b7243ab20e9a64662af.NewAddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ChartsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChartsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChartsRequestBuilder) {
    m := &ChartsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/charts{?top,skip,search,filter,count,orderby,select,expand}";
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\microsoft.graph.count()
func (m *ChartsRequestBuilder) Count()(*ic7726546589b142e02e65ac275d1115f2d5bfc64e6e571013e1eb18144d10bda.CountRequestBuilder) {
    return ic7726546589b142e02e65ac275d1115f2d5bfc64e6e571013e1eb18144d10bda.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ChartsRequestBuilder) CreateGetRequestInformation(q func (value *ChartsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ChartsRequestBuilderGetQueryParameters)
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
func (m *ChartsRequestBuilder) CreatePostRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ChartsRequestBuilder) Get(q func (value *ChartsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*ChartsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChartsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*ChartsResponse), nil
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\microsoft.graph.itemAt(index={index})
// Parameters:
//  - index : Usage: index={index}
func (m *ChartsRequestBuilder) ItemAtWithIndex(index *int32)(*ied7b61575feaa2e13a72ce52d62b8a82b8bb62ec4ff1bfa69cb45de337dc7c7b.ItemAtWithIndexRequestBuilder) {
    return ied7b61575feaa2e13a72ce52d62b8a82b8bb62ec4ff1bfa69cb45de337dc7c7b.NewItemAtWithIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter, index);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\microsoft.graph.item(name='{name}')
// Parameters:
//  - name : Usage: name={name}
func (m *ChartsRequestBuilder) ItemWithName(name *string)(*i350ab6d962bbc4771226932eb9316a81de3fa793ff6bba0bbc11fedea9ab01ca.ItemWithNameRequestBuilder) {
    return i350ab6d962bbc4771226932eb9316a81de3fa793ff6bba0bbc11fedea9ab01ca.NewItemWithNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, name);
}
// Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ChartsRequestBuilder) Post(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookChart() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookChart), nil
}
