package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/acceptances"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/files"
    ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/file"
    ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/acceptances/item"
    id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/files/item"
)

// Builds and executes requests for operations under \agreements\{agreement-id}
type AgreementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AgreementRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AgreementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AgreementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from agreements by key
type AgreementRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AgreementRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Agreement;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AgreementRequestBuilder) Acceptances()(*i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb.AcceptancesRequestBuilder) {
    return i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb.NewAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.agreements.item.acceptances.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AgreementRequestBuilder) AcceptancesById(id string)(*ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991.AgreementAcceptanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance_id"] = id
    }
    return ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991.NewAgreementAcceptanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new AgreementRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAgreementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AgreementRequestBuilder) {
    m := &AgreementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/agreements/{agreement_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AgreementRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAgreementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AgreementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAgreementRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from agreements
// Parameters:
//  - options : Options for the request
func (m *AgreementRequestBuilder) CreateDeleteRequestInformation(options *AgreementRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from agreements by key
// Parameters:
//  - options : Options for the request
func (m *AgreementRequestBuilder) CreateGetRequestInformation(options *AgreementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update entity in agreements
// Parameters:
//  - options : Options for the request
func (m *AgreementRequestBuilder) CreatePatchRequestInformation(options *AgreementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from agreements
// Parameters:
//  - options : Options for the request
func (m *AgreementRequestBuilder) Delete(options *AgreementRequestBuilderDeleteOptions)(error) {
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
func (m *AgreementRequestBuilder) File()(*ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92.FileRequestBuilder) {
    return ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92.NewFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AgreementRequestBuilder) Files()(*i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b.FilesRequestBuilder) {
    return i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b.NewFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.agreements.item.files.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AgreementRequestBuilder) FilesById(id string)(*id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94.AgreementFileLocalizationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileLocalization_id"] = id
    }
    return id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94.NewAgreementFileLocalizationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get entity from agreements by key
// Parameters:
//  - options : Options for the request
func (m *AgreementRequestBuilder) Get(options *AgreementRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Agreement, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAgreement() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Agreement), nil
}
// Update entity in agreements
// Parameters:
//  - options : Options for the request
func (m *AgreementRequestBuilder) Patch(options *AgreementRequestBuilderPatchOptions)(error) {
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
