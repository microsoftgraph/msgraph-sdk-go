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

// WorkbookTableColumnRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}
type WorkbookTableColumnRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookTableColumnRequestBuilderDeleteOptions options for Delete
type WorkbookTableColumnRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookTableColumnRequestBuilderGetOptions options for Get
type WorkbookTableColumnRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookTableColumnRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookTableColumnRequestBuilderGetQueryParameters represents a collection of all the columns in the table. Read-only.
type WorkbookTableColumnRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookTableColumnRequestBuilderPatchOptions options for Patch
type WorkbookTableColumnRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTableColumn;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWorkbookTableColumnRequestBuilderInternal instantiates a new WorkbookTableColumnRequestBuilder and sets the default values.
func NewWorkbookTableColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableColumnRequestBuilder) {
    m := &WorkbookTableColumnRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/tables/{workbookTable_id}/columns/{workbookTableColumn_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookTableColumnRequestBuilder instantiates a new WorkbookTableColumnRequestBuilder and sets the default values.
func NewWorkbookTableColumnRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableColumnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableColumnRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) CreateDeleteRequestInformation(options *WorkbookTableColumnRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) CreateGetRequestInformation(options *WorkbookTableColumnRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) CreatePatchRequestInformation(options *WorkbookTableColumnRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// DataBodyRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableColumnRequestBuilder) DataBodyRange()(*i85372ccc5bf5004cf12c66e8057ebebecedee646133b74c50d2b8fd5fae27608.DataBodyRangeRequestBuilder) {
    return i85372ccc5bf5004cf12c66e8057ebebecedee646133b74c50d2b8fd5fae27608.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) Delete(options *WorkbookTableColumnRequestBuilderDeleteOptions)(error) {
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
func (m *WorkbookTableColumnRequestBuilder) Filter()(*ib31774a7d3dfb4d38f0650300e458faa3dec158286b2e28886298a1c243592f6.FilterRequestBuilder) {
    return ib31774a7d3dfb4d38f0650300e458faa3dec158286b2e28886298a1c243592f6.NewFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) Get(options *WorkbookTableColumnRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTableColumn, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookTableColumn() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTableColumn), nil
}
// HeaderRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableColumnRequestBuilder) HeaderRowRange()(*i417f2f22343fff332d5faa2a4cf128476b7ae0389d00d126b84d861a813bafdf.HeaderRowRangeRequestBuilder) {
    return i417f2f22343fff332d5faa2a4cf128476b7ae0389d00d126b84d861a813bafdf.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTableColumnRequestBuilder) Patch(options *WorkbookTableColumnRequestBuilderPatchOptions)(error) {
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.range()
func (m *WorkbookTableColumnRequestBuilder) Range_escaped()(*i5fd7d490a985a4da1cf7bc57b83e554d25057b81e511484d693f4c9d8a78723e.RangeRequestBuilder) {
    return i5fd7d490a985a4da1cf7bc57b83e554d25057b81e511484d693f4c9d8a78723e.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TotalRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableColumnRequestBuilder) TotalRowRange()(*ib58854552ea86887cd047323385dcb4323356cd252b268caa0de8f5235801948.TotalRowRangeRequestBuilder) {
    return ib58854552ea86887cd047323385dcb4323356cd252b268caa0de8f5235801948.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
