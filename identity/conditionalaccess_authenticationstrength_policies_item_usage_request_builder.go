package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder provides operations to call the usage method.
type ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilderInternal instantiates a new ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder) {
    m := &ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}/usage()", pathParameters),
    }
    return m
}
// NewConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder instantiates a new ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get allows the caller to see which Conditional Access policies reference a specified authentication strength policy. The policies are returned in two collections, one containing Conditional Access policies that require an MFA claim and the other containing Conditional Access policies that do not require such a claim. Policies in the former category are restricted in what kinds of changes may be made to them to prevent undermining the MFA requirement of those policies.
// returns a AuthenticationStrengthUsageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-usage?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthUsageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationStrengthUsageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationStrengthUsageable), nil
}
// ToGetRequestInformation allows the caller to see which Conditional Access policies reference a specified authentication strength policy. The policies are returned in two collections, one containing Conditional Access policies that require an MFA claim and the other containing Conditional Access policies that do not require such a claim. Policies in the former category are restricted in what kinds of changes may be made to them to prevent undermining the MFA requirement of those policies.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesItemUsageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
