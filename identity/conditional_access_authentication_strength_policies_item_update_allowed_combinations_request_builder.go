package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder provides operations to call the updateAllowedCombinations method.
type ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilderInternal instantiates a new UpdateAllowedCombinationsRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder) {
    m := &ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}/updateAllowedCombinations", pathParameters),
    }
    return m
}
// NewConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder instantiates a new UpdateAllowedCombinationsRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/authenticationstrengthpolicy-updateallowedcombinations?view=graph-rest-1.0
func (m *ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder) Post(ctx context.Context, body ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsPostRequestBodyable, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateAllowedCombinationsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUpdateAllowedCombinationsResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateAllowedCombinationsResultable), nil
}
// ToPostRequestInformation update the allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
func (m *ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsPostRequestBodyable, requestConfiguration *ConditionalAccessAuthenticationStrengthPoliciesItemUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
