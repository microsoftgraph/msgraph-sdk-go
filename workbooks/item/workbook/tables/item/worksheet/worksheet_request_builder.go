package worksheet

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i11a94f2cae31361a32b4d4240e4717921019962df2aae1544bdc4853b7ca8ee8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/usedrange"
    i203a12aa638160702197fd675c67a530d8d16abdd3529d5257fd2b182c0d6eef "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/rangewithaddress"
    i27f291d22a40e61bd12dfa8e9ac609688c0d7dead4890f8ae7762da3ebb85305 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/protection"
    i32b0fd7ebe7132dad7807d9f2c9f4d64ea0db6498e093f27b7c2a7970d41c316 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts"
    i5db80b38953d4e52efa904638be46d17f0a6349c59160837e45f9dba76a0d7f2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/cellwithrowwithcolumn"
    i9235b8550efd8b4cc6e220d2e09bdb5dd9f6ec34a62247f0b83f77b50469207a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/range_escpaped"
    ib0bbebe90b78f735acce602d7890c4ab74a3609f21a49969f37a5ae831c3dd0b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/names"
    id2fde7f1fde5e2425d9250857b0beef6916a17695e978b32fee20e1a29fe5911 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/pivottables"
    id65df021ece2097d3adb36a5d68520266eb1ddd1f7a72283f4c176d0321155c9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables"
    ief073dbf92fe3e28eaebe2bfefe5f35f24f7a0c0449015f426ef206250057e8e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/usedrangewithvaluesonly"
    i4161f9f172fd403629077bc212edba826ca8228dc729879aec8e114f8111ed65 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/pivottables/item"
    i5697b01121e5fbdfc18af996485f58b3a3149aab74676e3e4b3e7c6a1b215b75 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/names/item"
    ib84a8f93c2faece4c6d09fdb49006ff5ae86a0df36287113d222c2febc5e25bf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item"
    if6cb873ce125d821d991473e2a82753e671687c5cb7326aff4e8c0dc6166b0dc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item"
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
func (m *WorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i5db80b38953d4e52efa904638be46d17f0a6349c59160837e45f9dba76a0d7f2.CellWithRowWithColumnRequestBuilder) {
    return i5db80b38953d4e52efa904638be46d17f0a6349c59160837e45f9dba76a0d7f2.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorksheetRequestBuilder) Charts()(*i32b0fd7ebe7132dad7807d9f2c9f4d64ea0db6498e093f27b7c2a7970d41c316.ChartsRequestBuilder) {
    return i32b0fd7ebe7132dad7807d9f2c9f4d64ea0db6498e093f27b7c2a7970d41c316.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) ChartsById(id string)(*ib84a8f93c2faece4c6d09fdb49006ff5ae86a0df36287113d222c2febc5e25bf.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return ib84a8f93c2faece4c6d09fdb49006ff5ae86a0df36287113d222c2febc5e25bf.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    m := &WorksheetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet{?select,expand}";
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
func (m *WorksheetRequestBuilder) Names()(*ib0bbebe90b78f735acce602d7890c4ab74a3609f21a49969f37a5ae831c3dd0b.NamesRequestBuilder) {
    return ib0bbebe90b78f735acce602d7890c4ab74a3609f21a49969f37a5ae831c3dd0b.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) NamesById(id string)(*i5697b01121e5fbdfc18af996485f58b3a3149aab74676e3e4b3e7c6a1b215b75.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return i5697b01121e5fbdfc18af996485f58b3a3149aab74676e3e4b3e7c6a1b215b75.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *WorksheetRequestBuilder) PivotTables()(*id2fde7f1fde5e2425d9250857b0beef6916a17695e978b32fee20e1a29fe5911.PivotTablesRequestBuilder) {
    return id2fde7f1fde5e2425d9250857b0beef6916a17695e978b32fee20e1a29fe5911.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) PivotTablesById(id string)(*i4161f9f172fd403629077bc212edba826ca8228dc729879aec8e114f8111ed65.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return i4161f9f172fd403629077bc212edba826ca8228dc729879aec8e114f8111ed65.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Protection()(*i27f291d22a40e61bd12dfa8e9ac609688c0d7dead4890f8ae7762da3ebb85305.ProtectionRequestBuilder) {
    return i27f291d22a40e61bd12dfa8e9ac609688c0d7dead4890f8ae7762da3ebb85305.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Range_escpaped()(*i9235b8550efd8b4cc6e220d2e09bdb5dd9f6ec34a62247f0b83f77b50469207a.RangeRequestBuilder) {
    return i9235b8550efd8b4cc6e220d2e09bdb5dd9f6ec34a62247f0b83f77b50469207a.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) RangeWithAddress(address *string)(*i203a12aa638160702197fd675c67a530d8d16abdd3529d5257fd2b182c0d6eef.RangeWithAddressRequestBuilder) {
    return i203a12aa638160702197fd675c67a530d8d16abdd3529d5257fd2b182c0d6eef.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorksheetRequestBuilder) Tables()(*id65df021ece2097d3adb36a5d68520266eb1ddd1f7a72283f4c176d0321155c9.TablesRequestBuilder) {
    return id65df021ece2097d3adb36a5d68520266eb1ddd1f7a72283f4c176d0321155c9.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) TablesById(id string)(*if6cb873ce125d821d991473e2a82753e671687c5cb7326aff4e8c0dc6166b0dc.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookTable_id1"] = id
    }
    return if6cb873ce125d821d991473e2a82753e671687c5cb7326aff4e8c0dc6166b0dc.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) UsedRange()(*i11a94f2cae31361a32b4d4240e4717921019962df2aae1544bdc4853b7ca8ee8.UsedRangeRequestBuilder) {
    return i11a94f2cae31361a32b4d4240e4717921019962df2aae1544bdc4853b7ca8ee8.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ief073dbf92fe3e28eaebe2bfefe5f35f24f7a0c0449015f426ef206250057e8e.UsedRangeWithValuesOnlyRequestBuilder) {
    return ief073dbf92fe3e28eaebe2bfefe5f35f24f7a0c0449015f426ef206250057e8e.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
