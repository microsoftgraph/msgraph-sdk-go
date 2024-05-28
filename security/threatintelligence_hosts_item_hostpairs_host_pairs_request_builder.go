package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder provides operations to manage the hostPairs property of the microsoft.graph.security.host entity.
type ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderGetQueryParameters get the list of hostPair resources associated with a host, where that host is *either* the *parent* or the *child*.
type ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderGetQueryParameters struct {
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
// ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderGetQueryParameters
}
// ByHostPairId provides operations to manage the hostPairs property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemHostpairsHostPairItemRequestBuilder when successful
func (m *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) ByHostPairId(hostPairId string)(*ThreatintelligenceHostsItemHostpairsHostPairItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hostPairId != "" {
        urlTplParams["hostPair%2Did"] = hostPairId
    }
    return NewThreatintelligenceHostsItemHostpairsHostPairItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderInternal instantiates a new ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) {
    m := &ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/hostPairs{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder instantiates a new ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceHostsItemHostpairsCountRequestBuilder when successful
func (m *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) Count()(*ThreatintelligenceHostsItemHostpairsCountRequestBuilder) {
    return NewThreatintelligenceHostsItemHostpairsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the list of hostPair resources associated with a host, where that host is *either* the *parent* or the *child*.
// returns a HostPairCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-host-list-hostpairs?view=graph-rest-1.0
func (m *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostPairCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateHostPairCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostPairCollectionResponseable), nil
}
// ToGetRequestInformation get the list of hostPair resources associated with a host, where that host is *either* the *parent* or the *child*.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder when successful
func (m *ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder) {
    return NewThreatintelligenceHostsItemHostpairsHostPairsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
