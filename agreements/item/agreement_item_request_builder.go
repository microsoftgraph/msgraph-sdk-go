package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/acceptances"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/files"
    ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/file"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/acceptances/item"
    id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94 "github.com/microsoftgraph/msgraph-sdk-go/agreements/item/files/item"
)

// AgreementItemRequestBuilder provides operations to manage the collection of agreement entities.
type AgreementItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AgreementItemRequestBuilderDeleteOptions options for Delete
type AgreementItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AgreementItemRequestBuilderGetOptions options for Get
type AgreementItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AgreementItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AgreementItemRequestBuilderGetQueryParameters get entity from agreements by key
type AgreementItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// AgreementItemRequestBuilderPatchOptions options for Patch
type AgreementItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Agreementable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AgreementItemRequestBuilder) Acceptances()(*i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb.AcceptancesRequestBuilder) {
    return i340d5d9b1c93dce1b5a110cf5bbac8114c9a4b58e229c9166305cebe2c92e1cb.NewAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.agreements.item.acceptances.item collection
func (m *AgreementItemRequestBuilder) AcceptancesById(id string)(*ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991.AgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance_id"] = id
    }
    return ic5f55850233142ac7d28f19d006edf4c80dee8f569419b3a03d839aeeaba1991.NewAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAgreementItemRequestBuilderInternal instantiates a new AgreementItemRequestBuilder and sets the default values.
func NewAgreementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AgreementItemRequestBuilder) {
    m := &AgreementItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/agreements/{agreement_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAgreementItemRequestBuilder instantiates a new AgreementItemRequestBuilder and sets the default values.
func NewAgreementItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AgreementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAgreementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from agreements
func (m *AgreementItemRequestBuilder) CreateDeleteRequestInformation(options *AgreementItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from agreements by key
func (m *AgreementItemRequestBuilder) CreateGetRequestInformation(options *AgreementItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in agreements
func (m *AgreementItemRequestBuilder) CreatePatchRequestInformation(options *AgreementItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from agreements
func (m *AgreementItemRequestBuilder) Delete(options *AgreementItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
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
func (m *AgreementItemRequestBuilder) File()(*ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92.FileRequestBuilder) {
    return ibfafeebed2dd285e7305e7475c15f9e74a95731aa8cf08121e9f6a0d20ce2d92.NewFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AgreementItemRequestBuilder) Files()(*i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b.FilesRequestBuilder) {
    return i534405002975c09989779fd4d2b641edcec6af4615b71a008c4aad75a339f35b.NewFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.agreements.item.files.item collection
func (m *AgreementItemRequestBuilder) FilesById(id string)(*id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94.AgreementFileLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileLocalization_id"] = id
    }
    return id334b355b0e9c1a808648dc507c9c2722906e5d1665f782a1e2496b2a183da94.NewAgreementFileLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from agreements by key
func (m *AgreementItemRequestBuilder) Get(options *AgreementItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Agreementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateAgreementFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Agreementable), nil
}
// Patch update entity in agreements
func (m *AgreementItemRequestBuilder) Patch(options *AgreementItemRequestBuilderPatchOptions)(error) {
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
