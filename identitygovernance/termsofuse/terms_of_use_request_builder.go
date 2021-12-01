package termsofuse

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1ce056d01800dfdf0c63c27efbc64f7896ea760db3ba778ace708aeddd054d0f "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreementacceptances"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    iacf793fb78f5c5f75bf9b848a21075f45ec242f4bcf7744ebdca694da4a5802c "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreements"
    ied57f20b5e5c3beb8344ef3753e6a67a240318854397236c0cd83ef4d9597226 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreements/item"
    if58558aaf7d3a571a2b7b154eaf42dbeae3dfda757c15665c189420b3fc2cc8f "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/termsofuse/agreementacceptances/item"
)

// TermsOfUseRequestBuilder builds and executes requests for operations under \identityGovernance\termsOfUse
type TermsOfUseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermsOfUseRequestBuilderDeleteOptions options for Delete
type TermsOfUseRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsOfUseRequestBuilderGetOptions options for Get
type TermsOfUseRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermsOfUseRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsOfUseRequestBuilderGetQueryParameters get termsOfUse from identityGovernance
type TermsOfUseRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TermsOfUseRequestBuilderPatchOptions options for Patch
type TermsOfUseRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsOfUseContainer;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TermsOfUseRequestBuilder) AgreementAcceptances()(*i1ce056d01800dfdf0c63c27efbc64f7896ea760db3ba778ace708aeddd054d0f.AgreementAcceptancesRequestBuilder) {
    return i1ce056d01800dfdf0c63c27efbc64f7896ea760db3ba778ace708aeddd054d0f.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.termsOfUse.agreementAcceptances.item collection
func (m *TermsOfUseRequestBuilder) AgreementAcceptancesById(id string)(*if58558aaf7d3a571a2b7b154eaf42dbeae3dfda757c15665c189420b3fc2cc8f.AgreementAcceptanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance_id"] = id
    }
    return if58558aaf7d3a571a2b7b154eaf42dbeae3dfda757c15665c189420b3fc2cc8f.NewAgreementAcceptanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermsOfUseRequestBuilder) Agreements()(*iacf793fb78f5c5f75bf9b848a21075f45ec242f4bcf7744ebdca694da4a5802c.AgreementsRequestBuilder) {
    return iacf793fb78f5c5f75bf9b848a21075f45ec242f4bcf7744ebdca694da4a5802c.NewAgreementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.termsOfUse.agreements.item collection
func (m *TermsOfUseRequestBuilder) AgreementsById(id string)(*ied57f20b5e5c3beb8344ef3753e6a67a240318854397236c0cd83ef4d9597226.AgreementRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreement_id"] = id
    }
    return ied57f20b5e5c3beb8344ef3753e6a67a240318854397236c0cd83ef4d9597226.NewAgreementRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermsOfUseRequestBuilderInternal instantiates a new TermsOfUseRequestBuilder and sets the default values.
func NewTermsOfUseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsOfUseRequestBuilder) {
    m := &TermsOfUseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/termsOfUse{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermsOfUseRequestBuilder instantiates a new TermsOfUseRequestBuilder and sets the default values.
func NewTermsOfUseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsOfUseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsOfUseRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property termsOfUse for identityGovernance
func (m *TermsOfUseRequestBuilder) CreateDeleteRequestInformation(options *TermsOfUseRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get termsOfUse from identityGovernance
func (m *TermsOfUseRequestBuilder) CreateGetRequestInformation(options *TermsOfUseRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property termsOfUse in identityGovernance
func (m *TermsOfUseRequestBuilder) CreatePatchRequestInformation(options *TermsOfUseRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property termsOfUse for identityGovernance
func (m *TermsOfUseRequestBuilder) Delete(options *TermsOfUseRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get termsOfUse from identityGovernance
func (m *TermsOfUseRequestBuilder) Get(options *TermsOfUseRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsOfUseContainer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTermsOfUseContainer() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsOfUseContainer), nil
}
// Patch update the navigation property termsOfUse in identityGovernance
func (m *TermsOfUseRequestBuilder) Patch(options *TermsOfUseRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
