package security

import (
    i13bd20348dd8380732465deeb248cdc3f18bf1b888ea708c2518379965331b7d "github.com/microsoftgraph/msgraph-sdk-go/security/securescores"
    ic55f4c3693b3458d37f169972d41697131ecfdb31a625c4fae6e024f02167222 "github.com/microsoftgraph/msgraph-sdk-go/security/alerts"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie4328a6e5437e25d06c50bee3076698796ad30481e9e114d704c631aad9b6f5b "github.com/microsoftgraph/msgraph-sdk-go/security/securescorecontrolprofiles"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i341438a13fe09ad4284380c5f498c991d5debd415f26d3ad4445315402edc168 "github.com/microsoftgraph/msgraph-sdk-go/security/securescorecontrolprofiles/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i935eca4ee55c30c68a1ba6558e201935fcaf58fd82c8f8076c82de8f806aac5b "github.com/microsoftgraph/msgraph-sdk-go/security/securescores/item"
    i9f4cc7fe9832e758c23f51e8095ff0946d511066f5a7d110c1662df7efadb000 "github.com/microsoftgraph/msgraph-sdk-go/security/alerts/item"
)

type SecurityRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SecurityRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *SecurityRequestBuilder) Alerts()(*ic55f4c3693b3458d37f169972d41697131ecfdb31a625c4fae6e024f02167222.AlertsRequestBuilder) {
    return ic55f4c3693b3458d37f169972d41697131ecfdb31a625c4fae6e024f02167222.NewAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SecurityRequestBuilder) AlertsById(id string)(*i9f4cc7fe9832e758c23f51e8095ff0946d511066f5a7d110c1662df7efadb000.AlertRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["alert_id"] = id
    }
    return i9f4cc7fe9832e758c23f51e8095ff0946d511066f5a7d110c1662df7efadb000.NewAlertRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SecurityRequestBuilder) {
    m := &SecurityRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/security{?select,expand}";
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
func NewSecurityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SecurityRequestBuilder) CreateGetRequestInformation(q func (value *SecurityRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SecurityRequestBuilderGetQueryParameters)
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
func (m *SecurityRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Security, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SecurityRequestBuilder) Get(q func (value *SecurityRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Security, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSecurity() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Security), nil
}
func (m *SecurityRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Security, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SecurityRequestBuilder) SecureScoreControlProfiles()(*ie4328a6e5437e25d06c50bee3076698796ad30481e9e114d704c631aad9b6f5b.SecureScoreControlProfilesRequestBuilder) {
    return ie4328a6e5437e25d06c50bee3076698796ad30481e9e114d704c631aad9b6f5b.NewSecureScoreControlProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SecurityRequestBuilder) SecureScoreControlProfilesById(id string)(*i341438a13fe09ad4284380c5f498c991d5debd415f26d3ad4445315402edc168.SecureScoreControlProfileRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["secureScoreControlProfile_id"] = id
    }
    return i341438a13fe09ad4284380c5f498c991d5debd415f26d3ad4445315402edc168.NewSecureScoreControlProfileRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) SecureScores()(*i13bd20348dd8380732465deeb248cdc3f18bf1b888ea708c2518379965331b7d.SecureScoresRequestBuilder) {
    return i13bd20348dd8380732465deeb248cdc3f18bf1b888ea708c2518379965331b7d.NewSecureScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SecurityRequestBuilder) SecureScoresById(id string)(*i935eca4ee55c30c68a1ba6558e201935fcaf58fd82c8f8076c82de8f806aac5b.SecureScoreRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["secureScore_id"] = id
    }
    return i935eca4ee55c30c68a1ba6558e201935fcaf58fd82c8f8076c82de8f806aac5b.NewSecureScoreRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
