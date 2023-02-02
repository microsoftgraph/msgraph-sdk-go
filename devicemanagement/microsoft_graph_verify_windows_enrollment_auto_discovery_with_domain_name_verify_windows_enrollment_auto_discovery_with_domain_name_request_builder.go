package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
type MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewMicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, domainName *string)(*MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    m := &MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.verifyWindowsEnrollmentAutoDiscovery(domainName='{domainName}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if domainName != nil {
        urlTplParams["domainName"] = *domainName
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder instantiates a new VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewMicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function verifyWindowsEnrollmentAutoDiscovery
func (m *MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration)(MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable), nil
}
// ToGetRequestInformation invoke function verifyWindowsEnrollmentAutoDiscovery
func (m *MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
