package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder provides operations to call the asHierarchy method.
type CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderGetQueryParameters list eDiscovery review tags with the tag hierarchy shown.
type CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderGetQueryParameters
}
// NewCasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder) {
    m := &CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/tags/microsoft.graph.security.asHierarchy(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder instantiates a new CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list eDiscovery review tags with the tag hierarchy shown.
// Deprecated: This method is obsolete. Use GetAsAsHierarchyGetResponse instead.
// returns a CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyAsHierarchyResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration)(CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyAsHierarchyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyAsHierarchyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyAsHierarchyResponseable), nil
}
// GetAsAsHierarchyGetResponse list eDiscovery review tags with the tag hierarchy shown.
// returns a CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyAsHierarchyGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder) GetAsAsHierarchyGetResponse(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration)(CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyAsHierarchyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyAsHierarchyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyAsHierarchyGetResponseable), nil
}
// ToGetRequestInformation list eDiscovery review tags with the tag hierarchy shown.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder when successful
func (m *CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder) {
    return NewCasesEdiscoverycasesItemTagsMicrosoftgraphsecurityashierarchyMicrosoftGraphSecurityAsHierarchyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
