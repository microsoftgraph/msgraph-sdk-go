package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder provides operations to call the getManagedAppPolicies method.
type ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderGetQueryParameters gets app restrictions for a given user.
type ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderGetQueryParameters struct {
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
// ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderGetQueryParameters
}
// NewItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderInternal instantiates a new ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder and sets the default values.
func NewItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder) {
    m := &ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/getManagedAppPolicies(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder instantiates a new ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder and sets the default values.
func NewItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets app restrictions for a given user.
// Deprecated: This method is obsolete. Use GetAsGetManagedAppPoliciesGetResponse instead.
// returns a ItemGetmanagedapppoliciesGetManagedAppPoliciesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-user-getmanagedapppolicies?view=graph-rest-1.0
func (m *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderGetRequestConfiguration)(ItemGetmanagedapppoliciesGetManagedAppPoliciesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmanagedapppoliciesGetManagedAppPoliciesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmanagedapppoliciesGetManagedAppPoliciesResponseable), nil
}
// GetAsGetManagedAppPoliciesGetResponse gets app restrictions for a given user.
// returns a ItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-user-getmanagedapppolicies?view=graph-rest-1.0
func (m *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder) GetAsGetManagedAppPoliciesGetResponse(ctx context.Context, requestConfiguration *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderGetRequestConfiguration)(ItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponseable), nil
}
// ToGetRequestInformation gets app restrictions for a given user.
// returns a *RequestInformation when successful
func (m *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder when successful
func (m *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder) {
    return NewItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
