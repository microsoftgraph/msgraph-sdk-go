package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
type AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters settings that may be used to require specific types or instances of an authentication method to be used when authenticating with a specified combination of authentication methods.
type AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters
}
// AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal instantiates a new AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder and sets the default values.
func NewAuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    m := &AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationStrengthPolicies/{authenticationStrengthPolicy%2Did}/combinationConfigurations/{authenticationCombinationConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder instantiates a new AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder and sets the default values.
func NewAuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property combinationConfigurations for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get settings that may be used to require specific types or instances of an authentication method to be used when authenticating with a specified combination of authentication methods.
// returns a AuthenticationCombinationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationCombinationConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable), nil
}
// Patch update the navigation property combinationConfigurations in policies
// returns a AuthenticationCombinationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, requestConfiguration *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationCombinationConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property combinationConfigurations for policies
// returns a *RequestInformation when successful
func (m *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation settings that may be used to require specific types or instances of an authentication method to be used when authenticating with a specified combination of authentication methods.
// returns a *RequestInformation when successful
func (m *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property combinationConfigurations in policies
// returns a *RequestInformation when successful
func (m *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, requestConfiguration *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder when successful
func (m *AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    return NewAuthenticationstrengthpoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
