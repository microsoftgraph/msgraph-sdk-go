package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
type ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters settings that may be used to require specific types or instances of an authentication method to be used when authenticating with a specified combination of authentication methods.
type ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters
}
// ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal instantiates a new AuthenticationCombinationConfigurationItemRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    m := &ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}/combinationConfigurations/{authenticationCombinationConfiguration%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder instantiates a new AuthenticationCombinationConfigurationItemRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property combinationConfigurations for identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get settings that may be used to require specific types or instances of an authentication method to be used when authenticating with a specified combination of authentication methods.
func (m *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property combinationConfigurations in identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete navigation property combinationConfigurations for identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation settings that may be used to require specific types or instances of an authentication method to be used when authenticating with a specified combination of authentication methods.
func (m *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property combinationConfigurations in identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
