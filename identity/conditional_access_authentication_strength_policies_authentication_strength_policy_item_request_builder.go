package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder provides operations to manage the policies property of the microsoft.graph.authenticationStrengthRoot entity.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters
}
// ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CombinationConfigurations provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CombinationConfigurations()(*ConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthPoliciesItemCombinationConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal instantiates a new AuthenticationStrengthPolicyItemRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    m := &ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder instantiates a new AuthenticationStrengthPolicyItemRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property policies for identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationStrengthPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable), nil
}
// Patch update the navigation property policies in identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationStrengthPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property policies for identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property policies in identity
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateAllowedCombinations provides operations to call the updateAllowedCombinations method.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) UpdateAllowedCombinations()(*ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Usage provides operations to call the usage method.
func (m *ConditionalAccessAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Usage()(*ConditionalAccessAuthenticationStrengthPoliciesItemUsageRequestBuilder) {
    return NewConditionalAccessAuthenticationStrengthPoliciesItemUsageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
