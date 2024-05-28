package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
type RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderGetQueryParameters get the details of all role management policy assignments made in PIM for Microsoft Entra roles and PIM for groups.
type RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderGetQueryParameters struct {
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
// RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderGetQueryParameters
}
// RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleManagementPolicyAssignmentId provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
// returns a *RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder when successful
func (m *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) ByUnifiedRoleManagementPolicyAssignmentId(unifiedRoleManagementPolicyAssignmentId string)(*RolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleManagementPolicyAssignmentId != "" {
        urlTplParams["unifiedRoleManagementPolicyAssignment%2Did"] = unifiedRoleManagementPolicyAssignmentId
    }
    return NewRolemanagementpolicyassignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderInternal instantiates a new RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder and sets the default values.
func NewRolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) {
    m := &RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/roleManagementPolicyAssignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder instantiates a new RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder and sets the default values.
func NewRolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RolemanagementpolicyassignmentsCountRequestBuilder when successful
func (m *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) Count()(*RolemanagementpolicyassignmentsCountRequestBuilder) {
    return NewRolemanagementpolicyassignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the details of all role management policy assignments made in PIM for Microsoft Entra roles and PIM for groups.
// returns a UnifiedRoleManagementPolicyAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/policyroot-list-rolemanagementpolicyassignments?view=graph-rest-1.0
func (m *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleManagementPolicyAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentCollectionResponseable), nil
}
// Post create new navigation property to roleManagementPolicyAssignments for policies
// returns a UnifiedRoleManagementPolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable, requestConfiguration *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get the details of all role management policy assignments made in PIM for Microsoft Entra roles and PIM for groups.
// returns a *RequestInformation when successful
func (m *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleManagementPolicyAssignments for policies
// returns a *RequestInformation when successful
func (m *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleManagementPolicyAssignmentable, requestConfiguration *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder when successful
func (m *RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) WithUrl(rawUrl string)(*RolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder) {
    return NewRolemanagementpolicyassignmentsRoleManagementPolicyAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
