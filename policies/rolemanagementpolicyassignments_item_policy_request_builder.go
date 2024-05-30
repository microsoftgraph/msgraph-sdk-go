package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RolemanagementpolicyassignmentsItemPolicyRequestBuilder provides operations to manage the policy property of the microsoft.graph.unifiedRoleManagementPolicyAssignment entity.
type RolemanagementpolicyassignmentsItemPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementpolicyassignmentsItemPolicyRequestBuilderGetQueryParameters the policy that's associated with a policy assignment. Supports $expand and a nested $expand of the rules and effectiveRules relationships for the policy.
type RolemanagementpolicyassignmentsItemPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementpolicyassignmentsItemPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpolicyassignmentsItemPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementpolicyassignmentsItemPolicyRequestBuilderGetQueryParameters
}
// NewRolemanagementpolicyassignmentsItemPolicyRequestBuilderInternal instantiates a new RolemanagementpolicyassignmentsItemPolicyRequestBuilder and sets the default values.
func NewRolemanagementpolicyassignmentsItemPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpolicyassignmentsItemPolicyRequestBuilder) {
    m := &RolemanagementpolicyassignmentsItemPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/roleManagementPolicyAssignments/{unifiedRoleManagementPolicyAssignment%2Did}/policy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementpolicyassignmentsItemPolicyRequestBuilder instantiates a new RolemanagementpolicyassignmentsItemPolicyRequestBuilder and sets the default values.
func NewRolemanagementpolicyassignmentsItemPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpolicyassignmentsItemPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementpolicyassignmentsItemPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the policy that's associated with a policy assignment. Supports $expand and a nested $expand of the rules and effectiveRules relationships for the policy.
// returns a UnifiedRoleManagementPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpolicyassignmentsItemPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementpolicyassignmentsItemPolicyRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyable, error) {
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
// ToGetRequestInformation the policy that's associated with a policy assignment. Supports $expand and a nested $expand of the rules and effectiveRules relationships for the policy.
// returns a *RequestInformation when successful
func (m *RolemanagementpolicyassignmentsItemPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpolicyassignmentsItemPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RolemanagementpolicyassignmentsItemPolicyRequestBuilder when successful
func (m *RolemanagementpolicyassignmentsItemPolicyRequestBuilder) WithUrl(rawUrl string)(*RolemanagementpolicyassignmentsItemPolicyRequestBuilder) {
    return NewRolemanagementpolicyassignmentsItemPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
