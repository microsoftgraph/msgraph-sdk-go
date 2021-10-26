package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i417f2f22343fff332d5faa2a4cf128476b7ae0389d00d126b84d861a813bafdf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/headerrowrange"
    i5fd7d490a985a4da1cf7bc57b83e554d25057b81e511484d693f4c9d8a78723e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/range_escaped"
    i85372ccc5bf5004cf12c66e8057ebebecedee646133b74c50d2b8fd5fae27608 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/databodyrange"
    ib31774a7d3dfb4d38f0650300e458faa3dec158286b2e28886298a1c243592f6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter"
    ib58854552ea86887cd047323385dcb4323356cd252b268caa0de8f5235801948 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/totalrowrange"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}
type WorkbookTableColumnRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Represents a collection of all the columns in the table. Read-only.
type WorkbookTableColumnRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new WorkbookTableColumnRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookTableColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableColumnRequestBuilder) {
    m := &WorkbookTableColumnRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/tables/{workbookTable_id}/columns/{workbookTableColumn_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorkbookTableColumnRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookTableColumnRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableColumnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableColumnRequestBuilderInternal(urlParams, requestAdapter)
}
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *WorkbookTableColumnRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *WorkbookTableColumnRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookTableColumnRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookTableColumnRequestBuilderGetQueryParameters)
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
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *WorkbookTableColumnRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTableColumn, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableColumnRequestBuilder) DataBodyRange()(*i85372ccc5bf5004cf12c66e8057ebebecedee646133b74c50d2b8fd5fae27608.DataBodyRangeRequestBuilder) {
    return i85372ccc5bf5004cf12c66e8057ebebecedee646133b74c50d2b8fd5fae27608.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookTableColumnRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookTableColumnRequestBuilder) Filter()(*ib31774a7d3dfb4d38f0650300e458faa3dec158286b2e28886298a1c243592f6.FilterRequestBuilder) {
    return ib31774a7d3dfb4d38f0650300e458faa3dec158286b2e28886298a1c243592f6.NewFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookTableColumnRequestBuilder) Get(q func (value *WorkbookTableColumnRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTableColumn, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookTableColumn() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTableColumn), nil
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableColumnRequestBuilder) HeaderRowRange()(*i417f2f22343fff332d5faa2a4cf128476b7ae0389d00d126b84d861a813bafdf.HeaderRowRangeRequestBuilder) {
    return i417f2f22343fff332d5faa2a4cf128476b7ae0389d00d126b84d861a813bafdf.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookTableColumnRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTableColumn, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.range()
func (m *WorkbookTableColumnRequestBuilder) Range_escaped()(*i5fd7d490a985a4da1cf7bc57b83e554d25057b81e511484d693f4c9d8a78723e.RangeRequestBuilder) {
    return i5fd7d490a985a4da1cf7bc57b83e554d25057b81e511484d693f4c9d8a78723e.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableColumnRequestBuilder) TotalRowRange()(*ib58854552ea86887cd047323385dcb4323356cd252b268caa0de8f5235801948.TotalRowRangeRequestBuilder) {
    return ib58854552ea86887cd047323385dcb4323356cd252b268caa0de8f5235801948.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
