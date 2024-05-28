package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder provides operations to manage the passiveDnsReverse property of the microsoft.graph.security.host entity.
type ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderGetQueryParameters get a collection of passiveDnsRecord resources from a reverse passive DNS retrieval.  A reverse DNS lookup queries the hostname of a host using an IP address.
type ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderGetQueryParameters struct {
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
// ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderGetQueryParameters
}
// ByPassiveDnsRecordId provides operations to manage the passiveDnsReverse property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemPassivednsreversePassiveDnsRecordItemRequestBuilder when successful
func (m *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) ByPassiveDnsRecordId(passiveDnsRecordId string)(*ThreatintelligenceHostsItemPassivednsreversePassiveDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if passiveDnsRecordId != "" {
        urlTplParams["passiveDnsRecord%2Did"] = passiveDnsRecordId
    }
    return NewThreatintelligenceHostsItemPassivednsreversePassiveDnsRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderInternal instantiates a new ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) {
    m := &ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/passiveDnsReverse{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder instantiates a new ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceHostsItemPassivednsreverseCountRequestBuilder when successful
func (m *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) Count()(*ThreatintelligenceHostsItemPassivednsreverseCountRequestBuilder) {
    return NewThreatintelligenceHostsItemPassivednsreverseCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a collection of passiveDnsRecord resources from a reverse passive DNS retrieval.  A reverse DNS lookup queries the hostname of a host using an IP address.
// returns a PassiveDnsRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-host-list-passivednsreverse?view=graph-rest-1.0
func (m *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PassiveDnsRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreatePassiveDnsRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PassiveDnsRecordCollectionResponseable), nil
}
// ToGetRequestInformation get a collection of passiveDnsRecord resources from a reverse passive DNS retrieval.  A reverse DNS lookup queries the hostname of a host using an IP address.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder when successful
func (m *ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder) {
    return NewThreatintelligenceHostsItemPassivednsreversePassiveDnsReverseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
