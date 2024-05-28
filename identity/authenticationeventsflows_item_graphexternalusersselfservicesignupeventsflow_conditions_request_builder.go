package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder builds and executes requests for operations under \identity\authenticationEventsFlows\{authenticationEventsFlow-id}\graph.externalUsersSelfServiceSignUpEventsFlow\conditions
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderGetQueryParameters the conditions representing the context of the authentication request that's used to decide whether the events policy is invoked.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderGetQueryParameters
}
// Applications the applications property
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder) Applications()(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder) {
    m := &AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/conditions{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the conditions representing the context of the authentication request that's used to decide whether the events policy is invoked.
// returns a AuthenticationConditionsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationConditionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionsable), nil
}
// ToGetRequestInformation the conditions representing the context of the authentication request that's used to decide whether the events policy is invoked.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
