package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder provides operations to manage the identityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
type B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderGetQueryParameters get the identity providers in a b2xIdentityUserFlow object.
type B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderGetQueryParameters struct {
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
// B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderGetQueryParameters
}
// ByIdentityProviderId provides operations to manage the identityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a *B2xuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder when successful
func (m *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) ByIdentityProviderId(identityProviderId string)(*B2xuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityProviderId != "" {
        urlTplParams["identityProvider%2Did"] = identityProviderId
    }
    return NewB2xuserflowsItemIdentityprovidersIdentityProviderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewB2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderInternal instantiates a new B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder and sets the default values.
func NewB2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) {
    m := &B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/identityProviders{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder instantiates a new B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder and sets the default values.
func NewB2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *B2xuserflowsItemIdentityprovidersCountRequestBuilder when successful
func (m *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) Count()(*B2xuserflowsItemIdentityprovidersCountRequestBuilder) {
    return NewB2xuserflowsItemIdentityprovidersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the identity providers in a b2xIdentityUserFlow object.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a IdentityProviderCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/b2xidentityuserflow-list-identityproviders?view=graph-rest-1.0
func (m *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityProviderCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderCollectionResponseable), nil
}
// ToGetRequestInformation get the identity providers in a b2xIdentityUserFlow object.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder when successful
func (m *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) {
    return NewB2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
