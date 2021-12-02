package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0d2116063140655ff0006524223221da2722f8f17d4414ae09101cec373fe186 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/clearfilters"
    i238e8100dc643d44c111ef9be2599bdc0060f631fe3c10611a6f0e5346e396e7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/converttorange"
    i270bfaf5c08f7645bb2c14589e9208ad650c36c4a6d239bc82777d08389de087 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/databodyrange"
    i28f2607281f955f64aa9e9733df14c92c47abde67626112b192fe83f14aae759 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/range_escaped"
    i384735f7afb53e99e9096748feb737883aa5b60a9568c919f64c4949b1b5867a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/totalrowrange"
    i58379b9a8db83ef4948c8d336e90caf23e3d3c7f50698713f0388a2f226fc775 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/worksheet"
    i6839e7105f301fcb6cb28af21d542cc1211ee920851a40c465ab1e66dc5c9a5f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/sort"
    ib3aa54bb6604235e02a0288b9307f15cbe47dc03be9fffd87fe4c3adc785ade4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns"
    ic6691b48d97464c740fc9517d89d95bbdf8c93b576e53fb357a86507c494cb1b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/reapplyfilters"
    ida7822f4a94d7c5aaf542b9a2f938c61cf17041eae3311f29f2f1256b7f6f36f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/rows"
    iee936b72318db6b19eb4c29081edece437e74a2bd4aaef15e4b725616c1ab775 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/headerrowrange"
    i0680c4a7ba6d95a69c5ede76b8a5f53c0811779779e91bf325f5f69ba6aea10e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item"
    i98904eb4dd77c5865218010d32120f22f8e1b959e48427e914c1402811e8f4ba "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/rows/item"
)

// WorkbookTableRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}
type WorkbookTableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookTableRequestBuilderDeleteOptions options for Delete
type WorkbookTableRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookTableRequestBuilderGetOptions options for Get
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
// WorkbookTableRequestBuilderGetQueryParameters collection of tables that are part of the worksheet. Read-only.
type WorkbookTableRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookTableRequestBuilderPatchOptions options for Patch
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
func (m *WorkbookTableRequestBuilder) ClearFilters()(*i0d2116063140655ff0006524223221da2722f8f17d4414ae09101cec373fe186.ClearFiltersRequestBuilder) {
    return i0d2116063140655ff0006524223221da2722f8f17d4414ae09101cec373fe186.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Columns()(*ib3aa54bb6604235e02a0288b9307f15cbe47dc03be9fffd87fe4c3adc785ade4.ColumnsRequestBuilder) {
    return ib3aa54bb6604235e02a0288b9307f15cbe47dc03be9fffd87fe4c3adc785ade4.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item.tables.item.columns.item collection
func (m *WorkbookTableRequestBuilder) ColumnsById(id string)(*i0680c4a7ba6d95a69c5ede76b8a5f53c0811779779e91bf325f5f69ba6aea10e.WorkbookTableColumnRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableColumn_id"] = id
    }
    return i0680c4a7ba6d95a69c5ede76b8a5f53c0811779779e91bf325f5f69ba6aea10e.NewWorkbookTableColumnRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWorkbookTableRequestBuilderInternal instantiates a new WorkbookTableRequestBuilder and sets the default values.
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/tables/{workbookTable_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookTableRequestBuilder instantiates a new WorkbookTableRequestBuilder and sets the default values.
func NewWorkbookTableRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*i238e8100dc643d44c111ef9be2599bdc0060f631fe3c10611a6f0e5346e396e7.ConvertToRangeRequestBuilder) {
    return i238e8100dc643d44c111ef9be2599bdc0060f631fe3c10611a6f0e5346e396e7.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation collection of tables that are part of the worksheet. Read-only.
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
// CreateGetRequestInformation collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookTableRequestBuilder) CreateGetRequestInformation(options *WorkbookTableRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation collection of tables that are part of the worksheet. Read-only.
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
// DataBodyRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*i270bfaf5c08f7645bb2c14589e9208ad650c36c4a6d239bc82777d08389de087.DataBodyRangeRequestBuilder) {
    return i270bfaf5c08f7645bb2c14589e9208ad650c36c4a6d239bc82777d08389de087.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete collection of tables that are part of the worksheet. Read-only.
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
// Get collection of tables that are part of the worksheet. Read-only.
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
// HeaderRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*iee936b72318db6b19eb4c29081edece437e74a2bd4aaef15e4b725616c1ab775.HeaderRowRangeRequestBuilder) {
    return iee936b72318db6b19eb4c29081edece437e74a2bd4aaef15e4b725616c1ab775.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch collection of tables that are part of the worksheet. Read-only.
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\microsoft.graph.range()
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i28f2607281f955f64aa9e9733df14c92c47abde67626112b192fe83f14aae759.RangeRequestBuilder) {
    return i28f2607281f955f64aa9e9733df14c92c47abde67626112b192fe83f14aae759.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*ic6691b48d97464c740fc9517d89d95bbdf8c93b576e53fb357a86507c494cb1b.ReapplyFiltersRequestBuilder) {
    return ic6691b48d97464c740fc9517d89d95bbdf8c93b576e53fb357a86507c494cb1b.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Rows()(*ida7822f4a94d7c5aaf542b9a2f938c61cf17041eae3311f29f2f1256b7f6f36f.RowsRequestBuilder) {
    return ida7822f4a94d7c5aaf542b9a2f938c61cf17041eae3311f29f2f1256b7f6f36f.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RowsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item.tables.item.rows.item collection
func (m *WorkbookTableRequestBuilder) RowsById(id string)(*i98904eb4dd77c5865218010d32120f22f8e1b959e48427e914c1402811e8f4ba.WorkbookTableRowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableRow_id"] = id
    }
    return i98904eb4dd77c5865218010d32120f22f8e1b959e48427e914c1402811e8f4ba.NewWorkbookTableRowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Sort()(*i6839e7105f301fcb6cb28af21d542cc1211ee920851a40c465ab1e66dc5c9a5f.SortRequestBuilder) {
    return i6839e7105f301fcb6cb28af21d542cc1211ee920851a40c465ab1e66dc5c9a5f.NewSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TotalRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*i384735f7afb53e99e9096748feb737883aa5b60a9568c919f64c4949b1b5867a.TotalRowRangeRequestBuilder) {
    return i384735f7afb53e99e9096748feb737883aa5b60a9568c919f64c4949b1b5867a.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Worksheet()(*i58379b9a8db83ef4948c8d336e90caf23e3d3c7f50698713f0388a2f226fc775.WorksheetRequestBuilder) {
    return i58379b9a8db83ef4948c8d336e90caf23e3d3c7f50698713f0388a2f226fc775.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
