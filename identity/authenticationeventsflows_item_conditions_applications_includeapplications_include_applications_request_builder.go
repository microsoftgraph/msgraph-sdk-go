package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder provides operations to manage the includeApplications property of the microsoft.graph.authenticationConditionsApplications entity.
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetQueryParameters list the applications linked to an authenticationEventsFlow. These are the applications for which the authentication experience defined by the user flow is enabled.
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetQueryParameters struct {
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
// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationConditionApplicationAppId provides operations to manage the includeApplications property of the microsoft.graph.authenticationConditionsApplications entity.
// returns a *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) ByAuthenticationConditionApplicationAppId(authenticationConditionApplicationAppId string)(*AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationConditionApplicationAppId != "" {
        urlTplParams["authenticationConditionApplication%2DappId"] = authenticationConditionApplicationAppId
    }
    return NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) {
    m := &AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/conditions/applications/includeApplications{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder instantiates a new AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsCountRequestBuilder when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) Count()(*AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsCountRequestBuilder) {
    return NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the applications linked to an authenticationEventsFlow. These are the applications for which the authentication experience defined by the user flow is enabled.
// returns a AuthenticationConditionApplicationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationconditionsapplications-list-includeapplications?view=graph-rest-1.0
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationConditionApplicationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationCollectionResponseable), nil
}
// Post add or link an application to a user flow, or authenticationEventsFlow. This enables the authentication experience defined by the user flow to be enabled for the application. An application can only be linked to one user flow.
// returns a AuthenticationConditionApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationconditionsapplications-post-includeapplications?view=graph-rest-1.0
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationConditionApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable), nil
}
// ToGetRequestInformation list the applications linked to an authenticationEventsFlow. These are the applications for which the authentication experience defined by the user flow is enabled.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add or link an application to a user flow, or authenticationEventsFlow. This enables the authentication experience defined by the user flow to be enabled for the application. An application can only be linked to one user flow.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, requestConfiguration *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder when successful
func (m *AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) {
    return NewAuthenticationeventsflowsItemConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
