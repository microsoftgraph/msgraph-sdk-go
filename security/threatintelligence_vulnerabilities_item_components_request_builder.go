package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder provides operations to manage the components property of the microsoft.graph.security.vulnerability entity.
type ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderGetQueryParameters get a list of vulnerabilityComponent objects related to a known vulnerability.
type ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderGetQueryParameters struct {
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
// ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderGetQueryParameters
}
// ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByVulnerabilityComponentId provides operations to manage the components property of the microsoft.graph.security.vulnerability entity.
// returns a *ThreatintelligenceVulnerabilitiesItemComponentsVulnerabilityComponentItemRequestBuilder when successful
func (m *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) ByVulnerabilityComponentId(vulnerabilityComponentId string)(*ThreatintelligenceVulnerabilitiesItemComponentsVulnerabilityComponentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if vulnerabilityComponentId != "" {
        urlTplParams["vulnerabilityComponent%2Did"] = vulnerabilityComponentId
    }
    return NewThreatintelligenceVulnerabilitiesItemComponentsVulnerabilityComponentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderInternal instantiates a new ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder and sets the default values.
func NewThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) {
    m := &ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/vulnerabilities/{vulnerability%2Did}/components{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder instantiates a new ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder and sets the default values.
func NewThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceVulnerabilitiesItemComponentsCountRequestBuilder when successful
func (m *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) Count()(*ThreatintelligenceVulnerabilitiesItemComponentsCountRequestBuilder) {
    return NewThreatintelligenceVulnerabilitiesItemComponentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of vulnerabilityComponent objects related to a known vulnerability.
// returns a VulnerabilityComponentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-vulnerability-list-components?view=graph-rest-1.0
func (m *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.VulnerabilityComponentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateVulnerabilityComponentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.VulnerabilityComponentCollectionResponseable), nil
}
// Post create new navigation property to components for security
// returns a VulnerabilityComponentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) Post(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.VulnerabilityComponentable, requestConfiguration *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.VulnerabilityComponentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateVulnerabilityComponentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.VulnerabilityComponentable), nil
}
// ToGetRequestInformation get a list of vulnerabilityComponent objects related to a known vulnerability.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to components for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.VulnerabilityComponentable, requestConfiguration *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder when successful
func (m *ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder) {
    return NewThreatintelligenceVulnerabilitiesItemComponentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
