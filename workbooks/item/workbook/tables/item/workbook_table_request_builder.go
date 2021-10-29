package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i064d0a112d6201dafdb8646f9ef7e5951acfad71e2f9e632ba227ddef677c6be "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/range_escaped"
    i516bf2eb8c27608c2b04ce423ef8ef76e04b9c97b2a4f5e7e3ead91d5d8b4ef3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/clearfilters"
    i678b2c47d821b5f0d3f5fe6f2245a15c54354ab8a367d46054a24327f75877ea "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns"
    i7f0a3e516ec50983722a7099a82a2669d35c41a9e9324e0b3d656bf9722ce6e8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/converttorange"
    ia60762e61ed894b0e550f8a2ae2966f24117658540a0970cba5c094b8f782673 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/headerrowrange"
    ib91060f169b8bd39146db4c07fa37532518cdd9147a5ff6e0c87a23aaa611799 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet"
    ic30e45ed9a3e16cf3233ad8b9a57ae5ea9eb9f128feb9a1cf797e693c0b0ccfd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/totalrowrange"
    icb8ff613f65f2c408b61aa652389279e4d69e53763a05b613109b8413b65d08e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/sort"
    icd5cdc5eab7434bd1b95eb7aa9d3721b201ded3217d27e6d44f478765dbc52fc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/rows"
    id2470789c713b1909ac240c12c8cd9a64f10554749b6ffdb07cf991e18c42552 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/reapplyfilters"
    ie35ed0aab635507d1cd375d72774cf81706889de9213964781ee2761f2f8c015 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/databodyrange"
    i566a5c7dda8f122b01a2ee5cf811e54ebfaaa820024445d7de321c8f73ae485d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item"
    ica9449d8925b6e3e2256625a3ef2607a71dc8f95b414fc729ef5e2761a6f56e7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/rows/item"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}
type WorkbookTableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type WorkbookTableRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type WorkbookTableRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookTableRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Represents a collection of tables associated with the workbook. Read-only.
type WorkbookTableRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type WorkbookTableRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WorkbookTableRequestBuilder) ClearFilters()(*i516bf2eb8c27608c2b04ce423ef8ef76e04b9c97b2a4f5e7e3ead91d5d8b4ef3.ClearFiltersRequestBuilder) {
    return i516bf2eb8c27608c2b04ce423ef8ef76e04b9c97b2a4f5e7e3ead91d5d8b4ef3.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Columns()(*i678b2c47d821b5f0d3f5fe6f2245a15c54354ab8a367d46054a24327f75877ea.ColumnsRequestBuilder) {
    return i678b2c47d821b5f0d3f5fe6f2245a15c54354ab8a367d46054a24327f75877ea.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.workbooks.item.workbook.tables.item.columns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookTableRequestBuilder) ColumnsById(id string)(*i566a5c7dda8f122b01a2ee5cf811e54ebfaaa820024445d7de321c8f73ae485d.WorkbookTableColumnRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableColumn_id"] = id
    }
    return i566a5c7dda8f122b01a2ee5cf811e54ebfaaa820024445d7de321c8f73ae485d.NewWorkbookTableColumnRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new WorkbookTableRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorkbookTableRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookTableRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*i7f0a3e516ec50983722a7099a82a2669d35c41a9e9324e0b3d656bf9722ce6e8.ConvertToRangeRequestBuilder) {
    return i7f0a3e516ec50983722a7099a82a2669d35c41a9e9324e0b3d656bf9722ce6e8.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of tables associated with the workbook. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) CreateDeleteRequestInformation(options *WorkbookTableRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Represents a collection of tables associated with the workbook. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) CreateGetRequestInformation(options *WorkbookTableRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Represents a collection of tables associated with the workbook. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) CreatePatchRequestInformation(options *WorkbookTableRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*ie35ed0aab635507d1cd375d72774cf81706889de9213964781ee2761f2f8c015.DataBodyRangeRequestBuilder) {
    return ie35ed0aab635507d1cd375d72774cf81706889de9213964781ee2761f2f8c015.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of tables associated with the workbook. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) Delete(options *WorkbookTableRequestBuilderDeleteOptions)(error) {
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
// Represents a collection of tables associated with the workbook. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) Get(options *WorkbookTableRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookTable() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable), nil
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*ia60762e61ed894b0e550f8a2ae2966f24117658540a0970cba5c094b8f782673.HeaderRowRangeRequestBuilder) {
    return ia60762e61ed894b0e550f8a2ae2966f24117658540a0970cba5c094b8f782673.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Represents a collection of tables associated with the workbook. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorkbookTableRequestBuilder) Patch(options *WorkbookTableRequestBuilderPatchOptions)(error) {
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\microsoft.graph.range()
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i064d0a112d6201dafdb8646f9ef7e5951acfad71e2f9e632ba227ddef677c6be.RangeRequestBuilder) {
    return i064d0a112d6201dafdb8646f9ef7e5951acfad71e2f9e632ba227ddef677c6be.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*id2470789c713b1909ac240c12c8cd9a64f10554749b6ffdb07cf991e18c42552.ReapplyFiltersRequestBuilder) {
    return id2470789c713b1909ac240c12c8cd9a64f10554749b6ffdb07cf991e18c42552.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Rows()(*icd5cdc5eab7434bd1b95eb7aa9d3721b201ded3217d27e6d44f478765dbc52fc.RowsRequestBuilder) {
    return icd5cdc5eab7434bd1b95eb7aa9d3721b201ded3217d27e6d44f478765dbc52fc.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.workbooks.item.workbook.tables.item.rows.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookTableRequestBuilder) RowsById(id string)(*ica9449d8925b6e3e2256625a3ef2607a71dc8f95b414fc729ef5e2761a6f56e7.WorkbookTableRowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableRow_id"] = id
    }
    return ica9449d8925b6e3e2256625a3ef2607a71dc8f95b414fc729ef5e2761a6f56e7.NewWorkbookTableRowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Sort()(*icb8ff613f65f2c408b61aa652389279e4d69e53763a05b613109b8413b65d08e.SortRequestBuilder) {
    return icb8ff613f65f2c408b61aa652389279e4d69e53763a05b613109b8413b65d08e.NewSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*ic30e45ed9a3e16cf3233ad8b9a57ae5ea9eb9f128feb9a1cf797e693c0b0ccfd.TotalRowRangeRequestBuilder) {
    return ic30e45ed9a3e16cf3233ad8b9a57ae5ea9eb9f128feb9a1cf797e693c0b0ccfd.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Worksheet()(*ib91060f169b8bd39146db4c07fa37532518cdd9147a5ff6e0c87a23aaa611799.WorksheetRequestBuilder) {
    return ib91060f169b8bd39146db4c07fa37532518cdd9147a5ff6e0c87a23aaa611799.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
