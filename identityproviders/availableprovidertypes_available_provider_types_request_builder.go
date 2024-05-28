package identityproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AvailableprovidertypesAvailableProviderTypesRequestBuilder provides operations to call the availableProviderTypes method.
type AvailableprovidertypesAvailableProviderTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters retrieves all identity provider types available in a directory.
type AvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters
}
// NewAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal instantiates a new AvailableprovidertypesAvailableProviderTypesRequestBuilder and sets the default values.
func NewAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    m := &AvailableprovidertypesAvailableProviderTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProviders/availableProviderTypes(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAvailableprovidertypesAvailableProviderTypesRequestBuilder instantiates a new AvailableprovidertypesAvailableProviderTypesRequestBuilder and sets the default values.
func NewAvailableprovidertypesAvailableProviderTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves all identity provider types available in a directory.
// Deprecated: This method is obsolete. Use GetAsAvailableProviderTypesGetResponse instead.
// returns a AvailableprovidertypesAvailableProviderTypesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityprovider-list-availableprovidertypes?view=graph-rest-1.0
func (m *AvailableprovidertypesAvailableProviderTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *AvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(AvailableprovidertypesAvailableProviderTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAvailableprovidertypesAvailableProviderTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AvailableprovidertypesAvailableProviderTypesResponseable), nil
}
// GetAsAvailableProviderTypesGetResponse retrieves all identity provider types available in a directory.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a AvailableprovidertypesAvailableProviderTypesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityprovider-list-availableprovidertypes?view=graph-rest-1.0
func (m *AvailableprovidertypesAvailableProviderTypesRequestBuilder) GetAsAvailableProviderTypesGetResponse(ctx context.Context, requestConfiguration *AvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(AvailableprovidertypesAvailableProviderTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAvailableprovidertypesAvailableProviderTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AvailableprovidertypesAvailableProviderTypesGetResponseable), nil
}
// ToGetRequestInformation retrieves all identity provider types available in a directory.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a *RequestInformation when successful
func (m *AvailableprovidertypesAvailableProviderTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AvailableprovidertypesAvailableProviderTypesRequestBuilder when successful
func (m *AvailableprovidertypesAvailableProviderTypesRequestBuilder) WithUrl(rawUrl string)(*AvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    return NewAvailableprovidertypesAvailableProviderTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
