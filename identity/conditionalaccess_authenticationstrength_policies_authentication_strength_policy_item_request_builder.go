package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder provides operations to manage the policies property of the microsoft.graph.authenticationStrengthRoot entity.
type ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
type ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetQueryParameters
}
// ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CombinationConfigurations provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) CombinationConfigurations()(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal instantiates a new ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    m := &ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder instantiates a new ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property policies for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
// returns a AuthenticationStrengthPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a AuthenticationStrengthPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of authentication strength policies that exist for this tenant, including both built-in and custom policies.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property policies in identity
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthPolicyable, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateAllowedCombinations provides operations to call the updateAllowedCombinations method.
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) UpdateAllowedCombinations()(*ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Usage provides operations to call the usage method.
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) Usage()(*ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
