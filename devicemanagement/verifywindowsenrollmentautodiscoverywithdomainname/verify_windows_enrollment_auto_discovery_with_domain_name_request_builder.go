package verifywindowsenrollmentautodiscoverywithdomainname

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
type VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetOptions options for Get
type VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, domainName *string)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
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
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function verifyWindowsEnrollmentAutoDiscovery
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) CreateGetRequestInformation(options *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function verifyWindowsEnrollmentAutoDiscovery
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) Get(options *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetOptions)(VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable), nil
}
