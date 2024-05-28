package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsItemIdentityprovidersCountRequestBuilder provides operations to count the resources in the collection.
type B2xuserflowsItemIdentityprovidersCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemIdentityprovidersCountRequestBuilderGetQueryParameters get the number of the resource
type B2xuserflowsItemIdentityprovidersCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// B2xuserflowsItemIdentityprovidersCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemIdentityprovidersCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsItemIdentityprovidersCountRequestBuilderGetQueryParameters
}
// NewB2xuserflowsItemIdentityprovidersCountRequestBuilderInternal instantiates a new B2xuserflowsItemIdentityprovidersCountRequestBuilder and sets the default values.
func NewB2xuserflowsItemIdentityprovidersCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemIdentityprovidersCountRequestBuilder) {
    m := &B2xuserflowsItemIdentityprovidersCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/identityProviders/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemIdentityprovidersCountRequestBuilder instantiates a new B2xuserflowsItemIdentityprovidersCountRequestBuilder and sets the default values.
func NewB2xuserflowsItemIdentityprovidersCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemIdentityprovidersCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemIdentityprovidersCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemIdentityprovidersCountRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemIdentityprovidersCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemIdentityprovidersCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemIdentityprovidersCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The identityProvider API is deprecated and will stop returning data on March 2023. Please use the new identityProviderBase API. as of 2021-05/identityProvider
// returns a *B2xuserflowsItemIdentityprovidersCountRequestBuilder when successful
func (m *B2xuserflowsItemIdentityprovidersCountRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemIdentityprovidersCountRequestBuilder) {
    return NewB2xuserflowsItemIdentityprovidersCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
