package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder provides operations to manage the assignedUsers property of the microsoft.graph.cloudPcProvisioningPolicyAssignment entity.
type VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderGetQueryParameters the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. Read-only. Supports$expand.
type VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderGetQueryParameters struct {
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
// VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderGetQueryParameters
}
// ByUserId provides operations to manage the assignedUsers property of the microsoft.graph.cloudPcProvisioningPolicyAssignment entity.
// returns a *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersUserItemRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) ByUserId(userId string)(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userId != "" {
        urlTplParams["user%2Did"] = userId
    }
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderInternal instantiates a new VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder and sets the default values.
func NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) {
    m := &VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy%2Did}/assignments/{cloudPcProvisioningPolicyAssignment%2Did}/assignedUsers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder instantiates a new VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder and sets the default values.
func NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersCountRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) Count()(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersCountRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. Read-only. Supports$expand.
// returns a UserCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserCollectionResponseable), nil
}
// ToGetRequestInformation the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. Read-only. Supports$expand.
// returns a *RequestInformation when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder when successful
func (m *VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesItemAssignmentsItemAssignedusersAssignedUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
