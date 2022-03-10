package security

import (
    i13bd20348dd8380732465deeb248cdc3f18bf1b888ea708c2518379965331b7d "github.com/microsoftgraph/msgraph-sdk-go/security/securescores"
    ic55f4c3693b3458d37f169972d41697131ecfdb31a625c4fae6e024f02167222 "github.com/microsoftgraph/msgraph-sdk-go/security/alerts"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie4328a6e5437e25d06c50bee3076698796ad30481e9e114d704c631aad9b6f5b "github.com/microsoftgraph/msgraph-sdk-go/security/securescorecontrolprofiles"
    i341438a13fe09ad4284380c5f498c991d5debd415f26d3ad4445315402edc168 "github.com/microsoftgraph/msgraph-sdk-go/security/securescorecontrolprofiles/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i935eca4ee55c30c68a1ba6558e201935fcaf58fd82c8f8076c82de8f806aac5b "github.com/microsoftgraph/msgraph-sdk-go/security/securescores/item"
    i9f4cc7fe9832e758c23f51e8095ff0946d511066f5a7d110c1662df7efadb000 "github.com/microsoftgraph/msgraph-sdk-go/security/alerts/item"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
)

// SecurityRequestBuilder provides operations to manage the security singleton.
type SecurityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SecurityRequestBuilderGetOptions options for Get
type SecurityRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SecurityRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SecurityRequestBuilderGetQueryParameters get security
type SecurityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SecurityRequestBuilderPatchOptions options for Patch
type SecurityRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Securityable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SecurityRequestBuilder) Alerts()(*ic55f4c3693b3458d37f169972d41697131ecfdb31a625c4fae6e024f02167222.AlertsRequestBuilder) {
    return ic55f4c3693b3458d37f169972d41697131ecfdb31a625c4fae6e024f02167222.NewAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.security.alerts.item collection
func (m *SecurityRequestBuilder) AlertsById(id string)(*i9f4cc7fe9832e758c23f51e8095ff0946d511066f5a7d110c1662df7efadb000.AlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert_id"] = id
    }
    return i9f4cc7fe9832e758c23f51e8095ff0946d511066f5a7d110c1662df7efadb000.NewAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSecurityRequestBuilderInternal instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SecurityRequestBuilder) {
    m := &SecurityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityRequestBuilder instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get security
func (m *SecurityRequestBuilder) CreateGetRequestInformation(options *SecurityRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update security
func (m *SecurityRequestBuilder) CreatePatchRequestInformation(options *SecurityRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get security
func (m *SecurityRequestBuilder) Get(options *SecurityRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Securityable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateSecurityFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Securityable), nil
}
// Patch update security
func (m *SecurityRequestBuilder) Patch(options *SecurityRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SecurityRequestBuilder) SecureScoreControlProfiles()(*ie4328a6e5437e25d06c50bee3076698796ad30481e9e114d704c631aad9b6f5b.SecureScoreControlProfilesRequestBuilder) {
    return ie4328a6e5437e25d06c50bee3076698796ad30481e9e114d704c631aad9b6f5b.NewSecureScoreControlProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecureScoreControlProfilesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.security.secureScoreControlProfiles.item collection
func (m *SecurityRequestBuilder) SecureScoreControlProfilesById(id string)(*i341438a13fe09ad4284380c5f498c991d5debd415f26d3ad4445315402edc168.SecureScoreControlProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secureScoreControlProfile_id"] = id
    }
    return i341438a13fe09ad4284380c5f498c991d5debd415f26d3ad4445315402edc168.NewSecureScoreControlProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SecurityRequestBuilder) SecureScores()(*i13bd20348dd8380732465deeb248cdc3f18bf1b888ea708c2518379965331b7d.SecureScoresRequestBuilder) {
    return i13bd20348dd8380732465deeb248cdc3f18bf1b888ea708c2518379965331b7d.NewSecureScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecureScoresById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.security.secureScores.item collection
func (m *SecurityRequestBuilder) SecureScoresById(id string)(*i935eca4ee55c30c68a1ba6558e201935fcaf58fd82c8f8076c82de8f806aac5b.SecureScoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secureScore_id"] = id
    }
    return i935eca4ee55c30c68a1ba6558e201935fcaf58fd82c8f8076c82de8f806aac5b.NewSecureScoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
