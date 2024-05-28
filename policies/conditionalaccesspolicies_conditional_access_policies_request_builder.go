package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
type ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetQueryParameters the custom rules that define an access scenario.
type ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetQueryParameters
}
// ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByConditionalAccessPolicyId provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
// returns a *ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder when successful
func (m *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) ByConditionalAccessPolicyId(conditionalAccessPolicyId string)(*ConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if conditionalAccessPolicyId != "" {
        urlTplParams["conditionalAccessPolicy%2Did"] = conditionalAccessPolicyId
    }
    return NewConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderInternal instantiates a new ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder and sets the default values.
func NewConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) {
    m := &ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/conditionalAccessPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder instantiates a new ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder and sets the default values.
func NewConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConditionalaccesspoliciesCountRequestBuilder when successful
func (m *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) Count()(*ConditionalaccesspoliciesCountRequestBuilder) {
    return NewConditionalaccesspoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the custom rules that define an access scenario.
// returns a ConditionalAccessPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConditionalAccessPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyCollectionResponseable), nil
}
// Post create new navigation property to conditionalAccessPolicies for policies
// returns a ConditionalAccessPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the custom rules that define an access scenario.
// returns a *RequestInformation when successful
func (m *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to conditionalAccessPolicies for policies
// returns a *RequestInformation when successful
func (m *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessPolicyable, requestConfiguration *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder when successful
func (m *ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) {
    return NewConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
