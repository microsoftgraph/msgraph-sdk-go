package verifywindowsenrollmentautodiscoverywithdomainname

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// verifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder builds and executes requests for operations under \deviceManagement\microsoft.graph.verifyWindowsEnrollmentAutoDiscovery(domainName='{domainName}')
type VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetOptions options for Get
type VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, domainName *string)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    m := &VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.verifyWindowsEnrollmentAutoDiscovery(domainName='{domainName}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if domainName != nil {
        urlTplParams["domainName"] = *domainName
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function verifyWindowsEnrollmentAutoDiscovery
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) CreateGetRequestInformation(options *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function verifyWindowsEnrollmentAutoDiscovery
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) Get(options *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetOptions)(*bool, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "bool", nil)
    if err != nil {
        return nil, err
    }
    return res.(*bool), nil
}
