package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
type RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters retrieve the details of a role management policy.
type RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters
}
// RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderInternal instantiates a new RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder and sets the default values.
func NewRolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) {
    m := &RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder instantiates a new RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder and sets the default values.
func NewRolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleManagementPolicies for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// EffectiveRules provides operations to manage the effectiveRules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
// returns a *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder when successful
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) EffectiveRules()(*RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) {
    return NewRolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the details of a role management policy.
// returns a UnifiedRoleManagementPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedrolemanagementpolicy-get?view=graph-rest-1.0
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable), nil
}
// Patch update the navigation property roleManagementPolicies in policies
// returns a UnifiedRoleManagementPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, requestConfiguration *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable), nil
}
// Rules provides operations to manage the rules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
// returns a *RolemanagementpoliciesItemRulesRequestBuilder when successful
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) Rules()(*RolemanagementpoliciesItemRulesRequestBuilder) {
    return NewRolemanagementpoliciesItemRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleManagementPolicies for policies
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the details of a role management policy.
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleManagementPolicies in policies
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, requestConfiguration *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder when successful
func (m *RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) WithUrl(rawUrl string)(*RolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder) {
    return NewRolemanagementpoliciesUnifiedRoleManagementPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
