package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder provides operations to manage the effectiveRules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
type RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetQueryParameters the list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
type RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetQueryParameters
}
// RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal instantiates a new RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder and sets the default values.
func NewRolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    m := &RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy%2Did}/effectiveRules/{unifiedRoleManagementPolicyRule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder instantiates a new RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder and sets the default values.
func NewRolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property effectiveRules for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
// returns a UnifiedRoleManagementPolicyRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, error) {
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
// Patch update the navigation property effectiveRules in policies
// returns a UnifiedRoleManagementPolicyRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, requestConfiguration *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, error) {
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
// ToDeleteRequestInformation delete navigation property effectiveRules for policies
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property effectiveRules in policies
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, requestConfiguration *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder when successful
func (m *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) WithUrl(rawUrl string)(*RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    return NewRolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
