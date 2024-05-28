package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceSubdomainsItemHostRequestBuilder provides operations to manage the host property of the microsoft.graph.security.subdomain entity.
type ThreatintelligenceSubdomainsItemHostRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceSubdomainsItemHostRequestBuilderGetQueryParameters the host of the subdomain.
type ThreatintelligenceSubdomainsItemHostRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceSubdomainsItemHostRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceSubdomainsItemHostRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceSubdomainsItemHostRequestBuilderGetQueryParameters
}
// NewThreatintelligenceSubdomainsItemHostRequestBuilderInternal instantiates a new ThreatintelligenceSubdomainsItemHostRequestBuilder and sets the default values.
func NewThreatintelligenceSubdomainsItemHostRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceSubdomainsItemHostRequestBuilder) {
    m := &ThreatintelligenceSubdomainsItemHostRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/subdomains/{subdomain%2Did}/host{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceSubdomainsItemHostRequestBuilder instantiates a new ThreatintelligenceSubdomainsItemHostRequestBuilder and sets the default values.
func NewThreatintelligenceSubdomainsItemHostRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceSubdomainsItemHostRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceSubdomainsItemHostRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the host of the subdomain.
// returns a Hostable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceSubdomainsItemHostRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceSubdomainsItemHostRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.Hostable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateHostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.Hostable), nil
}
// ToGetRequestInformation the host of the subdomain.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceSubdomainsItemHostRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceSubdomainsItemHostRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ThreatintelligenceSubdomainsItemHostRequestBuilder when successful
func (m *ThreatintelligenceSubdomainsItemHostRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceSubdomainsItemHostRequestBuilder) {
    return NewThreatintelligenceSubdomainsItemHostRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
