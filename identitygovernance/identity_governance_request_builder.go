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

type IdentityGovernanceRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type IdentityGovernanceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *IdentityGovernanceRequestBuilder) AccessReviews()(*i1cda3b7fdbc22bd911cf3362278cc718553e52184a22fc83dfca175c3aa68d68.AccessReviewsRequestBuilder) {
    return i1cda3b7fdbc22bd911cf3362278cc718553e52184a22fc83dfca175c3aa68d68.NewAccessReviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IdentityGovernanceRequestBuilder) AppConsent()(*i3533004f4d40fc23d0e3aa97288bb203746287e045fd17464a9ddfc34309ec67.AppConsentRequestBuilder) {
    return i3533004f4d40fc23d0e3aa97288bb203746287e045fd17464a9ddfc34309ec67.NewAppConsentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewIdentityGovernanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    m := &IdentityGovernanceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/identityGovernance{?select,expand}";
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
func NewIdentityGovernanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceRequestBuilderInternal(urlParams, requestAdapter)
}
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
