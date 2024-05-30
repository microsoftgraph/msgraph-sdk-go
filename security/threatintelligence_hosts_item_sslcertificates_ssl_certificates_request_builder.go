package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder provides operations to manage the sslCertificates property of the microsoft.graph.security.host entity.
type ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderGetQueryParameters get a list of hostSslCertificate objects from the host navigation property.
type ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderGetQueryParameters
}
// ByHostSslCertificateId provides operations to manage the sslCertificates property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemSslcertificatesHostSslCertificateItemRequestBuilder when successful
func (m *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) ByHostSslCertificateId(hostSslCertificateId string)(*ThreatintelligenceHostsItemSslcertificatesHostSslCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hostSslCertificateId != "" {
        urlTplParams["hostSslCertificate%2Did"] = hostSslCertificateId
    }
    return NewThreatintelligenceHostsItemSslcertificatesHostSslCertificateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderInternal instantiates a new ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) {
    m := &ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/sslCertificates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder instantiates a new ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceHostsItemSslcertificatesCountRequestBuilder when successful
func (m *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) Count()(*ThreatintelligenceHostsItemSslcertificatesCountRequestBuilder) {
    return NewThreatintelligenceHostsItemSslcertificatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of hostSslCertificate objects from the host navigation property.
// returns a HostSslCertificateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-host-list-sslcertificates?view=graph-rest-1.0
func (m *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostSslCertificateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateHostSslCertificateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostSslCertificateCollectionResponseable), nil
}
// ToGetRequestInformation get a list of hostSslCertificate objects from the host navigation property.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder when successful
func (m *ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder) {
    return NewThreatintelligenceHostsItemSslcertificatesSslCertificatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
