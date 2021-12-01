package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0fdf1c6304d8b6b67aba227723509d9ef681f97a2697787683eef2cdff5f042e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/totalrowrange"
    i34fd2b2262d7579bd8008ec288aa0a0de3d5a6045093a9b407d08854f9c1a674 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/range_escaped"
    i425e9c30f36ad5cb5e38bc2f2ae84c910bc698531a411328cdfda0e3bdd7263c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/headerrowrange"
    i7b385d4829cebac01b3252b90b6ecce372487bba36ae4726d82c1245573ceb8d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/clearfilters"
    ia5dedf9b9c77ff8d4c7cdd7f3af15354ab3792056aa38dd33de77a60cd08e68f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/databodyrange"
    ic00900eb2dec6fd41a403adbce7605af998680a6cfa0155ac689437d06cd1295 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/reapplyfilters"
    id2b71aba0df462bf56ee33a24217293160b9c77886c3873c1cc980c2a899a694 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item/converttorange"
)

// WorkbookTableRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}
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
func (m *WorkbookTableRequestBuilder) ClearFilters()(*i7b385d4829cebac01b3252b90b6ecce372487bba36ae4726d82c1245573ceb8d.ClearFiltersRequestBuilder) {
    return i7b385d4829cebac01b3252b90b6ecce372487bba36ae4726d82c1245573ceb8d.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookTableRequestBuilderInternal instantiates a new WorkbookTableRequestBuilder and sets the default values.
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet/tables/{workbookTable_id1}{?select,expand}";
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
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*id2b71aba0df462bf56ee33a24217293160b9c77886c3873c1cc980c2a899a694.ConvertToRangeRequestBuilder) {
    return id2b71aba0df462bf56ee33a24217293160b9c77886c3873c1cc980c2a899a694.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// DataBodyRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*ia5dedf9b9c77ff8d4c7cdd7f3af15354ab3792056aa38dd33de77a60cd08e68f.DataBodyRangeRequestBuilder) {
    return ia5dedf9b9c77ff8d4c7cdd7f3af15354ab3792056aa38dd33de77a60cd08e68f.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// HeaderRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}\microsoft.graph.headerRowRange()
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*i425e9c30f36ad5cb5e38bc2f2ae84c910bc698531a411328cdfda0e3bdd7263c.HeaderRowRangeRequestBuilder) {
    return i425e9c30f36ad5cb5e38bc2f2ae84c910bc698531a411328cdfda0e3bdd7263c.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Range_escaped builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}\microsoft.graph.range()
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i34fd2b2262d7579bd8008ec288aa0a0de3d5a6045093a9b407d08854f9c1a674.RangeRequestBuilder) {
    return i34fd2b2262d7579bd8008ec288aa0a0de3d5a6045093a9b407d08854f9c1a674.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*ic00900eb2dec6fd41a403adbce7605af998680a6cfa0155ac689437d06cd1295.ReapplyFiltersRequestBuilder) {
    return ic00900eb2dec6fd41a403adbce7605af998680a6cfa0155ac689437d06cd1295.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TotalRowRange builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\tables\{workbookTable-id1}\microsoft.graph.totalRowRange()
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*i0fdf1c6304d8b6b67aba227723509d9ef681f97a2697787683eef2cdff5f042e.TotalRowRangeRequestBuilder) {
    return i0fdf1c6304d8b6b67aba227723509d9ef681f97a2697787683eef2cdff5f042e.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
