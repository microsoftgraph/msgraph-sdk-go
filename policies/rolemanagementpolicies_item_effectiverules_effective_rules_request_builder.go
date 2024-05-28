package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder provides operations to manage the effectiveRules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
type RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderGetQueryParameters the list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
type RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderGetQueryParameters struct {
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
// RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderGetQueryParameters
}
// RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleManagementPolicyRuleId provides operations to manage the effectiveRules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
// returns a *RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder when successful
func (m *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) ByUnifiedRoleManagementPolicyRuleId(unifiedRoleManagementPolicyRuleId string)(*RolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleManagementPolicyRuleId != "" {
        urlTplParams["unifiedRoleManagementPolicyRule%2Did"] = unifiedRoleManagementPolicyRuleId
    }
    return NewRolemanagementpoliciesItemEffectiverulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderInternal instantiates a new RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder and sets the default values.
func NewRolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) {
    m := &RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy%2Did}/effectiveRules{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder instantiates a new RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder and sets the default values.
func NewRolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RolemanagementpoliciesItemEffectiverulesCountRequestBuilder when successful
func (m *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) Count()(*RolemanagementpoliciesItemEffectiverulesCountRequestBuilder) {
    return NewRolemanagementpoliciesItemEffectiverulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
// returns a UnifiedRoleManagementPolicyRuleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyRuleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleCollectionResponseable), nil
}
// Post create new navigation property to effectiveRules for policies
// returns a UnifiedRoleManagementPolicyRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, requestConfiguration *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to effectiveRules for policies
// returns a *RequestInformation when successful
func (m *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyRuleable, requestConfiguration *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder when successful
func (m *RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) WithUrl(rawUrl string)(*RolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder) {
    return NewRolemanagementpoliciesItemEffectiverulesEffectiveRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
