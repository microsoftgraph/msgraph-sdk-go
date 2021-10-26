package verifywindowsenrollmentautodiscoverywithdomainname

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// Builds and executes requests for operations under \deviceManagement\microsoft.graph.verifyWindowsEnrollmentAutoDiscovery(domainName='{domainName}')
type VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
// Parameters:
//  - domainName : Usage: domainName={domainName}
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, domainName *string)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    m := &VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/deviceManagement/microsoft.graph.verifyWindowsEnrollmentAutoDiscovery(domainName='{domainName}')";
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
// Instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Invoke function verifyWindowsEnrollmentAutoDiscovery
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) CreateGetRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Invoke function verifyWindowsEnrollmentAutoDiscovery
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) Get(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*bool, error) {
    requestInfo, err := m.CreateGetRequestInformation(h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "bool", responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*bool), nil
}
