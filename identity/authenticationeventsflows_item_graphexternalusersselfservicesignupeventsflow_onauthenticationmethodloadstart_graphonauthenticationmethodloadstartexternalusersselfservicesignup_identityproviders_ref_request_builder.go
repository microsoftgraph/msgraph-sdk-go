package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder provides operations to manage the collection of identityContainer entities.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderDeleteQueryParameters delete ref of navigation property identityProviders for identity
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderDeleteQueryParameters struct {
    // The delete Uri
    Id *string `uriparametername:"%40id"`
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderDeleteQueryParameters
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderGetQueryParameters get the identity providers that are defined for an external identities self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object type.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) {
    m := &AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/graph.onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/$ref?@id={%40id}{&%24count,%24filter,%24orderby,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property identityProviders for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the identity providers that are defined for an external identities self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object type.
// returns a StringCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onauthenticationmethodloadstartexternalusersselfservicesignup-list-identityproviders?view=graph-rest-1.0
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateStringCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringCollectionResponseable), nil
}
// Post add an identity provider to an external identities self-service user flow represented by an externalUsersSelfServiceSignupEventsFlow object type. The identity provider must first be configured in the tenant.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onauthenticationmethodloadstartexternalusersselfservicesignup-post-identityproviders?view=graph-rest-1.0
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceCreateable, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete ref of navigation property identityProviders for identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/graph.onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/$ref?@id={%40id}", m.BaseRequestBuilder.PathParameters)
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
// ToGetRequestInformation get the identity providers that are defined for an external identities self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object type.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/graph.onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/$ref{?%24count,%24filter,%24orderby,%24search,%24skip,%24top}", m.BaseRequestBuilder.PathParameters)
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
// ToPostRequestInformation add an identity provider to an external identities self-service user flow represented by an externalUsersSelfServiceSignupEventsFlow object type. The identity provider must first be configured in the tenant.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceCreateable, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/graph.onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/$ref", m.BaseRequestBuilder.PathParameters)
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
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnauthenticationmethodloadstartGraphonauthenticationmethodloadstartexternalusersselfservicesignupIdentityprovidersRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
