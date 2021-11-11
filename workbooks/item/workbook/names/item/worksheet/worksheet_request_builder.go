package worksheet

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i261678c5f5a15694493963fe717eabe43c4934b5327279953b86ae94225380e4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/range_escaped"
    i64ccfbae56914e7f7a124e3677fd78e3db5b08b88d24815c9b41646d85c93fb1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/cellwithrowwithcolumn"
    i85317c503f2f27185e070ade5b9624ae780d770e19305074b565d7a00d47a3da "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/protection"
    i9ea4a635326d648cea4fd3aa2cd53ff2e732c3d50c1fffe56270a60b299c0eac "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables"
    ic81ca612ad48c0f47f947a769ecba949b1d9dc349a3380bf7198be7180b1de93 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/usedrangewithvaluesonly"
    id9cac5cb49a69b54bf5d8a221ad0026911c5b716d04bd5168a55b29dc6cbe016 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/usedrange"
    ie90ffca7b164359c033847e1f150328fd6989ad1b4e7cf52fa5d53ce6a228106 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/pivottables"
    iebe192e817f0989a5ac93e7753114363de65420f7846aab0afa7c1b51c50e215 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts"
    if6887ddb306e8a1aee896bc8347e27bdd4d399126be76020ab257042ad9d47d2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/rangewithaddress"
    if92dc1160b5acc1b3cff97bed31a29c0cd468493dc248cfa779cdae1a1d84751 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/names"
    i1aaf19edcdc701f5f7ba1313e8a52938d4a98f5ee081bd238323f84d780a8220 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/names/item"
    i38632c08ea0ceaf6e28d20b457b81cf5c5600eb7bd7e290d1f5e6fd03843dbe8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item"
    ia5ecdd25c9e61362a7cab70fef1a86855b8f4b1b4d141525c20606e0826fb816 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item"
    id098e77a8f63642f05aa6cbfc83f6060b05dd7ecdc18b30cb291f7904c1efd00 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/pivottables/item"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet
type WorksheetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type WorksheetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type WorksheetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorksheetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
type WorksheetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type WorksheetRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\microsoft.graph.cell(row={row},column={column})
// Parameters:
//  - column : Usage: column={column}
//  - row : Usage: row={row}
func (m *WorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i64ccfbae56914e7f7a124e3677fd78e3db5b08b88d24815c9b41646d85c93fb1.CellWithRowWithColumnRequestBuilder) {
    return i64ccfbae56914e7f7a124e3677fd78e3db5b08b88d24815c9b41646d85c93fb1.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorksheetRequestBuilder) Charts()(*iebe192e817f0989a5ac93e7753114363de65420f7846aab0afa7c1b51c50e215.ChartsRequestBuilder) {
    return iebe192e817f0989a5ac93e7753114363de65420f7846aab0afa7c1b51c50e215.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.names.item.worksheet.charts.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorksheetRequestBuilder) ChartsById(id string)(*ia5ecdd25c9e61362a7cab70fef1a86855b8f4b1b4d141525c20606e0826fb816.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return ia5ecdd25c9e61362a7cab70fef1a86855b8f4b1b4d141525c20606e0826fb816.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new WorksheetRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    m := &WorksheetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorksheetRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
// Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorksheetRequestBuilder) CreateDeleteRequestInformation(options *WorksheetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorksheetRequestBuilder) CreateGetRequestInformation(options *WorksheetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
// Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorksheetRequestBuilder) CreatePatchRequestInformation(options *WorksheetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorksheetRequestBuilder) Delete(options *WorksheetRequestBuilderDeleteOptions)(error) {
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
// Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorksheetRequestBuilder) Get(options *WorksheetRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookWorksheet() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet), nil
}
func (m *WorksheetRequestBuilder) Names()(*if92dc1160b5acc1b3cff97bed31a29c0cd468493dc248cfa779cdae1a1d84751.NamesRequestBuilder) {
    return if92dc1160b5acc1b3cff97bed31a29c0cd468493dc248cfa779cdae1a1d84751.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.names.item.worksheet.names.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorksheetRequestBuilder) NamesById(id string)(*i1aaf19edcdc701f5f7ba1313e8a52938d4a98f5ee081bd238323f84d780a8220.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id1"] = id
    }
    return i1aaf19edcdc701f5f7ba1313e8a52938d4a98f5ee081bd238323f84d780a8220.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only.
// Parameters:
//  - options : Options for the request
func (m *WorksheetRequestBuilder) Patch(options *WorksheetRequestBuilderPatchOptions)(error) {
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
func (m *WorksheetRequestBuilder) PivotTables()(*ie90ffca7b164359c033847e1f150328fd6989ad1b4e7cf52fa5d53ce6a228106.PivotTablesRequestBuilder) {
    return ie90ffca7b164359c033847e1f150328fd6989ad1b4e7cf52fa5d53ce6a228106.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.names.item.worksheet.pivotTables.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorksheetRequestBuilder) PivotTablesById(id string)(*id098e77a8f63642f05aa6cbfc83f6060b05dd7ecdc18b30cb291f7904c1efd00.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return id098e77a8f63642f05aa6cbfc83f6060b05dd7ecdc18b30cb291f7904c1efd00.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Protection()(*i85317c503f2f27185e070ade5b9624ae780d770e19305074b565d7a00d47a3da.ProtectionRequestBuilder) {
    return i85317c503f2f27185e070ade5b9624ae780d770e19305074b565d7a00d47a3da.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\microsoft.graph.range()
func (m *WorksheetRequestBuilder) Range_escaped()(*i261678c5f5a15694493963fe717eabe43c4934b5327279953b86ae94225380e4.RangeRequestBuilder) {
    return i261678c5f5a15694493963fe717eabe43c4934b5327279953b86ae94225380e4.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\microsoft.graph.range(address='{address}')
// Parameters:
//  - address : Usage: address={address}
func (m *WorksheetRequestBuilder) RangeWithAddress(address *string)(*if6887ddb306e8a1aee896bc8347e27bdd4d399126be76020ab257042ad9d47d2.RangeWithAddressRequestBuilder) {
    return if6887ddb306e8a1aee896bc8347e27bdd4d399126be76020ab257042ad9d47d2.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorksheetRequestBuilder) Tables()(*i9ea4a635326d648cea4fd3aa2cd53ff2e732c3d50c1fffe56270a60b299c0eac.TablesRequestBuilder) {
    return i9ea4a635326d648cea4fd3aa2cd53ff2e732c3d50c1fffe56270a60b299c0eac.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.names.item.worksheet.tables.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorksheetRequestBuilder) TablesById(id string)(*i38632c08ea0ceaf6e28d20b457b81cf5c5600eb7bd7e290d1f5e6fd03843dbe8.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return i38632c08ea0ceaf6e28d20b457b81cf5c5600eb7bd7e290d1f5e6fd03843dbe8.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\microsoft.graph.usedRange()
func (m *WorksheetRequestBuilder) UsedRange()(*id9cac5cb49a69b54bf5d8a221ad0026911c5b716d04bd5168a55b29dc6cbe016.UsedRangeRequestBuilder) {
    return id9cac5cb49a69b54bf5d8a221ad0026911c5b716d04bd5168a55b29dc6cbe016.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\microsoft.graph.usedRange(valuesOnly={valuesOnly})
// Parameters:
//  - valuesOnly : Usage: valuesOnly={valuesOnly}
func (m *WorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ic81ca612ad48c0f47f947a769ecba949b1d9dc349a3380bf7198be7180b1de93.UsedRangeWithValuesOnlyRequestBuilder) {
    return ic81ca612ad48c0f47f947a769ecba949b1d9dc349a3380bf7198be7180b1de93.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
