package filter

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2a7b4d6e4c81fd73aef3fb00786c5d8b4920eddd112a6692dcb9fc5f7755e865 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/clear"
    i2e892a857ad30b23dda69f466fe02db1576088a7228f286056bb892a27a838d4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applycellcolorfilter"
    i3f6bc6b1827929cb35cfd1e5fc1ce6b5ded778701675242b26d3059a6aadfdc1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applycustomfilter"
    i4f3cfff0406ecafa47b2fd52eb01c436741f7e5122d505f43e634f600535c88d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applybottompercentfilter"
    i93837311e18e3fb667ae64a69773c9460ef3f60afcf9eed96dfe741c45748d27 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/apply"
    i9da386aa60ed9b59089e863522decf31163b5b96c9abdab308eb1347945cf94f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applydynamicfilter"
    ia81b26cf8447fd9b3f3c09b3792b3f9c384f2c842a27a61b41492fc06d73b5a6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applyiconfilter"
    ibc0c9b43720d6f889afa73d05dd257eaeda38c7ca354bf41009d28087988aeb4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applyfontcolorfilter"
    id1c5801f4112dbc4f529b101e8e92b4057ede33ca73d25533d0c6235a2c80e96 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applytopitemsfilter"
    id9e2dcd2398d7ca98e57c8ba30d767da57504e2f77de32be43673e8aa4ec0bc5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applyvaluesfilter"
    idf0e5e48868706c7f65e84c1747a4a9f8b9a96608cc52d66642f341afdae6320 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applytoppercentfilter"
    ief3f39300cd914f872cc06ec1196c76744107ebd8aad6ea752d0c38028d7682c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item/columns/item/filter/applybottomitemsfilter"
)

type FilterRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type FilterRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *FilterRequestBuilder) Apply()(*i93837311e18e3fb667ae64a69773c9460ef3f60afcf9eed96dfe741c45748d27.ApplyRequestBuilder) {
    return i93837311e18e3fb667ae64a69773c9460ef3f60afcf9eed96dfe741c45748d27.NewApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomItemsFilter()(*ief3f39300cd914f872cc06ec1196c76744107ebd8aad6ea752d0c38028d7682c.ApplyBottomItemsFilterRequestBuilder) {
    return ief3f39300cd914f872cc06ec1196c76744107ebd8aad6ea752d0c38028d7682c.NewApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomPercentFilter()(*i4f3cfff0406ecafa47b2fd52eb01c436741f7e5122d505f43e634f600535c88d.ApplyBottomPercentFilterRequestBuilder) {
    return i4f3cfff0406ecafa47b2fd52eb01c436741f7e5122d505f43e634f600535c88d.NewApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCellColorFilter()(*i2e892a857ad30b23dda69f466fe02db1576088a7228f286056bb892a27a838d4.ApplyCellColorFilterRequestBuilder) {
    return i2e892a857ad30b23dda69f466fe02db1576088a7228f286056bb892a27a838d4.NewApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCustomFilter()(*i3f6bc6b1827929cb35cfd1e5fc1ce6b5ded778701675242b26d3059a6aadfdc1.ApplyCustomFilterRequestBuilder) {
    return i3f6bc6b1827929cb35cfd1e5fc1ce6b5ded778701675242b26d3059a6aadfdc1.NewApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyDynamicFilter()(*i9da386aa60ed9b59089e863522decf31163b5b96c9abdab308eb1347945cf94f.ApplyDynamicFilterRequestBuilder) {
    return i9da386aa60ed9b59089e863522decf31163b5b96c9abdab308eb1347945cf94f.NewApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyFontColorFilter()(*ibc0c9b43720d6f889afa73d05dd257eaeda38c7ca354bf41009d28087988aeb4.ApplyFontColorFilterRequestBuilder) {
    return ibc0c9b43720d6f889afa73d05dd257eaeda38c7ca354bf41009d28087988aeb4.NewApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyIconFilter()(*ia81b26cf8447fd9b3f3c09b3792b3f9c384f2c842a27a61b41492fc06d73b5a6.ApplyIconFilterRequestBuilder) {
    return ia81b26cf8447fd9b3f3c09b3792b3f9c384f2c842a27a61b41492fc06d73b5a6.NewApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopItemsFilter()(*id1c5801f4112dbc4f529b101e8e92b4057ede33ca73d25533d0c6235a2c80e96.ApplyTopItemsFilterRequestBuilder) {
    return id1c5801f4112dbc4f529b101e8e92b4057ede33ca73d25533d0c6235a2c80e96.NewApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopPercentFilter()(*idf0e5e48868706c7f65e84c1747a4a9f8b9a96608cc52d66642f341afdae6320.ApplyTopPercentFilterRequestBuilder) {
    return idf0e5e48868706c7f65e84c1747a4a9f8b9a96608cc52d66642f341afdae6320.NewApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyValuesFilter()(*id9e2dcd2398d7ca98e57c8ba30d767da57504e2f77de32be43673e8aa4ec0bc5.ApplyValuesFilterRequestBuilder) {
    return id9e2dcd2398d7ca98e57c8ba30d767da57504e2f77de32be43673e8aa4ec0bc5.NewApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) Clear()(*i2a7b4d6e4c81fd73aef3fb00786c5d8b4920eddd112a6692dcb9fc5f7755e865.ClearRequestBuilder) {
    return i2a7b4d6e4c81fd73aef3fb00786c5d8b4920eddd112a6692dcb9fc5f7755e865.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    m := &FilterRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/columns/{workbookTableColumn_id}/filter{?select,expand}";
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
func NewFilterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilterRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *FilterRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *FilterRequestBuilder) CreateGetRequestInformation(q func (value *FilterRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(FilterRequestBuilderGetQueryParameters)
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
func (m *FilterRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFilter, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *FilterRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *FilterRequestBuilder) Get(q func (value *FilterRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFilter, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookFilter() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFilter), nil
}
func (m *FilterRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFilter, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
