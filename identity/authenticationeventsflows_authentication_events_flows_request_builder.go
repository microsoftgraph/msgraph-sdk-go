package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder provides operations to manage the authenticationEventsFlows property of the microsoft.graph.identityContainer entity.
type AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderGetQueryParameters get a collection of authentication events policies that are derived from authenticationEventsFlow. The following derived subtypes are supported: - externalUsersSelfServiceSignupEventsFlow
type AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderGetQueryParameters struct {
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
// AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationEventsFlowId provides operations to manage the authenticationEventsFlows property of the microsoft.graph.identityContainer entity.
// returns a *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) ByAuthenticationEventsFlowId(authenticationEventsFlowId string)(*AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationEventsFlowId != "" {
        urlTplParams["authenticationEventsFlow%2Did"] = authenticationEventsFlowId
    }
    return NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderInternal instantiates a new AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) {
    m := &AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder instantiates a new AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationeventsflowsCountRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) Count()(*AuthenticationeventsflowsCountRequestBuilder) {
    return NewAuthenticationeventsflowsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a collection of authentication events policies that are derived from authenticationEventsFlow. The following derived subtypes are supported: - externalUsersSelfServiceSignupEventsFlow
// returns a AuthenticationEventsFlowCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitycontainer-list-authenticationeventsflows?view=graph-rest-1.0
func (m *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationEventsFlowCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowCollectionResponseable), nil
}
// GraphExternalUsersSelfServiceSignUpEventsFlow casts the previous resource to externalUsersSelfServiceSignUpEventsFlow.
// returns a *AuthenticationeventsflowsGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) GraphExternalUsersSelfServiceSignUpEventsFlow()(*AuthenticationeventsflowsGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    return NewAuthenticationeventsflowsGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new authenticationEventsFlow object that is of the type specified in the request body. The following derived subtypes are supported:- externalUsersSelfServiceSignupEventsFlow object type.
// returns a AuthenticationEventsFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitycontainer-post-authenticationeventsflows?view=graph-rest-1.0
func (m *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowable, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationEventsFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowable), nil
}
// ToGetRequestInformation get a collection of authentication events policies that are derived from authenticationEventsFlow. The following derived subtypes are supported: - externalUsersSelfServiceSignupEventsFlow
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new authenticationEventsFlow object that is of the type specified in the request body. The following derived subtypes are supported:- externalUsersSelfServiceSignupEventsFlow object type.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowable, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder) {
    return NewAuthenticationeventsflowsAuthenticationEventsFlowsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
