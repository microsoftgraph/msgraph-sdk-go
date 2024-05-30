package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder provides operations to manage the authenticationEventsFlows property of the microsoft.graph.identityContainer entity.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a specific authenticationEventsFlow object by ID. The @odata.type property in the response object indicates the type of the object, which can be one of the following derived subtypes:- externalUsersSelfServiceSignupEventsFlow
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Conditions the conditions property
// returns a *AuthenticationeventsflowsItemConditionsRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) Conditions()(*AuthenticationeventsflowsItemConditionsRequestBuilder) {
    return NewAuthenticationeventsflowsItemConditionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderInternal instantiates a new AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) {
    m := &AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder instantiates a new AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a specific authenticationEventsFlow resource by ID. This operation also removes or unlinks all applications from the flow, which disables the customized authentication experience defined for the application.  The following derived subtypes are supported:- externalUsersSelfServiceSignupEventsFlow
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationeventsflow-delete?view=graph-rest-1.0
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties and relationships of a specific authenticationEventsFlow object by ID. The @odata.type property in the response object indicates the type of the object, which can be one of the following derived subtypes:- externalUsersSelfServiceSignupEventsFlow
// returns a AuthenticationEventsFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationeventsflow-get?view=graph-rest-1.0
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// GraphExternalUsersSelfServiceSignUpEventsFlow casts the previous resource to externalUsersSelfServiceSignUpEventsFlow.
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) GraphExternalUsersSelfServiceSignUpEventsFlow()(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowGraphExternalUsersSelfServiceSignUpEventsFlowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an authenticationEventsFlow object by ID. You must specify the @odata.type property and the value of the authenticationEventsFlow object type to update. The following derived subtypes are supported:- externalUsersSelfServiceSignupEventsFlow
// returns a AuthenticationEventsFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationeventsflow-update?view=graph-rest-1.0
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowable, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete a specific authenticationEventsFlow resource by ID. This operation also removes or unlinks all applications from the flow, which disables the customized authentication experience defined for the application.  The following derived subtypes are supported:- externalUsersSelfServiceSignupEventsFlow
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a specific authenticationEventsFlow object by ID. The @odata.type property in the response object indicates the type of the object, which can be one of the following derived subtypes:- externalUsersSelfServiceSignupEventsFlow
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an authenticationEventsFlow object by ID. You must specify the @odata.type property and the value of the authenticationEventsFlow object type to update. The following derived subtypes are supported:- externalUsersSelfServiceSignupEventsFlow
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationEventsFlowable, requestConfiguration *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder when successful
func (m *AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder) {
    return NewAuthenticationeventsflowsAuthenticationEventsFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
