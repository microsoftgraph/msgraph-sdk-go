package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
type VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal instantiates a new VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, domainName *string)(*VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    m := &VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/verifyWindowsEnrollmentAutoDiscovery(domainName='{domainName}')", pathParameters),
    }
    if domainName != nil {
        m.BaseRequestBuilder.PathParameters["domainName"] = *domainName
    }
    return m
}
// NewVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder instantiates a new VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder and sets the default values.
func NewVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function verifyWindowsEnrollmentAutoDiscovery
// Deprecated: This method is obsolete. Use GetAsVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponse instead.
// returns a VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) Get(ctx context.Context, requestConfiguration *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration)(VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseable), nil
}
// GetAsVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponse invoke function verifyWindowsEnrollmentAutoDiscovery
// returns a VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) GetAsVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponse(ctx context.Context, requestConfiguration *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration)(VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameGetResponseable), nil
}
// ToGetRequestInformation invoke function verifyWindowsEnrollmentAutoDiscovery
// returns a *RequestInformation when successful
func (m *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder when successful
func (m *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) WithUrl(rawUrl string)(*VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return NewVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
