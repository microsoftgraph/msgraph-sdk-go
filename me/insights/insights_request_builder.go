package insights

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i29d2c155f681081063d0b2c0ab92796f0e2b78c469f71af45a16d8240b348ebf "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i8893ecb7e9d38ec8961e185f29391303f74779021473d1c37caeafd1abebef9a "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending"
    i99fad623e5b2b4a7ae85326fa692745befba2f6b4fb179d4572c3e7219fbf3ac "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used"
    i7eb746d3d91c620af79236bba2af4b708c6b65d5f4ec8f9cc42e5d4305f1c396 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item"
    ic47fb69b7ecc13ce7f81973ffdde08efe7b70445311d1e83964064ef79497098 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item"
)

type InsightsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type InsightsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewInsightsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InsightsRequestBuilder) {
    m := &InsightsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights{?select,expand}";
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
func NewInsightsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InsightsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInsightsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *InsightsRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *InsightsRequestBuilder) CreateGetRequestInformation(q func (value *InsightsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(InsightsRequestBuilderGetQueryParameters)
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
func (m *InsightsRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsights, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *InsightsRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *InsightsRequestBuilder) Get(q func (value *InsightsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsights, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOfficeGraphInsights() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsights), nil
}
func (m *InsightsRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsights, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *InsightsRequestBuilder) Shared()(*i29d2c155f681081063d0b2c0ab92796f0e2b78c469f71af45a16d8240b348ebf.SharedRequestBuilder) {
    return i29d2c155f681081063d0b2c0ab92796f0e2b78c469f71af45a16d8240b348ebf.NewSharedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InsightsRequestBuilder) SharedById(id string)(*i7eb746d3d91c620af79236bba2af4b708c6b65d5f4ec8f9cc42e5d4305f1c396.SharedInsightRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["sharedInsight_id"] = id
    }
    return i7eb746d3d91c620af79236bba2af4b708c6b65d5f4ec8f9cc42e5d4305f1c396.NewSharedInsightRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InsightsRequestBuilder) Trending()(*i8893ecb7e9d38ec8961e185f29391303f74779021473d1c37caeafd1abebef9a.TrendingRequestBuilder) {
    return i8893ecb7e9d38ec8961e185f29391303f74779021473d1c37caeafd1abebef9a.NewTrendingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InsightsRequestBuilder) TrendingById(id string)(*i8893ecb7e9d38ec8961e185f29391303f74779021473d1c37caeafd1abebef9a.TrendingRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["trending_id"] = id
    }
    return i8893ecb7e9d38ec8961e185f29391303f74779021473d1c37caeafd1abebef9a.NewTrendingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InsightsRequestBuilder) Used()(*i99fad623e5b2b4a7ae85326fa692745befba2f6b4fb179d4572c3e7219fbf3ac.UsedRequestBuilder) {
    return i99fad623e5b2b4a7ae85326fa692745befba2f6b4fb179d4572c3e7219fbf3ac.NewUsedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InsightsRequestBuilder) UsedById(id string)(*ic47fb69b7ecc13ce7f81973ffdde08efe7b70445311d1e83964064ef79497098.UsedInsightRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["usedInsight_id"] = id
    }
    return ic47fb69b7ecc13ce7f81973ffdde08efe7b70445311d1e83964064ef79497098.NewUsedInsightRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
