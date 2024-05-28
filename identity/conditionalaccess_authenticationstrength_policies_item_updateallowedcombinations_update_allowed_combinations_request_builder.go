package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder provides operations to call the updateAllowedCombinations method.
type ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderInternal instantiates a new ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) {
    m := &ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}/updateAllowedCombinations", pathParameters),
    }
    return m
}
// NewConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder instantiates a new ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the allowedCombinations property of an authenticationStrengthPolicy object. To update other properties of an authenticationStrengthPolicy object, use the Update authenticationStrengthPolicy method.
// returns a UpdateAllowedCombinationsResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-updateallowedcombinations?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) Post(ctx context.Context, body ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsPostRequestBodyable, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateAllowedCombinationsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsPostRequestBodyable, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesItemUpdateallowedcombinationsUpdateAllowedCombinationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
