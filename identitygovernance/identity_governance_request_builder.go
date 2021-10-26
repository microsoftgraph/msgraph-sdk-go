package identitygovernance

import (
    i19843957d5edfe5fb320ce126acd252770b9df2ea52037acd5d5f2308496e3fb "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse"
    i1cda3b7fdbc22bd911cf3362278cc718553e52184a22fc83dfca175c3aa68d68 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/accessreviews"
    i3533004f4d40fc23d0e3aa97288bb203746287e045fd17464a9ddfc34309ec67 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/appconsent"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ife915c8a8f434e475a7b17efdcb242db3e0f1f1ef327268189874360082827ec "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \identityGovernance
type IdentityGovernanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get identityGovernance
type IdentityGovernanceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *IdentityGovernanceRequestBuilder) AccessReviews()(*i1cda3b7fdbc22bd911cf3362278cc718553e52184a22fc83dfca175c3aa68d68.AccessReviewsRequestBuilder) {
    return i1cda3b7fdbc22bd911cf3362278cc718553e52184a22fc83dfca175c3aa68d68.NewAccessReviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IdentityGovernanceRequestBuilder) AppConsent()(*i3533004f4d40fc23d0e3aa97288bb203746287e045fd17464a9ddfc34309ec67.AppConsentRequestBuilder) {
    return i3533004f4d40fc23d0e3aa97288bb203746287e045fd17464a9ddfc34309ec67.NewAppConsentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewIdentityGovernanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    m := &IdentityGovernanceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/identityGovernance{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewIdentityGovernanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Get identityGovernance
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *IdentityGovernanceRequestBuilder) CreateGetRequestInformation(q func (value *IdentityGovernanceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(IdentityGovernanceRequestBuilderGetQueryParameters)
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
// Update identityGovernance
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *IdentityGovernanceRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityGovernance, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *IdentityGovernanceRequestBuilder) EntitlementManagement()(*ife915c8a8f434e475a7b17efdcb242db3e0f1f1ef327268189874360082827ec.EntitlementManagementRequestBuilder) {
    return ife915c8a8f434e475a7b17efdcb242db3e0f1f1ef327268189874360082827ec.NewEntitlementManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get identityGovernance
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *IdentityGovernanceRequestBuilder) Get(q func (value *IdentityGovernanceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityGovernance, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewIdentityGovernance() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityGovernance), nil
}
// Update identityGovernance
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *IdentityGovernanceRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityGovernance, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *IdentityGovernanceRequestBuilder) TermsOfUse()(*i19843957d5edfe5fb320ce126acd252770b9df2ea52037acd5d5f2308496e3fb.TermsOfUseRequestBuilder) {
    return i19843957d5edfe5fb320ce126acd252770b9df2ea52037acd5d5f2308496e3fb.NewTermsOfUseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
