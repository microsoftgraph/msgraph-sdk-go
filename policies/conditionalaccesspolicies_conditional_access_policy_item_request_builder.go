package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
type ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetQueryParameters the custom rules that define an access scenario.
type ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetQueryParameters
}
// ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderInternal instantiates a new ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder and sets the default values.
func NewConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) {
    m := &ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/conditionalAccessPolicies/{conditionalAccessPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder instantiates a new ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder and sets the default values.
func NewConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property conditionalAccessPolicies for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the custom rules that define an access scenario.
// returns a ConditionalAccessPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConditionalAccessPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable), nil
}
// Patch update the navigation property conditionalAccessPolicies in policies
// returns a ConditionalAccessPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConditionalAccessPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property conditionalAccessPolicies for policies
// returns a *RequestInformation when successful
func (m *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the custom rules that define an access scenario.
// returns a *RequestInformation when successful
func (m *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property conditionalAccessPolicies in policies
// returns a *RequestInformation when successful
func (m *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder when successful
func (m *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) {
    return NewConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
