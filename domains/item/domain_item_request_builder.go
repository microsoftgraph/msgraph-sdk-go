package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1881cafcb69c2a03b070eab982c3f40a50f879626e4f5a433adf3d53764e612f "github.com/microsoftgraph/msgraph-sdk-go/domains/item/forcedelete"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    iacad0a6834b5884a8d686d9e187dfb93e979024e3471c842fc86db3dd765352a "github.com/microsoftgraph/msgraph-sdk-go/domains/item/verify"
    icf5b8e1a8b1c94340467eae996f377e8f09f9616367b7d5309238180fac55b36 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/verificationdnsrecords"
    id6628f59c761c8d4beb30bc461b4a9033a596bf14d71905302fee3144bffe677 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/domainnamereferences"
    if6c517fa3faff0b607727add22bace8751787df881a7f48b830e8ee8854493e1 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/serviceconfigurationrecords"
    i4eeff4aea5c653c30cf74b43987a1c46ab8d4723858e572edb5f60ff78363be0 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/domainnamereferences/item"
    i920c7da1a3bc0cc91d66e6345a870d4afe1c3ff7eadb815aed8f74ac233de19e "github.com/microsoftgraph/msgraph-sdk-go/domains/item/verificationdnsrecords/item"
    if238474bb023ee1ae7b1e221337d6649adb30bb8a532c530686588c4d8f149e4 "github.com/microsoftgraph/msgraph-sdk-go/domains/item/serviceconfigurationrecords/item"
)

// DomainItemRequestBuilder provides operations to manage the collection of domain entities.
type DomainItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DomainItemRequestBuilderDeleteOptions options for Delete
type DomainItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DomainItemRequestBuilderGetOptions options for Get
type DomainItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DomainItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DomainItemRequestBuilderGetQueryParameters get entity from domains by key
type DomainItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DomainItemRequestBuilderPatchOptions options for Patch
type DomainItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Domainable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDomainItemRequestBuilderInternal instantiates a new DomainItemRequestBuilder and sets the default values.
func NewDomainItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DomainItemRequestBuilder) {
    m := &DomainItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/domains/{domain_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDomainItemRequestBuilder instantiates a new DomainItemRequestBuilder and sets the default values.
func NewDomainItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DomainItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDomainItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from domains
func (m *DomainItemRequestBuilder) CreateDeleteRequestInformation(options *DomainItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from domains by key
func (m *DomainItemRequestBuilder) CreateGetRequestInformation(options *DomainItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in domains
func (m *DomainItemRequestBuilder) CreatePatchRequestInformation(options *DomainItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from domains
func (m *DomainItemRequestBuilder) Delete(options *DomainItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DomainItemRequestBuilder) DomainNameReferences()(*id6628f59c761c8d4beb30bc461b4a9033a596bf14d71905302fee3144bffe677.DomainNameReferencesRequestBuilder) {
    return id6628f59c761c8d4beb30bc461b4a9033a596bf14d71905302fee3144bffe677.NewDomainNameReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainNameReferencesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.domains.item.domainNameReferences.item collection
func (m *DomainItemRequestBuilder) DomainNameReferencesById(id string)(*i4eeff4aea5c653c30cf74b43987a1c46ab8d4723858e572edb5f60ff78363be0.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i4eeff4aea5c653c30cf74b43987a1c46ab8d4723858e572edb5f60ff78363be0.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DomainItemRequestBuilder) ForceDelete()(*i1881cafcb69c2a03b070eab982c3f40a50f879626e4f5a433adf3d53764e612f.ForceDeleteRequestBuilder) {
    return i1881cafcb69c2a03b070eab982c3f40a50f879626e4f5a433adf3d53764e612f.NewForceDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from domains by key
func (m *DomainItemRequestBuilder) Get(options *DomainItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Domainable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDomainFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Domainable), nil
}
// Patch update entity in domains
func (m *DomainItemRequestBuilder) Patch(options *DomainItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DomainItemRequestBuilder) ServiceConfigurationRecords()(*if6c517fa3faff0b607727add22bace8751787df881a7f48b830e8ee8854493e1.ServiceConfigurationRecordsRequestBuilder) {
    return if6c517fa3faff0b607727add22bace8751787df881a7f48b830e8ee8854493e1.NewServiceConfigurationRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServiceConfigurationRecordsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.domains.item.serviceConfigurationRecords.item collection
func (m *DomainItemRequestBuilder) ServiceConfigurationRecordsById(id string)(*if238474bb023ee1ae7b1e221337d6649adb30bb8a532c530686588c4d8f149e4.DomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord_id"] = id
    }
    return if238474bb023ee1ae7b1e221337d6649adb30bb8a532c530686588c4d8f149e4.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DomainItemRequestBuilder) VerificationDnsRecords()(*icf5b8e1a8b1c94340467eae996f377e8f09f9616367b7d5309238180fac55b36.VerificationDnsRecordsRequestBuilder) {
    return icf5b8e1a8b1c94340467eae996f377e8f09f9616367b7d5309238180fac55b36.NewVerificationDnsRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VerificationDnsRecordsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.domains.item.verificationDnsRecords.item collection
func (m *DomainItemRequestBuilder) VerificationDnsRecordsById(id string)(*i920c7da1a3bc0cc91d66e6345a870d4afe1c3ff7eadb815aed8f74ac233de19e.DomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord_id"] = id
    }
    return i920c7da1a3bc0cc91d66e6345a870d4afe1c3ff7eadb815aed8f74ac233de19e.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DomainItemRequestBuilder) Verify()(*iacad0a6834b5884a8d686d9e187dfb93e979024e3471c842fc86db3dd765352a.VerifyRequestBuilder) {
    return iacad0a6834b5884a8d686d9e187dfb93e979024e3471c842fc86db3dd765352a.NewVerifyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
