package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder provides operations to manage the authenticationEventListeners property of the microsoft.graph.identityContainer entity.
type AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderGetQueryParameters get a list of the authenticationEventListener objects and their properties. The following derived types are supported:
type AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderGetQueryParameters struct {
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
// AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderGetQueryParameters
}
// AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationEventListenerId provides operations to manage the authenticationEventListeners property of the microsoft.graph.identityContainer entity.
// returns a *AuthenticationeventlistenersAuthenticationEventListenerItemRequestBuilder when successful
func (m *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) ByAuthenticationEventListenerId(authenticationEventListenerId string)(*AuthenticationeventlistenersAuthenticationEventListenerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationEventListenerId != "" {
        urlTplParams["authenticationEventListener%2Did"] = authenticationEventListenerId
    }
    return NewAuthenticationeventlistenersAuthenticationEventListenerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationeventlistenersAuthenticationEventListenersRequestBuilderInternal instantiates a new AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder and sets the default values.
func NewAuthenticationeventlistenersAuthenticationEventListenersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) {
    m := &AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventListeners{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationeventlistenersAuthenticationEventListenersRequestBuilder instantiates a new AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder and sets the default values.
func NewAuthenticationeventlistenersAuthenticationEventListenersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventlistenersAuthenticationEventListenersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationeventlistenersCountRequestBuilder when successful
func (m *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) Count()(*AuthenticationeventlistenersCountRequestBuilder) {
    return NewAuthenticationeventlistenersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the authenticationEventListener objects and their properties. The following derived types are supported:
// returns a AuthenticationEventListenerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitycontainer-list-authenticationeventlisteners?view=graph-rest-1.0
func (m *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventListenerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationEventListenerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventListenerCollectionResponseable), nil
}
// Post create a new authenticationEventListener object. You can create one of the following subtypes that are derived from authenticationEventListener.
// returns a AuthenticationEventListenerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitycontainer-post-authenticationeventlisteners?view=graph-rest-1.0
func (m *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventListenerable, requestConfiguration *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventListenerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationEventListenerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventListenerable), nil
}
// ToGetRequestInformation get a list of the authenticationEventListener objects and their properties. The following derived types are supported:
// returns a *RequestInformation when successful
func (m *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new authenticationEventListener object. You can create one of the following subtypes that are derived from authenticationEventListener.
// returns a *RequestInformation when successful
func (m *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventListenerable, requestConfiguration *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder when successful
func (m *AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventlistenersAuthenticationEventListenersRequestBuilder) {
    return NewAuthenticationeventlistenersAuthenticationEventListenersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
