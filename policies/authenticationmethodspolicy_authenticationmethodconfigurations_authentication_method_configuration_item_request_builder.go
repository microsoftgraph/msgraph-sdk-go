package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder provides operations to manage the authenticationMethodConfigurations property of the microsoft.graph.authenticationMethodsPolicy entity.
type AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
type AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters
}
// AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal instantiates a new AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder and sets the default values.
func NewAuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) {
    m := &AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/{authenticationMethodConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder instantiates a new AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder and sets the default values.
func NewAuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property authenticationMethodConfigurations for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
// returns a AuthenticationMethodConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationMethodConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable), nil
}
// Patch update the navigation property authenticationMethodConfigurations in policies
// returns a AuthenticationMethodConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable, requestConfiguration *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationMethodConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property authenticationMethodConfigurations for policies
// returns a *RequestInformation when successful
func (m *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
// returns a *RequestInformation when successful
func (m *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property authenticationMethodConfigurations in policies
// returns a *RequestInformation when successful
func (m *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable, requestConfiguration *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder when successful
func (m *AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder) {
    return NewAuthenticationmethodspolicyAuthenticationmethodconfigurationsAuthenticationMethodConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
