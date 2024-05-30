package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder provides operations to manage the cookies property of the microsoft.graph.security.host entity.
type ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderGetQueryParameters the hostCookies that are associated with this host.
type ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderGetQueryParameters
}
// NewThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderInternal instantiates a new ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder) {
    m := &ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/cookies/{hostCookie%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder instantiates a new ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the hostCookies that are associated with this host.
// returns a HostCookieable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostCookieable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateHostCookieFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostCookieable), nil
}
// ToGetRequestInformation the hostCookies that are associated with this host.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder when successful
func (m *ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder) {
    return NewThreatintelligenceHostsItemCookiesHostCookieItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
