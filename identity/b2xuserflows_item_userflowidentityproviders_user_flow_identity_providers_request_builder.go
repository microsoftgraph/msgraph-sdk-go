package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
type B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderGetQueryParameters get userFlowIdentityProviders from identity
type B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderGetQueryParameters struct {
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
// B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderGetQueryParameters
}
// ByIdentityProviderBaseId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.userFlowIdentityProviders.item collection
// returns a *B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder when successful
func (m *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) ByIdentityProviderBaseId(identityProviderBaseId string)(*B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityProviderBaseId != "" {
        urlTplParams["identityProviderBase%2Did"] = identityProviderBaseId
    }
    return NewB2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewB2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderInternal instantiates a new B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) {
    m := &B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userFlowIdentityProviders{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder instantiates a new B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *B2xuserflowsItemUserflowidentityprovidersCountRequestBuilder when successful
func (m *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) Count()(*B2xuserflowsItemUserflowidentityprovidersCountRequestBuilder) {
    return NewB2xuserflowsItemUserflowidentityprovidersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get userFlowIdentityProviders from identity
// returns a IdentityProviderBaseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityProviderBaseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseCollectionResponseable), nil
}
// Ref provides operations to manage the collection of identityContainer entities.
// returns a *B2xuserflowsItemUserflowidentityprovidersRefRequestBuilder when successful
func (m *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) Ref()(*B2xuserflowsItemUserflowidentityprovidersRefRequestBuilder) {
    return NewB2xuserflowsItemUserflowidentityprovidersRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get userFlowIdentityProviders from identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder when successful
func (m *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) {
    return NewB2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
