package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FederationconfigurationsFederationConfigurationsRequestBuilder provides operations to manage the federationConfigurations property of the microsoft.graph.directory entity.
type FederationconfigurationsFederationConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FederationconfigurationsFederationConfigurationsRequestBuilderGetQueryParameters configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
type FederationconfigurationsFederationConfigurationsRequestBuilderGetQueryParameters struct {
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
// FederationconfigurationsFederationConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FederationconfigurationsFederationConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FederationconfigurationsFederationConfigurationsRequestBuilderGetQueryParameters
}
// FederationconfigurationsFederationConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FederationconfigurationsFederationConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AvailableProviderTypes provides operations to call the availableProviderTypes method.
// returns a *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder when successful
func (m *FederationconfigurationsFederationConfigurationsRequestBuilder) AvailableProviderTypes()(*FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    return NewFederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdentityProviderBaseId provides operations to manage the federationConfigurations property of the microsoft.graph.directory entity.
// returns a *FederationconfigurationsIdentityProviderBaseItemRequestBuilder when successful
func (m *FederationconfigurationsFederationConfigurationsRequestBuilder) ByIdentityProviderBaseId(identityProviderBaseId string)(*FederationconfigurationsIdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityProviderBaseId != "" {
        urlTplParams["identityProviderBase%2Did"] = identityProviderBaseId
    }
    return NewFederationconfigurationsIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFederationconfigurationsFederationConfigurationsRequestBuilderInternal instantiates a new FederationconfigurationsFederationConfigurationsRequestBuilder and sets the default values.
func NewFederationconfigurationsFederationConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FederationconfigurationsFederationConfigurationsRequestBuilder) {
    m := &FederationconfigurationsFederationConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/federationConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFederationconfigurationsFederationConfigurationsRequestBuilder instantiates a new FederationconfigurationsFederationConfigurationsRequestBuilder and sets the default values.
func NewFederationconfigurationsFederationConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FederationconfigurationsFederationConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFederationconfigurationsFederationConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FederationconfigurationsCountRequestBuilder when successful
func (m *FederationconfigurationsFederationConfigurationsRequestBuilder) Count()(*FederationconfigurationsCountRequestBuilder) {
    return NewFederationconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
// returns a IdentityProviderBaseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FederationconfigurationsFederationConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *FederationconfigurationsFederationConfigurationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseCollectionResponseable, error) {
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
// Post create new navigation property to federationConfigurations for directory
// returns a IdentityProviderBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FederationconfigurationsFederationConfigurationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseable, requestConfiguration *FederationconfigurationsFederationConfigurationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityProviderBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseable), nil
}
// ToGetRequestInformation configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
// returns a *RequestInformation when successful
func (m *FederationconfigurationsFederationConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FederationconfigurationsFederationConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to federationConfigurations for directory
// returns a *RequestInformation when successful
func (m *FederationconfigurationsFederationConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseable, requestConfiguration *FederationconfigurationsFederationConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FederationconfigurationsFederationConfigurationsRequestBuilder when successful
func (m *FederationconfigurationsFederationConfigurationsRequestBuilder) WithUrl(rawUrl string)(*FederationconfigurationsFederationConfigurationsRequestBuilder) {
    return NewFederationconfigurationsFederationConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
