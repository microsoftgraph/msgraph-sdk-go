package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0ffc856a99ab076eacab92d71f01358a27ce7f82ca9cea1eeef71ccc0b690b04 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/protection"
    i100ec4503074fbce609d04c7af9f3038ecbe146a28dd1e9696d968fa800e43fb "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/usedrange"
    i63bfade034059e0bafdeccc55d1d84a9578056cd0b4452d244d1ea48d219bfd5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables"
    i9c65a5e0511223ed4b7aadfa985f3cecba8763d0234d695c0eba9470b08105fd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/pivottables"
    i9e565a8c9d35090b91ddaf8942fefd178ab4d8274ecd66a46cebd61471579ac8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/names"
    ia27bd953af5b50bd89e1ac4a64ebdd5fc2bc7d28f43a08e7f1892e8d22f44739 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/range_escaped"
    ia4b371526ab2f408171fc46a799dd775d49573d38da9d1f1459fa6b1e8cc5242 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/usedrangewithvaluesonly"
    iad940f8124c7334170f4bc84a78fc081f07f84f05e2bbcdad79b8c17146992f2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts"
    ib8d36af0717b2669dc8b8889287197f9036d4251ba8a717b88522dcf0bc194d8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/rangewithaddress"
    ie4b7c6c1c8b3e34e3e51127f9ebb673f0a04c788986d8229f1a710411d5f225d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/cellwithrowwithcolumn"
    i2b7eb38de9f32e0a84e0f4d15c3dfa5d1787f849e6127d28c3c84d7fb2c68f22 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/pivottables/item"
    i3ecee7d2013b82649254a2abcecff10c453e98848525efb8661b8198de377cbd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item"
    i8dc6bf9c3cd413c8da2a0e6e6c76573ab0a8670c078efc72a8efa32a1950ad10 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/names/item"
    iaf6e80887a038f65ff678934a224c20738b5518de1ef6bd042c68d64e3eb866a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/charts/item"
)

// WorkbookWorksheetRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}
type WorkbookWorksheetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookWorksheetRequestBuilderDeleteOptions options for Delete
type WorkbookWorksheetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookWorksheetRequestBuilderGetOptions options for Get
type WorkbookWorksheetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookWorksheetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookWorksheetRequestBuilderGetQueryParameters represents a collection of worksheets associated with the workbook. Read-only.
type WorkbookWorksheetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// WorkbookWorksheetRequestBuilderPatchOptions options for Patch
type WorkbookWorksheetRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CellWithRowWithColumn builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.cell(row={row},column={column})
func (m *WorkbookWorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*ie4b7c6c1c8b3e34e3e51127f9ebb673f0a04c788986d8229f1a710411d5f225d.CellWithRowWithColumnRequestBuilder) {
    return ie4b7c6c1c8b3e34e3e51127f9ebb673f0a04c788986d8229f1a710411d5f225d.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookWorksheetRequestBuilder) Charts()(*iad940f8124c7334170f4bc84a78fc081f07f84f05e2bbcdad79b8c17146992f2.ChartsRequestBuilder) {
    return iad940f8124c7334170f4bc84a78fc081f07f84f05e2bbcdad79b8c17146992f2.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChartsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item.charts.item collection
func (m *WorkbookWorksheetRequestBuilder) ChartsById(id string)(*iaf6e80887a038f65ff678934a224c20738b5518de1ef6bd042c68d64e3eb866a.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return iaf6e80887a038f65ff678934a224c20738b5518de1ef6bd042c68d64e3eb866a.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWorkbookWorksheetRequestBuilderInternal instantiates a new WorkbookWorksheetRequestBuilder and sets the default values.
func NewWorkbookWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetRequestBuilder) {
    m := &WorkbookWorksheetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookWorksheetRequestBuilder instantiates a new WorkbookWorksheetRequestBuilder and sets the default values.
func NewWorkbookWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookWorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) CreateDeleteRequestInformation(options *WorkbookWorksheetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) CreateGetRequestInformation(options *WorkbookWorksheetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) CreatePatchRequestInformation(options *WorkbookWorksheetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) Delete(options *WorkbookWorksheetRequestBuilderDeleteOptions)(error) {
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
// Get represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) Get(options *WorkbookWorksheetRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet, error) {
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
func (m *WorkbookWorksheetRequestBuilder) Names()(*i9e565a8c9d35090b91ddaf8942fefd178ab4d8274ecd66a46cebd61471579ac8.NamesRequestBuilder) {
    return i9e565a8c9d35090b91ddaf8942fefd178ab4d8274ecd66a46cebd61471579ac8.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item.names.item collection
func (m *WorkbookWorksheetRequestBuilder) NamesById(id string)(*i8dc6bf9c3cd413c8da2a0e6e6c76573ab0a8670c078efc72a8efa32a1950ad10.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return i8dc6bf9c3cd413c8da2a0e6e6c76573ab0a8670c078efc72a8efa32a1950ad10.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch represents a collection of worksheets associated with the workbook. Read-only.
func (m *WorkbookWorksheetRequestBuilder) Patch(options *WorkbookWorksheetRequestBuilderPatchOptions)(error) {
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
func (m *WorkbookWorksheetRequestBuilder) PivotTables()(*i9c65a5e0511223ed4b7aadfa985f3cecba8763d0234d695c0eba9470b08105fd.PivotTablesRequestBuilder) {
    return i9c65a5e0511223ed4b7aadfa985f3cecba8763d0234d695c0eba9470b08105fd.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PivotTablesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item.pivotTables.item collection
func (m *WorkbookWorksheetRequestBuilder) PivotTablesById(id string)(*i2b7eb38de9f32e0a84e0f4d15c3dfa5d1787f849e6127d28c3c84d7fb2c68f22.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return i2b7eb38de9f32e0a84e0f4d15c3dfa5d1787f849e6127d28c3c84d7fb2c68f22.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookWorksheetRequestBuilder) Protection()(*i0ffc856a99ab076eacab92d71f01358a27ce7f82ca9cea1eeef71ccc0b690b04.ProtectionRequestBuilder) {
    return i0ffc856a99ab076eacab92d71f01358a27ce7f82ca9cea1eeef71ccc0b690b04.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.range()
func (m *WorkbookWorksheetRequestBuilder) Range_escaped()(*ia27bd953af5b50bd89e1ac4a64ebdd5fc2bc7d28f43a08e7f1892e8d22f44739.RangeRequestBuilder) {
    return ia27bd953af5b50bd89e1ac4a64ebdd5fc2bc7d28f43a08e7f1892e8d22f44739.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RangeWithAddress builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.range(address='{address}')
func (m *WorkbookWorksheetRequestBuilder) RangeWithAddress(address *string)(*ib8d36af0717b2669dc8b8889287197f9036d4251ba8a717b88522dcf0bc194d8.RangeWithAddressRequestBuilder) {
    return ib8d36af0717b2669dc8b8889287197f9036d4251ba8a717b88522dcf0bc194d8.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorkbookWorksheetRequestBuilder) Tables()(*i63bfade034059e0bafdeccc55d1d84a9578056cd0b4452d244d1ea48d219bfd5.TablesRequestBuilder) {
    return i63bfade034059e0bafdeccc55d1d84a9578056cd0b4452d244d1ea48d219bfd5.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TablesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item.tables.item collection
func (m *WorkbookWorksheetRequestBuilder) TablesById(id string)(*i3ecee7d2013b82649254a2abcecff10c453e98848525efb8661b8198de377cbd.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return i3ecee7d2013b82649254a2abcecff10c453e98848525efb8661b8198de377cbd.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsedRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.usedRange()
func (m *WorkbookWorksheetRequestBuilder) UsedRange()(*i100ec4503074fbce609d04c7af9f3038ecbe146a28dd1e9696d968fa800e43fb.UsedRangeRequestBuilder) {
    return i100ec4503074fbce609d04c7af9f3038ecbe146a28dd1e9696d968fa800e43fb.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRangeWithValuesOnly builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\microsoft.graph.usedRange(valuesOnly={valuesOnly})
func (m *WorkbookWorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ia4b371526ab2f408171fc46a799dd775d49573d38da9d1f1459fa6b1e8cc5242.UsedRangeWithValuesOnlyRequestBuilder) {
    return ia4b371526ab2f408171fc46a799dd775d49573d38da9d1f1459fa6b1e8cc5242.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
