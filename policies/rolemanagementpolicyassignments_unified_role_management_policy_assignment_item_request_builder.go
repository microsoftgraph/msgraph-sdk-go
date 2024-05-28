package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
type RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderGetQueryParameters get the details of a policy assignment in PIM that's assigned to Microsoft Entra roles or group membership or ownership.
type RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderGetQueryParameters
}
// RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderInternal instantiates a new RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder and sets the default values.
func NewRolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) {
    m := &RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/roleManagementPolicyAssignments/{unifiedRoleManagementPolicyAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder instantiates a new RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder and sets the default values.
func NewRolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleManagementPolicyAssignments for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the details of a policy assignment in PIM that's assigned to Microsoft Entra roles or group membership or ownership.
// returns a UnifiedRoleManagementPolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedrolemanagementpolicyassignment-get?view=graph-rest-1.0
func (m *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable), nil
}
// Patch update the navigation property roleManagementPolicyAssignments in policies
// returns a UnifiedRoleManagementPolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable, requestConfiguration *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.unifiedRoleManagementPolicyAssignment entity.
// returns a *RolemanagementpolicyassignmentsItemPolicyRequestBuilder when successful
func (m *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) Policy()(*RolemanagementpolicyassignmentsItemPolicyRequestBuilder) {
    return NewRolemanagementpolicyassignmentsItemPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleManagementPolicyAssignments for policies
// returns a *RequestInformation when successful
func (m *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the details of a policy assignment in PIM that's assigned to Microsoft Entra roles or group membership or ownership.
// returns a *RequestInformation when successful
func (m *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleManagementPolicyAssignments in policies
// returns a *RequestInformation when successful
func (m *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable, requestConfiguration *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder when successful
func (m *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) {
    return NewRolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
