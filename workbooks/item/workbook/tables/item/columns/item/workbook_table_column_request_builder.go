package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3c120d46a4d11e47149cfb1ddec041bc8d2e8c51e13907f4a0d57087cbef3d6d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/totalrowrange"
    i54f19cf38ed40d26b1e52134c586733a2a3d0b3952000975fec091576106c3e6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter"
    i9ae6712c6e9e104af266f1e20d1800b94454640530be83ee79e59d11b52b60f3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/databodyrange"
    icd831b29ae1fc7c749712141fecdca6f6c69a244d8fd66a0be9f21880e152d6d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/headerrowrange"
    iecbd4b151a3cae3f0062357a5baa64c3b6782a57e0a6059726627b21e2b6ac5b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/range_escaped"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}
type WorkbookTableColumnRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type WorkbookTableColumnRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// Represents a collection of all the columns in the table. Read-only.
type WorkbookTableColumnRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
// Instantiates a new WorkbookTableColumnRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookTableColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableColumnRequestBuilder) {
    m := &WorkbookTableColumnRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/columns/{workbookTableColumn_id}{?select,expand}";
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
//  - options : Options for the request
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
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - options : Options for the request
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
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - options : Options for the request
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableColumnRequestBuilder) DataBodyRange()(*i9ae6712c6e9e104af266f1e20d1800b94454640530be83ee79e59d11b52b60f3.DataBodyRangeRequestBuilder) {
    return i9ae6712c6e9e104af266f1e20d1800b94454640530be83ee79e59d11b52b60f3.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - options : Options for the request
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
func (m *WorkbookTableColumnRequestBuilder) Filter()(*i54f19cf38ed40d26b1e52134c586733a2a3d0b3952000975fec091576106c3e6.FilterRequestBuilder) {
    return i54f19cf38ed40d26b1e52134c586733a2a3d0b3952000975fec091576106c3e6.NewFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - options : Options for the request
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableColumnRequestBuilder) HeaderRowRange()(*icd831b29ae1fc7c749712141fecdca6f6c69a244d8fd66a0be9f21880e152d6d.HeaderRowRangeRequestBuilder) {
    return icd831b29ae1fc7c749712141fecdca6f6c69a244d8fd66a0be9f21880e152d6d.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - options : Options for the request
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.range()
func (m *WorkbookTableColumnRequestBuilder) Range_escaped()(*iecbd4b151a3cae3f0062357a5baa64c3b6782a57e0a6059726627b21e2b6ac5b.RangeRequestBuilder) {
    return iecbd4b151a3cae3f0062357a5baa64c3b6782a57e0a6059726627b21e2b6ac5b.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableColumnRequestBuilder) TotalRowRange()(*i3c120d46a4d11e47149cfb1ddec041bc8d2e8c51e13907f4a0d57087cbef3d6d.TotalRowRangeRequestBuilder) {
    return i3c120d46a4d11e47149cfb1ddec041bc8d2e8c51e13907f4a0d57087cbef3d6d.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
