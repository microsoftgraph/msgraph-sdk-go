package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceHostsHostItemRequestBuilder provides operations to manage the hosts property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceHostsHostItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceHostsHostItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostsHostItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligenceHostsHostItemRequestBuilderGetQueryParameters read the properties and relationships of a host object. The host resource is the abstract base type that returns an implementation. A host can be of one of the following types:
type ThreatintelligenceHostsHostItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceHostsHostItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostsHostItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceHostsHostItemRequestBuilderGetQueryParameters
}
// ThreatintelligenceHostsHostItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostsHostItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChildHostPairs provides operations to manage the childHostPairs property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemChildhostpairsChildHostPairsRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) ChildHostPairs()(*ThreatintelligenceHostsItemChildhostpairsChildHostPairsRequestBuilder) {
    return NewThreatintelligenceHostsItemChildhostpairsChildHostPairsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Components provides operations to manage the components property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemComponentsRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Components()(*ThreatintelligenceHostsItemComponentsRequestBuilder) {
    return NewThreatintelligenceHostsItemComponentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceHostsHostItemRequestBuilderInternal instantiates a new ThreatintelligenceHostsHostItemRequestBuilder and sets the default values.
func NewThreatintelligenceHostsHostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsHostItemRequestBuilder) {
    m := &ThreatintelligenceHostsHostItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceHostsHostItemRequestBuilder instantiates a new ThreatintelligenceHostsHostItemRequestBuilder and sets the default values.
func NewThreatintelligenceHostsHostItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsHostItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceHostsHostItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Cookies provides operations to manage the cookies property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemCookiesRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Cookies()(*ThreatintelligenceHostsItemCookiesRequestBuilder) {
    return NewThreatintelligenceHostsItemCookiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property hosts for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligenceHostsHostItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a host object. The host resource is the abstract base type that returns an implementation. A host can be of one of the following types:
// returns a Hostable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-host-get?view=graph-rest-1.0
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceHostsHostItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.Hostable, error) {
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
// HostPairs provides operations to manage the hostPairs property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) HostPairs()(*ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) {
    return NewThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentHostPairs provides operations to manage the parentHostPairs property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemParenthostpairsParentHostPairsRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) ParentHostPairs()(*ThreatintelligenceHostsItemParenthostpairsParentHostPairsRequestBuilder) {
    return NewThreatintelligenceHostsItemParenthostpairsParentHostPairsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PassiveDns provides operations to manage the passiveDns property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) PassiveDns()(*ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) {
    return NewThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PassiveDnsReverse provides operations to manage the passiveDnsReverse property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) PassiveDnsReverse()(*ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) {
    return NewThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property hosts in security
// returns a Hostable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.Hostable, requestConfiguration *ThreatintelligenceHostsHostItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.Hostable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Ports provides operations to manage the ports property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemPortsRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Ports()(*ThreatintelligenceHostsItemPortsRequestBuilder) {
    return NewThreatintelligenceHostsItemPortsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reputation provides operations to manage the reputation property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemReputationRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Reputation()(*ThreatintelligenceHostsItemReputationRequestBuilder) {
    return NewThreatintelligenceHostsItemReputationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SslCertificates provides operations to manage the sslCertificates property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) SslCertificates()(*ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) {
    return NewThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subdomains provides operations to manage the subdomains property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemSubdomainsRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Subdomains()(*ThreatintelligenceHostsItemSubdomainsRequestBuilder) {
    return NewThreatintelligenceHostsItemSubdomainsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property hosts for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostsHostItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a host object. The host resource is the abstract base type that returns an implementation. A host can be of one of the following types:
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostsHostItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hosts in security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.Hostable, requestConfiguration *ThreatintelligenceHostsHostItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Trackers provides operations to manage the trackers property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemTrackersRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Trackers()(*ThreatintelligenceHostsItemTrackersRequestBuilder) {
    return NewThreatintelligenceHostsItemTrackersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Whois provides operations to manage the whois property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemWhoisRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) Whois()(*ThreatintelligenceHostsItemWhoisRequestBuilder) {
    return NewThreatintelligenceHostsItemWhoisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ThreatintelligenceHostsHostItemRequestBuilder when successful
func (m *ThreatintelligenceHostsHostItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceHostsHostItemRequestBuilder) {
    return NewThreatintelligenceHostsHostItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
