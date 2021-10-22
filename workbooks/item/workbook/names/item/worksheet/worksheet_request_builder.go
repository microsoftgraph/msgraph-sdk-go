package worksheet

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i62d52fd4ee533434279ebba67b0c72dbb43d7ad1fd44b2316953eb5e96f5d2ac "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/range_escpaped"
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

type WorksheetRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorksheetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *WorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i64ccfbae56914e7f7a124e3677fd78e3db5b08b88d24815c9b41646d85c93fb1.CellWithRowWithColumnRequestBuilder) {
    return i64ccfbae56914e7f7a124e3677fd78e3db5b08b88d24815c9b41646d85c93fb1.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorksheetRequestBuilder) Charts()(*iebe192e817f0989a5ac93e7753114363de65420f7846aab0afa7c1b51c50e215.ChartsRequestBuilder) {
    return iebe192e817f0989a5ac93e7753114363de65420f7846aab0afa7c1b51c50e215.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) ChartsById(id string)(*ia5ecdd25c9e61362a7cab70fef1a86855b8f4b1b4d141525c20606e0826fb816.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return ia5ecdd25c9e61362a7cab70fef1a86855b8f4b1b4d141525c20606e0826fb816.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    m := &WorksheetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorksheetRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorksheetRequestBuilder) CreateGetRequestInformation(q func (value *WorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorksheetRequestBuilderGetQueryParameters)
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
func (m *WorksheetRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorksheetRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorksheetRequestBuilder) Get(q func (value *WorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookWorksheet() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet), nil
}
func (m *WorksheetRequestBuilder) Names()(*if92dc1160b5acc1b3cff97bed31a29c0cd468493dc248cfa779cdae1a1d84751.NamesRequestBuilder) {
    return if92dc1160b5acc1b3cff97bed31a29c0cd468493dc248cfa779cdae1a1d84751.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) NamesById(id string)(*i1aaf19edcdc701f5f7ba1313e8a52938d4a98f5ee081bd238323f84d780a8220.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id1"] = id
    }
    return i1aaf19edcdc701f5f7ba1313e8a52938d4a98f5ee081bd238323f84d780a8220.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorksheetRequestBuilder) PivotTables()(*ie90ffca7b164359c033847e1f150328fd6989ad1b4e7cf52fa5d53ce6a228106.PivotTablesRequestBuilder) {
    return ie90ffca7b164359c033847e1f150328fd6989ad1b4e7cf52fa5d53ce6a228106.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) PivotTablesById(id string)(*id098e77a8f63642f05aa6cbfc83f6060b05dd7ecdc18b30cb291f7904c1efd00.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return id098e77a8f63642f05aa6cbfc83f6060b05dd7ecdc18b30cb291f7904c1efd00.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Protection()(*i85317c503f2f27185e070ade5b9624ae780d770e19305074b565d7a00d47a3da.ProtectionRequestBuilder) {
    return i85317c503f2f27185e070ade5b9624ae780d770e19305074b565d7a00d47a3da.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Range_escpaped()(*i62d52fd4ee533434279ebba67b0c72dbb43d7ad1fd44b2316953eb5e96f5d2ac.RangeRequestBuilder) {
    return i62d52fd4ee533434279ebba67b0c72dbb43d7ad1fd44b2316953eb5e96f5d2ac.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) RangeWithAddress(address *string)(*if6887ddb306e8a1aee896bc8347e27bdd4d399126be76020ab257042ad9d47d2.RangeWithAddressRequestBuilder) {
    return if6887ddb306e8a1aee896bc8347e27bdd4d399126be76020ab257042ad9d47d2.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorksheetRequestBuilder) Tables()(*i9ea4a635326d648cea4fd3aa2cd53ff2e732c3d50c1fffe56270a60b299c0eac.TablesRequestBuilder) {
    return i9ea4a635326d648cea4fd3aa2cd53ff2e732c3d50c1fffe56270a60b299c0eac.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) TablesById(id string)(*i38632c08ea0ceaf6e28d20b457b81cf5c5600eb7bd7e290d1f5e6fd03843dbe8.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return i38632c08ea0ceaf6e28d20b457b81cf5c5600eb7bd7e290d1f5e6fd03843dbe8.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) UsedRange()(*id9cac5cb49a69b54bf5d8a221ad0026911c5b716d04bd5168a55b29dc6cbe016.UsedRangeRequestBuilder) {
    return id9cac5cb49a69b54bf5d8a221ad0026911c5b716d04bd5168a55b29dc6cbe016.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ic81ca612ad48c0f47f947a769ecba949b1d9dc349a3380bf7198be7180b1de93.UsedRangeWithValuesOnlyRequestBuilder) {
    return ic81ca612ad48c0f47f947a769ecba949b1d9dc349a3380bf7198be7180b1de93.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
