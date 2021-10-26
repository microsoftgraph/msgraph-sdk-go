package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2c0ab814aea96def24c02ddbce61c67236d6f3dc665ebf502636d1ea6d4b1b45 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/clearfilters"
    i443382fea9319c90cc99f60949a1d5f85be11225dd45485edeaad348b2c1b3f5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/databodyrange"
    i4f0cba09d219fefd73eecd41f4f0c201a0d85158be9c3776af7775e36132b063 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/converttorange"
    i5dca63b6ba47e35a3fefaca406be81af4b75af498c5cb8b00e4a5401d9d59361 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/range_escaped"
    i7b6b696f6b76a3a880399b152362c2fd0f5c09bf83303acd04e3e9ef98083497 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns"
    i830b5bc0705e8462f5bdeae1338f4905baded81c0fbedfcf61dd445f6c1e639c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/reapplyfilters"
    i848b1facf5d9fd52c83dce6488c8523bdf2e958f5fbd6663a61bb1aafec6ea64 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/worksheet"
    i9669039f20a687e0054ea8b178095595904c2e89b486c3a620ea386cf59f0f64 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/rows"
    ia5322d7a8ad4593dbe7a5362bff7cd2be2d49beac945be242026420037d0710b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/headerrowrange"
    idbb7bab9a027efb2e87c98397f441f66d86375733cf13667c9c3c214eecbe67e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/totalrowrange"
    if6178cdc51d000e197840471e00259677f71d5981acaa7e66279cf0f685deaf9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/sort"
    i6d4d72b88b7b91ea45a618a16dab58940867ae4272bb7d7c49e69d55b8f337d3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item"
    i8e07163a4645b025b152b0df6a62496804418eac520030aa9528172f878b0eef "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/rows/item"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}
type WorkbookTableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Collection of tables that are part of the worksheet. Read-only.
type WorkbookTableRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *WorkbookTableRequestBuilder) ClearFilters()(*i2c0ab814aea96def24c02ddbce61c67236d6f3dc665ebf502636d1ea6d4b1b45.ClearFiltersRequestBuilder) {
    return i2c0ab814aea96def24c02ddbce61c67236d6f3dc665ebf502636d1ea6d4b1b45.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Columns()(*i7b6b696f6b76a3a880399b152362c2fd0f5c09bf83303acd04e3e9ef98083497.ColumnsRequestBuilder) {
    return i7b6b696f6b76a3a880399b152362c2fd0f5c09bf83303acd04e3e9ef98083497.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.names.item.worksheet.tables.item.columns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookTableRequestBuilder) ColumnsById(id string)(*i6d4d72b88b7b91ea45a618a16dab58940867ae4272bb7d7c49e69d55b8f337d3.WorkbookTableColumnRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableColumn_id"] = id
    }
    return i6d4d72b88b7b91ea45a618a16dab58940867ae4272bb7d7c49e69d55b8f337d3.NewWorkbookTableColumnRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new WorkbookTableRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/tables/{workbookTable_id}{?select,expand}";
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
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*i4f0cba09d219fefd73eecd41f4f0c201a0d85158be9c3776af7775e36132b063.ConvertToRangeRequestBuilder) {
    return i4f0cba09d219fefd73eecd41f4f0c201a0d85158be9c3776af7775e36132b063.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *WorkbookTableRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *WorkbookTableRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookTableRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookTableRequestBuilderGetQueryParameters)
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
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *WorkbookTableRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\microsoft.graph.dataBodyRange()
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*i443382fea9319c90cc99f60949a1d5f85be11225dd45485edeaad348b2c1b3f5.DataBodyRangeRequestBuilder) {
    return i443382fea9319c90cc99f60949a1d5f85be11225dd45485edeaad348b2c1b3f5.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookTableRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookTableRequestBuilder) Get(q func (value *WorkbookTableRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookTable() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable), nil
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\microsoft.graph.headerRowRange()
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*ia5322d7a8ad4593dbe7a5362bff7cd2be2d49beac945be242026420037d0710b.HeaderRowRangeRequestBuilder) {
    return ia5322d7a8ad4593dbe7a5362bff7cd2be2d49beac945be242026420037d0710b.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *WorkbookTableRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookTable, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\microsoft.graph.range()
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i5dca63b6ba47e35a3fefaca406be81af4b75af498c5cb8b00e4a5401d9d59361.RangeRequestBuilder) {
    return i5dca63b6ba47e35a3fefaca406be81af4b75af498c5cb8b00e4a5401d9d59361.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*i830b5bc0705e8462f5bdeae1338f4905baded81c0fbedfcf61dd445f6c1e639c.ReapplyFiltersRequestBuilder) {
    return i830b5bc0705e8462f5bdeae1338f4905baded81c0fbedfcf61dd445f6c1e639c.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Rows()(*i9669039f20a687e0054ea8b178095595904c2e89b486c3a620ea386cf59f0f64.RowsRequestBuilder) {
    return i9669039f20a687e0054ea8b178095595904c2e89b486c3a620ea386cf59f0f64.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.names.item.worksheet.tables.item.rows.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookTableRequestBuilder) RowsById(id string)(*i8e07163a4645b025b152b0df6a62496804418eac520030aa9528172f878b0eef.WorkbookTableRowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableRow_id"] = id
    }
    return i8e07163a4645b025b152b0df6a62496804418eac520030aa9528172f878b0eef.NewWorkbookTableRowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Sort()(*if6178cdc51d000e197840471e00259677f71d5981acaa7e66279cf0f685deaf9.SortRequestBuilder) {
    return if6178cdc51d000e197840471e00259677f71d5981acaa7e66279cf0f685deaf9.NewSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\microsoft.graph.totalRowRange()
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*idbb7bab9a027efb2e87c98397f441f66d86375733cf13667c9c3c214eecbe67e.TotalRowRangeRequestBuilder) {
    return idbb7bab9a027efb2e87c98397f441f66d86375733cf13667c9c3c214eecbe67e.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Worksheet()(*i848b1facf5d9fd52c83dce6488c8523bdf2e958f5fbd6663a61bb1aafec6ea64.WorksheetRequestBuilder) {
    return i848b1facf5d9fd52c83dce6488c8523bdf2e958f5fbd6663a61bb1aafec6ea64.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
