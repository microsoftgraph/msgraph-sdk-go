package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder provides operations to manage the rules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
type RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetQueryParameters retrieve a rule or settings defined for a role management policy. The rule can be one of the following types that are derived from the unifiedRoleManagementPolicyRule object:
type RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetQueryParameters
}
// RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal instantiates a new RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder and sets the default values.
func NewRolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    m := &RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy%2Did}/rules/{unifiedRoleManagementPolicyRule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder instantiates a new RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder and sets the default values.
func NewRolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property rules for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve a rule or settings defined for a role management policy. The rule can be one of the following types that are derived from the unifiedRoleManagementPolicyRule object:
// returns a UnifiedRoleManagementPolicyRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedrolemanagementpolicyrule-get?view=graph-rest-1.0
func (m *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable), nil
}
// Patch update a rule defined for a role management policy. The rule can be one of the following types that are derived from the unifiedRoleManagementPolicyRule object: For more information about rules for Microsoft Entra roles and examples of updating rules, see the following articles:
// returns a UnifiedRoleManagementPolicyRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedrolemanagementpolicyrule-update?view=graph-rest-1.0
func (m *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, requestConfiguration *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable), nil
}
// ToDeleteRequestInformation delete navigation property rules for policies
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a rule or settings defined for a role management policy. The rule can be one of the following types that are derived from the unifiedRoleManagementPolicyRule object:
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update a rule defined for a role management policy. The rule can be one of the following types that are derived from the unifiedRoleManagementPolicyRule object: For more information about rules for Microsoft Entra roles and examples of updating rules, see the following articles:
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, requestConfiguration *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder when successful
func (m *RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) WithUrl(rawUrl string)(*RolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    return NewRolemanagementpoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
