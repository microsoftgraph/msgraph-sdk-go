package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3a3e8e75f2c335a1ade1f4b7244f0e83a16d386f968a09c4fdbe0552226fd238 "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/item/history"
    ibab4d69cc173d2d5537e2f7f752b3b1830737eb656e2025616be95afe72a46f1 "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/item/history/item"
)

// RiskyUserItemRequestBuilder builds and executes requests for operations under \identityProtection\riskyUsers\{riskyUser-id}
type RiskyUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RiskyUserItemRequestBuilderDeleteOptions options for Delete
type RiskyUserItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RiskyUserItemRequestBuilderGetOptions options for Get
type RiskyUserItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RiskyUserItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RiskyUserItemRequestBuilderGetQueryParameters users that are flagged as at-risk by Azure AD Identity Protection.
type RiskyUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RiskyUserItemRequestBuilderPatchOptions options for Patch
type RiskyUserItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RiskyUser;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRiskyUserItemRequestBuilderInternal instantiates a new RiskyUserItemRequestBuilder and sets the default values.
func NewRiskyUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RiskyUserItemRequestBuilder) {
    m := &RiskyUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyUsers/{riskyUser_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRiskyUserItemRequestBuilder instantiates a new RiskyUserItemRequestBuilder and sets the default values.
func NewRiskyUserItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RiskyUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUserItemRequestBuilder) CreateDeleteRequestInformation(options *RiskyUserItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUserItemRequestBuilder) CreateGetRequestInformation(options *RiskyUserItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUserItemRequestBuilder) CreatePatchRequestInformation(options *RiskyUserItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUserItemRequestBuilder) Delete(options *RiskyUserItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUserItemRequestBuilder) Get(options *RiskyUserItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RiskyUser, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewRiskyUser() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RiskyUser), nil
}
func (m *RiskyUserItemRequestBuilder) History()(*i3a3e8e75f2c335a1ade1f4b7244f0e83a16d386f968a09c4fdbe0552226fd238.HistoryRequestBuilder) {
    return i3a3e8e75f2c335a1ade1f4b7244f0e83a16d386f968a09c4fdbe0552226fd238.NewHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HistoryById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityProtection.riskyUsers.item.history.item collection
func (m *RiskyUserItemRequestBuilder) HistoryById(id string)(*ibab4d69cc173d2d5537e2f7f752b3b1830737eb656e2025616be95afe72a46f1.RiskyUserHistoryItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyUserHistoryItem_id"] = id
    }
    return ibab4d69cc173d2d5537e2f7f752b3b1830737eb656e2025616be95afe72a46f1.NewRiskyUserHistoryItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUserItemRequestBuilder) Patch(options *RiskyUserItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
