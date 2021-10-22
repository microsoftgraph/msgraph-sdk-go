package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0a0310fe84d81d83996a2aef7d5c59643f0d37b37fa399938d492a62b1c755ea "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/columns"
    i374ab2c07ee8853772ee041c66fd22fdbba3306fff11238ef92a75a562224a39 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/drive"
    i637397bfe35b12432f6a7dfdf6a733639d17bb6c7ec32033be4f1a179262a9e3 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes"
    icbe1aadc69be395af894f2fcffa73621ad8b93f3e237737513917536cab07824 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items"
    ie5663fa3ec2f5c7aef7ca0c99a034c879cba05b632b6afd20e7ad1c08ff1c437 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/subscriptions"
    i0befc12065c1f523f268ef23fb2b480c96f7d18d72934e0563b3b4fbe977f82d "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/columns/item"
    i52367d7cec2b5a644a4aac2a5d8d592f9e281e7194c330d53c34d2163639d04d "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item"
    i61ab027a9179fbc39656243c6dcf797fd209a10d47964360d5075fc3811230e8 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/subscriptions/item"
    ieaf637fc737fb74672cfa54c30a3a26b4b735b92361be7df20275f2854de5963 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/items/item"
)

type ListRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ListRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ListRequestBuilder) Columns()(*i0a0310fe84d81d83996a2aef7d5c59643f0d37b37fa399938d492a62b1c755ea.ColumnsRequestBuilder) {
    return i0a0310fe84d81d83996a2aef7d5c59643f0d37b37fa399938d492a62b1c755ea.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ColumnsById(id string)(*i0befc12065c1f523f268ef23fb2b480c96f7d18d72934e0563b3b4fbe977f82d.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i0befc12065c1f523f268ef23fb2b480c96f7d18d72934e0563b3b4fbe977f82d.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/drives/{drive_id}/list{?select,expand}";
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
func NewListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListRequestBuilder) ContentTypes()(*i637397bfe35b12432f6a7dfdf6a733639d17bb6c7ec32033be4f1a179262a9e3.ContentTypesRequestBuilder) {
    return i637397bfe35b12432f6a7dfdf6a733639d17bb6c7ec32033be4f1a179262a9e3.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ContentTypesById(id string)(*i52367d7cec2b5a644a4aac2a5d8d592f9e281e7194c330d53c34d2163639d04d.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i52367d7cec2b5a644a4aac2a5d8d592f9e281e7194c330d53c34d2163639d04d.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListRequestBuilder) CreateGetRequestInformation(q func (value *ListRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ListRequestBuilderGetQueryParameters)
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
func (m *ListRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListRequestBuilder) Drive()(*i374ab2c07ee8853772ee041c66fd22fdbba3306fff11238ef92a75a562224a39.DriveRequestBuilder) {
    return i374ab2c07ee8853772ee041c66fd22fdbba3306fff11238ef92a75a562224a39.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) Get(q func (value *ListRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewList() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List), nil
}
func (m *ListRequestBuilder) Items()(*icbe1aadc69be395af894f2fcffa73621ad8b93f3e237737513917536cab07824.ItemsRequestBuilder) {
    return icbe1aadc69be395af894f2fcffa73621ad8b93f3e237737513917536cab07824.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ItemsById(id string)(*ieaf637fc737fb74672cfa54c30a3a26b4b735b92361be7df20275f2854de5963.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return ieaf637fc737fb74672cfa54c30a3a26b4b735b92361be7df20275f2854de5963.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.List, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListRequestBuilder) Subscriptions()(*ie5663fa3ec2f5c7aef7ca0c99a034c879cba05b632b6afd20e7ad1c08ff1c437.SubscriptionsRequestBuilder) {
    return ie5663fa3ec2f5c7aef7ca0c99a034c879cba05b632b6afd20e7ad1c08ff1c437.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) SubscriptionsById(id string)(*i61ab027a9179fbc39656243c6dcf797fd209a10d47964360d5075fc3811230e8.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i61ab027a9179fbc39656243c6dcf797fd209a10d47964360d5075fc3811230e8.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
